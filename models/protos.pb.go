// Code generated by protoc-gen-go.
// source: protos.proto
// DO NOT EDIT!

/*
Package models is a generated protocol buffer package.

It is generated from these files:
	protos.proto

It has these top-level messages:
	AddRequest
	AddReturns
	SetCheckedRequest
	SetCheckedResponse
	RemoveRequest
	RemoveResponse
	ItemsRequest
	ItemStreamReturns
	Item
*/
package models

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AddRequest struct {
	Title string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
}

func (m *AddRequest) Reset()                    { *m = AddRequest{} }
func (m *AddRequest) String() string            { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()               {}
func (*AddRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type AddReturns struct {
	Item *Item `protobuf:"bytes,1,opt,name=item" json:"item,omitempty"`
}

func (m *AddReturns) Reset()                    { *m = AddReturns{} }
func (m *AddReturns) String() string            { return proto.CompactTextString(m) }
func (*AddReturns) ProtoMessage()               {}
func (*AddReturns) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddReturns) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

type SetCheckedRequest struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Checked bool   `protobuf:"varint,2,opt,name=checked" json:"checked,omitempty"`
}

func (m *SetCheckedRequest) Reset()                    { *m = SetCheckedRequest{} }
func (m *SetCheckedRequest) String() string            { return proto.CompactTextString(m) }
func (*SetCheckedRequest) ProtoMessage()               {}
func (*SetCheckedRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SetCheckedRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SetCheckedRequest) GetChecked() bool {
	if m != nil {
		return m.Checked
	}
	return false
}

type SetCheckedResponse struct {
}

