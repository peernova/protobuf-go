// Code generated by protoc-gen-go. DO NOT EDIT.
// source: extensions/proto3/ext3.proto

package proto3

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	prototype "github.com/golang/protobuf/v2/reflect/prototype"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Enum int32

const (
	Enum_ZERO Enum = 0
)

func (e Enum) Type() protoreflect.EnumType {
	return xxx_Ext3_ProtoFile_EnumTypes[0]
}
func (e Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

var Enum_name = map[int32]string{
	0: "ZERO",
}

var Enum_value = map[string]int32{
	"ZERO": 0,
}

func (x Enum) String() string {
	return proto.EnumName(Enum_name, int32(x))
}

func (Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3db31bb248c8865e, []int{0}
}

type Message struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type xxx_Message struct{ m *Message }

func (m *Message) ProtoReflect() protoreflect.Message {
	return xxx_Message{m}
}
func (m xxx_Message) Type() protoreflect.MessageType {
	return xxx_Ext3_ProtoFile_MessageTypes[0].Type
}
func (m xxx_Message) KnownFields() protoreflect.KnownFields {
	return xxx_Ext3_ProtoFile_MessageTypes[0].KnownFieldsOf(m.m)
}
func (m xxx_Message) UnknownFields() protoreflect.UnknownFields {
	return xxx_Ext3_ProtoFile_MessageTypes[0].UnknownFieldsOf(m.m)
}
func (m xxx_Message) Interface() protoreflect.ProtoMessage {
	return m.m
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_3db31bb248c8865e, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

var E_ExtensionBool = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         1001,
	Name:          "goproto.protoc.extension.proto3.extension_bool",
	Tag:           "varint,1001,opt,name=extension_bool",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_ExtensionEnum = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*Enum)(nil),
	Field:         1002,
	Name:          "goproto.protoc.extension.proto3.extension_enum",
	Tag:           "varint,1002,opt,name=extension_enum,enum=goproto.protoc.extension.proto3.Enum",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_ExtensionInt32 = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*int32)(nil),
	Field:         1003,
	Name:          "goproto.protoc.extension.proto3.extension_int32",
	Tag:           "varint,1003,opt,name=extension_int32",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_ExtensionSint32 = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*int32)(nil),
	Field:         1004,
	Name:          "goproto.protoc.extension.proto3.extension_sint32",
	Tag:           "zigzag32,1004,opt,name=extension_sint32",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_ExtensionUint32 = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*uint32)(nil),
	Field:         1005,
	Name:          "goproto.protoc.extension.proto3.extension_uint32",
	Tag:           "varint,1005,opt,name=extension_uint32",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_ExtensionInt64 = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*int64)(nil),
	Field:         1006,
	Name:          "goproto.protoc.extension.proto3.extension_int64",
	Tag:           "varint,1006,opt,name=extension_int64",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_ExtensionSint64 = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*int64)(nil),
	Field:         1007,
	Name:          "goproto.protoc.extension.proto3.extension_sint64",
	Tag:           "zigzag64,1007,opt,name=extension_sint64",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_ExtensionUint64 = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*uint64)(nil),
	Field:         1008,
	Name:          "goproto.protoc.extension.proto3.extension_uint64",
	Tag:           "varint,1008,opt,name=extension_uint64",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_ExtensionSfixed32 = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*int32)(nil),
	Field:         1009,
	Name:          "goproto.protoc.extension.proto3.extension_sfixed32",
	Tag:           "fixed32,1009,opt,name=extension_sfixed32",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_ExtensionFixed32 = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*uint32)(nil),
	Field:         1010,
	Name:          "goproto.protoc.extension.proto3.extension_fixed32",
	Tag:           "fixed32,1010,opt,name=extension_fixed32",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_ExtensionFloat = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*float32)(nil),
	Field:         1011,
	Name:          "goproto.protoc.extension.proto3.extension_float",
	Tag:           "fixed32,1011,opt,name=extension_float",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_ExtensionSfixed64 = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*int64)(nil),
	Field:         1012,
	Name:          "goproto.protoc.extension.proto3.extension_sfixed64",
	Tag:           "fixed64,1012,opt,name=extension_sfixed64",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_ExtensionFixed64 = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*uint64)(nil),
	Field:         1013,
	Name:          "goproto.protoc.extension.proto3.extension_fixed64",
	Tag:           "fixed64,1013,opt,name=extension_fixed64",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_ExtensionDouble = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*float64)(nil),
	Field:         1014,
	Name:          "goproto.protoc.extension.proto3.extension_double",
	Tag:           "fixed64,1014,opt,name=extension_double",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_ExtensionString = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         1015,
	Name:          "goproto.protoc.extension.proto3.extension_string",
	Tag:           "bytes,1015,opt,name=extension_string",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_ExtensionBytes = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: ([]byte)(nil),
	Field:         1016,
	Name:          "goproto.protoc.extension.proto3.extension_bytes",
	Tag:           "bytes,1016,opt,name=extension_bytes",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_Extension_Message = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*Message)(nil),
	Field:         1017,
	Name:          "goproto.protoc.extension.proto3.extension_Message",
	Tag:           "bytes,1017,opt,name=extension_Message",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_RepeatedExtensionBool = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: ([]bool)(nil),
	Field:         2001,
	Name:          "goproto.protoc.extension.proto3.repeated_extension_bool",
	Tag:           "varint,2001,rep,name=repeated_extension_bool",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_RepeatedExtensionEnum = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: ([]Enum)(nil),
	Field:         2002,
	Name:          "goproto.protoc.extension.proto3.repeated_extension_enum",
	Tag:           "varint,2002,rep,name=repeated_extension_enum,enum=goproto.protoc.extension.proto3.Enum",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_RepeatedExtensionInt32 = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: ([]int32)(nil),
	Field:         2003,
	Name:          "goproto.protoc.extension.proto3.repeated_extension_int32",
	Tag:           "varint,2003,rep,name=repeated_extension_int32",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_RepeatedExtensionSint32 = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: ([]int32)(nil),
	Field:         2004,
	Name:          "goproto.protoc.extension.proto3.repeated_extension_sint32",
	Tag:           "zigzag32,2004,rep,name=repeated_extension_sint32",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_RepeatedExtensionUint32 = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: ([]uint32)(nil),
	Field:         2005,
	Name:          "goproto.protoc.extension.proto3.repeated_extension_uint32",
	Tag:           "varint,2005,rep,name=repeated_extension_uint32",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_RepeatedExtensionInt64 = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: ([]int64)(nil),
	Field:         2006,
	Name:          "goproto.protoc.extension.proto3.repeated_extension_int64",
	Tag:           "varint,2006,rep,name=repeated_extension_int64",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_RepeatedExtensionSint64 = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: ([]int64)(nil),
	Field:         2007,
	Name:          "goproto.protoc.extension.proto3.repeated_extension_sint64",
	Tag:           "zigzag64,2007,rep,name=repeated_extension_sint64",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_RepeatedExtensionUint64 = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: ([]uint64)(nil),
	Field:         2008,
	Name:          "goproto.protoc.extension.proto3.repeated_extension_uint64",
	Tag:           "varint,2008,rep,name=repeated_extension_uint64",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_RepeatedExtensionSfixed32 = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: ([]int32)(nil),
	Field:         2009,
	Name:          "goproto.protoc.extension.proto3.repeated_extension_sfixed32",
	Tag:           "fixed32,2009,rep,name=repeated_extension_sfixed32",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_RepeatedExtensionFixed32 = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: ([]uint32)(nil),
	Field:         2010,
	Name:          "goproto.protoc.extension.proto3.repeated_extension_fixed32",
	Tag:           "fixed32,2010,rep,name=repeated_extension_fixed32",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_RepeatedExtensionFloat = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: ([]float32)(nil),
	Field:         2011,
	Name:          "goproto.protoc.extension.proto3.repeated_extension_float",
	Tag:           "fixed32,2011,rep,name=repeated_extension_float",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_RepeatedExtensionSfixed64 = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: ([]int64)(nil),
	Field:         2012,
	Name:          "goproto.protoc.extension.proto3.repeated_extension_sfixed64",
	Tag:           "fixed64,2012,rep,name=repeated_extension_sfixed64",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_RepeatedExtensionFixed64 = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: ([]uint64)(nil),
	Field:         2013,
	Name:          "goproto.protoc.extension.proto3.repeated_extension_fixed64",
	Tag:           "fixed64,2013,rep,name=repeated_extension_fixed64",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_RepeatedExtensionDouble = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: ([]float64)(nil),
	Field:         2014,
	Name:          "goproto.protoc.extension.proto3.repeated_extension_double",
	Tag:           "fixed64,2014,rep,name=repeated_extension_double",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_RepeatedExtensionString = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: ([]string)(nil),
	Field:         2015,
	Name:          "goproto.protoc.extension.proto3.repeated_extension_string",
	Tag:           "bytes,2015,rep,name=repeated_extension_string",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_RepeatedExtensionBytes = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: ([][]byte)(nil),
	Field:         2016,
	Name:          "goproto.protoc.extension.proto3.repeated_extension_bytes",
	Tag:           "bytes,2016,rep,name=repeated_extension_bytes",
	Filename:      "extensions/proto3/ext3.proto",
}

