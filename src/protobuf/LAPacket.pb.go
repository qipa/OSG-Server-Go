// Code generated by protoc-gen-go.
// source: LAPacket.proto
// DO NOT EDIT!

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	LAPacket.proto

It has these top-level messages:
	LA_CheckAccount
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

type LA_CheckAccount struct {
	Uid              *string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
	Account          *string `protobuf:"bytes,2,req,name=account" json:"account,omitempty"`
	Password         *string `protobuf:"bytes,3,req,name=password" json:"password,omitempty"`
	CreateTime       *uint32 `protobuf:"varint,4,opt,name=create_time" json:"create_time,omitempty"`
	Option           *string `protobuf:"bytes,5,opt,name=option" json:"option,omitempty"`
	Language         *uint32 `protobuf:"varint,6,opt,name=language" json:"language,omitempty"`
	Udid             *string `protobuf:"bytes,7,opt,name=udid" json:"udid,omitempty"`
	SessionKey       *string `protobuf:"bytes,8,opt,name=sessionKey" json:"sessionKey,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LA_CheckAccount) Reset()         { *m = LA_CheckAccount{} }
func (m *LA_CheckAccount) String() string { return proto.CompactTextString(m) }
func (*LA_CheckAccount) ProtoMessage()    {}

func (m *LA_CheckAccount) GetUid() string {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return ""
}

func (m *LA_CheckAccount) GetAccount() string {
	if m != nil && m.Account != nil {
		return *m.Account
	}
	return ""
}

func (m *LA_CheckAccount) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *LA_CheckAccount) GetCreateTime() uint32 {
	if m != nil && m.CreateTime != nil {
		return *m.CreateTime
	}
	return 0
}

func (m *LA_CheckAccount) GetOption() string {
	if m != nil && m.Option != nil {
		return *m.Option
	}
	return ""
}

func (m *LA_CheckAccount) GetLanguage() uint32 {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return 0
}

func (m *LA_CheckAccount) GetUdid() string {
	if m != nil && m.Udid != nil {
		return *m.Udid
	}
	return ""
}

func (m *LA_CheckAccount) GetSessionKey() string {
	if m != nil && m.SessionKey != nil {
		return *m.SessionKey
	}
	return ""
}

func (m *LA_CheckAccount) SetUid(value string) {
	if m != nil {
		m.Uid = &value
	}
}

func (m *LA_CheckAccount) SetAccount(value string) {
	if m != nil {
		m.Account = &value
	}
}

func (m *LA_CheckAccount) SetPassword(value string) {
	if m != nil {
		m.Password = &value
	}
}

func (m *LA_CheckAccount) SetCreateTime(value uint32) {
	if m != nil {
		m.CreateTime = &value
	}
}

func (m *LA_CheckAccount) SetOption(value string) {
	if m != nil {
		m.Option = &value
	}
}

func (m *LA_CheckAccount) SetLanguage(value uint32) {
	if m != nil {
		m.Language = &value
	}
}

func (m *LA_CheckAccount) SetUdid(value string) {
	if m != nil {
		m.Udid = &value
	}
}

func (m *LA_CheckAccount) SetSessionKey(value string) {
	if m != nil {
		m.SessionKey = &value
	}
}

func init() {
}
