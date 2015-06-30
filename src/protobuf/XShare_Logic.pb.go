// Code generated by protoc-gen-go.
// source: XShare_Logic.proto
// DO NOT EDIT!

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	XShare_Logic.proto

It has these top-level messages:
	StatusInfo
	PlayerBaseInfo
*/
package protobuf

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type StatusInfo struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Level            *uint32 `protobuf:"varint,2,opt,name=level" json:"level,omitempty"`
	Experience       *uint32 `protobuf:"varint,3,opt,name=experience" json:"experience,omitempty"`
	Gender           *uint32 `protobuf:"varint,4,opt,name=gender" json:"gender,omitempty"`
	HeadIcon         *uint32 `protobuf:"varint,5,opt,name=headIcon" json:"headIcon,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StatusInfo) Reset()         { *m = StatusInfo{} }
func (m *StatusInfo) String() string { return proto.CompactTextString(m) }
func (*StatusInfo) ProtoMessage()    {}

func (m *StatusInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *StatusInfo) GetLevel() uint32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *StatusInfo) GetExperience() uint32 {
	if m != nil && m.Experience != nil {
		return *m.Experience
	}
	return 0
}

func (m *StatusInfo) GetGender() uint32 {
	if m != nil && m.Gender != nil {
		return *m.Gender
	}
	return 0
}

func (m *StatusInfo) GetHeadIcon() uint32 {
	if m != nil && m.HeadIcon != nil {
		return *m.HeadIcon
	}
	return 0
}

func (m *StatusInfo) SetName(value string) {
	if m != nil {
		m.Name = &value
	}
}

func (m *StatusInfo) SetLevel(value uint32) {
	if m != nil {
		m.Level = &value
	}
}

func (m *StatusInfo) SetExperience(value uint32) {
	if m != nil {
		m.Experience = &value
	}
}

func (m *StatusInfo) SetGender(value uint32) {
	if m != nil {
		m.Gender = &value
	}
}

func (m *StatusInfo) SetHeadIcon(value uint32) {
	if m != nil {
		m.HeadIcon = &value
	}
}

type PlayerBaseInfo struct {
	Uid              *string     `protobuf:"bytes,1,req,name=uid" json:"uid,omitempty"`
	Stat             *StatusInfo `protobuf:"bytes,2,req,name=stat" json:"stat,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *PlayerBaseInfo) Reset()         { *m = PlayerBaseInfo{} }
func (m *PlayerBaseInfo) String() string { return proto.CompactTextString(m) }
func (*PlayerBaseInfo) ProtoMessage()    {}

func (m *PlayerBaseInfo) GetUid() string {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return ""
}

func (m *PlayerBaseInfo) GetStat() *StatusInfo {
	if m != nil {
		return m.Stat
	}
	return nil
}

func (m *PlayerBaseInfo) SetUid(value string) {
	if m != nil {
		m.Uid = &value
	}
}

func (m *PlayerBaseInfo) SetStat(value *StatusInfo) {
	if m != nil {
		m.Stat = value
	}
}

func init() {
}