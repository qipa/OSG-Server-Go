// Code generated by protoc-gen-go.
// source: CSPacket.proto
// DO NOT EDIT!

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	CSPacket.proto

It has these top-level messages:
	CS_CheckSession
*/
package protobuf

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// discarding unused import protobuf1 "PB_PacketDefine.pb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type CS_CheckSession struct {
	Uid              *string `protobuf:"bytes,1,req,name=uid" json:"uid,omitempty"`
	SessionKey       *string `protobuf:"bytes,2,req,name=sessionKey" json:"sessionKey,omitempty"`
	CreateTime       *uint32 `protobuf:"varint,3,opt,name=create_time" json:"create_time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CS_CheckSession) Reset()         { *m = CS_CheckSession{} }
func (m *CS_CheckSession) String() string { return proto.CompactTextString(m) }
func (*CS_CheckSession) ProtoMessage()    {}

func (m *CS_CheckSession) GetUid() string {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return ""
}

func (m *CS_CheckSession) GetSessionKey() string {
	if m != nil && m.SessionKey != nil {
		return *m.SessionKey
	}
	return ""
}

func (m *CS_CheckSession) GetCreateTime() uint32 {
	if m != nil && m.CreateTime != nil {
		return *m.CreateTime
	}
	return 0
}

func (m *CS_CheckSession) SetUid(value string) {
	if m != nil {
		m.Uid = &value
	}
}

func (m *CS_CheckSession) SetSessionKey(value string) {
	if m != nil {
		m.SessionKey = &value
	}
}

func (m *CS_CheckSession) SetCreateTime(value uint32) {
	if m != nil {
		m.CreateTime = &value
	}
}

func init() {
}
