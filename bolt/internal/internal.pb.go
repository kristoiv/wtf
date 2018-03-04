// Code generated by protoc-gen-go.
// source: internal.proto
// DO NOT EDIT!

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	internal.proto

It has these top-level messages:
	Item
*/
package internal

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Item struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Created int64  `protobuf:"varint,3,opt,name=created" json:"created,omitempty"`
	Changed int64  `protobuf:"varint,4,opt,name=changed" json:"changed,omitempty"`
	Checked bool   `protobuf:"varint,5,opt,name=checked" json:"checked,omitempty"`
}

func (m *Item) Reset()                    { *m = Item{} }
func (m *Item) String() string            { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()               {}
func (*Item) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Item) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Item) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Item) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Item) GetChanged() int64 {
	if m != nil {
		return m.Changed
	}
	return 0
}

func (m *Item) GetChecked() bool {
	if m != nil {
		return m.Checked
	}
	return false
}

func init() {
	proto.RegisterType((*Item)(nil), "internal.Item")
}

func init() { proto.RegisterFile("internal.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0xaa,
	0xb8, 0x58, 0x3c, 0x4b, 0x52, 0x73, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0x98, 0x32, 0x53, 0x84, 0x44, 0xb8, 0x58, 0x4b, 0x32, 0x4b, 0x72, 0x52, 0x25,
	0x98, 0xc0, 0x42, 0x10, 0x8e, 0x90, 0x04, 0x17, 0x7b, 0x72, 0x51, 0x6a, 0x62, 0x49, 0x6a, 0x8a,
	0x04, 0xb3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x8c, 0x0b, 0x96, 0xc9, 0x48, 0xcc, 0x4b, 0x4f, 0x4d,
	0x91, 0x60, 0x81, 0xca, 0x40, 0xb8, 0x10, 0x99, 0xd4, 0xe4, 0xec, 0xd4, 0x14, 0x09, 0x56, 0x05,
	0x46, 0x0d, 0x8e, 0x20, 0x18, 0x37, 0x89, 0x0d, 0xec, 0x18, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x28, 0xfb, 0xba, 0x64, 0x9e, 0x00, 0x00, 0x00,
}
