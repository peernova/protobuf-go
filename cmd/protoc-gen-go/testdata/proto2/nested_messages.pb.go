// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto2/nested_messages.proto

package proto2

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	prototype "github.com/golang/protobuf/v2/reflect/prototype"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Layer1 struct {
	L2                   *Layer1_Layer2        `protobuf:"bytes,1,opt,name=l2" json:"l2,omitempty"`
	L3                   *Layer1_Layer2_Layer3 `protobuf:"bytes,2,opt,name=l3" json:"l3,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

type xxx_Layer1 struct{ m *Layer1 }

func (m *Layer1) ProtoReflect() protoreflect.Message {
	return xxx_Layer1{m}
}
func (m xxx_Layer1) Type() protoreflect.MessageType {
	return xxx_NestedMessages_ProtoFile_MessageTypes[0].Type
}
func (m xxx_Layer1) KnownFields() protoreflect.KnownFields {
	return xxx_NestedMessages_ProtoFile_MessageTypes[0].KnownFieldsOf(m.m)
}
func (m xxx_Layer1) UnknownFields() protoreflect.UnknownFields {
	return xxx_NestedMessages_ProtoFile_MessageTypes[0].UnknownFieldsOf(m.m)
}
func (m xxx_Layer1) Interface() protoreflect.ProtoMessage {
	return m.m
}

func (m *Layer1) Reset()         { *m = Layer1{} }
func (m *Layer1) String() string { return proto.CompactTextString(m) }
func (*Layer1) ProtoMessage()    {}
func (*Layer1) Descriptor() ([]byte, []int) {
	return fileDescriptor_7417ee157699d191, []int{0}
}

func (m *Layer1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Layer1.Unmarshal(m, b)
}
func (m *Layer1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Layer1.Marshal(b, m, deterministic)
}
func (m *Layer1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Layer1.Merge(m, src)
}
func (m *Layer1) XXX_Size() int {
	return xxx_messageInfo_Layer1.Size(m)
}
func (m *Layer1) XXX_DiscardUnknown() {
	xxx_messageInfo_Layer1.DiscardUnknown(m)
}

var xxx_messageInfo_Layer1 proto.InternalMessageInfo

func (m *Layer1) GetL2() *Layer1_Layer2 {
	if m != nil {
		return m.L2
	}
	return nil
}

func (m *Layer1) GetL3() *Layer1_Layer2_Layer3 {
	if m != nil {
		return m.L3
	}
	return nil
}

type Layer1_Layer2 struct {
	L3                   *Layer1_Layer2_Layer3 `protobuf:"bytes,1,opt,name=l3" json:"l3,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

type xxx_Layer1_Layer2 struct{ m *Layer1_Layer2 }

func (m *Layer1_Layer2) ProtoReflect() protoreflect.Message {
	return xxx_Layer1_Layer2{m}
}
func (m xxx_Layer1_Layer2) Type() protoreflect.MessageType {
	return xxx_NestedMessages_ProtoFile_MessageTypes[1].Type
}
func (m xxx_Layer1_Layer2) KnownFields() protoreflect.KnownFields {
	return xxx_NestedMessages_ProtoFile_MessageTypes[1].KnownFieldsOf(m.m)
}
func (m xxx_Layer1_Layer2) UnknownFields() protoreflect.UnknownFields {
	return xxx_NestedMessages_ProtoFile_MessageTypes[1].UnknownFieldsOf(m.m)
}
func (m xxx_Layer1_Layer2) Interface() protoreflect.ProtoMessage {
	return m.m
}

func (m *Layer1_Layer2) Reset()         { *m = Layer1_Layer2{} }
func (m *Layer1_Layer2) String() string { return proto.CompactTextString(m) }
func (*Layer1_Layer2) ProtoMessage()    {}
func (*Layer1_Layer2) Descriptor() ([]byte, []int) {
	return fileDescriptor_7417ee157699d191, []int{0, 0}
}

func (m *Layer1_Layer2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Layer1_Layer2.Unmarshal(m, b)
}
func (m *Layer1_Layer2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Layer1_Layer2.Marshal(b, m, deterministic)
}
func (m *Layer1_Layer2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Layer1_Layer2.Merge(m, src)
}
func (m *Layer1_Layer2) XXX_Size() int {
	return xxx_messageInfo_Layer1_Layer2.Size(m)
}
func (m *Layer1_Layer2) XXX_DiscardUnknown() {
	xxx_messageInfo_Layer1_Layer2.DiscardUnknown(m)
}

var xxx_messageInfo_Layer1_Layer2 proto.InternalMessageInfo

func (m *Layer1_Layer2) GetL3() *Layer1_Layer2_Layer3 {
	if m != nil {
		return m.L3
	}
	return nil
}

type Layer1_Layer2_Layer3 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type xxx_Layer1_Layer2_Layer3 struct{ m *Layer1_Layer2_Layer3 }

func (m *Layer1_Layer2_Layer3) ProtoReflect() protoreflect.Message {
	return xxx_Layer1_Layer2_Layer3{m}
}
func (m xxx_Layer1_Layer2_Layer3) Type() protoreflect.MessageType {
	return xxx_NestedMessages_ProtoFile_MessageTypes[2].Type
}
func (m xxx_Layer1_Layer2_Layer3) KnownFields() protoreflect.KnownFields {
	return xxx_NestedMessages_ProtoFile_MessageTypes[2].KnownFieldsOf(m.m)
}
func (m xxx_Layer1_Layer2_Layer3) UnknownFields() protoreflect.UnknownFields {
	return xxx_NestedMessages_ProtoFile_MessageTypes[2].UnknownFieldsOf(m.m)
}
func (m xxx_Layer1_Layer2_Layer3) Interface() protoreflect.ProtoMessage {
	return m.m
}

func (m *Layer1_Layer2_Layer3) Reset()         { *m = Layer1_Layer2_Layer3{} }
func (m *Layer1_Layer2_Layer3) String() string { return proto.CompactTextString(m) }
func (*Layer1_Layer2_Layer3) ProtoMessage()    {}
func (*Layer1_Layer2_Layer3) Descriptor() ([]byte, []int) {
	return fileDescriptor_7417ee157699d191, []int{0, 0, 0}
}

func (m *Layer1_Layer2_Layer3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Layer1_Layer2_Layer3.Unmarshal(m, b)
}
func (m *Layer1_Layer2_Layer3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Layer1_Layer2_Layer3.Marshal(b, m, deterministic)
}
func (m *Layer1_Layer2_Layer3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Layer1_Layer2_Layer3.Merge(m, src)
}
func (m *Layer1_Layer2_Layer3) XXX_Size() int {
	return xxx_messageInfo_Layer1_Layer2_Layer3.Size(m)
}
func (m *Layer1_Layer2_Layer3) XXX_DiscardUnknown() {
	xxx_messageInfo_Layer1_Layer2_Layer3.DiscardUnknown(m)
}

var xxx_messageInfo_Layer1_Layer2_Layer3 proto.InternalMessageInfo

func init() {
	proto.RegisterFile("proto2/nested_messages.proto", fileDescriptor_7417ee157699d191)
	proto.RegisterType((*Layer1)(nil), "goproto.protoc.proto2.Layer1")
	proto.RegisterType((*Layer1_Layer2)(nil), "goproto.protoc.proto2.Layer1.Layer2")
	proto.RegisterType((*Layer1_Layer2_Layer3)(nil), "goproto.protoc.proto2.Layer1.Layer2.Layer3")
}

var fileDescriptor_7417ee157699d191 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0x37, 0xd2, 0xcf, 0x4b, 0x2d, 0x2e, 0x49, 0x4d, 0x89, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c,
	0x4f, 0x2d, 0xd6, 0x03, 0x0b, 0x0b, 0x89, 0xa6, 0xe7, 0x83, 0x19, 0x10, 0x6e, 0x32, 0x84, 0x32,
	0x52, 0x3a, 0xc3, 0xc8, 0xc5, 0xe6, 0x93, 0x58, 0x99, 0x5a, 0x64, 0x28, 0x64, 0xc2, 0xc5, 0x94,
	0x63, 0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0xa4, 0xa2, 0x87, 0x55, 0xb9, 0x1e, 0x44, 0x29,
	0x84, 0x32, 0x0a, 0x62, 0xca, 0x31, 0x12, 0xb2, 0xe6, 0x62, 0xca, 0x31, 0x96, 0x60, 0x02, 0xeb,
	0xd2, 0x26, 0x46, 0x17, 0x84, 0x32, 0x0e, 0x62, 0xca, 0x31, 0x96, 0xf2, 0x87, 0x5a, 0x0e, 0x33,
	0x86, 0x91, 0x3c, 0x63, 0x38, 0xa0, 0xc6, 0x18, 0x3b, 0x39, 0x46, 0xd9, 0xa7, 0x67, 0x96, 0x64,
	0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xe7, 0x24, 0xe6, 0xa5, 0xeb, 0x83, 0xb5,
	0x27, 0x95, 0xa6, 0xe9, 0x97, 0x19, 0xe9, 0x27, 0xe7, 0xa6, 0x40, 0xf8, 0xc9, 0xba, 0xe9, 0xa9,
	0x79, 0xba, 0xe9, 0xf9, 0xfa, 0x25, 0xa9, 0xc5, 0x25, 0x29, 0x89, 0x25, 0x89, 0x10, 0x61, 0x23,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x11, 0x21, 0xb0, 0x06, 0x47, 0x01, 0x00, 0x00,
}