var E_RepeatedExtension_Message = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: ([]*Message)(nil),
	Field:         2017,
	Name:          "goproto.protoc.extension.proto3.repeated_extension_Message",
	Tag:           "bytes,2017,rep,name=repeated_extension_Message",
	Filename:      "extensions/proto3/ext3.proto",
}

func init() {
	proto.RegisterFile("extensions/proto3/ext3.proto", fileDescriptor_3db31bb248c8865e)
	proto.RegisterEnum("goproto.protoc.extension.proto3.Enum", Enum_name, Enum_value)
	proto.RegisterType((*Message)(nil), "goproto.protoc.extension.proto3.Message")
	proto.RegisterExtension(E_ExtensionBool)
	proto.RegisterExtension(E_ExtensionEnum)
	proto.RegisterExtension(E_ExtensionInt32)
	proto.RegisterExtension(E_ExtensionSint32)
	proto.RegisterExtension(E_ExtensionUint32)
	proto.RegisterExtension(E_ExtensionInt64)
	proto.RegisterExtension(E_ExtensionSint64)
	proto.RegisterExtension(E_ExtensionUint64)
	proto.RegisterExtension(E_ExtensionSfixed32)
	proto.RegisterExtension(E_ExtensionFixed32)
	proto.RegisterExtension(E_ExtensionFloat)
	proto.RegisterExtension(E_ExtensionSfixed64)
	proto.RegisterExtension(E_ExtensionFixed64)
	proto.RegisterExtension(E_ExtensionDouble)
	proto.RegisterExtension(E_ExtensionString)
	proto.RegisterExtension(E_ExtensionBytes)
	proto.RegisterExtension(E_Extension_Message)
	proto.RegisterExtension(E_RepeatedExtensionBool)
	proto.RegisterExtension(E_RepeatedExtensionEnum)
	proto.RegisterExtension(E_RepeatedExtensionInt32)
	proto.RegisterExtension(E_RepeatedExtensionSint32)
	proto.RegisterExtension(E_RepeatedExtensionUint32)
	proto.RegisterExtension(E_RepeatedExtensionInt64)
	proto.RegisterExtension(E_RepeatedExtensionSint64)
	proto.RegisterExtension(E_RepeatedExtensionUint64)
	proto.RegisterExtension(E_RepeatedExtensionSfixed32)
	proto.RegisterExtension(E_RepeatedExtensionFixed32)
	proto.RegisterExtension(E_RepeatedExtensionFloat)
	proto.RegisterExtension(E_RepeatedExtensionSfixed64)
	proto.RegisterExtension(E_RepeatedExtensionFixed64)
	proto.RegisterExtension(E_RepeatedExtensionDouble)
	proto.RegisterExtension(E_RepeatedExtensionString)
	proto.RegisterExtension(E_RepeatedExtensionBytes)
	proto.RegisterExtension(E_RepeatedExtension_Message)
}

