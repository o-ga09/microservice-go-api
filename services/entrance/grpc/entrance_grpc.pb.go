// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: services/entrance/grpc/entrance.proto

package grpc

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

// EntranceServiceClient is the client API for EntranceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EntranceServiceClient interface {
	CheckIn(ctx context.Context, in *CheckInRequest, opts ...grpc.CallOption) (*CheckInResponse, error)
	CheckOut(ctx context.Context, in *CheckOutRequest, opts ...grpc.CallOption) (*CheckOutResponse, error)
}

type entranceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEntranceServiceClient(cc grpc.ClientConnInterface) EntranceServiceClient {
	return &entranceServiceClient{cc}
}

func (c *entranceServiceClient) CheckIn(ctx context.Context, in *CheckInRequest, opts ...grpc.CallOption) (*CheckInResponse, error) {
	out := new(CheckInResponse)
	err := c.cc.Invoke(ctx, "/entrance.EntranceService/CheckIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entranceServiceClient) CheckOut(ctx context.Context, in *CheckOutRequest, opts ...grpc.CallOption) (*CheckOutResponse, error) {
	out := new(CheckOutResponse)
	err := c.cc.Invoke(ctx, "/entrance.EntranceService/CheckOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntranceServiceServer is the server API for EntranceService service.
// All implementations must embed UnimplementedEntranceServiceServer
// for forward compatibility
type EntranceServiceServer interface {
	CheckIn(context.Context, *CheckInRequest) (*CheckInResponse, error)
	CheckOut(context.Context, *CheckOutRequest) (*CheckOutResponse, error)
	mustEmbedUnimplementedEntranceServiceServer()
}

// UnimplementedEntranceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEntranceServiceServer struct {
}

func (UnimplementedEntranceServiceServer) CheckIn(context.Context, *CheckInRequest) (*CheckInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIn not implemented")
}
func (UnimplementedEntranceServiceServer) CheckOut(context.Context, *CheckOutRequest) (*CheckOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckOut not implemented")
}
func (UnimplementedEntranceServiceServer) mustEmbedUnimplementedEntranceServiceServer() {}

// UnsafeEntranceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EntranceServiceServer will
// result in compilation errors.
type UnsafeEntranceServiceServer interface {
	mustEmbedUnimplementedEntranceServiceServer()
}

func RegisterEntranceServiceServer(s grpc.ServiceRegistrar, srv EntranceServiceServer) {
	s.RegisterService(&EntranceService_ServiceDesc, srv)
}

func _EntranceService_CheckIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntranceServiceServer).CheckIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entrance.EntranceService/CheckIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntranceServiceServer).CheckIn(ctx, req.(*CheckInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntranceService_CheckOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntranceServiceServer).CheckOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entrance.EntranceService/CheckOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntranceServiceServer).CheckOut(ctx, req.(*CheckOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EntranceService_ServiceDesc is the grpc.ServiceDesc for EntranceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EntranceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "entrance.EntranceService",
	HandlerType: (*EntranceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckIn",
			Handler:    _EntranceService_CheckIn_Handler,
		},
		{
			MethodName: "CheckOut",
			Handler:    _EntranceService_CheckOut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/entrance/grpc/entrance.proto",
}
