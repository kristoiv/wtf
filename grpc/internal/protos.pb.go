// Code generated by protoc-gen-go.
// source: protos.proto
// DO NOT EDIT!

/*
Package internal is a generated protocol buffer package.

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
package internal

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
	proto.RegisterType((*AddRequest)(nil), "internal.AddRequest")
	proto.RegisterType((*AddReturns)(nil), "internal.AddReturns")
	proto.RegisterType((*SetCheckedRequest)(nil), "internal.SetCheckedRequest")
	proto.RegisterType((*SetCheckedResponse)(nil), "internal.SetCheckedResponse")
	proto.RegisterType((*RemoveRequest)(nil), "internal.RemoveRequest")
	proto.RegisterType((*RemoveResponse)(nil), "internal.RemoveResponse")
	proto.RegisterType((*ItemsRequest)(nil), "internal.ItemsRequest")
	proto.RegisterType((*ItemStreamReturns)(nil), "internal.ItemStreamReturns")
	proto.RegisterType((*Item)(nil), "internal.Item")
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
	err := grpc.Invoke(ctx, "/internal.Grpc/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcClient) SetChecked(ctx context.Context, in *SetCheckedRequest, opts ...grpc.CallOption) (*SetCheckedResponse, error) {
	out := new(SetCheckedResponse)
	err := grpc.Invoke(ctx, "/internal.Grpc/SetChecked", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	out := new(RemoveResponse)
	err := grpc.Invoke(ctx, "/internal.Grpc/Remove", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcClient) Items(ctx context.Context, in *ItemsRequest, opts ...grpc.CallOption) (Grpc_ItemsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Grpc_serviceDesc.Streams[0], c.cc, "/internal.Grpc/Items", opts...)
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
		FullMethod: "/internal.Grpc/Add",
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
		FullMethod: "/internal.Grpc/SetChecked",
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
		FullMethod: "/internal.Grpc/Remove",
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
	ServiceName: "internal.Grpc",
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
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0xa5, 0x1f, 0x20, 0x8e, 0x48, 0x64, 0x43, 0xb4, 0x29, 0x26, 0x92, 0x3d, 0x71, 0x22, 0x04,
	0x6e, 0x26, 0x24, 0x12, 0x0f, 0x86, 0xeb, 0xf2, 0x0b, 0x6a, 0x77, 0xa2, 0x8d, 0x74, 0x8b, 0xed,
	0x60, 0x8c, 0x7f, 0xc0, 0xbf, 0x6d, 0xba, 0xdb, 0xd2, 0x16, 0xf4, 0xe0, 0xa9, 0x79, 0x33, 0xef,
	0xcd, 0xee, 0xbc, 0xd7, 0x85, 0xde, 0x2e, 0x4d, 0x28, 0xc9, 0xa6, 0xfa, 0xc3, 0xba, 0x91, 0x22,
	0x4c, 0x55, 0xb0, 0xe5, 0x1c, 0x60, 0x25, 0xa5, 0xc0, 0xf7, 0x3d, 0x66, 0xc4, 0x86, 0xd0, 0xa6,
	0x88, 0xb6, 0xe8, 0x59, 0x63, 0x6b, 0x72, 0x2e, 0x0c, 0xe0, 0xb3, 0x82, 0x43, 0xfb, 0x54, 0x65,
	0x8c, 0x83, 0x1b, 0x11, 0xc6, 0x9a, 0x72, 0x31, 0xef, 0x4f, 0xcb, 0x51, 0xd3, 0x35, 0x61, 0x2c,
	0x74, 0x8f, 0x2f, 0x61, 0xb0, 0x41, 0x7a, 0x7c, 0xc5, 0xf0, 0x0d, 0x0f, 0xc3, 0xfb, 0x60, 0x47,
	0xb2, 0x98, 0x6c, 0x47, 0x92, 0x79, 0x70, 0x16, 0x1a, 0x86, 0x67, 0x8f, 0xad, 0x49, 0x57, 0x94,
	0x90, 0x0f, 0x81, 0xd5, 0xe5, 0xd9, 0x2e, 0x51, 0x19, 0xf2, 0x3b, 0xb8, 0x14, 0x18, 0x27, 0x1f,
	0xf8, 0xc7, 0x40, 0x7e, 0x05, 0xfd, 0x92, 0x50, 0x48, 0xee, 0xa1, 0x97, 0xdf, 0x2a, 0xab, 0xed,
	0x17, 0x29, 0x89, 0x9f, 0x5a, 0xe4, 0x08, 0x03, 0xf2, 0x6a, 0x98, 0xec, 0x15, 0xe9, 0x6b, 0x38,
	0xc2, 0x00, 0x1e, 0xc2, 0x20, 0xd7, 0x6e, 0x28, 0xc5, 0x20, 0xfe, 0xc7, 0xf2, 0xd5, 0x21, 0xf6,
	0xd1, 0x21, 0x94, 0x50, 0xb0, 0xf5, 0x1c, 0x53, 0xd5, 0x80, 0x7f, 0x81, 0x9b, 0x2b, 0x4f, 0xbc,
	0x39, 0x04, 0x61, 0xd7, 0x82, 0xd0, 0x8e, 0xa5, 0x18, 0x10, 0x4a, 0x3d, 0xa5, 0x27, 0x4a, 0x68,
	0xbc, 0x0c, 0xd4, 0x0b, 0x4a, 0xcf, 0x2d, 0x3a, 0x06, 0xd6, 0x5d, 0x6e, 0x37, 0x5c, 0x9e, 0x7f,
	0xdb, 0xe0, 0x3e, 0xa5, 0xbb, 0x90, 0x2d, 0xc0, 0x59, 0x49, 0xc9, 0x86, 0xd5, 0x36, 0xd5, 0x2f,
	0xe1, 0x1f, 0x57, 0xb5, 0x0f, 0xbc, 0xc5, 0xd6, 0x00, 0x55, 0x46, 0x6c, 0x54, 0xb1, 0x4e, 0x82,
	0xf7, 0x6f, 0x7f, 0x6f, 0x16, 0x19, 0xb5, 0xd8, 0x12, 0x3a, 0x26, 0x37, 0x76, 0x53, 0x31, 0x1b,
	0x51, 0xfb, 0xde, 0x69, 0xe3, 0x20, 0x7f, 0x80, 0xb6, 0x0e, 0x99, 0x5d, 0x37, 0xe3, 0x28, 0x53,
	0xf7, 0x47, 0xcd, 0x7a, 0x23, 0x51, 0xde, 0x9a, 0x59, 0xcf, 0x1d, 0xfd, 0x2a, 0x16, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x65, 0x16, 0xcf, 0x67, 0x25, 0x03, 0x00, 0x00,
}