var fileDescriptor_3db31bb248c8865e = []byte{
	// 758 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x97, 0xdb, 0x4b, 0xdc, 0x5e,
	0x10, 0xc7, 0x7f, 0xab, 0xab, 0xab, 0xe7, 0xa7, 0xee, 0xba, 0xd0, 0x7a, 0x69, 0x41, 0x29, 0x14,
	0xa4, 0x60, 0x02, 0xbb, 0x21, 0x0f, 0x79, 0x94, 0x6a, 0xb1, 0xad, 0x08, 0x29, 0x85, 0x22, 0xa5,
	0x76, 0xb3, 0x39, 0xa6, 0x81, 0x98, 0xb3, 0x6c, 0x4e, 0x5a, 0xfb, 0xd4, 0xd7, 0xfe, 0x7b, 0xbd,
	0xdf, 0x6f, 0x8f, 0xbd, 0xdf, 0x6f, 0x6f, 0x65, 0xf6, 0x64, 0x4d, 0x36, 0x39, 0x31, 0xe3, 0x93,
	0x49, 0xcc, 0x7c, 0x66, 0xe6, 0x3b, 0x67, 0xbf, 0x43, 0xc8, 0x71, 0xba, 0xc7, 0xa9, 0x1f, 0xb8,
	0xcc, 0x0f, 0xd4, 0x4e, 0x97, 0x71, 0xd6, 0x54, 0xe9, 0x1e, 0x6f, 0x2a, 0xbd, 0xeb, 0xfa, 0x82,
	0xc3, 0x7a, 0x17, 0xe2, 0xb6, 0xad, 0xec, 0xbf, 0x2c, 0x1e, 0x34, 0xe7, 0x17, 0x1d, 0xc6, 0x1c,
	0x8f, 0x8a, 0x50, 0x2b, 0xdc, 0x51, 0x6d, 0x1a, 0xb4, 0xbb, 0x6e, 0x87, 0xb3, 0xae, 0x78, 0xe5,
	0xc4, 0x38, 0xa9, 0x6c, 0xd0, 0x20, 0x68, 0x39, 0xf4, 0x54, 0x8d, 0x94, 0x57, 0xfd, 0x70, 0xb7,
	0x3e, 0x46, 0xca, 0x5b, 0xab, 0xe6, 0x66, 0xed, 0x3f, 0xe3, 0x0c, 0x99, 0xda, 0x47, 0x6e, 0x5b,
	0x8c, 0x79, 0xf5, 0x05, 0x45, 0x10, 0x95, 0x3e, 0x51, 0x89, 0xa2, 0x37, 0x3b, 0x1c, 0x8a, 0x9c,
	0x7d, 0x5b, 0x59, 0x2c, 0x2d, 0x8d, 0x99, 0x93, 0xfb, 0x71, 0x2b, 0x8c, 0x79, 0x86, 0x9f, 0x04,
	0x51, 0x48, 0x52, 0x08, 0x7a, 0x07, 0xa0, 0xa9, 0xc6, 0x49, 0xa5, 0xa0, 0x47, 0x05, 0x6a, 0x4e,
	0xe4, 0x83, 0x5b, 0x63, 0x9d, 0x54, 0xe3, 0x7c, 0xae, 0xcf, 0x9b, 0x8d, 0xe2, 0x84, 0xef, 0x21,
	0xe1, 0x88, 0x19, 0x17, 0xba, 0x0e, 0x71, 0xc6, 0x39, 0x52, 0x8b, 0x51, 0x01, 0x92, 0xf5, 0x01,
	0x58, 0xd3, 0x66, 0x5c, 0xc4, 0x05, 0x37, 0x0b, 0x0b, 0x91, 0xb0, 0x8f, 0x00, 0x9b, 0x4c, 0xc0,
	0x2e, 0x0a, 0x58, 0xba, 0x49, 0x5d, 0x2b, 0x66, 0x7d, 0x02, 0xd6, 0xf0, 0x60, 0x93, 0xba, 0x96,
	0x6d, 0x12, 0xc3, 0xfa, 0x0c, 0xac, 0x7a, 0xaa, 0xc9, 0x34, 0x2c, 0x44, 0xc2, 0xbe, 0x00, 0xac,
	0x9c, 0x6a, 0x52, 0xd7, 0x8c, 0x4d, 0x52, 0x4f, 0x54, 0xb6, 0xe3, 0xee, 0x51, 0x1b, 0xa3, 0xd9,
	0x57, 0xc0, 0x55, 0xcd, 0xe9, 0xb8, 0xb6, 0x28, 0xd4, 0xd8, 0x20, 0xf1, 0xc3, 0x6d, 0x34, 0xef,
	0x1b, 0xf0, 0x2a, 0x66, 0xdc, 0xd8, 0x5a, 0x84, 0x1b, 0x18, 0xc2, 0x8e, 0xc7, 0x5a, 0xbc, 0x18,
	0xf6, 0x1d, 0x60, 0x43, 0x89, 0x21, 0xac, 0x41, 0x9c, 0xac, 0x55, 0x8c, 0x72, 0x3f, 0x80, 0x56,
	0xcb, 0xb4, 0xaa, 0x6b, 0x92, 0x56, 0x31, 0xbc, 0x9f, 0xc0, 0x1b, 0x4d, 0xb7, 0x9a, 0x9e, 0xab,
	0xcd, 0x42, 0xcb, 0xa3, 0xc5, 0xb4, 0x5f, 0x40, 0x2b, 0x25, 0xe6, 0x7a, 0xba, 0x17, 0x98, 0x3a,
	0x71, 0xbc, 0xeb, 0xfa, 0x4e, 0x31, 0xec, 0x37, 0xc0, 0xc6, 0x93, 0x27, 0xae, 0x17, 0x38, 0x38,
	0x04, 0xeb, 0x26, 0xa7, 0x41, 0x31, 0xeb, 0x0f, 0xb0, 0x26, 0x12, 0x43, 0x58, 0x81, 0x38, 0xe3,
	0x46, 0x52, 0xb3, 0x28, 0xa4, 0x18, 0xf6, 0x17, 0x60, 0xff, 0x37, 0x96, 0x0a, 0xcd, 0x2a, 0x8a,
	0x4b, 0xa8, 0x1b, 0x3d, 0x31, 0x2e, 0x91, 0x99, 0x2e, 0xed, 0xd0, 0x16, 0xa7, 0xf6, 0xf6, 0x61,
	0x4d, 0xf7, 0x4e, 0x75, 0x71, 0x78, 0x69, 0xcc, 0x3c, 0xd2, 0x07, 0xac, 0x0e, 0x98, 0xef, 0x2d,
	0x29, 0x19, 0xe7, 0xc2, 0x77, 0x81, 0x8c, 0x76, 0xe1, 0x6c, 0x01, 0x3d, 0x37, 0xde, 0x22, 0xb3,
	0x92, 0x02, 0x90, 0xee, 0x77, 0x0f, 0x2a, 0x18, 0x31, 0x8f, 0x66, 0xd0, 0xc2, 0x9e, 0x2f, 0x93,
	0x39, 0x09, 0x1b, 0xeb, 0xd3, 0xf7, 0x01, 0x3e, 0x6d, 0xce, 0x64, 0xe0, 0x91, 0x5f, 0xcb, 0xe9,
	0x58, 0xe3, 0x7e, 0x00, 0xf4, 0x49, 0x09, 0x3d, 0x32, 0xf0, 0x5c, 0x5d, 0x30, 0x3f, 0xd3, 0x87,
	0x00, 0x1f, 0x96, 0xeb, 0xa2, 0x6b, 0x07, 0xe8, 0x82, 0x81, 0x3f, 0x02, 0x78, 0x3d, 0x47, 0x97,
	0x5c, 0x3a, 0xd6, 0xeb, 0x1f, 0x03, 0xbd, 0x9c, 0xa3, 0x8b, 0xae, 0x19, 0x57, 0xc9, 0x31, 0x59,
	0xed, 0x68, 0xb3, 0x7e, 0x02, 0xfc, 0xaa, 0x39, 0x97, 0xad, 0xbe, 0xbf, 0x04, 0xae, 0x90, 0x79,
	0x49, 0x06, 0x74, 0x82, 0xa7, 0x90, 0xa0, 0x62, 0xce, 0x66, 0x12, 0xf4, 0xb7, 0x82, 0x7c, 0xb2,
	0xc8, 0xf5, 0xf0, 0x0c, 0xe8, 0x43, 0x92, 0xc9, 0x8a, 0x35, 0x71, 0x90, 0x3a, 0x18, 0xf5, 0x9f,
	0x03, 0xbe, 0x96, 0xab, 0x8e, 0xae, 0x1d, 0xa4, 0x0e, 0x26, 0xc1, 0x0b, 0x48, 0x30, 0x9a, 0xa7,
	0x4e, 0xee, 0xe9, 0xc1, 0x6e, 0x94, 0x97, 0x80, 0x2f, 0x49, 0x4e, 0x4f, 0xb4, 0x59, 0x72, 0x4e,
	0x3e, 0x72, 0xc5, 0xbc, 0x02, 0xfa, 0xb8, 0xec, 0xe4, 0x8b, 0x55, 0x23, 0x9f, 0x2c, 0x72, 0xe7,
	0xbc, 0x06, 0xf8, 0x84, 0x64, 0xb2, 0x62, 0xf7, 0xdc, 0x2e, 0x49, 0x85, 0x47, 0x6f, 0xa1, 0x37,
	0x80, 0x3f, 0xcc, 0x16, 0xca, 0x8e, 0x28, 0xfa, 0xcf, 0xca, 0xf9, 0xad, 0xb3, 0x8e, 0xcb, 0xaf,
	0x85, 0x96, 0xd2, 0x66, 0xbb, 0xaa, 0xc3, 0xbc, 0x96, 0xef, 0xc4, 0x5f, 0x11, 0xd7, 0x1b, 0x6a,
	0x7b, 0xd7, 0x16, 0xf7, 0xed, 0x65, 0x87, 0xfa, 0xcb, 0x0e, 0x53, 0x39, 0x0d, 0xb8, 0xdd, 0xe2,
	0x2d, 0x35, 0xf3, 0xc9, 0x62, 0x8d, 0x8a, 0xbf, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x65, 0x92,
	0xaa, 0xed, 0xce, 0x0c, 0x00, 0x00,
}