func init() {
	xxx_NestedMessages_ProtoFile_FileDesc.Messages = xxx_NestedMessages_ProtoFile_MessageDescs[0:1]
	xxx_NestedMessages_ProtoFile_MessageDescs[0].Messages = xxx_NestedMessages_ProtoFile_MessageDescs[1:2]
	xxx_NestedMessages_ProtoFile_MessageDescs[1].Messages = xxx_NestedMessages_ProtoFile_MessageDescs[2:3]
	xxx_NestedMessages_ProtoFile_MessageDescs[0].Fields[0].MessageType = xxx_NestedMessages_ProtoFile_MessageTypes[1].Type
	xxx_NestedMessages_ProtoFile_MessageDescs[0].Fields[1].MessageType = xxx_NestedMessages_ProtoFile_MessageTypes[2].Type
	xxx_NestedMessages_ProtoFile_MessageDescs[1].Fields[0].MessageType = xxx_NestedMessages_ProtoFile_MessageTypes[2].Type
	var err error
	NestedMessages_ProtoFile, err = prototype.NewFile(&xxx_NestedMessages_ProtoFile_FileDesc)
	if err != nil {
		panic(err)
	}
}

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var NestedMessages_ProtoFile protoreflect.FileDescriptor

var xxx_NestedMessages_ProtoFile_FileDesc = prototype.File{
	Syntax:  protoreflect.Proto2,
	Path:    "proto2/nested_messages.proto",
	Package: "goproto.protoc.proto2",
}
var xxx_NestedMessages_ProtoFile_MessageTypes = [3]protoimpl.MessageType{
	{Type: prototype.GoMessage(
		xxx_NestedMessages_ProtoFile_MessageDescs[0].Reference(),
		func(protoreflect.MessageType) protoreflect.Message {
			return xxx_Layer1{new(Layer1)}
		},
	)},
	{Type: prototype.GoMessage(
		xxx_NestedMessages_ProtoFile_MessageDescs[1].Reference(),
		func(protoreflect.MessageType) protoreflect.Message {
			return xxx_Layer1_Layer2{new(Layer1_Layer2)}
		},
	)},
	{Type: prototype.GoMessage(
		xxx_NestedMessages_ProtoFile_MessageDescs[2].Reference(),
		func(protoreflect.MessageType) protoreflect.Message {
			return xxx_Layer1_Layer2_Layer3{new(Layer1_Layer2_Layer3)}
		},
	)},
}
var xxx_NestedMessages_ProtoFile_MessageDescs = [3]prototype.Message{
	{
		Name: "Layer1",
		Fields: []prototype.Field{
			{
				Name:        "l2",
				Number:      1,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.MessageKind,
				JSONName:    "l2",
				IsPacked:    prototype.False,
			},
			{
				Name:        "l3",
				Number:      2,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.MessageKind,
				JSONName:    "l3",
				IsPacked:    prototype.False,
			},
		},
	},
	{
		Name: "Layer2",
		Fields: []prototype.Field{
			{
				Name:        "l3",
				Number:      1,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.MessageKind,
				JSONName:    "l3",
				IsPacked:    prototype.False,
			},
		},
	},
	{
		Name: "Layer3",
	},
}