func (m *SetCheckedResponse) Reset()                    { *m = SetCheckedResponse{} }
func (m *SetCheckedResponse) String() string            { return proto.CompactTextString(m) }
func (*SetCheckedResponse) ProtoMessage()               {}
func (*SetCheckedResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type RemoveRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *RemoveRequest) Reset()                    { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()               {}
func (*RemoveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RemoveRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RemoveResponse struct {
}

func (m *RemoveResponse) Reset()                    { *m = RemoveResponse{} }
func (m *RemoveResponse) String() string            { return proto.CompactTextString(m) }
func (*RemoveResponse) ProtoMessage()               {}
func (*RemoveResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ItemsRequest struct {
	Index int64 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	Count int64 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
}

func (m *ItemsRequest) Reset()                    { *m = ItemsRequest{} }
func (m *ItemsRequest) String() string            { return proto.CompactTextString(m) }
func (*ItemsRequest) ProtoMessage()               {}
func (*ItemsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ItemsRequest) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ItemsRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ItemStreamReturns struct {
	Item  *Item `protobuf:"bytes,1,opt,name=item" json:"item,omitempty"`
	Index int64 `protobuf:"varint,2,opt,name=index" json:"index,omitempty"`
	Total int64 `protobuf:"varint,3,opt,name=total" json:"total,omitempty"`
}

func (m *ItemStreamReturns) Reset()                    { *m = ItemStreamReturns{} }
func (m *ItemStreamReturns) String() string            { return proto.CompactTextString(m) }
func (*ItemStreamReturns) ProtoMessage()               {}
func (*ItemStreamReturns) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ItemStreamReturns) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *ItemStreamReturns) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ItemStreamReturns) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type Item struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Created []byte `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
	Changed []byte `protobuf:"bytes,4,opt,name=changed,proto3" json:"changed,omitempty"`
	Checked bool   `protobuf:"varint,5,opt,name=checked" json:"checked,omitempty"`
}

func (m *Item) Reset()                    { *m = Item{} }
func (m *Item) String() string            { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()               {}
func (*Item) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

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

func (m *Item) GetCreated() []byte {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Item) GetChanged() []byte {
	if m != nil {
		return m.Changed
	}
	return nil
}

func (m *Item) GetChecked() bool {
	if m != nil {
		return m.Checked
	}
	return false
}

func init() {
	proto.RegisterType((*AddRequest)(nil), "models.AddRequest")
	proto.RegisterType((*AddReturns)(nil), "models.AddReturns")
	proto.RegisterType((*SetCheckedRequest)(nil), "models.SetCheckedRequest")
	proto.RegisterType((*SetCheckedResponse)(nil), "models.SetCheckedResponse")
	proto.RegisterType((*RemoveRequest)(nil), "models.RemoveRequest")
	proto.RegisterType((*RemoveResponse)(nil), "models.RemoveResponse")
	proto.RegisterType((*ItemsRequest)(nil), "models.ItemsRequest")
	proto.RegisterType((*ItemStreamReturns)(nil), "models.ItemStreamReturns")
	proto.RegisterType((*Item)(nil), "models.Item")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Grpc service

type GrpcClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReturns, error)
	SetChecked(ctx context.Context, in *SetCheckedRequest, opts ...grpc.CallOption) (*SetCheckedResponse, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error)
	Items(ctx context.Context, in *ItemsRequest, opts ...grpc.CallOption) (Grpc_ItemsClient, error)
}

type grpcClient struct {
	cc *grpc.ClientConn
}

func NewGrpcClient(cc *grpc.ClientConn) GrpcClient {
	return &grpcClient{cc}
}

func (c *grpcClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReturns, error) {
	out := new(AddReturns)
	err := grpc.Invoke(ctx, "/models.Grpc/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcClient) SetChecked(ctx context.Context, in *SetCheckedRequest, opts ...grpc.CallOption) (*SetCheckedResponse, error) {
	out := new(SetCheckedResponse)
	err := grpc.Invoke(ctx, "/models.Grpc/SetChecked", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	out := new(RemoveResponse)
	err := grpc.Invoke(ctx, "/models.Grpc/Remove", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcClient) Items(ctx context.Context, in *ItemsRequest, opts ...grpc.CallOption) (Grpc_ItemsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Grpc_serviceDesc.Streams[0], c.cc, "/models.Grpc/Items", opts...)
	if err != nil {
		return nil, err
	}
	x := &grpcItemsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Grpc_ItemsClient interface {
	Recv() (*ItemStreamReturns, error)
	grpc.ClientStream
}

type grpcItemsClient struct {
	grpc.ClientStream
}

func (x *grpcItemsClient) Recv() (*ItemStreamReturns, error) {
	m := new(ItemStreamReturns)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Grpc service

type GrpcServer interface {
	Add(context.Context, *AddRequest) (*AddReturns, error)
	SetChecked(context.Context, *SetCheckedRequest) (*SetCheckedResponse, error)
	Remove(context.Context, *RemoveRequest) (*RemoveResponse, error)
	Items(*ItemsRequest, Grpc_ItemsServer) error
}

func RegisterGrpcServer(s *grpc.Server, srv GrpcServer) {
	s.RegisterService(&_Grpc_serviceDesc, srv)
}

func _Grpc_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.Grpc/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grpc_SetChecked_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetCheckedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).SetChecked(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.Grpc/SetChecked",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).SetChecked(ctx, req.(*SetCheckedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grpc_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.Grpc/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grpc_Items_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ItemsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GrpcServer).Items(m, &grpcItemsServer{stream})
}

type Grpc_ItemsServer interface {
	Send(*ItemStreamReturns) error
	grpc.ServerStream
}

type grpcItemsServer struct {
	grpc.ServerStream
}

func (x *grpcItemsServer) Send(m *ItemStreamReturns) error {
	return x.ServerStream.SendMsg(m)
}

var _Grpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "models.Grpc",
	HandlerType: (*GrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Grpc_Add_Handler,
		},
		{
			MethodName: "SetChecked",
			Handler:    _Grpc_SetChecked_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Grpc_Remove_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Items",
			Handler:       _Grpc_Items_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos.proto",
}

func init() { proto.RegisterFile("protos.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x4e, 0xf3, 0x30,
	0x10, 0x6c, 0x7e, 0xda, 0xef, 0x63, 0x09, 0x15, 0xb5, 0x0a, 0x4a, 0x73, 0xa1, 0xf2, 0xa9, 0xa7,
	0x80, 0xca, 0x09, 0x04, 0x87, 0x0a, 0x21, 0xc4, 0xd5, 0x7d, 0x82, 0x10, 0xaf, 0x20, 0xa2, 0x89,
	0x4b, 0xec, 0x22, 0xc4, 0x8b, 0x73, 0x45, 0xb1, 0xe3, 0x26, 0x69, 0x41, 0xe2, 0x64, 0xcd, 0xee,
	0xcc, 0xda, 0x3b, 0x63, 0x08, 0xd6, 0xa5, 0x50, 0x42, 0xc6, 0xfa, 0x20, 0x83, 0x5c, 0x70, 0x5c,
	0x49, 0x4a, 0x01, 0x16, 0x9c, 0x33, 0x7c, 0xdb, 0xa0, 0x54, 0x64, 0x0c, 0x7d, 0x95, 0xa9, 0x15,
	0x86, 0xce, 0xd4, 0x99, 0x1d, 0x30, 0x03, 0x68, 0x5c, 0x73, 0xd4, 0xa6, 0x2c, 0x24, 0x99, 0x82,
	0x9f, 0x29, 0xcc, 0x35, 0xe5, 0x70, 0x1e, 0xc4, 0x66, 0x50, 0xfc, 0xa8, 0x30, 0x67, 0xba, 0x43,
	0x6f, 0x61, 0xb4, 0x44, 0x75, 0xf7, 0x82, 0xe9, 0x2b, 0x6e, 0x47, 0x0f, 0xc1, 0xcd, 0x78, 0x3d,
	0xd7, 0xcd, 0x38, 0x09, 0xe1, 0x5f, 0x6a, 0x18, 0xa1, 0x3b, 0x75, 0x66, 0xff, 0x99, 0x85, 0x74,
	0x0c, 0xa4, 0x2d, 0x97, 0x6b, 0x51, 0x48, 0xa4, 0x67, 0x70, 0xc4, 0x30, 0x17, 0xef, 0xf8, 0xcb,
	0x40, 0x7a, 0x0c, 0x43, 0x4b, 0xa8, 0x25, 0xd7, 0x10, 0x54, 0xaf, 0x92, 0xad, 0xed, 0xb2, 0x82,
	0xe3, 0x87, 0x16, 0x79, 0xcc, 0x80, 0xaa, 0x9a, 0x8a, 0x4d, 0xa1, 0xf4, 0x33, 0x3c, 0x66, 0x00,
	0x4d, 0x60, 0x54, 0x69, 0x97, 0xaa, 0xc4, 0x24, 0xff, 0xf3, 0xea, 0xcd, 0x15, 0xee, 0xce, 0x15,
	0x4a, 0xa8, 0x64, 0x15, 0x7a, 0xa6, 0xaa, 0x01, 0xfd, 0x04, 0xbf, 0x52, 0xee, 0x39, 0xb3, 0x0d,
	0xc1, 0x6d, 0x85, 0xa0, 0xfd, 0x2a, 0x31, 0x51, 0xc8, 0xf5, 0x94, 0x80, 0x59, 0x68, 0x9c, 0x4c,
	0x8a, 0x67, 0xe4, 0xa1, 0x5f, 0x77, 0x0c, 0x6c, 0x7b, 0xdc, 0xef, 0x78, 0x3c, 0xff, 0x72, 0xc0,
	0x7f, 0x28, 0xd7, 0x29, 0x39, 0x07, 0x6f, 0xc1, 0x39, 0x21, 0x76, 0x97, 0xe6, 0x33, 0x44, 0xdd,
	0x9a, 0x76, 0x80, 0xf6, 0xc8, 0x3d, 0x40, 0x93, 0x0e, 0x99, 0x58, 0xce, 0x5e, 0xe0, 0x51, 0xf4,
	0x53, 0xab, 0x4e, 0xa6, 0x47, 0xae, 0x60, 0x60, 0xd2, 0x22, 0x27, 0x96, 0xd7, 0x89, 0x37, 0x3a,
	0xdd, 0x2d, 0x6f, 0xa5, 0x37, 0xd0, 0xd7, 0xb1, 0x92, 0x71, 0x3b, 0x00, 0x9b, 0x72, 0x34, 0x69,
	0x57, 0x3b, 0xf9, 0xd1, 0xde, 0x85, 0xf3, 0x34, 0xd0, 0xff, 0xff, 0xf2, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0xda, 0x18, 0x16, 0x9b, 0x0f, 0x03, 0x00, 0x00,
}