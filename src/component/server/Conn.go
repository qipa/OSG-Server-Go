package server

import (
	"bytes"
	"common/logger"
	"protobuf"
	"common/timer"
	"encoding/binary"
	"fmt"
	"github.com/gorilla/websocket"
	"net"
	"reflect"
	"time"
	"errors"
	"io"
	"sync"
	"strings"
	"github.com/golang/protobuf/proto"
)

const (
	ConnReadTimeOut  = 5e9
	ConnWriteTimeOut = 5e9
)

type Conn struct {
	web_socket  websocket.Conn
	tcp_socket  net.Conn
	client_type uint32						// 1 tcpsocket 2 websocket
}

func (c Conn)SetWriteDeadline(t time.Time) error {
	switch c.client_type {
		case 1:
			return c.tcp_socket.SetWriteDeadline(t)
		case 2:
			return c.web_socket.SetWriteDeadline(t)
		default:
			return errors.New("ProtoBufConn: err client_type unknown")
	}
}

func (c Conn)RemoteAddr() net.Addr {
	switch c.client_type {
		case 1:
		return c.tcp_socket.RemoteAddr()
		case 2:
		return c.web_socket.RemoteAddr()
		default:
		return nil
	}
}

func (c Conn)SetReadDeadline(t time.Time) error {
	switch c.client_type {
		case 1:
		return c.tcp_socket.SetReadDeadline(t)
		case 2:
		return c.web_socket.SetReadDeadline(t)
		default:
		return errors.New("ProtoBufConn: err client_type unknown")
	}
}

func (c Conn)Close() error {

	switch c.client_type {
		case 1:
		return c.tcp_socket.Close()
		case 2:
		return c.web_socket.Close()
		default:
		return errors.New("ProtoBufConn: err client_type unknown")
	}
}

func (c Conn) WriteMessage(data []byte) error {
	switch c.client_type {
		case 1:
		_, err := c.tcp_socket.Write(data)
		return err
		case 2:
		return c.web_socket.WriteMessage(websocket.BinaryMessage, data)
		default:
		return errors.New("ProtoBufConn: err client_type unknown")
	}
}

func (c Conn) ReadMessage() ([]byte, error) {

	var size uint32

	switch c.client_type {
		case 1:
			err := binary.Read(c.tcp_socket, binary.LittleEndian, &size)

			if err != nil {
				logger.Debug("ReadMessage err 1: %v", err.Error())
				return nil, err
			}

			buf := make([]byte, size)

			c.SetReadDeadline(time.Now().Add(ConnReadTimeOut))

			_, err = io.ReadFull(c.tcp_socket, buf)
			if err != nil {
				logger.Debug("ReadMessage err 2: %v", err.Error())
				return nil, err
			}

			return buf, err
		case 2:
			_, buf, err := c.web_socket.ReadMessage()

			dstBuffer := bytes.NewBuffer(buf)

			c.SetReadDeadline(time.Now().Add(ConnReadTimeOut))

			err = binary.Read(dstBuffer, binary.LittleEndian, &size)

			return dstBuffer.Bytes(), err
		default:
			return nil, errors.New("ProtoBufConn: err client_type unknown")
	}
}

type ProtoBufConn struct {
	id          uint64
	msg_id      uint64
	c			Conn
	send        chan *protobuf.Packet
	t           *timer.Timer
	exit        chan bool
	last_time   int64
	time_out    uint32
	lockForClose sync.Mutex
	is_closed   bool
	sync.Mutex
	connMgr     *Server
	resultServer	string
	protocol	map[string]uint32
}

func NewWebSocketConn(server *Server, c websocket.Conn, size int32, k uint32, t uint32) (conn RpcConn) {
	pbc := &ProtoBufConn{
		send:        make(chan *protobuf.Packet, size),
		exit:        make(chan bool, 1),
		last_time:   time.Now().Unix(),
		time_out:    k,
		connMgr:     server,
		resultServer:	"",
		protocol: 	  make(map[string]uint32),
	}

	pbc.c = Conn{
		web_socket:		c,
		client_type:	t,
	}

	if k > 0 {
		pbc.t = timer.NewTimer(time.Duration(k) * time.Second)
		pbc.t.Start(
			func() {
				pbc.OnCheck()
			},
		)
	}

	go pbc.mux()
	return pbc
}

func NewTCPSocketConn(server *Server, c net.Conn, size int32, k uint32, t uint32) (conn RpcConn) {
	pbc := &ProtoBufConn{
		send:        make(chan *protobuf.Packet, size),
		exit:        make(chan bool, 1),
		last_time:   time.Now().Unix(),
		time_out:    k,
		connMgr:     server,
		resultServer:	"",
	}

	pbc.c = Conn{
		tcp_socket:  c,
		client_type: t,
	}

	if k > 0 {
		pbc.t = timer.NewTimer(time.Duration(k) * time.Second)
		pbc.t.Start(
			func() {
				pbc.OnCheck()
			},
		)
	}

	go pbc.mux()
	return pbc
}

func (conn *ProtoBufConn) ApplyProtocol(protocal map[string]int32) {
	logger.Debug("ApplyProtocol")
	for key, value := range protocal {
		protocalMethod := strings.Split(key, "_")
		if len(protocalMethod) != 2 {
			logger.Error("rpc: ApplyProtocol ill-formed: %v , no '_' for split key" + key)
		}
		conn.protocol[protocalMethod[1]] = uint32(value)
	}
}

