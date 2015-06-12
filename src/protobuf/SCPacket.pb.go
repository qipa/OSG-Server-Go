// Code generated by protoc-gen-go.
// source: SCPacket.proto
// DO NOT EDIT!

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	SCPacket.proto

It has these top-level messages:
	SC_CheckSessionResult
	CS_PingResult
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

type SC_CheckSessionResult_Result int32

const (
	SC_CheckSessionResult_OK           SC_CheckSessionResult_Result = 0
	SC_CheckSessionResult_SERVERERROR  SC_CheckSessionResult_Result = 1
	SC_CheckSessionResult_USERNOTFOUND SC_CheckSessionResult_Result = 2
	SC_CheckSessionResult_AUTH_FAILED  SC_CheckSessionResult_Result = 3
	SC_CheckSessionResult_ISONFIRE     SC_CheckSessionResult_Result = 4
)

var SC_CheckSessionResult_Result_name = map[int32]string{
	0: "OK",
	1: "SERVERERROR",
	2: "USERNOTFOUND",
	3: "AUTH_FAILED",
	4: "ISONFIRE",
}
var SC_CheckSessionResult_Result_value = map[string]int32{
	"OK":           0,
	"SERVERERROR":  1,
	"USERNOTFOUND": 2,
	"AUTH_FAILED":  3,
	"ISONFIRE":     4,
}

func (x SC_CheckSessionResult_Result) Enum() *SC_CheckSessionResult_Result {
	p := new(SC_CheckSessionResult_Result)
	*p = x
	return p
}
func (x SC_CheckSessionResult_Result) String() string {
	return proto.EnumName(SC_CheckSessionResult_Result_name, int32(x))
}
func (x SC_CheckSessionResult_Result) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *SC_CheckSessionResult_Result) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SC_CheckSessionResult_Result_value, data, "SC_CheckSessionResult_Result")
	if err != nil {
		return err
	}
	*x = SC_CheckSessionResult_Result(value)
	return nil
}

type SC_CheckSessionResult struct {
	Result           *SC_CheckSessionResult_Result `protobuf:"varint,1,req,name=result,enum=protobuf.SC_CheckSessionResult_Result,def=0" json:"result,omitempty"`
	ServerTime       *uint32                       `protobuf:"varint,2,opt,name=server_time" json:"server_time,omitempty"`
	Errmsg           *string                       `protobuf:"bytes,3,opt,name=errmsg" json:"errmsg,omitempty"`
	XXX_unrecognized []byte                        `json:"-"`
}

func (m *SC_CheckSessionResult) Reset()         { *m = SC_CheckSessionResult{} }
func (m *SC_CheckSessionResult) String() string { return proto.CompactTextString(m) }
func (*SC_CheckSessionResult) ProtoMessage()    {}

const Default_SC_CheckSessionResult_Result SC_CheckSessionResult_Result = SC_CheckSessionResult_OK

func (m *SC_CheckSessionResult) GetResult() SC_CheckSessionResult_Result {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return Default_SC_CheckSessionResult_Result
}

func (m *SC_CheckSessionResult) GetServerTime() uint32 {
	if m != nil && m.ServerTime != nil {
		return *m.ServerTime
	}
	return 0
}

func (m *SC_CheckSessionResult) GetErrmsg() string {
	if m != nil && m.Errmsg != nil {
		return *m.Errmsg
	}
	return ""
}

func (m *SC_CheckSessionResult) SetResult(value SC_CheckSessionResult_Result) {
	if m != nil {
		m.Result = &value
	}
}

func (m *SC_CheckSessionResult) SetServerTime(value uint32) {
	if m != nil {
		m.ServerTime = &value
	}
}

func (m *SC_CheckSessionResult) SetErrmsg(value string) {
	if m != nil {
		m.Errmsg = &value
	}
}

type CS_PingResult struct {
	ServerTime       *uint32 `protobuf:"varint,1,req,name=server_time" json:"server_time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CS_PingResult) Reset()         { *m = CS_PingResult{} }
func (m *CS_PingResult) String() string { return proto.CompactTextString(m) }
func (*CS_PingResult) ProtoMessage()    {}

func (m *CS_PingResult) GetServerTime() uint32 {
	if m != nil && m.ServerTime != nil {
		return *m.ServerTime
	}
	return 0
}

func (m *CS_PingResult) SetServerTime(value uint32) {
	if m != nil {
		m.ServerTime = &value
	}
}

func init() {
	proto.RegisterEnum("protobuf.SC_CheckSessionResult_Result", SC_CheckSessionResult_Result_name, SC_CheckSessionResult_Result_value)
}
