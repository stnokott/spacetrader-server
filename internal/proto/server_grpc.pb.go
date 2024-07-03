// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.20.1
// source: server.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	SpaceTradersService_Ping_FullMethodName            = "/proto.SpaceTradersService/Ping"
	SpaceTradersService_GetServerStatus_FullMethodName = "/proto.SpaceTradersService/GetServerStatus"
)

// SpaceTradersServiceClient is the client API for SpaceTradersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpaceTradersServiceClient interface {
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetServerStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ServerStatusReply, error)
}

type spaceTradersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSpaceTradersServiceClient(cc grpc.ClientConnInterface) SpaceTradersServiceClient {
	return &spaceTradersServiceClient{cc}
}

func (c *spaceTradersServiceClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SpaceTradersService_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceTradersServiceClient) GetServerStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ServerStatusReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ServerStatusReply)
	err := c.cc.Invoke(ctx, SpaceTradersService_GetServerStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpaceTradersServiceServer is the server API for SpaceTradersService service.
// All implementations must embed UnimplementedSpaceTradersServiceServer
// for forward compatibility
type SpaceTradersServiceServer interface {
	Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	GetServerStatus(context.Context, *emptypb.Empty) (*ServerStatusReply, error)
	mustEmbedUnimplementedSpaceTradersServiceServer()
}

// UnimplementedSpaceTradersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSpaceTradersServiceServer struct {
}

func (UnimplementedSpaceTradersServiceServer) Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedSpaceTradersServiceServer) GetServerStatus(context.Context, *emptypb.Empty) (*ServerStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerStatus not implemented")
}
func (UnimplementedSpaceTradersServiceServer) mustEmbedUnimplementedSpaceTradersServiceServer() {}

// UnsafeSpaceTradersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpaceTradersServiceServer will
// result in compilation errors.
type UnsafeSpaceTradersServiceServer interface {
	mustEmbedUnimplementedSpaceTradersServiceServer()
}

func RegisterSpaceTradersServiceServer(s grpc.ServiceRegistrar, srv SpaceTradersServiceServer) {
	s.RegisterService(&SpaceTradersService_ServiceDesc, srv)
}

func _SpaceTradersService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceTradersServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceTradersService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceTradersServiceServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceTradersService_GetServerStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceTradersServiceServer).GetServerStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceTradersService_GetServerStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceTradersServiceServer).GetServerStatus(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// SpaceTradersService_ServiceDesc is the grpc.ServiceDesc for SpaceTradersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpaceTradersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SpaceTradersService",
	HandlerType: (*SpaceTradersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _SpaceTradersService_Ping_Handler,
		},
		{
			MethodName: "GetServerStatus",
			Handler:    _SpaceTradersService_GetServerStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
