// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: services/healthcheck/proto/health.proto

package gen

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

// HealthCheckServiceClient is the client API for HealthCheckService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthCheckServiceClient interface {
	HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	HealthCheckDB(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type healthCheckServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthCheckServiceClient(cc grpc.ClientConnInterface) HealthCheckServiceClient {
	return &healthCheckServiceClient{cc}
}

func (c *healthCheckServiceClient) HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/healthcheck.HealthCheckService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthCheckServiceClient) HealthCheckDB(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/healthcheck.HealthCheckService/HealthCheckDB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthCheckServiceServer is the server API for HealthCheckService service.
// All implementations must embed UnimplementedHealthCheckServiceServer
// for forward compatibility
type HealthCheckServiceServer interface {
	HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	HealthCheckDB(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	mustEmbedUnimplementedHealthCheckServiceServer()
}

// UnimplementedHealthCheckServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHealthCheckServiceServer struct {
}

func (UnimplementedHealthCheckServiceServer) HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedHealthCheckServiceServer) HealthCheckDB(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheckDB not implemented")
}
func (UnimplementedHealthCheckServiceServer) mustEmbedUnimplementedHealthCheckServiceServer() {}

// UnsafeHealthCheckServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthCheckServiceServer will
// result in compilation errors.
type UnsafeHealthCheckServiceServer interface {
	mustEmbedUnimplementedHealthCheckServiceServer()
}

func RegisterHealthCheckServiceServer(s grpc.ServiceRegistrar, srv HealthCheckServiceServer) {
	s.RegisterService(&HealthCheckService_ServiceDesc, srv)
}

func _HealthCheckService_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthCheckServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/healthcheck.HealthCheckService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthCheckServiceServer).HealthCheck(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthCheckService_HealthCheckDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthCheckServiceServer).HealthCheckDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/healthcheck.HealthCheckService/HealthCheckDB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthCheckServiceServer).HealthCheckDB(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HealthCheckService_ServiceDesc is the grpc.ServiceDesc for HealthCheckService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HealthCheckService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "healthcheck.HealthCheckService",
	HandlerType: (*HealthCheckServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _HealthCheckService_HealthCheck_Handler,
		},
		{
			MethodName: "HealthCheckDB",
			Handler:    _HealthCheckService_HealthCheckDB_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/healthcheck/proto/health.proto",
}
