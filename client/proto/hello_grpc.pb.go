// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0--rc1
// source: hello.proto

package service

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

const (
	Grace_Welcome_FullMethodName = "/Grace/Welcome"
)

// GraceClient is the client API for Grace service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GraceClient interface {
	Welcome(ctx context.Context, in *WelcomeRequest, opts ...grpc.CallOption) (*WelcomeResponse, error)
}

type graceClient struct {
	cc grpc.ClientConnInterface
}

func NewGraceClient(cc grpc.ClientConnInterface) GraceClient {
	return &graceClient{cc}
}

func (c *graceClient) Welcome(ctx context.Context, in *WelcomeRequest, opts ...grpc.CallOption) (*WelcomeResponse, error) {
	out := new(WelcomeResponse)
	err := c.cc.Invoke(ctx, Grace_Welcome_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GraceServer is the server API for Grace service.
// All implementations must embed UnimplementedGraceServer
// for forward compatibility
type GraceServer interface {
	Welcome(context.Context, *WelcomeRequest) (*WelcomeResponse, error)
	mustEmbedUnimplementedGraceServer()
}

// UnimplementedGraceServer must be embedded to have forward compatible implementations.
type UnimplementedGraceServer struct {
}

func (UnimplementedGraceServer) Welcome(context.Context, *WelcomeRequest) (*WelcomeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Welcome not implemented")
}
func (UnimplementedGraceServer) mustEmbedUnimplementedGraceServer() {}

// UnsafeGraceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GraceServer will
// result in compilation errors.
type UnsafeGraceServer interface {
	mustEmbedUnimplementedGraceServer()
}

func RegisterGraceServer(s grpc.ServiceRegistrar, srv GraceServer) {
	s.RegisterService(&Grace_ServiceDesc, srv)
}

func _Grace_Welcome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WelcomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraceServer).Welcome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Grace_Welcome_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraceServer).Welcome(ctx, req.(*WelcomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Grace_ServiceDesc is the grpc.ServiceDesc for Grace service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Grace_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Grace",
	HandlerType: (*GraceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Welcome",
			Handler:    _Grace_Welcome_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}
