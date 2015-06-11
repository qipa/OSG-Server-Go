// Code generated by protoc-gen-go.
// source: SLPacket.proto
// DO NOT EDIT!

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	SLPacket.proto

It has these top-level messages:
	SL_UpdatePlayerCount
*/
package protobuf

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// discarding unused import protobuf2 "PB_PacketServerDefine.pb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type SL_UpdatePlayerCount struct {
	ServerId         *uint32 `protobuf:"varint,1,req" json:"ServerId,omitempty"`
	PlayerCount      *uint32 `protobuf:"varint,2,req" json:"PlayerCount,omitempty"`
	TcpServerIp      *string `protobuf:"bytes,3,req" json:"TcpServerIp,omitempty"`
	HttpServerIp     *string `protobuf:"bytes,4,req" json:"HttpServerIp,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SL_UpdatePlayerCount) Reset()         { *m = SL_UpdatePlayerCount{} }
func (m *SL_UpdatePlayerCount) String() string { return proto.CompactTextString(m) }
func (*SL_UpdatePlayerCount) ProtoMessage()    {}

func (m *SL_UpdatePlayerCount) GetServerId() uint32 {
	if m != nil && m.ServerId != nil {
		return *m.ServerId
	}
	return 0
}

func (m *SL_UpdatePlayerCount) GetPlayerCount() uint32 {
	if m != nil && m.PlayerCount != nil {
		return *m.PlayerCount
	}
	return 0
}

func (m *SL_UpdatePlayerCount) GetTcpServerIp() string {
	if m != nil && m.TcpServerIp != nil {
		return *m.TcpServerIp
	}
	return ""
}

func (m *SL_UpdatePlayerCount) GetHttpServerIp() string {
	if m != nil && m.HttpServerIp != nil {
		return *m.HttpServerIp
	}
	return ""
}

func (m *SL_UpdatePlayerCount) SetServerId(value uint32) {
	if m != nil {
		m.ServerId = &value
	}
}

func (m *SL_UpdatePlayerCount) SetPlayerCount(value uint32) {
	if m != nil {
		m.PlayerCount = &value
	}
}

func (m *SL_UpdatePlayerCount) SetTcpServerIp(value string) {
	if m != nil {
		m.TcpServerIp = &value
	}
}

func (m *SL_UpdatePlayerCount) SetHttpServerIp(value string) {
	if m != nil {
		m.HttpServerIp = &value
	}
}

func init() {
}
