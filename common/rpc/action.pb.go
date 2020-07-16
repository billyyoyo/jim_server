// Code generated by protoc-gen-go. DO NOT EDIT.
// source: action.proto

package rpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ActType int32

const (
	ActType_ACT_UNKNOWN  ActType = 0
	ActType_ACT_JOIN     ActType = 1
	ActType_ACT_QUIT     ActType = 2
	ActType_ACT_WITHDRAW ActType = 3
	ActType_ACT_SYNC     ActType = 4
	ActType_ACT_RENAME   ActType = 5
)

var ActType_name = map[int32]string{
	0: "ACT_UNKNOWN",
	1: "ACT_JOIN",
	2: "ACT_QUIT",
	3: "ACT_WITHDRAW",
	4: "ACT_SYNC",
	5: "ACT_RENAME",
}

var ActType_value = map[string]int32{
	"ACT_UNKNOWN":  0,
	"ACT_JOIN":     1,
	"ACT_QUIT":     2,
	"ACT_WITHDRAW": 3,
	"ACT_SYNC":     4,
	"ACT_RENAME":   5,
}

func (x ActType) String() string {
	return proto.EnumName(ActType_name, int32(x))
}

func (ActType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_59885c909ad4dfd3, []int{0}
}

type Action struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	DeviceId             int64    `protobuf:"varint,2,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	RequestId            int64    `protobuf:"varint,3,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Time                 int64    `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`
	Type                 ActType  `protobuf:"varint,5,opt,name=type,proto3,enum=rpc.ActType" json:"type,omitempty"`
	Excute               string   `protobuf:"bytes,6,opt,name=excute,proto3" json:"excute,omitempty"`
	Data                 []byte   `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Action) Reset()         { *m = Action{} }
func (m *Action) String() string { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()    {}
func (*Action) Descriptor() ([]byte, []int) {
	return fileDescriptor_59885c909ad4dfd3, []int{0}
}

func (m *Action) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Action.Unmarshal(m, b)
}
func (m *Action) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Action.Marshal(b, m, deterministic)
}
func (m *Action) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Action.Merge(m, src)
}
func (m *Action) XXX_Size() int {
	return xxx_messageInfo_Action.Size(m)
}
func (m *Action) XXX_DiscardUnknown() {
	xxx_messageInfo_Action.DiscardUnknown(m)
}

var xxx_messageInfo_Action proto.InternalMessageInfo

func (m *Action) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Action) GetDeviceId() int64 {
	if m != nil {
		return m.DeviceId
	}
	return 0
}

func (m *Action) GetRequestId() int64 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *Action) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Action) GetType() ActType {
	if m != nil {
		return m.Type
	}
	return ActType_ACT_UNKNOWN
}

func (m *Action) GetExcute() string {
	if m != nil {
		return m.Excute
	}
	return ""
}

func (m *Action) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("rpc.ActType", ActType_name, ActType_value)
	proto.RegisterType((*Action)(nil), "rpc.Action")
}

func init() { proto.RegisterFile("action.proto", fileDescriptor_59885c909ad4dfd3) }

var fileDescriptor_59885c909ad4dfd3 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0xcd, 0xfa, 0x67, 0xdb, 0x6b, 0x99, 0xe1, 0x3d, 0x48, 0x10, 0x0f, 0xc5, 0x53, 0xf1,
	0xd0, 0x83, 0x7e, 0x82, 0x3a, 0x07, 0x46, 0x31, 0xc3, 0xd8, 0x51, 0xf4, 0x22, 0x35, 0x8d, 0xd0,
	0x83, 0xb6, 0x66, 0xa9, 0xb8, 0xef, 0xe6, 0x87, 0x93, 0x84, 0x4e, 0x6f, 0xcf, 0xef, 0xf9, 0x91,
	0xf0, 0xf2, 0x40, 0x52, 0x2b, 0xdb, 0x76, 0x1f, 0x79, 0x6f, 0x3a, 0xdb, 0x61, 0x60, 0x7a, 0x75,
	0xf6, 0x43, 0x20, 0x2e, 0x7c, 0x8b, 0xc7, 0x10, 0x0f, 0x5b, 0x6d, 0x78, 0xc3, 0x48, 0x4a, 0xb2,
	0x40, 0x8e, 0x84, 0x27, 0x30, 0x6b, 0xf4, 0x57, 0xab, 0x34, 0x6f, 0xd8, 0xc4, 0x9b, 0x3f, 0xc6,
	0x53, 0x98, 0x1b, 0xfd, 0x39, 0xe8, 0xad, 0xe5, 0x0d, 0x0b, 0xbc, 0xfc, 0x2f, 0x10, 0x21, 0xb4,
	0xed, 0xbb, 0x66, 0xa1, 0x17, 0x3e, 0x63, 0x0a, 0xa1, 0xdd, 0xf5, 0x9a, 0x45, 0x29, 0xc9, 0x16,
	0x17, 0x49, 0x6e, 0x7a, 0x95, 0x17, 0xca, 0x96, 0xbb, 0x5e, 0x4b, 0x6f, 0xdc, 0x1d, 0xfa, 0x5b,
	0x0d, 0x56, 0xb3, 0x38, 0x25, 0xd9, 0x5c, 0x8e, 0xe4, 0x7e, 0x6b, 0x6a, 0x5b, 0xb3, 0x69, 0x4a,
	0xb2, 0x44, 0xfa, 0x7c, 0xfe, 0x06, 0xd3, 0xf1, 0x31, 0x1e, 0xc1, 0x61, 0xb1, 0x2c, 0x5f, 0x36,
	0xe2, 0x4e, 0xac, 0x2b, 0x41, 0x0f, 0x30, 0x81, 0x99, 0x2b, 0x6e, 0xd7, 0x5c, 0x50, 0xb2, 0xa7,
	0x87, 0x0d, 0x2f, 0xe9, 0x04, 0x29, 0x24, 0x8e, 0x2a, 0x5e, 0xde, 0x5c, 0xcb, 0xa2, 0xa2, 0xc1,
	0xde, 0x3f, 0x3e, 0x89, 0x25, 0x0d, 0x71, 0x01, 0xe0, 0x48, 0xae, 0x44, 0x71, 0xbf, 0xa2, 0xd1,
	0x55, 0xf4, 0xec, 0xd6, 0x7a, 0x8d, 0xfd, 0x72, 0x97, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4f,
	0x2d, 0x48, 0x78, 0x49, 0x01, 0x00, 0x00,
}