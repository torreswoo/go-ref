// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateway.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	gateway.proto

It has these top-level messages:
	EmptyRequest
	CountResponse
	StringResponse
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc1 "google.golang.org/grpc"
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

type EmptyRequest struct {
}

func (m *EmptyRequest) Reset()                    { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string            { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()               {}
func (*EmptyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CountResponse struct {
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *CountResponse) Reset()                    { *m = CountResponse{} }
func (m *CountResponse) String() string            { return proto.CompactTextString(m) }
func (*CountResponse) ProtoMessage()               {}
func (*CountResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CountResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StringResponse struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *StringResponse) Reset()                    { *m = StringResponse{} }
func (m *StringResponse) String() string            { return proto.CompactTextString(m) }
func (*StringResponse) ProtoMessage()               {}
func (*StringResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StringResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*EmptyRequest)(nil), "grpc.EmptyRequest")
	proto.RegisterType((*CountResponse)(nil), "grpc.CountResponse")
	proto.RegisterType((*StringResponse)(nil), "grpc.StringResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for Gateway service

type GatewayClient interface {
	SubscribeTotalAccessCount(ctx context.Context, in *EmptyRequest, opts ...grpc1.CallOption) (Gateway_SubscribeTotalAccessCountClient, error)
	SubscribeCurrentUserCount(ctx context.Context, in *EmptyRequest, opts ...grpc1.CallOption) (Gateway_SubscribeCurrentUserCountClient, error)
	SubscribeCurrentNodeCount(ctx context.Context, in *EmptyRequest, opts ...grpc1.CallOption) (Gateway_SubscribeCurrentNodeCountClient, error)
	SubscribeCurrentMasterIdentifier(ctx context.Context, in *EmptyRequest, opts ...grpc1.CallOption) (Gateway_SubscribeCurrentMasterIdentifierClient, error)
}

type gatewayClient struct {
	cc *grpc1.ClientConn
}

func NewGatewayClient(cc *grpc1.ClientConn) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) SubscribeTotalAccessCount(ctx context.Context, in *EmptyRequest, opts ...grpc1.CallOption) (Gateway_SubscribeTotalAccessCountClient, error) {
	stream, err := grpc1.NewClientStream(ctx, &_Gateway_serviceDesc.Streams[0], c.cc, "/grpc.Gateway/SubscribeTotalAccessCount", opts...)
	if err != nil {
		return nil, err
	}
	x := &gatewaySubscribeTotalAccessCountClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Gateway_SubscribeTotalAccessCountClient interface {
	Recv() (*CountResponse, error)
	grpc1.ClientStream
}

type gatewaySubscribeTotalAccessCountClient struct {
	grpc1.ClientStream
}

func (x *gatewaySubscribeTotalAccessCountClient) Recv() (*CountResponse, error) {
	m := new(CountResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gatewayClient) SubscribeCurrentUserCount(ctx context.Context, in *EmptyRequest, opts ...grpc1.CallOption) (Gateway_SubscribeCurrentUserCountClient, error) {
	stream, err := grpc1.NewClientStream(ctx, &_Gateway_serviceDesc.Streams[1], c.cc, "/grpc.Gateway/SubscribeCurrentUserCount", opts...)
	if err != nil {
		return nil, err
	}
	x := &gatewaySubscribeCurrentUserCountClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Gateway_SubscribeCurrentUserCountClient interface {
	Recv() (*CountResponse, error)
	grpc1.ClientStream
}

type gatewaySubscribeCurrentUserCountClient struct {
	grpc1.ClientStream
}

func (x *gatewaySubscribeCurrentUserCountClient) Recv() (*CountResponse, error) {
	m := new(CountResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gatewayClient) SubscribeCurrentNodeCount(ctx context.Context, in *EmptyRequest, opts ...grpc1.CallOption) (Gateway_SubscribeCurrentNodeCountClient, error) {
	stream, err := grpc1.NewClientStream(ctx, &_Gateway_serviceDesc.Streams[2], c.cc, "/grpc.Gateway/SubscribeCurrentNodeCount", opts...)
	if err != nil {
		return nil, err
	}
	x := &gatewaySubscribeCurrentNodeCountClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Gateway_SubscribeCurrentNodeCountClient interface {
	Recv() (*CountResponse, error)
	grpc1.ClientStream
}

type gatewaySubscribeCurrentNodeCountClient struct {
	grpc1.ClientStream
}

func (x *gatewaySubscribeCurrentNodeCountClient) Recv() (*CountResponse, error) {
	m := new(CountResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gatewayClient) SubscribeCurrentMasterIdentifier(ctx context.Context, in *EmptyRequest, opts ...grpc1.CallOption) (Gateway_SubscribeCurrentMasterIdentifierClient, error) {
	stream, err := grpc1.NewClientStream(ctx, &_Gateway_serviceDesc.Streams[3], c.cc, "/grpc.Gateway/SubscribeCurrentMasterIdentifier", opts...)
	if err != nil {
		return nil, err
	}
	x := &gatewaySubscribeCurrentMasterIdentifierClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Gateway_SubscribeCurrentMasterIdentifierClient interface {
	Recv() (*StringResponse, error)
	grpc1.ClientStream
}

type gatewaySubscribeCurrentMasterIdentifierClient struct {
	grpc1.ClientStream
}

func (x *gatewaySubscribeCurrentMasterIdentifierClient) Recv() (*StringResponse, error) {
	m := new(StringResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Gateway service

type GatewayServer interface {
	SubscribeTotalAccessCount(*EmptyRequest, Gateway_SubscribeTotalAccessCountServer) error
	SubscribeCurrentUserCount(*EmptyRequest, Gateway_SubscribeCurrentUserCountServer) error
	SubscribeCurrentNodeCount(*EmptyRequest, Gateway_SubscribeCurrentNodeCountServer) error
	SubscribeCurrentMasterIdentifier(*EmptyRequest, Gateway_SubscribeCurrentMasterIdentifierServer) error
}

func RegisterGatewayServer(s *grpc1.Server, srv GatewayServer) {
	s.RegisterService(&_Gateway_serviceDesc, srv)
}

func _Gateway_SubscribeTotalAccessCount_Handler(srv interface{}, stream grpc1.ServerStream) error {
	m := new(EmptyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GatewayServer).SubscribeTotalAccessCount(m, &gatewaySubscribeTotalAccessCountServer{stream})
}

type Gateway_SubscribeTotalAccessCountServer interface {
	Send(*CountResponse) error
	grpc1.ServerStream
}

type gatewaySubscribeTotalAccessCountServer struct {
	grpc1.ServerStream
}

func (x *gatewaySubscribeTotalAccessCountServer) Send(m *CountResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Gateway_SubscribeCurrentUserCount_Handler(srv interface{}, stream grpc1.ServerStream) error {
	m := new(EmptyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GatewayServer).SubscribeCurrentUserCount(m, &gatewaySubscribeCurrentUserCountServer{stream})
}

type Gateway_SubscribeCurrentUserCountServer interface {
	Send(*CountResponse) error
	grpc1.ServerStream
}

type gatewaySubscribeCurrentUserCountServer struct {
	grpc1.ServerStream
}

func (x *gatewaySubscribeCurrentUserCountServer) Send(m *CountResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Gateway_SubscribeCurrentNodeCount_Handler(srv interface{}, stream grpc1.ServerStream) error {
	m := new(EmptyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GatewayServer).SubscribeCurrentNodeCount(m, &gatewaySubscribeCurrentNodeCountServer{stream})
}

type Gateway_SubscribeCurrentNodeCountServer interface {
	Send(*CountResponse) error
	grpc1.ServerStream
}

type gatewaySubscribeCurrentNodeCountServer struct {
	grpc1.ServerStream
}

func (x *gatewaySubscribeCurrentNodeCountServer) Send(m *CountResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Gateway_SubscribeCurrentMasterIdentifier_Handler(srv interface{}, stream grpc1.ServerStream) error {
	m := new(EmptyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GatewayServer).SubscribeCurrentMasterIdentifier(m, &gatewaySubscribeCurrentMasterIdentifierServer{stream})
}

type Gateway_SubscribeCurrentMasterIdentifierServer interface {
	Send(*StringResponse) error
	grpc1.ServerStream
}

type gatewaySubscribeCurrentMasterIdentifierServer struct {
	grpc1.ServerStream
}

func (x *gatewaySubscribeCurrentMasterIdentifierServer) Send(m *StringResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Gateway_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods:     []grpc1.MethodDesc{},
	Streams: []grpc1.StreamDesc{
		{
			StreamName:    "SubscribeTotalAccessCount",
			Handler:       _Gateway_SubscribeTotalAccessCount_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeCurrentUserCount",
			Handler:       _Gateway_SubscribeCurrentUserCount_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeCurrentNodeCount",
			Handler:       _Gateway_SubscribeCurrentNodeCount_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeCurrentMasterIdentifier",
			Handler:       _Gateway_SubscribeCurrentMasterIdentifier_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gateway.proto",
}

func init() { proto.RegisterFile("gateway.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4f, 0x2c, 0x49,
	0x2d, 0x4f, 0xac, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0x2f, 0x2a, 0x48, 0x56,
	0xe2, 0xe3, 0xe2, 0x71, 0xcd, 0x2d, 0x28, 0xa9, 0x0c, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x51,
	0x52, 0xe5, 0xe2, 0x75, 0xce, 0x2f, 0xcd, 0x2b, 0x09, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e,
	0x15, 0x12, 0xe1, 0x62, 0x4d, 0x06, 0x09, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0x41, 0x38,
	0x4a, 0x6a, 0x5c, 0x7c, 0xc1, 0x25, 0x45, 0x99, 0x79, 0xe9, 0xc8, 0xea, 0xca, 0x12, 0x73, 0x4a,
	0x53, 0xc1, 0xea, 0x38, 0x83, 0x20, 0x1c, 0xa3, 0x9d, 0x4c, 0x5c, 0xec, 0xee, 0x10, 0x6b, 0x85,
	0x3c, 0xb8, 0x24, 0x83, 0x4b, 0x93, 0x8a, 0x93, 0x8b, 0x32, 0x93, 0x52, 0x43, 0xf2, 0x4b, 0x12,
	0x73, 0x1c, 0x93, 0x93, 0x53, 0x8b, 0x8b, 0xc1, 0xd6, 0x09, 0x09, 0xe9, 0x81, 0x9c, 0xa3, 0x87,
	0xec, 0x16, 0x29, 0x61, 0x88, 0x18, 0x8a, 0x7b, 0x94, 0x18, 0x0c, 0x18, 0x51, 0x4c, 0x72, 0x2e,
	0x2d, 0x2a, 0x4a, 0xcd, 0x2b, 0x09, 0x2d, 0x4e, 0x2d, 0xa2, 0x8e, 0x49, 0x7e, 0xf9, 0x29, 0xa9,
	0x64, 0x98, 0x14, 0xc0, 0xa5, 0x80, 0x6e, 0x92, 0x6f, 0x62, 0x71, 0x49, 0x6a, 0x91, 0x67, 0x4a,
	0x6a, 0x5e, 0x49, 0x66, 0x5a, 0x66, 0x6a, 0x11, 0x56, 0x03, 0x45, 0x20, 0x62, 0xa8, 0xa1, 0x09,
	0x32, 0x31, 0x89, 0x0d, 0x1c, 0x4f, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x59, 0x45,
	0xef, 0xb8, 0x01, 0x00, 0x00,
}