func init() {
	xxx_Ext3_ProtoFile_FileDesc.Enums = xxx_Ext3_ProtoFile_EnumDescs[0:1]
	xxx_Ext3_ProtoFile_FileDesc.Messages = xxx_Ext3_ProtoFile_MessageDescs[0:1]
	var err error
	Ext3_ProtoFile, err = prototype.NewFile(&xxx_Ext3_ProtoFile_FileDesc)
	if err != nil {
		panic(err)
	}
}

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var Ext3_ProtoFile protoreflect.FileDescriptor

var xxx_Ext3_ProtoFile_FileDesc = prototype.File{
	Syntax:  protoreflect.Proto3,
	Path:    "extensions/proto3/ext3.proto",
	Package: "goproto.protoc.extension.proto3",
	Imports: []protoreflect.FileImport{
		{FileDescriptor: prototype.PlaceholderFile("google/protobuf/descriptor.proto", "google.protobuf")},
	},
}
var xxx_Ext3_ProtoFile_EnumTypes = [1]protoreflect.EnumType{
	prototype.GoEnum(
		xxx_Ext3_ProtoFile_EnumDescs[0].Reference(),
		func(_ protoreflect.EnumType, n protoreflect.EnumNumber) protoreflect.Enum {
			return Enum(n)
		},
	),
}
var xxx_Ext3_ProtoFile_EnumDescs = [1]prototype.Enum{
	{
		Name: "Enum",
		Values: []prototype.EnumValue{
			{Name: "ZERO", Number: 0},
		},
	},
}
var xxx_Ext3_ProtoFile_MessageTypes = [1]protoimpl.MessageType{
	{Type: prototype.GoMessage(
		xxx_Ext3_ProtoFile_MessageDescs[0].Reference(),
		func(protoreflect.MessageType) protoreflect.Message {
			return xxx_Message{new(Message)}
		},
	)},
}
var xxx_Ext3_ProtoFile_MessageDescs = [1]prototype.Message{
	{
		Name: "Message",
	},
}
