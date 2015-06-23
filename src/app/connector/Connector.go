package connector

import (
	"common/config"
	"common/logger"
	"protobuf"
	"component/db"
	"component/rpc"
	"component/server"
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"time"
	"runtime/debug"
	"sync"
	"common"
	"fmt"
)

type serverInfo struct {
	PlayerCount uint16
	ServerIp    string
}

type Connector struct {
	m            map[uint32]serverInfo
	stableServer string
	authserver  *rpc.Client
	gateserver   *rpc.Client
	FsMgr        FServerConnMgr
	rpcServer    *server.Server
	players         map[uint64]*Player
	playersbyid     map[string]*Player
	l               sync.RWMutex
	id              uint32
	listenTcpIp        string
	listenHttpIp        string
}

var pConnector *Connector

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsServeConnHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintln(w, "rpc: error")
		return
	}

	logger.Debug("wsServeConnHandler : %v", r.FormValue("method"))

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Info("Upgrade:", err.Error())
		conn, bufrw, err := w.(http.Hijacker).Hijack()
		if err != nil {
			logger.Debug("rpc hijacking %v : %v", r.RemoteAddr, err.Error())
			return
		}else {

			httpConn := server.NewTCPSocketConn(pConnector.rpcServer, conn, 1, 1, 1)
			logger.Debug("rpc hijacking %v : %v",  r.RemoteAddr, r.FormValue("method"))

			fmt.Fprintln(w, "rpc: hello")

			buf := make([]byte, 10)
			for i := 0; i < len(buf); i++ {
				buf[i] = byte(i)
			}

			_, err = bufrw.Write(buf)
			if err == nil {
				err = bufrw.Flush()
			}
			if err != nil {
				fmt.Printf("ResponseError: %s\\n", err)
			} else {
				fmt.Println("Bye, Jack!")
			}

			logger.Debug("httpConn WriteObj %v", httpConn.GetRemoteIp())
		}
		return
	}

	rpcConn := server.NewWebSocketConn(pConnector.rpcServer, *conn, 128, 45, 2)
	defer func() {
		rpcConn.Close() // 客户端退出减去计数
	}()

	pConnector.rpcServer.ServeConn(rpcConn)
}

func CreateConnectorServerForClient(cfg *config.SvrConfig) *Connector {

	db.Init()

	var authCfg config.AuthConfig
	if err := config.ReadConfig("etc/authserver.json", &authCfg); err != nil {
		logger.Fatal("load config failed, error is: %v", err)
	}
	authConn, err := net.Dial("tcp", authCfg.AuthHost)
	if err != nil {
		logger.Fatal("connect logserver failed %s", err.Error())
	}

	var gsCfg config.GateConfig
	if err = config.ReadConfig("etc/gateserver.json", &gsCfg); err != nil {
		logger.Fatal("load config failed, error is: %v", err)
	}
	gsConn, err := net.Dial("tcp", gsCfg.GateHost)
	if err != nil {
		logger.Fatal("%s", err.Error())
	}

	pConnector = &Connector{
		m:           make(map[uint32]serverInfo),
		authserver: rpc.NewClient(authConn),
		gateserver: rpc.NewClient(gsConn),
		rpcServer:   server.NewServer(),
		players:       make(map[uint64]*Player),
		playersbyid:   make(map[string]*Player),
	}

	pConnector.rpcServer.Register(pConnector)

	pConnector.rpcServer.RegCallBackOnConn(
	func(conn server.RpcConn) {
		pConnector.onConn(conn)
	},
	)

	pConnector.rpcServer.RegCallBackOnDisConn(
	func(conn server.RpcConn) {
		pConnector.onDisConn(conn)
	},
	)

	pConnector.rpcServer.RegCallBackOnCallBefore(
	func(conn server.RpcConn) {
		conn.Lock()
	},
	)

	pConnector.rpcServer.RegCallBackOnCallAfter(
	func(conn server.RpcConn) {
		conn.Unlock()
	},
	)

	//开始对fightserver的RPC服务
	pConnector.FsMgr.Init(pConnector.rpcServer, cfg)

	listener, err := net.Listen("tcp", cfg.TcpHost)
	if err != nil {
		logger.Fatal("net.Listen: %s", err.Error())
	}

	pConnector.id = cfg.ServerID
	pConnector.listenTcpIp = cfg.TcpHost
	pConnector.listenHttpIp = cfg.HttpHost

	pConnector.sendPlayerCountToGateServer()

	go func() {
		for {
			//For Client/////////////////////////////
			time.Sleep(time.Millisecond * 5)
			conn, err := listener.Accept()

			if err != nil {
				logger.Error("cns StartServices %s", err.Error())
				break
			}

			go func() {
				rpcConn := server.NewTCPSocketConn(pConnector.rpcServer, conn, 128, 45, 1)
				defer func() {
					if r := recover(); r != nil {
						logger.Error("player rpc runtime error begin:", r)
						debug.PrintStack()
						rpcConn.Close()

						logger.Error("player rpc runtime error end ")
					}
				}()

				pConnector.rpcServer.ServeConn(rpcConn)
			}()
		}
	}()

	http.HandleFunc("/", wsServeConnHandler)
	http.ListenAndServe(cfg.HttpHost, nil)

	return pConnector
}

