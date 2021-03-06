// Code generated by protoc-gen-go.
// source: ALPacket.proto
// DO NOT EDIT!

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	ALPacket.proto
	CLPacket.proto
	CSPacket.proto
	LAPacket.proto
	LCPacket.proto
	LSPacket.proto
	PB_PacketCommon.proto
	PB_PacketDefine.proto
	PB_PacketServerDefine.proto
	SCPacket.proto
	SLPacket.proto
	XShare_Logic.proto
	XShare_Server.proto

It has these top-level messages:
	AL_CheckAccountResult
	CL_CheckAccount
	CS_CheckSession
	CS_Ping
	LA_CheckAccount
	LC_CheckAccountResult
	LS_UpdatePlayerCountResult
	Packet
	RpcErrorResponse
	SC_CheckSessionResult
	SC_PingResult
	SL_UpdatePlayerCount
	StatusInfo
	PlayerBaseInfo
	AccountInfo
*/
package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type AL_CheckAccountResult_Result int32

const (
	AL_CheckAccountResult_SUCCESS      AL_CheckAccountResult_Result = 0
	AL_CheckAccountResult_OK           AL_CheckAccountResult_Result = 1
	AL_CheckAccountResult_SERVERERROR  AL_CheckAccountResult_Result = 2
	AL_CheckAccountResult_USERNOTFOUND AL_CheckAccountResult_Result = 3
	AL_CheckAccountResult_AUTH_FAILED  AL_CheckAccountResult_Result = 4
	AL_CheckAccountResult_ISONFIRE     AL_CheckAccountResult_Result = 5
)

var AL_CheckAccountResult_Result_name = map[int32]string{
	0: "SUCCESS",
	1: "OK",
	2: "SERVERERROR",
	3: "USERNOTFOUND",
	4: "AUTH_FAILED",
	5: "ISONFIRE",
}
var AL_CheckAccountResult_Result_value = map[string]int32{
	"SUCCESS":      0,
	"OK":           1,
	"SERVERERROR":  2,
	"USERNOTFOUND": 3,
	"AUTH_FAILED":  4,
	"ISONFIRE":     5,
}

func (x AL_CheckAccountResult_Result) String() string {
	return proto.EnumName(AL_CheckAccountResult_Result_name, int32(x))
}
func (AL_CheckAccountResult_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

type AL_CheckAccountResult struct {
	Result     AL_CheckAccountResult_Result `protobuf:"varint,1,opt,name=result,enum=protobuf.AL_CheckAccountResult_Result" json:"result,omitempty"`
	SessionKey string                       `protobuf:"bytes,2,opt,name=sessionKey" json:"sessionKey,omitempty"`
	Uid        string                       `protobuf:"bytes,3,opt,name=uid" json:"uid,omitempty"`
	Errmsg     string                       `protobuf:"bytes,4,opt,name=errmsg" json:"errmsg,omitempty"`
}

func (m *AL_CheckAccountResult) Reset()                    { *m = AL_CheckAccountResult{} }
func (m *AL_CheckAccountResult) String() string            { return proto.CompactTextString(m) }
func (*AL_CheckAccountResult) ProtoMessage()               {}
func (*AL_CheckAccountResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*AL_CheckAccountResult)(nil), "protobuf.AL_CheckAccountResult")
	proto.RegisterEnum("protobuf.AL_CheckAccountResult_Result", AL_CheckAccountResult_Result_name, AL_CheckAccountResult_Result_value)
}

var fileDescriptor0 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x73, 0xf4, 0x09, 0x48,
	0x4c, 0xce, 0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x49, 0xa5,
	0x69, 0x4a, 0x4d, 0x4c, 0x5c, 0xa2, 0x8e, 0x3e, 0xf1, 0xce, 0x19, 0xa9, 0xc9, 0xd9, 0x8e, 0xc9,
	0xc9, 0xf9, 0xa5, 0x79, 0x25, 0x41, 0xa9, 0xc5, 0xa5, 0x39, 0x25, 0x42, 0x76, 0x5c, 0x6c, 0x45,
	0x60, 0x96, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x9f, 0x91, 0x9a, 0x1e, 0x4c, 0x93, 0x1e, 0x56, 0x0d,
	0x7a, 0x10, 0x2a, 0x08, 0xaa, 0x4b, 0x48, 0x8e, 0x8b, 0xab, 0x38, 0xb5, 0xb8, 0x38, 0x33, 0x3f,
	0xcf, 0x3b, 0xb5, 0x52, 0x82, 0x09, 0x68, 0x06, 0x67, 0x10, 0x92, 0x88, 0x90, 0x00, 0x17, 0x73,
	0x69, 0x66, 0x8a, 0x04, 0x33, 0x58, 0x02, 0xc4, 0x14, 0x12, 0xe3, 0x62, 0x4b, 0x2d, 0x2a, 0xca,
	0x2d, 0x4e, 0x97, 0x60, 0x01, 0x0b, 0x42, 0x79, 0x4a, 0xf1, 0x5c, 0x6c, 0x50, 0x37, 0x71, 0x73,
	0xb1, 0x07, 0x87, 0x3a, 0x3b, 0xbb, 0x06, 0x07, 0x0b, 0x30, 0x08, 0xb1, 0x71, 0x31, 0xf9, 0x7b,
	0x0b, 0x30, 0x0a, 0xf1, 0x73, 0x71, 0x07, 0xbb, 0x06, 0x85, 0xb9, 0x06, 0xb9, 0x06, 0x05, 0xf9,
	0x07, 0x09, 0x30, 0x01, 0x4d, 0xe6, 0x09, 0x05, 0x8a, 0xf8, 0xf9, 0x87, 0xb8, 0xf9, 0x87, 0xfa,
	0xb9, 0x08, 0x30, 0x83, 0x94, 0x38, 0x86, 0x86, 0x78, 0xc4, 0xbb, 0x39, 0x7a, 0xfa, 0xb8, 0xba,
	0x08, 0xb0, 0x08, 0xf1, 0x70, 0x71, 0x78, 0x06, 0xfb, 0xfb, 0xb9, 0x79, 0x06, 0xb9, 0x0a, 0xb0,
	0x3a, 0x31, 0x79, 0x30, 0x26, 0xb1, 0x81, 0x7d, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x6f,
	0xfe, 0x3f, 0x4d, 0x2b, 0x01, 0x00, 0x00,
}
