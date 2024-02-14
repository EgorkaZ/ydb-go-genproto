// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: ydb_monitoring_v1.proto

package Ydb_Monitoring_V1

import (
	context "context"
	Ydb_Monitoring "github.com/EgorkaZ/ydb-go-genproto/protos/Ydb_Monitoring"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MonitoringService_SelfCheck_FullMethodName = "/Ydb.Monitoring.V1.MonitoringService/SelfCheck"
	MonitoringService_NodeCheck_FullMethodName = "/Ydb.Monitoring.V1.MonitoringService/NodeCheck"
)

// MonitoringServiceClient is the client API for MonitoringService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MonitoringServiceClient interface {
	// Gets the health status of the database.
	SelfCheck(ctx context.Context, in *Ydb_Monitoring.SelfCheckRequest, opts ...grpc.CallOption) (*Ydb_Monitoring.SelfCheckResponse, error)
	// Checks current node health
	NodeCheck(ctx context.Context, in *Ydb_Monitoring.NodeCheckRequest, opts ...grpc.CallOption) (*Ydb_Monitoring.NodeCheckResponse, error)
}

type monitoringServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMonitoringServiceClient(cc grpc.ClientConnInterface) MonitoringServiceClient {
	return &monitoringServiceClient{cc}
}

func (c *monitoringServiceClient) SelfCheck(ctx context.Context, in *Ydb_Monitoring.SelfCheckRequest, opts ...grpc.CallOption) (*Ydb_Monitoring.SelfCheckResponse, error) {
	out := new(Ydb_Monitoring.SelfCheckResponse)
	err := c.cc.Invoke(ctx, MonitoringService_SelfCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringServiceClient) NodeCheck(ctx context.Context, in *Ydb_Monitoring.NodeCheckRequest, opts ...grpc.CallOption) (*Ydb_Monitoring.NodeCheckResponse, error) {
	out := new(Ydb_Monitoring.NodeCheckResponse)
	err := c.cc.Invoke(ctx, MonitoringService_NodeCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonitoringServiceServer is the server API for MonitoringService service.
// All implementations must embed UnimplementedMonitoringServiceServer
// for forward compatibility
type MonitoringServiceServer interface {
	// Gets the health status of the database.
	SelfCheck(context.Context, *Ydb_Monitoring.SelfCheckRequest) (*Ydb_Monitoring.SelfCheckResponse, error)
	// Checks current node health
	NodeCheck(context.Context, *Ydb_Monitoring.NodeCheckRequest) (*Ydb_Monitoring.NodeCheckResponse, error)
	mustEmbedUnimplementedMonitoringServiceServer()
}

// UnimplementedMonitoringServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMonitoringServiceServer struct {
}

func (UnimplementedMonitoringServiceServer) SelfCheck(context.Context, *Ydb_Monitoring.SelfCheckRequest) (*Ydb_Monitoring.SelfCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelfCheck not implemented")
}
func (UnimplementedMonitoringServiceServer) NodeCheck(context.Context, *Ydb_Monitoring.NodeCheckRequest) (*Ydb_Monitoring.NodeCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeCheck not implemented")
}
func (UnimplementedMonitoringServiceServer) mustEmbedUnimplementedMonitoringServiceServer() {}

// UnsafeMonitoringServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MonitoringServiceServer will
// result in compilation errors.
type UnsafeMonitoringServiceServer interface {
	mustEmbedUnimplementedMonitoringServiceServer()
}

func RegisterMonitoringServiceServer(s grpc.ServiceRegistrar, srv MonitoringServiceServer) {
	s.RegisterService(&MonitoringService_ServiceDesc, srv)
}

func _MonitoringService_SelfCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Monitoring.SelfCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServiceServer).SelfCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MonitoringService_SelfCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServiceServer).SelfCheck(ctx, req.(*Ydb_Monitoring.SelfCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitoringService_NodeCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Monitoring.NodeCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServiceServer).NodeCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MonitoringService_NodeCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServiceServer).NodeCheck(ctx, req.(*Ydb_Monitoring.NodeCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MonitoringService_ServiceDesc is the grpc.ServiceDesc for MonitoringService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MonitoringService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Ydb.Monitoring.V1.MonitoringService",
	HandlerType: (*MonitoringServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SelfCheck",
			Handler:    _MonitoringService_SelfCheck_Handler,
		},
		{
			MethodName: "NodeCheck",
			Handler:    _MonitoringService_NodeCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ydb_monitoring_v1.proto",
}
