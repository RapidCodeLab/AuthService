// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: proto/auth.proto

package auth_grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthGPRCClient is the client API for AuthGPRC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthGPRCClient interface {
	GetPublicKey(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
}

type authGPRCClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthGPRCClient(cc grpc.ClientConnInterface) AuthGPRCClient {
	return &authGPRCClient{cc}
}

func (c *authGPRCClient) GetPublicKey(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/auth_grpc.AuthGPRC/GetPublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthGPRCServer is the server API for AuthGPRC service.
// All implementations must embed UnimplementedAuthGPRCServer
// for forward compatibility
type AuthGPRCServer interface {
	GetPublicKey(context.Context, *Empty) (*Response, error)
	mustEmbedUnimplementedAuthGPRCServer()
}

// UnimplementedAuthGPRCServer must be embedded to have forward compatible implementations.
type UnimplementedAuthGPRCServer struct {
}

func (UnimplementedAuthGPRCServer) GetPublicKey(context.Context, *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicKey not implemented")
}
func (UnimplementedAuthGPRCServer) mustEmbedUnimplementedAuthGPRCServer() {}

// UnsafeAuthGPRCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthGPRCServer will
// result in compilation errors.
type UnsafeAuthGPRCServer interface {
	mustEmbedUnimplementedAuthGPRCServer()
}

func RegisterAuthGPRCServer(s grpc.ServiceRegistrar, srv AuthGPRCServer) {
	s.RegisterService(&AuthGPRC_ServiceDesc, srv)
}

func _AuthGPRC_GetPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthGPRCServer).GetPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_grpc.AuthGPRC/GetPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthGPRCServer).GetPublicKey(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthGPRC_ServiceDesc is the grpc.ServiceDesc for AuthGPRC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthGPRC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth_grpc.AuthGPRC",
	HandlerType: (*AuthGPRCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPublicKey",
			Handler:    _AuthGPRC_GetPublicKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/auth.proto",
}
