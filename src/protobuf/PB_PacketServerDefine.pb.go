// Code generated by protoc-gen-go.
// source: PB_PacketServerDefine.proto
// DO NOT EDIT!

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	PB_PacketServerDefine.proto

It has these top-level messages:
*/
package protobuf

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type LA_Protocol int32

const (
	LA_Protocol_eLA_PacketBegin LA_Protocol = 11000
	// ----------------------------
	LA_Protocol_eLA_Connected    LA_Protocol = 11000
	LA_Protocol_eLA_Disconnected LA_Protocol = 11001
	LA_Protocol_eLA_CheckAccount LA_Protocol = 11002
	// ----------------------------
	LA_Protocol_eLA_PacketEnd LA_Protocol = 12000
)

var LA_Protocol_name = map[int32]string{
	11000: "eLA_PacketBegin",
	// Duplicate value: 11000: "eLA_Connected",
	11001: "eLA_Disconnected",
	11002: "eLA_CheckAccount",
	12000: "eLA_PacketEnd",
}
var LA_Protocol_value = map[string]int32{
	"eLA_PacketBegin":  11000,
	"eLA_Connected":    11000,
	"eLA_Disconnected": 11001,
	"eLA_CheckAccount": 11002,
	"eLA_PacketEnd":    12000,
}

func (x LA_Protocol) Enum() *LA_Protocol {
	p := new(LA_Protocol)
	*p = x
	return p
}
func (x LA_Protocol) String() string {
	return proto.EnumName(LA_Protocol_name, int32(x))
}
func (x LA_Protocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *LA_Protocol) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LA_Protocol_value, data, "LA_Protocol")
	if err != nil {
		return err
	}
	*x = LA_Protocol(value)
	return nil
}

type AL_Protocol int32

const (
	AL_Protocol_eAL_PacketBegin AL_Protocol = 12000
	// ----------------------------
	AL_Protocol_eAL_Connected          AL_Protocol = 12000
	AL_Protocol_eAL_Disconnected       AL_Protocol = 12001
	AL_Protocol_eAL_CheckAccountResult AL_Protocol = 12002
	// ----------------------------
	AL_Protocol_eAL_PacketEnd AL_Protocol = 13000
)

var AL_Protocol_name = map[int32]string{
	12000: "eAL_PacketBegin",
	// Duplicate value: 12000: "eAL_Connected",
	12001: "eAL_Disconnected",
	12002: "eAL_CheckAccountResult",
	13000: "eAL_PacketEnd",
}
var AL_Protocol_value = map[string]int32{
	"eAL_PacketBegin":        12000,
	"eAL_Connected":          12000,
	"eAL_Disconnected":       12001,
	"eAL_CheckAccountResult": 12002,
	"eAL_PacketEnd":          13000,
}

func (x AL_Protocol) Enum() *AL_Protocol {
	p := new(AL_Protocol)
	*p = x
	return p
}
func (x AL_Protocol) String() string {
	return proto.EnumName(AL_Protocol_name, int32(x))
}
func (x AL_Protocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *AL_Protocol) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AL_Protocol_value, data, "AL_Protocol")
	if err != nil {
		return err
	}
	*x = AL_Protocol(value)
	return nil
}

type LS_Protocol int32

const (
	LS_Protocol_eLS_PacketBegin LS_Protocol = 13000
	// ----------------------------
	LS_Protocol_eLS_Connected               LS_Protocol = 13000
	LS_Protocol_eLS_Disconnected            LS_Protocol = 13001
	LS_Protocol_eLS_UpdatePlayerCountResult LS_Protocol = 13002
	// ----------------------------
	LS_Protocol_eSLS_PacketEnd LS_Protocol = 14000
)

var LS_Protocol_name = map[int32]string{
	13000: "eLS_PacketBegin",
	// Duplicate value: 13000: "eLS_Connected",
	13001: "eLS_Disconnected",
	13002: "eLS_UpdatePlayerCountResult",
	14000: "eSLS_PacketEnd",
}
var LS_Protocol_value = map[string]int32{
	"eLS_PacketBegin":             13000,
	"eLS_Connected":               13000,
	"eLS_Disconnected":            13001,
	"eLS_UpdatePlayerCountResult": 13002,
	"eSLS_PacketEnd":              14000,
}

func (x LS_Protocol) Enum() *LS_Protocol {
	p := new(LS_Protocol)
	*p = x
	return p
}
func (x LS_Protocol) String() string {
	return proto.EnumName(LS_Protocol_name, int32(x))
}
func (x LS_Protocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *LS_Protocol) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LS_Protocol_value, data, "LS_Protocol")
	if err != nil {
		return err
	}
	*x = LS_Protocol(value)
	return nil
}

type SL_Protocol int32

const (
	SL_Protocol_eSL_PacketBegin SL_Protocol = 14000
	// ----------------------------
	SL_Protocol_eSL_Connected         SL_Protocol = 14000
	SL_Protocol_eSL_Disconnected      SL_Protocol = 14001
	SL_Protocol_eSL_UpdatePlayerCount SL_Protocol = 14002
	// ----------------------------
	SL_Protocol_eSL_PacketEnd SL_Protocol = 20000
)

var SL_Protocol_name = map[int32]string{
	14000: "eSL_PacketBegin",
	// Duplicate value: 14000: "eSL_Connected",
	14001: "eSL_Disconnected",
	14002: "eSL_UpdatePlayerCount",
	20000: "eSL_PacketEnd",
}
var SL_Protocol_value = map[string]int32{
	"eSL_PacketBegin":       14000,
	"eSL_Connected":         14000,
	"eSL_Disconnected":      14001,
	"eSL_UpdatePlayerCount": 14002,
	"eSL_PacketEnd":         20000,
}

func (x SL_Protocol) Enum() *SL_Protocol {
	p := new(SL_Protocol)
	*p = x
	return p
}
func (x SL_Protocol) String() string {
	return proto.EnumName(SL_Protocol_name, int32(x))
}
func (x SL_Protocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *SL_Protocol) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SL_Protocol_value, data, "SL_Protocol")
	if err != nil {
		return err
	}
	*x = SL_Protocol(value)
	return nil
}

func init() {
	proto.RegisterEnum("protobuf.LA_Protocol", LA_Protocol_name, LA_Protocol_value)
	proto.RegisterEnum("protobuf.AL_Protocol", AL_Protocol_name, AL_Protocol_value)
	proto.RegisterEnum("protobuf.LS_Protocol", LS_Protocol_name, LS_Protocol_value)
	proto.RegisterEnum("protobuf.SL_Protocol", SL_Protocol_name, SL_Protocol_value)
}