func (self *Connector) sendPlayerCountToGateServer() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				logger.Info("sendPlayerCountToGateServer runtime error:", r)

				debug.PrintStack()
			}
		}()

		for {

			time.Sleep(5 * time.Second)

			self.l.RLock()
			playerCount := uint32(len(self.players))
			self.l.RUnlock()

			var ret protobuf.ConnectorInfoResult
			req := protobuf.ConnectorInfo{}
			req.SetServerId(self.id)
			req.SetPlayerCount(playerCount)
			req.SetTcpServerIp(self.listenTcpIp)
			req.SetHttpServerIp(self.listenHttpIp)

			//logger.Debug("playerCount %v", playerCount)

			err := self.gateserver.Call("GateServices.UpdateCnsPlayerCount", &req, &ret)

			if err != nil {
				logger.Error("Error On GateServices.UpdateCnsPlayerCount : %s", err.Error())
				return
			}

		}

	}()
}

func (self *Connector) onConn(conn server.RpcConn) {
}

func (self *Connector) onDisConn(conn server.RpcConn) {
	logger.Info("Connector:onDisConn  %v", conn.GetId())
	self.delPlayer(conn.GetId())
}

//添加玩家到全局表中
func (self *Connector) addPlayer(connId uint64, p *Player) {
	logger.Info("Connector:addPlayer %v, %v", connId, p.GetUid())

	self.l.Lock()
	defer self.l.Unlock()

	//进入服务器全局表
	self.players[connId] = p
	self.playersbyid[p.GetUid()] = p
}

//销毁玩家
func (self *Connector) delPlayer(connId uint64) {
	logger.Info("Connector:delPlayer %v", connId)

	p, exist := self.players[connId]
	if exist {
		p.OnQuit()

		self.l.Lock()
		delete(self.players, connId)
		delete(self.playersbyid, p.GetUid())
		self.l.Unlock()
	}
}

func WriteResult(conn server.RpcConn, value interface{}) bool {
	err := conn.WriteObj(value)
	if err != nil {
		logger.Info("WriteResult Error %s", err.Error())
		return false
	}
	return true
}

func (self *Connector) Login(conn server.RpcConn, login protobuf.Login) error {

	rep := protobuf.LoginResult{}
	uid := login.GetUid()
	if uid == "" {
		uid = common.GenUUID(login.GetAccount())
		login.SetUid(uid)
	}

	self.authserver.Call("AuthServer.Login", &login, &rep)

	logger.Debug("LoginResult %v", rep)

	rep.SetResult(rep.GetResult())

	if rep.GetResult() == protobuf.LoginResult_OK {
		WriteResult(conn, &rep)
		if p, ok := self.playersbyid[login.GetUid()]; ok {
			if err := p.conn.Close(); err == nil {
				logger.Info("kick the online player")
			}
		}

		var base protobuf.PlayerBaseInfo
		logger.Info("query db : %v", rep.GetUid())
		result, err :=db.Query("playerbase", rep.GetUid(), &base)
		if result == false {
			base = protobuf.PlayerBaseInfo{}
			base.SetUid(login.GetUid())

			stat := &protobuf.StatusInfo{}
			stat.SetName(login.GetAccount())
			trans:= protobuf.Transform{}
			vec3:= protobuf.Vector3{}
			vec3.SetX(0)
			vec3.SetY(0)
			vec3.SetZ(0)
			quat:= protobuf.Quaternion{}
			quat.SetX(0)
			quat.SetY(0)
			quat.SetZ(0)
			quat.SetW(1)
			trans.SetPosition(&vec3)
			trans.SetRotation(&quat)
			trans.SetScale(&vec3)
			stat.SetTransform(&trans)
			stat.SetLevel(1)
			stat.SetExperience(0)
			stat.SetHP(100)
			stat.SetMP(100)
			stat.SetRage(100)
			base.SetStat(stat)
			prop := &protobuf.PropertyInfo{}
			prop.SetMaxHP(100)
			prop.SetMaxMP(100)
			base.SetProperty(prop)
			db.Write("playerbase", rep.GetUid(), &base)
			logger.Info("playerbase create %v", rep.GetUid())
		}else {
			if err != nil {
				logger.Info("err query db : %v", err)
				return err
			}
			logger.Info("playerbase find")
		}

		p := &Player{PlayerBaseInfo: &base, conn: conn}

		p.SetUid(uid)

		//进入服务器全局表

		self.addPlayer(conn.GetId(), p)

	}else {
		rep.SetSessionKey("")
		rep.SetUid("")
		WriteResult(conn, &rep)

		go func() {
			time.Sleep(time.Millisecond * 1000)
			defer func() {
				conn.Close()
			}()
		}()
	}

	return nil
}

func (self *Connector) Ping(conn server.RpcConn, login protobuf.Ping) error {

	rep := protobuf.PingResult{}
	rep.SetServerTime(uint32(time.Now().Unix()))

	WriteResult(conn, &rep)
	return nil
}
