// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.0
// source: proto/bacotell_component.proto

package bacotellpb

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
	Component_CustomId_FullMethodName = "/bacotell.Component/CustomId"
	Component_Handle_FullMethodName   = "/bacotell.Component/Handle"
)

// ComponentClient is the client API for Component service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComponentClient interface {
	CustomId(ctx context.Context, in *CustomIdRequest, opts ...grpc.CallOption) (*CustomIdResponse, error)
	Handle(ctx context.Context, in *HandleRequest, opts ...grpc.CallOption) (*HandleResponse, error)
}

type componentClient struct {
	cc grpc.ClientConnInterface
}

func NewComponentClient(cc grpc.ClientConnInterface) ComponentClient {
	return &componentClient{cc}
}

func (c *componentClient) CustomId(ctx context.Context, in *CustomIdRequest, opts ...grpc.CallOption) (*CustomIdResponse, error) {
	out := new(CustomIdResponse)
	err := c.cc.Invoke(ctx, Component_CustomId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *componentClient) Handle(ctx context.Context, in *HandleRequest, opts ...grpc.CallOption) (*HandleResponse, error) {
	out := new(HandleResponse)
	err := c.cc.Invoke(ctx, Component_Handle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComponentServer is the server API for Component service.
// All implementations must embed UnimplementedComponentServer
// for forward compatibility
type ComponentServer interface {
	CustomId(context.Context, *CustomIdRequest) (*CustomIdResponse, error)
	Handle(context.Context, *HandleRequest) (*HandleResponse, error)
	mustEmbedUnimplementedComponentServer()
}

// UnimplementedComponentServer must be embedded to have forward compatible implementations.
type UnimplementedComponentServer struct {
}

func (UnimplementedComponentServer) CustomId(context.Context, *CustomIdRequest) (*CustomIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CustomId not implemented")
}
func (UnimplementedComponentServer) Handle(context.Context, *HandleRequest) (*HandleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}
func (UnimplementedComponentServer) mustEmbedUnimplementedComponentServer() {}

// UnsafeComponentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComponentServer will
// result in compilation errors.
type UnsafeComponentServer interface {
	mustEmbedUnimplementedComponentServer()
}

func RegisterComponentServer(s grpc.ServiceRegistrar, srv ComponentServer) {
	s.RegisterService(&Component_ServiceDesc, srv)
}

func _Component_CustomId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComponentServer).CustomId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Component_CustomId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComponentServer).CustomId(ctx, req.(*CustomIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Component_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComponentServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Component_Handle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComponentServer).Handle(ctx, req.(*HandleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Component_ServiceDesc is the grpc.ServiceDesc for Component service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Component_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bacotell.Component",
	HandlerType: (*ComponentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CustomId",
			Handler:    _Component_CustomId_Handler,
		},
		{
			MethodName: "Handle",
			Handler:    _Component_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/bacotell_component.proto",
}

const ()

// HandleProxyClient is the client API for HandleProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HandleProxyClient interface {
}

type handleProxyClient struct {
	cc grpc.ClientConnInterface
}

func NewHandleProxyClient(cc grpc.ClientConnInterface) HandleProxyClient {
	return &handleProxyClient{cc}
}

// HandleProxyServer is the server API for HandleProxy service.
// All implementations must embed UnimplementedHandleProxyServer
// for forward compatibility
type HandleProxyServer interface {
	mustEmbedUnimplementedHandleProxyServer()
}

// UnimplementedHandleProxyServer must be embedded to have forward compatible implementations.
type UnimplementedHandleProxyServer struct {
}

func (UnimplementedHandleProxyServer) mustEmbedUnimplementedHandleProxyServer() {}

// UnsafeHandleProxyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HandleProxyServer will
// result in compilation errors.
type UnsafeHandleProxyServer interface {
	mustEmbedUnimplementedHandleProxyServer()
}

func RegisterHandleProxyServer(s grpc.ServiceRegistrar, srv HandleProxyServer) {
	s.RegisterService(&HandleProxy_ServiceDesc, srv)
}

// HandleProxy_ServiceDesc is the grpc.ServiceDesc for HandleProxy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HandleProxy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bacotell.HandleProxy",
	HandlerType: (*HandleProxyServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "proto/bacotell_component.proto",
}