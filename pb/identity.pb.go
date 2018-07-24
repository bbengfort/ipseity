// Code generated by protoc-gen-go. DO NOT EDIT.
// source: identity.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	identity.proto

It has these top-level messages:
	IdentityRequest
	IdentityReply
*/
package pb

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

type IdentityRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *IdentityRequest) Reset()                    { *m = IdentityRequest{} }
func (m *IdentityRequest) String() string            { return proto.CompactTextString(m) }
func (*IdentityRequest) ProtoMessage()               {}
func (*IdentityRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IdentityRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type IdentityReply struct {
	Key      string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Identity int64  `protobuf:"varint,2,opt,name=identity" json:"identity,omitempty"`
}

func (m *IdentityReply) Reset()                    { *m = IdentityReply{} }
func (m *IdentityReply) String() string            { return proto.CompactTextString(m) }
func (*IdentityReply) ProtoMessage()               {}
func (*IdentityReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IdentityReply) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *IdentityReply) GetIdentity() int64 {
	if m != nil {
		return m.Identity
	}
	return 0
}

func init() {
	proto.RegisterType((*IdentityRequest)(nil), "pb.IdentityRequest")
	proto.RegisterType((*IdentityReply)(nil), "pb.IdentityReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Identity service

type IdentityClient interface {
	Next(ctx context.Context, in *IdentityRequest, opts ...grpc.CallOption) (*IdentityReply, error)
}

type identityClient struct {
	cc *grpc.ClientConn
}

func NewIdentityClient(cc *grpc.ClientConn) IdentityClient {
	return &identityClient{cc}
}

func (c *identityClient) Next(ctx context.Context, in *IdentityRequest, opts ...grpc.CallOption) (*IdentityReply, error) {
	out := new(IdentityReply)
	err := grpc.Invoke(ctx, "/pb.Identity/Next", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Identity service

type IdentityServer interface {
	Next(context.Context, *IdentityRequest) (*IdentityReply, error)
}

func RegisterIdentityServer(s *grpc.Server, srv IdentityServer) {
	s.RegisterService(&_Identity_serviceDesc, srv)
}

func _Identity_Next_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).Next(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Identity/Next",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).Next(ctx, req.(*IdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Identity_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Identity",
	HandlerType: (*IdentityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Next",
			Handler:    _Identity_Next_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identity.proto",
}

func init() { proto.RegisterFile("identity.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x4c, 0x49, 0xcd,
	0x2b, 0xc9, 0x2c, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x52,
	0xe6, 0xe2, 0xf7, 0x84, 0x8a, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09, 0x70, 0x31,
	0x67, 0xa7, 0x56, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x4a, 0xb6, 0x5c, 0xbc,
	0x08, 0x45, 0x05, 0x39, 0x95, 0x98, 0x4a, 0x84, 0xa4, 0xb8, 0x38, 0x60, 0xa6, 0x4b, 0x30, 0x29,
	0x30, 0x6a, 0x30, 0x07, 0xc1, 0xf9, 0x46, 0x36, 0x5c, 0x1c, 0x30, 0xed, 0x42, 0x06, 0x5c, 0x2c,
	0x7e, 0xa9, 0x15, 0x25, 0x42, 0xc2, 0x7a, 0x05, 0x49, 0x7a, 0x68, 0x36, 0x4b, 0x09, 0xa2, 0x0a,
	0x16, 0xe4, 0x54, 0x2a, 0x31, 0x24, 0xb1, 0x81, 0x1d, 0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x71, 0x5e, 0xe0, 0xd2, 0xbe, 0x00, 0x00, 0x00,
}
