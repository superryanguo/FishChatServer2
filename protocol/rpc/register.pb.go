// Code generated by protoc-gen-go.
// source: register.proto
// DO NOT EDIT!

package rpc

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

type RGPingReq struct {
	UID int64 `protobuf:"varint,1,opt,name=UID" json:"UID,omitempty"`
}

func (m *RGPingReq) Reset()                    { *m = RGPingReq{} }
func (m *RGPingReq) String() string            { return proto.CompactTextString(m) }
func (*RGPingReq) ProtoMessage()               {}
func (*RGPingReq) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type RGPingRes struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,2,opt,name=errStr" json:"errStr,omitempty"`
}

func (m *RGPingRes) Reset()                    { *m = RGPingRes{} }
func (m *RGPingRes) String() string            { return proto.CompactTextString(m) }
func (*RGPingRes) ProtoMessage()               {}
func (*RGPingRes) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

type RGOnlineReq struct {
	UID int64 `protobuf:"varint,1,opt,name=UID" json:"UID,omitempty"`
}

func (m *RGOnlineReq) Reset()                    { *m = RGOnlineReq{} }
func (m *RGOnlineReq) String() string            { return proto.CompactTextString(m) }
func (*RGOnlineReq) ProtoMessage()               {}
func (*RGOnlineReq) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

type RGOnlineRes struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,2,opt,name=errStr" json:"errStr,omitempty"`
	Online  bool   `protobuf:"varint,3,opt,name=online" json:"online,omitempty"`
}

func (m *RGOnlineRes) Reset()                    { *m = RGOnlineRes{} }
func (m *RGOnlineRes) String() string            { return proto.CompactTextString(m) }
func (*RGOnlineRes) ProtoMessage()               {}
func (*RGOnlineRes) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func init() {
	proto.RegisterType((*RGPingReq)(nil), "rpc.RGPingReq")
	proto.RegisterType((*RGPingRes)(nil), "rpc.RGPingRes")
	proto.RegisterType((*RGOnlineReq)(nil), "rpc.RGOnlineReq")
	proto.RegisterType((*RGOnlineRes)(nil), "rpc.RGOnlineRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for RegisterServerRPC service

type RegisterServerRPCClient interface {
	Ping(ctx context.Context, in *RGPingReq, opts ...grpc.CallOption) (*RGPingRes, error)
	Online(ctx context.Context, in *RGOnlineReq, opts ...grpc.CallOption) (*RGOnlineRes, error)
}

type registerServerRPCClient struct {
	cc *grpc.ClientConn
}

func NewRegisterServerRPCClient(cc *grpc.ClientConn) RegisterServerRPCClient {
	return &registerServerRPCClient{cc}
}

func (c *registerServerRPCClient) Ping(ctx context.Context, in *RGPingReq, opts ...grpc.CallOption) (*RGPingRes, error) {
	out := new(RGPingRes)
	err := grpc.Invoke(ctx, "/rpc.RegisterServerRPC/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registerServerRPCClient) Online(ctx context.Context, in *RGOnlineReq, opts ...grpc.CallOption) (*RGOnlineRes, error) {
	out := new(RGOnlineRes)
	err := grpc.Invoke(ctx, "/rpc.RegisterServerRPC/Online", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RegisterServerRPC service

type RegisterServerRPCServer interface {
	Ping(context.Context, *RGPingReq) (*RGPingRes, error)
	Online(context.Context, *RGOnlineReq) (*RGOnlineRes, error)
}

func RegisterRegisterServerRPCServer(s *grpc.Server, srv RegisterServerRPCServer) {
	s.RegisterService(&_RegisterServerRPC_serviceDesc, srv)
}

func _RegisterServerRPC_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RGPingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServerRPCServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.RegisterServerRPC/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServerRPCServer).Ping(ctx, req.(*RGPingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegisterServerRPC_Online_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RGOnlineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServerRPCServer).Online(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.RegisterServerRPC/Online",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServerRPCServer).Online(ctx, req.(*RGOnlineReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _RegisterServerRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.RegisterServerRPC",
	HandlerType: (*RegisterServerRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _RegisterServerRPC_Ping_Handler,
		},
		{
			MethodName: "Online",
			Handler:    _RegisterServerRPC_Online_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor3,
}

func init() { proto.RegisterFile("register.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4a, 0x4d, 0xcf,
	0x2c, 0x2e, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48, 0x56,
	0x92, 0xe5, 0xe2, 0x0c, 0x72, 0x0f, 0xc8, 0xcc, 0x4b, 0x0f, 0x4a, 0x2d, 0x14, 0x12, 0xe0, 0x62,
	0x0e, 0xf5, 0x74, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x02, 0x31, 0x95, 0x6c, 0x11, 0xd2,
	0xc5, 0x42, 0x12, 0x5c, 0xec, 0xa9, 0x45, 0x45, 0xce, 0xf9, 0x29, 0xa9, 0x60, 0x25, 0xbc, 0x41,
	0x30, 0xae, 0x90, 0x18, 0x17, 0x5b, 0x6a, 0x51, 0x51, 0x70, 0x49, 0x91, 0x04, 0x93, 0x02, 0xa3,
	0x06, 0x67, 0x10, 0x94, 0xa7, 0x24, 0xcf, 0xc5, 0x1d, 0xe4, 0xee, 0x9f, 0x97, 0x93, 0x99, 0x97,
	0x8a, 0xdd, 0xfc, 0x70, 0x64, 0x05, 0x64, 0xd8, 0x00, 0x12, 0xcf, 0x07, 0x6b, 0x97, 0x60, 0x56,
	0x60, 0xd4, 0xe0, 0x08, 0x82, 0xf2, 0x8c, 0x72, 0xb9, 0x04, 0x83, 0xa0, 0xde, 0x0d, 0x4e, 0x2d,
	0x2a, 0x4b, 0x2d, 0x0a, 0x0a, 0x70, 0x16, 0xd2, 0xe0, 0x62, 0x01, 0xf9, 0x45, 0x88, 0x4f, 0xaf,
	0xa8, 0x20, 0x59, 0x0f, 0xee, 0x6f, 0x29, 0x54, 0x7e, 0xb1, 0x12, 0x83, 0x90, 0x1e, 0x17, 0x1b,
	0xc4, 0x55, 0x42, 0x02, 0x50, 0x39, 0xb8, 0x2f, 0xa4, 0xd0, 0x45, 0x8a, 0x95, 0x18, 0x92, 0xd8,
	0xc0, 0x41, 0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x02, 0x9f, 0xff, 0x64, 0x01, 0x00,
	0x00,
}
