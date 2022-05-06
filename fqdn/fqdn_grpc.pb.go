// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: fqdn/fqdn.proto

package fqdn

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

// FQDNResolverClient is the client API for FQDNResolver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FQDNResolverClient interface {
	ResolveThis(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*ResolveReply, error)
}

type fQDNResolverClient struct {
	cc grpc.ClientConnInterface
}

func NewFQDNResolverClient(cc grpc.ClientConnInterface) FQDNResolverClient {
	return &fQDNResolverClient{cc}
}

func (c *fQDNResolverClient) ResolveThis(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*ResolveReply, error) {
	out := new(ResolveReply)
	err := c.cc.Invoke(ctx, "/fqdn.FQDNResolver/ResolveThis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FQDNResolverServer is the server API for FQDNResolver service.
// All implementations must embed UnimplementedFQDNResolverServer
// for forward compatibility
type FQDNResolverServer interface {
	ResolveThis(context.Context, *ResolveRequest) (*ResolveReply, error)
	mustEmbedUnimplementedFQDNResolverServer()
}

// UnimplementedFQDNResolverServer must be embedded to have forward compatible implementations.
type UnimplementedFQDNResolverServer struct {
}

func (UnimplementedFQDNResolverServer) ResolveThis(context.Context, *ResolveRequest) (*ResolveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveThis not implemented")
}
func (UnimplementedFQDNResolverServer) mustEmbedUnimplementedFQDNResolverServer() {}

// UnsafeFQDNResolverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FQDNResolverServer will
// result in compilation errors.
type UnsafeFQDNResolverServer interface {
	mustEmbedUnimplementedFQDNResolverServer()
}

func RegisterFQDNResolverServer(s grpc.ServiceRegistrar, srv FQDNResolverServer) {
	s.RegisterService(&FQDNResolver_ServiceDesc, srv)
}

func _FQDNResolver_ResolveThis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FQDNResolverServer).ResolveThis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fqdn.FQDNResolver/ResolveThis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FQDNResolverServer).ResolveThis(ctx, req.(*ResolveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FQDNResolver_ServiceDesc is the grpc.ServiceDesc for FQDNResolver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FQDNResolver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fqdn.FQDNResolver",
	HandlerType: (*FQDNResolverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ResolveThis",
			Handler:    _FQDNResolver_ResolveThis_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fqdn/fqdn.proto",
}
