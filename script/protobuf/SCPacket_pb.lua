-- Generated By protoc-gen-lua Do not Edit
local protobuf = require "protobuf"
local PB_PacketDefine_pb = require("PB_PacketDefine_pb")
module('SCPacket_pb')


local SC_CHECKSESSIONRESULT = protobuf.Descriptor();
local SC_CHECKSESSIONRESULT_RESULT = protobuf.EnumDescriptor();
local SC_CHECKSESSIONRESULT_RESULT_OK_ENUM = protobuf.EnumValueDescriptor();
local SC_CHECKSESSIONRESULT_RESULT_SERVERERROR_ENUM = protobuf.EnumValueDescriptor();
local SC_CHECKSESSIONRESULT_RESULT_USERNOTFOUND_ENUM = protobuf.EnumValueDescriptor();
local SC_CHECKSESSIONRESULT_RESULT_AUTH_FAILED_ENUM = protobuf.EnumValueDescriptor();
local SC_CHECKSESSIONRESULT_RESULT_ISONFIRE_ENUM = protobuf.EnumValueDescriptor();
local SC_CHECKSESSIONRESULT_RESULT_FIELD = protobuf.FieldDescriptor();
local SC_CHECKSESSIONRESULT_SERVER_TIME_FIELD = protobuf.FieldDescriptor();
local SC_CHECKSESSIONRESULT_ERRMSG_FIELD = protobuf.FieldDescriptor();

SC_CHECKSESSIONRESULT_RESULT_OK_ENUM.name = "OK"
SC_CHECKSESSIONRESULT_RESULT_OK_ENUM.index = 0
SC_CHECKSESSIONRESULT_RESULT_OK_ENUM.number = 0
SC_CHECKSESSIONRESULT_RESULT_SERVERERROR_ENUM.name = "SERVERERROR"
SC_CHECKSESSIONRESULT_RESULT_SERVERERROR_ENUM.index = 1
SC_CHECKSESSIONRESULT_RESULT_SERVERERROR_ENUM.number = 1
SC_CHECKSESSIONRESULT_RESULT_USERNOTFOUND_ENUM.name = "USERNOTFOUND"
SC_CHECKSESSIONRESULT_RESULT_USERNOTFOUND_ENUM.index = 2
SC_CHECKSESSIONRESULT_RESULT_USERNOTFOUND_ENUM.number = 2
SC_CHECKSESSIONRESULT_RESULT_AUTH_FAILED_ENUM.name = "AUTH_FAILED"
SC_CHECKSESSIONRESULT_RESULT_AUTH_FAILED_ENUM.index = 3
SC_CHECKSESSIONRESULT_RESULT_AUTH_FAILED_ENUM.number = 3
SC_CHECKSESSIONRESULT_RESULT_ISONFIRE_ENUM.name = "ISONFIRE"
SC_CHECKSESSIONRESULT_RESULT_ISONFIRE_ENUM.index = 4
SC_CHECKSESSIONRESULT_RESULT_ISONFIRE_ENUM.number = 4
SC_CHECKSESSIONRESULT_RESULT.name = "Result"
SC_CHECKSESSIONRESULT_RESULT.full_name = ".protobuf.SC_CheckSessionResult.Result"
SC_CHECKSESSIONRESULT_RESULT.values = {SC_CHECKSESSIONRESULT_RESULT_OK_ENUM,SC_CHECKSESSIONRESULT_RESULT_SERVERERROR_ENUM,SC_CHECKSESSIONRESULT_RESULT_USERNOTFOUND_ENUM,SC_CHECKSESSIONRESULT_RESULT_AUTH_FAILED_ENUM,SC_CHECKSESSIONRESULT_RESULT_ISONFIRE_ENUM}
SC_CHECKSESSIONRESULT_RESULT_FIELD.name = "result"
SC_CHECKSESSIONRESULT_RESULT_FIELD.full_name = ".protobuf.SC_CheckSessionResult.result"
SC_CHECKSESSIONRESULT_RESULT_FIELD.number = 1
SC_CHECKSESSIONRESULT_RESULT_FIELD.index = 0
SC_CHECKSESSIONRESULT_RESULT_FIELD.label = 2
SC_CHECKSESSIONRESULT_RESULT_FIELD.has_default_value = true
SC_CHECKSESSIONRESULT_RESULT_FIELD.default_value = OK
SC_CHECKSESSIONRESULT_RESULT_FIELD.enum_type = SC_CHECKSESSIONRESULT_RESULT
SC_CHECKSESSIONRESULT_RESULT_FIELD.type = 14
SC_CHECKSESSIONRESULT_RESULT_FIELD.cpp_type = 8

SC_CHECKSESSIONRESULT_SERVER_TIME_FIELD.name = "server_time"
SC_CHECKSESSIONRESULT_SERVER_TIME_FIELD.full_name = ".protobuf.SC_CheckSessionResult.server_time"
SC_CHECKSESSIONRESULT_SERVER_TIME_FIELD.number = 2
SC_CHECKSESSIONRESULT_SERVER_TIME_FIELD.index = 1
SC_CHECKSESSIONRESULT_SERVER_TIME_FIELD.label = 1
SC_CHECKSESSIONRESULT_SERVER_TIME_FIELD.has_default_value = false
SC_CHECKSESSIONRESULT_SERVER_TIME_FIELD.default_value = 0
SC_CHECKSESSIONRESULT_SERVER_TIME_FIELD.type = 13
SC_CHECKSESSIONRESULT_SERVER_TIME_FIELD.cpp_type = 3

SC_CHECKSESSIONRESULT_ERRMSG_FIELD.name = "errmsg"
SC_CHECKSESSIONRESULT_ERRMSG_FIELD.full_name = ".protobuf.SC_CheckSessionResult.errmsg"
SC_CHECKSESSIONRESULT_ERRMSG_FIELD.number = 3
SC_CHECKSESSIONRESULT_ERRMSG_FIELD.index = 2
SC_CHECKSESSIONRESULT_ERRMSG_FIELD.label = 1
SC_CHECKSESSIONRESULT_ERRMSG_FIELD.has_default_value = false
SC_CHECKSESSIONRESULT_ERRMSG_FIELD.default_value = ""
SC_CHECKSESSIONRESULT_ERRMSG_FIELD.type = 9
SC_CHECKSESSIONRESULT_ERRMSG_FIELD.cpp_type = 9

SC_CHECKSESSIONRESULT.name = "SC_CheckSessionResult"
SC_CHECKSESSIONRESULT.full_name = ".protobuf.SC_CheckSessionResult"
SC_CHECKSESSIONRESULT.nested_types = {}
SC_CHECKSESSIONRESULT.enum_types = {SC_CHECKSESSIONRESULT_RESULT}
SC_CHECKSESSIONRESULT.fields = {SC_CHECKSESSIONRESULT_RESULT_FIELD, SC_CHECKSESSIONRESULT_SERVER_TIME_FIELD, SC_CHECKSESSIONRESULT_ERRMSG_FIELD}
SC_CHECKSESSIONRESULT.is_extendable = false
SC_CHECKSESSIONRESULT.extensions = {}

SC_CheckSessionResult = protobuf.Message(SC_CHECKSESSIONRESULT)

