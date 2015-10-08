// Code generated by protoc-gen-go.
// source: auth.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	auth.proto

It has these top-level messages:
	Auth
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Auth_CertificateType int32

const (
	Auth_UUID     Auth_CertificateType = 0
	Auth_PLAIN    Auth_CertificateType = 1
	Auth_TOKEN    Auth_CertificateType = 2
	Auth_FACEBOOK Auth_CertificateType = 3
)

var Auth_CertificateType_name = map[int32]string{
	0: "UUID",
	1: "PLAIN",
	2: "TOKEN",
	3: "FACEBOOK",
}
var Auth_CertificateType_value = map[string]int32{
	"UUID":     0,
	"PLAIN":    1,
	"TOKEN":    2,
	"FACEBOOK": 3,
}

func (x Auth_CertificateType) String() string {
	return proto1.EnumName(Auth_CertificateType_name, int32(x))
}

type Auth struct {
}

func (m *Auth) Reset()         { *m = Auth{} }
func (m *Auth) String() string { return proto1.CompactTextString(m) }
func (*Auth) ProtoMessage()    {}

type Auth_Certificate struct {
	Type  Auth_CertificateType `protobuf:"varint,1,opt,name=Type,enum=proto.Auth_CertificateType" json:"Type,omitempty"`
	Proof []byte               `protobuf:"bytes,2,opt,name=Proof,proto3" json:"Proof,omitempty"`
}

func (m *Auth_Certificate) Reset()         { *m = Auth_Certificate{} }
func (m *Auth_Certificate) String() string { return proto1.CompactTextString(m) }
func (*Auth_Certificate) ProtoMessage()    {}

type Auth_Result struct {
	OK     bool   `protobuf:"varint,1,opt,name=OK" json:"OK,omitempty"`
	UserId int32  `protobuf:"varint,2,opt,name=UserId" json:"UserId,omitempty"`
	Body   []byte `protobuf:"bytes,3,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (m *Auth_Result) Reset()         { *m = Auth_Result{} }
func (m *Auth_Result) String() string { return proto1.CompactTextString(m) }
func (*Auth_Result) ProtoMessage()    {}

func init() {
	proto1.RegisterEnum("proto.Auth_CertificateType", Auth_CertificateType_name, Auth_CertificateType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for AuthService service

type AuthServiceClient interface {
	Auth(ctx context.Context, in *Auth_Certificate, opts ...grpc.CallOption) (*Auth_Result, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Auth(ctx context.Context, in *Auth_Certificate, opts ...grpc.CallOption) (*Auth_Result, error) {
	out := new(Auth_Result)
	err := grpc.Invoke(ctx, "/proto.AuthService/Auth", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthService service

type AuthServiceServer interface {
	Auth(context.Context, *Auth_Certificate) (*Auth_Result, error)
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Auth_Certificate)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AuthServiceServer).Auth(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _AuthService_Auth_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