func (conn *ProtoBufConn) SetResultServer(name string) {
	conn.resultServer = name;
}

func (conn *ProtoBufConn) IsWebConn() bool {
	return conn.c.client_type == 2
}

func (conn *ProtoBufConn) OnCheck() {
	time_diff := uint32(time.Now().Unix() - conn.last_time)
	if time_diff > conn.time_out<<1 {
		logger.Info("Conn %d TimeOut: %d", conn.GetId(), time_diff)
		conn.Close()
	}
}

func (conn *ProtoBufConn) mux() {
	for {
		select {
		case r := <-conn.send:

			//logger.Debug("writeRequest %v", r)

			buf, err := proto.Marshal(r)
			if err != nil {
				logger.Error("ProtoBufConn Marshal Error %s", err.Error())
				continue
			}

			//logger.Debug("        mux: %v", buf)

			//dst, err := snappy.Encode(nil, buf)
			dst := buf

			//logger.Debug("  dst   mux: %v", dst)

			if err != nil {
				logger.Error("ProtoBufConn snappy.Encode Error %s", err.Error())
				continue
			}

			conn.c.SetWriteDeadline(time.Now().Add(ConnWriteTimeOut))
			dstBuffer := new(bytes.Buffer)
			err = binary.Write(dstBuffer, binary.LittleEndian, int32(len(dst)))
			if err != nil {
				logger.Error("ProtoBufConn Write Error %s", err.Error())
				continue
			}

			conn.c.SetWriteDeadline(time.Now().Add(ConnWriteTimeOut))
			_, err = dstBuffer.Write(dst)
			if err != nil {
				logger.Error("ProtoBufConn Write Error %s", err.Error())
				continue
			}

			conn.c.SetWriteDeadline(time.Now().Add(ConnWriteTimeOut))
			conn.c.WriteMessage(dstBuffer.Bytes())
			if err != nil {
				logger.Error("ProtoBufConn Write Error %s", err.Error())
				continue
			}

		case <-conn.exit:

			return
		}
	}
}

func (conn *ProtoBufConn) GetRemoteIp() string {
	return conn.c.RemoteAddr().String()
}

func (conn *ProtoBufConn) ReadRequest(req *protobuf.Packet) error {

	conn.c.SetReadDeadline(time.Now().Add(ConnReadTimeOut))

	dst, err := conn.c.ReadMessage()
	if err != nil {
		logger.Debug("ReadRequest Read binary Err: %v", err)
		return err
	}

	//dst, err := snappy.Decode(nil, dstBuffer.Bytes())

	if err != nil {
		logger.Debug("ReadRequest Decode Err: %v", err)
		return err
	}

	conn.last_time = time.Now().Unix()

	//logger.Info("ReadRequest dst: %v", dst)

	err = proto.Unmarshal(dst, req)
	conn.msg_id = req.Id
	return err
}

func (conn *ProtoBufConn) writeRequest(r *protobuf.Packet) error {
	r.Id = conn.msg_id
	conn.send <- r
	return nil
}

func (conn *ProtoBufConn) Call(cmd uint32, args interface{}) (err error) {

	var msg proto.Message
	var buf []byte

	switch m := args.(type) {
	case proto.Message:
		msg = m

		buf, err = proto.Marshal(msg)
		if err != nil {
			return err
		}
	case []byte:
		buf = m
	default:
		return fmt.Errorf("Call args type error %v", args)
	}

	req := &protobuf.Packet{}
	req.Cmd = cmd
	req.SerializedData = buf

	return conn.writeRequest(req)
}

func (conn *ProtoBufConn) GetRequestBody(req *protobuf.Packet, body interface{}) error {
	if value, ok := body.(proto.Message); ok {
		return proto.Unmarshal(req.SerializedData, value)
	}

	return fmt.Errorf("GetRequestBody value type error %v", body)
}

func (conn *ProtoBufConn) WriteObj(value interface{}) error {

	var msg proto.Message

	switch m := value.(type) {
	case proto.Message:
		msg = m
	default:
		return fmt.Errorf("WriteObj value type error %v", value)
	}

	buf, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	req := &protobuf.Packet{}

	t := reflect.Indirect(reflect.ValueOf(msg)).Type()
	if conn.resultServer == "" {
		req.Cmd = conn.protocol[t.PkgPath() + "." + t.Name()]
	} else {
		req.Cmd = conn.protocol[conn.resultServer + "." + t.Name()]
	}
	req.SerializedData = buf
	return conn.writeRequest(req)
}

func (conn *ProtoBufConn) SetId(id uint64) {
	conn.id = id
}

func (conn *ProtoBufConn) GetId() uint64 {
	return conn.id
}

func (conn *ProtoBufConn) Close() (errret error) {
	conn.lockForClose.Lock()

	logger.Info("ProtoBufConn %d Close : %v", conn.id, errret)

	if conn.is_closed {
		conn.lockForClose.Unlock()
		return nil
	}

	if err := conn.c.Close(); err != nil {

		//再尝试一次
		time.Sleep(10 * time.Millisecond)
		if err := conn.c.Close(); err != nil {
			conn.lockForClose.Unlock()
			return err
		}
	}

	conn.is_closed = true

	if conn.t != nil {
		conn.t.Stop()
	}

	conn.exit <- true

	conn.lockForClose.Unlock()

	return nil
}
