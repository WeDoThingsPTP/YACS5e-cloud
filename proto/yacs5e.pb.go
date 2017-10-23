// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yacs5e.proto

/*
Package yacs5e is a generated protocol buffer package.

It is generated from these files:
	yacs5e.proto

It has these top-level messages:
	User
	Empty
*/
package yacs5e

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

type User struct {
	Login    string `protobuf:"bytes,1,opt,name=login" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterType((*Empty)(nil), "Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for YACS5E service

type YACS5EClient interface {
	// ERROR CODES:
	// 100: UNKNOWN ERROR
	// 101: INVALID LOGIN
	// 102: INVALID PASSWORD
	// 103: USER EXISTS
	Registration(ctx context.Context, in *User, opts ...grpc.CallOption) (*Empty, error)
	// ERROR CODES:
	// 110: UNKNOWN ERROR
	// 111: INVALID LOGIN
	// 112: INVALID PASSWORD
	// 113: USER EXISTS
	Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*Empty, error)
}

type yACS5EClient struct {
	cc *grpc.ClientConn
}

func NewYACS5EClient(cc *grpc.ClientConn) YACS5EClient {
	return &yACS5EClient{cc}
}

func (c *yACS5EClient) Registration(ctx context.Context, in *User, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/YACS5e/Registration", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yACS5EClient) Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/YACS5e/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for YACS5E service

type YACS5EServer interface {
	// ERROR CODES:
	// 100: UNKNOWN ERROR
	// 101: INVALID LOGIN
	// 102: INVALID PASSWORD
	// 103: USER EXISTS
	Registration(context.Context, *User) (*Empty, error)
	// ERROR CODES:
	// 110: UNKNOWN ERROR
	// 111: INVALID LOGIN
	// 112: INVALID PASSWORD
	// 113: USER EXISTS
	Login(context.Context, *User) (*Empty, error)
}

func RegisterYACS5EServer(s *grpc.Server, srv YACS5EServer) {
	s.RegisterService(&_YACS5E_serviceDesc, srv)
}

func _YACS5E_Registration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YACS5EServer).Registration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/YACS5e/Registration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YACS5EServer).Registration(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _YACS5E_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YACS5EServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/YACS5e/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YACS5EServer).Login(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _YACS5E_serviceDesc = grpc.ServiceDesc{
	ServiceName: "YACS5e",
	HandlerType: (*YACS5EServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Registration",
			Handler:    _YACS5E_Registration_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _YACS5E_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yacs5e.proto",
}

func init() { proto.RegisterFile("yacs5e.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xa9, 0x4c, 0x4c, 0x2e,
	0x36, 0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xb2, 0xe0, 0x62, 0x09, 0x2d, 0x4e, 0x2d,
	0x12, 0x12, 0xe1, 0x62, 0xcd, 0xc9, 0x4f, 0xcf, 0xcc, 0x93, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c,
	0x82, 0x70, 0x84, 0xa4, 0xb8, 0x38, 0x0a, 0x12, 0x8b, 0x8b, 0xcb, 0xf3, 0x8b, 0x52, 0x24, 0x98,
	0xc0, 0x12, 0x70, 0xbe, 0x12, 0x3b, 0x17, 0xab, 0x6b, 0x6e, 0x41, 0x49, 0xa5, 0x91, 0x3d, 0x17,
	0x5b, 0xa4, 0xa3, 0x73, 0xb0, 0x69, 0xaa, 0x90, 0x2c, 0x17, 0x4f, 0x50, 0x6a, 0x7a, 0x66, 0x71,
	0x49, 0x51, 0x62, 0x49, 0x66, 0x7e, 0x9e, 0x10, 0xab, 0x1e, 0xc8, 0x6c, 0x29, 0x36, 0x3d, 0xb0,
	0x42, 0x21, 0x31, 0x2e, 0x56, 0x1f, 0xb0, 0xb1, 0xa8, 0xe2, 0x49, 0x6c, 0x60, 0xa7, 0x18, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x8c, 0xeb, 0x6c, 0x24, 0x9a, 0x00, 0x00, 0x00,
}
