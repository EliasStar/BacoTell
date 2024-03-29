// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.0
// source: proto/bacotell_plugin.proto

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
	Plugin_Id_FullMethodName                  = "/bacotell.Plugin/Id"
	Plugin_ApplicationCommands_FullMethodName = "/bacotell.Plugin/ApplicationCommands"
	Plugin_MessageComponents_FullMethodName   = "/bacotell.Plugin/MessageComponents"
	Plugin_Modals_FullMethodName              = "/bacotell.Plugin/Modals"
)

// PluginClient is the client API for Plugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PluginClient interface {
	Id(ctx context.Context, in *PluginIdRequest, opts ...grpc.CallOption) (*PluginIdResponse, error)
	ApplicationCommands(ctx context.Context, in *PluginApplicationCommandsRequest, opts ...grpc.CallOption) (*PluginApplicationCommandsResponse, error)
	MessageComponents(ctx context.Context, in *PluginMessageComponentsRequest, opts ...grpc.CallOption) (*PluginMessageComponentsResponse, error)
	Modals(ctx context.Context, in *PluginModalsRequest, opts ...grpc.CallOption) (*PluginModalsResponse, error)
}

type pluginClient struct {
	cc grpc.ClientConnInterface
}

func NewPluginClient(cc grpc.ClientConnInterface) PluginClient {
	return &pluginClient{cc}
}

func (c *pluginClient) Id(ctx context.Context, in *PluginIdRequest, opts ...grpc.CallOption) (*PluginIdResponse, error) {
	out := new(PluginIdResponse)
	err := c.cc.Invoke(ctx, Plugin_Id_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) ApplicationCommands(ctx context.Context, in *PluginApplicationCommandsRequest, opts ...grpc.CallOption) (*PluginApplicationCommandsResponse, error) {
	out := new(PluginApplicationCommandsResponse)
	err := c.cc.Invoke(ctx, Plugin_ApplicationCommands_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) MessageComponents(ctx context.Context, in *PluginMessageComponentsRequest, opts ...grpc.CallOption) (*PluginMessageComponentsResponse, error) {
	out := new(PluginMessageComponentsResponse)
	err := c.cc.Invoke(ctx, Plugin_MessageComponents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Modals(ctx context.Context, in *PluginModalsRequest, opts ...grpc.CallOption) (*PluginModalsResponse, error) {
	out := new(PluginModalsResponse)
	err := c.cc.Invoke(ctx, Plugin_Modals_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginServer is the server API for Plugin service.
// All implementations must embed UnimplementedPluginServer
// for forward compatibility
type PluginServer interface {
	Id(context.Context, *PluginIdRequest) (*PluginIdResponse, error)
	ApplicationCommands(context.Context, *PluginApplicationCommandsRequest) (*PluginApplicationCommandsResponse, error)
	MessageComponents(context.Context, *PluginMessageComponentsRequest) (*PluginMessageComponentsResponse, error)
	Modals(context.Context, *PluginModalsRequest) (*PluginModalsResponse, error)
	mustEmbedUnimplementedPluginServer()
}

// UnimplementedPluginServer must be embedded to have forward compatible implementations.
type UnimplementedPluginServer struct {
}

func (UnimplementedPluginServer) Id(context.Context, *PluginIdRequest) (*PluginIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Id not implemented")
}
func (UnimplementedPluginServer) ApplicationCommands(context.Context, *PluginApplicationCommandsRequest) (*PluginApplicationCommandsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplicationCommands not implemented")
}
func (UnimplementedPluginServer) MessageComponents(context.Context, *PluginMessageComponentsRequest) (*PluginMessageComponentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MessageComponents not implemented")
}
func (UnimplementedPluginServer) Modals(context.Context, *PluginModalsRequest) (*PluginModalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Modals not implemented")
}
func (UnimplementedPluginServer) mustEmbedUnimplementedPluginServer() {}

// UnsafePluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PluginServer will
// result in compilation errors.
type UnsafePluginServer interface {
	mustEmbedUnimplementedPluginServer()
}

func RegisterPluginServer(s grpc.ServiceRegistrar, srv PluginServer) {
	s.RegisterService(&Plugin_ServiceDesc, srv)
}

func _Plugin_Id_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Id(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Plugin_Id_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Id(ctx, req.(*PluginIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_ApplicationCommands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginApplicationCommandsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).ApplicationCommands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Plugin_ApplicationCommands_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).ApplicationCommands(ctx, req.(*PluginApplicationCommandsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_MessageComponents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginMessageComponentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).MessageComponents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Plugin_MessageComponents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).MessageComponents(ctx, req.(*PluginMessageComponentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_Modals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginModalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Modals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Plugin_Modals_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Modals(ctx, req.(*PluginModalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Plugin_ServiceDesc is the grpc.ServiceDesc for Plugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Plugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bacotell.Plugin",
	HandlerType: (*PluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Id",
			Handler:    _Plugin_Id_Handler,
		},
		{
			MethodName: "ApplicationCommands",
			Handler:    _Plugin_ApplicationCommands_Handler,
		},
		{
			MethodName: "MessageComponents",
			Handler:    _Plugin_MessageComponents_Handler,
		},
		{
			MethodName: "Modals",
			Handler:    _Plugin_Modals_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/bacotell_plugin.proto",
}
