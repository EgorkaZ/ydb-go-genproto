// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: draft/ydb_dynamic_config_v1.proto

package Ydb_DynamicConfig_V1

import (
	context "context"
	Ydb_DynamicConfig "github.com/EgorkaZ/ydb-go-genproto/draft/protos/Ydb_DynamicConfig"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DynamicConfigService_SetConfig_FullMethodName            = "/Ydb.DynamicConfig.V1.DynamicConfigService/SetConfig"
	DynamicConfigService_ReplaceConfig_FullMethodName        = "/Ydb.DynamicConfig.V1.DynamicConfigService/ReplaceConfig"
	DynamicConfigService_GetMetadata_FullMethodName          = "/Ydb.DynamicConfig.V1.DynamicConfigService/GetMetadata"
	DynamicConfigService_GetConfig_FullMethodName            = "/Ydb.DynamicConfig.V1.DynamicConfigService/GetConfig"
	DynamicConfigService_DropConfig_FullMethodName           = "/Ydb.DynamicConfig.V1.DynamicConfigService/DropConfig"
	DynamicConfigService_AddVolatileConfig_FullMethodName    = "/Ydb.DynamicConfig.V1.DynamicConfigService/AddVolatileConfig"
	DynamicConfigService_RemoveVolatileConfig_FullMethodName = "/Ydb.DynamicConfig.V1.DynamicConfigService/RemoveVolatileConfig"
	DynamicConfigService_GetNodeLabels_FullMethodName        = "/Ydb.DynamicConfig.V1.DynamicConfigService/GetNodeLabels"
	DynamicConfigService_ResolveConfig_FullMethodName        = "/Ydb.DynamicConfig.V1.DynamicConfigService/ResolveConfig"
	DynamicConfigService_ResolveAllConfig_FullMethodName     = "/Ydb.DynamicConfig.V1.DynamicConfigService/ResolveAllConfig"
)

// DynamicConfigServiceClient is the client API for DynamicConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DynamicConfigServiceClient interface {
	// Apply new version of config.
	// Overwrites previous version.
	// This call is idempotent:
	//   - Two subsequent identical calls will return success,
	//   - After applying next version another identical call will result in an error.
	//
	// The field version within the YAML in the request must strictly be set to the current version increment by one.
	// The field cluster within the YAML should be identical to the one configured on the node used by the console tablet.
	SetConfig(ctx context.Context, in *Ydb_DynamicConfig.SetConfigRequest, opts ...grpc.CallOption) (*Ydb_DynamicConfig.SetConfigResponse, error)
	ReplaceConfig(ctx context.Context, in *Ydb_DynamicConfig.ReplaceConfigRequest, opts ...grpc.CallOption) (*Ydb_DynamicConfig.ReplaceConfigResponse, error)
	// Get current configs metadata.
	GetMetadata(ctx context.Context, in *Ydb_DynamicConfig.GetMetadataRequest, opts ...grpc.CallOption) (*Ydb_DynamicConfig.GetMetadataResponse, error)
	// Get current configs.
	GetConfig(ctx context.Context, in *Ydb_DynamicConfig.GetConfigRequest, opts ...grpc.CallOption) (*Ydb_DynamicConfig.GetConfigResponse, error)
	// Drop current config.
	// This call is idempotent.
	DropConfig(ctx context.Context, in *Ydb_DynamicConfig.DropConfigRequest, opts ...grpc.CallOption) (*Ydb_DynamicConfig.DropConfigResponse, error)
	// Add volatile config.
	AddVolatileConfig(ctx context.Context, in *Ydb_DynamicConfig.AddVolatileConfigRequest, opts ...grpc.CallOption) (*Ydb_DynamicConfig.AddVolatileConfigResponse, error)
	// Remove volatile config.
	RemoveVolatileConfig(ctx context.Context, in *Ydb_DynamicConfig.RemoveVolatileConfigRequest, opts ...grpc.CallOption) (*Ydb_DynamicConfig.RemoveVolatileConfigResponse, error)
	GetNodeLabels(ctx context.Context, in *Ydb_DynamicConfig.GetNodeLabelsRequest, opts ...grpc.CallOption) (*Ydb_DynamicConfig.GetNodeLabelsResponse, error)
	// Resolve config for particular labels.
	ResolveConfig(ctx context.Context, in *Ydb_DynamicConfig.ResolveConfigRequest, opts ...grpc.CallOption) (*Ydb_DynamicConfig.ResolveConfigResponse, error)
	// Resolve config for all possible labels combinations.
	ResolveAllConfig(ctx context.Context, in *Ydb_DynamicConfig.ResolveAllConfigRequest, opts ...grpc.CallOption) (*Ydb_DynamicConfig.ResolveAllConfigResponse, error)
}

type dynamicConfigServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDynamicConfigServiceClient(cc grpc.ClientConnInterface) DynamicConfigServiceClient {
	return &dynamicConfigServiceClient{cc}
}

func (c *dynamicConfigServiceClient) SetConfig(ctx context.Context, in *Ydb_DynamicConfig.SetConfigRequest, opts ...grpc.CallOption) (*Ydb_DynamicConfig.SetConfigResponse, error) {
	out := new(Ydb_DynamicConfig.SetConfigResponse)
	err := c.cc.Invoke(ctx, DynamicConfigService_SetConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dynamicConfigServiceClient) ReplaceConfig(ctx context.Context, in *Ydb_DynamicConfig.ReplaceConfigRequest, opts ...grpc.CallOption) (*Ydb_DynamicConfig.ReplaceConfigResponse, error) {
	out := new(Ydb_DynamicConfig.ReplaceConfigResponse)
	err := c.cc.Invoke(ctx, DynamicConfigService_ReplaceConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dynamicConfigServiceClient) GetMetadata(ctx context.Context, in *Ydb_DynamicConfig.GetMetadataRequest, opts ...grpc.CallOption) (*Ydb_DynamicConfig.GetMetadataResponse, error) {
	out := new(Ydb_DynamicConfig.GetMetadataResponse)
	err := c.cc.Invoke(ctx, DynamicConfigService_GetMetadata_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dynamicConfigServiceClient) GetConfig(ctx context.Context, in *Ydb_DynamicConfig.GetConfigRequest, opts ...grpc.CallOption) (*Ydb_DynamicConfig.GetConfigResponse, error) {
	out := new(Ydb_DynamicConfig.GetConfigResponse)
	err := c.cc.Invoke(ctx, DynamicConfigService_GetConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dynamicConfigServiceClient) DropConfig(ctx context.Context, in *Ydb_DynamicConfig.DropConfigRequest, opts ...grpc.CallOption) (*Ydb_DynamicConfig.DropConfigResponse, error) {
	out := new(Ydb_DynamicConfig.DropConfigResponse)
	err := c.cc.Invoke(ctx, DynamicConfigService_DropConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dynamicConfigServiceClient) AddVolatileConfig(ctx context.Context, in *Ydb_DynamicConfig.AddVolatileConfigRequest, opts ...grpc.CallOption) (*Ydb_DynamicConfig.AddVolatileConfigResponse, error) {
	out := new(Ydb_DynamicConfig.AddVolatileConfigResponse)
	err := c.cc.Invoke(ctx, DynamicConfigService_AddVolatileConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dynamicConfigServiceClient) RemoveVolatileConfig(ctx context.Context, in *Ydb_DynamicConfig.RemoveVolatileConfigRequest, opts ...grpc.CallOption) (*Ydb_DynamicConfig.RemoveVolatileConfigResponse, error) {
	out := new(Ydb_DynamicConfig.RemoveVolatileConfigResponse)
	err := c.cc.Invoke(ctx, DynamicConfigService_RemoveVolatileConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dynamicConfigServiceClient) GetNodeLabels(ctx context.Context, in *Ydb_DynamicConfig.GetNodeLabelsRequest, opts ...grpc.CallOption) (*Ydb_DynamicConfig.GetNodeLabelsResponse, error) {
	out := new(Ydb_DynamicConfig.GetNodeLabelsResponse)
	err := c.cc.Invoke(ctx, DynamicConfigService_GetNodeLabels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dynamicConfigServiceClient) ResolveConfig(ctx context.Context, in *Ydb_DynamicConfig.ResolveConfigRequest, opts ...grpc.CallOption) (*Ydb_DynamicConfig.ResolveConfigResponse, error) {
	out := new(Ydb_DynamicConfig.ResolveConfigResponse)
	err := c.cc.Invoke(ctx, DynamicConfigService_ResolveConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dynamicConfigServiceClient) ResolveAllConfig(ctx context.Context, in *Ydb_DynamicConfig.ResolveAllConfigRequest, opts ...grpc.CallOption) (*Ydb_DynamicConfig.ResolveAllConfigResponse, error) {
	out := new(Ydb_DynamicConfig.ResolveAllConfigResponse)
	err := c.cc.Invoke(ctx, DynamicConfigService_ResolveAllConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DynamicConfigServiceServer is the server API for DynamicConfigService service.
// All implementations must embed UnimplementedDynamicConfigServiceServer
// for forward compatibility
type DynamicConfigServiceServer interface {
	// Apply new version of config.
	// Overwrites previous version.
	// This call is idempotent:
	//   - Two subsequent identical calls will return success,
	//   - After applying next version another identical call will result in an error.
	//
	// The field version within the YAML in the request must strictly be set to the current version increment by one.
	// The field cluster within the YAML should be identical to the one configured on the node used by the console tablet.
	SetConfig(context.Context, *Ydb_DynamicConfig.SetConfigRequest) (*Ydb_DynamicConfig.SetConfigResponse, error)
	ReplaceConfig(context.Context, *Ydb_DynamicConfig.ReplaceConfigRequest) (*Ydb_DynamicConfig.ReplaceConfigResponse, error)
	// Get current configs metadata.
	GetMetadata(context.Context, *Ydb_DynamicConfig.GetMetadataRequest) (*Ydb_DynamicConfig.GetMetadataResponse, error)
	// Get current configs.
	GetConfig(context.Context, *Ydb_DynamicConfig.GetConfigRequest) (*Ydb_DynamicConfig.GetConfigResponse, error)
	// Drop current config.
	// This call is idempotent.
	DropConfig(context.Context, *Ydb_DynamicConfig.DropConfigRequest) (*Ydb_DynamicConfig.DropConfigResponse, error)
	// Add volatile config.
	AddVolatileConfig(context.Context, *Ydb_DynamicConfig.AddVolatileConfigRequest) (*Ydb_DynamicConfig.AddVolatileConfigResponse, error)
	// Remove volatile config.
	RemoveVolatileConfig(context.Context, *Ydb_DynamicConfig.RemoveVolatileConfigRequest) (*Ydb_DynamicConfig.RemoveVolatileConfigResponse, error)
	GetNodeLabels(context.Context, *Ydb_DynamicConfig.GetNodeLabelsRequest) (*Ydb_DynamicConfig.GetNodeLabelsResponse, error)
	// Resolve config for particular labels.
	ResolveConfig(context.Context, *Ydb_DynamicConfig.ResolveConfigRequest) (*Ydb_DynamicConfig.ResolveConfigResponse, error)
	// Resolve config for all possible labels combinations.
	ResolveAllConfig(context.Context, *Ydb_DynamicConfig.ResolveAllConfigRequest) (*Ydb_DynamicConfig.ResolveAllConfigResponse, error)
	mustEmbedUnimplementedDynamicConfigServiceServer()
}

// UnimplementedDynamicConfigServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDynamicConfigServiceServer struct {
}

func (UnimplementedDynamicConfigServiceServer) SetConfig(context.Context, *Ydb_DynamicConfig.SetConfigRequest) (*Ydb_DynamicConfig.SetConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConfig not implemented")
}
func (UnimplementedDynamicConfigServiceServer) ReplaceConfig(context.Context, *Ydb_DynamicConfig.ReplaceConfigRequest) (*Ydb_DynamicConfig.ReplaceConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplaceConfig not implemented")
}
func (UnimplementedDynamicConfigServiceServer) GetMetadata(context.Context, *Ydb_DynamicConfig.GetMetadataRequest) (*Ydb_DynamicConfig.GetMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetadata not implemented")
}
func (UnimplementedDynamicConfigServiceServer) GetConfig(context.Context, *Ydb_DynamicConfig.GetConfigRequest) (*Ydb_DynamicConfig.GetConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedDynamicConfigServiceServer) DropConfig(context.Context, *Ydb_DynamicConfig.DropConfigRequest) (*Ydb_DynamicConfig.DropConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropConfig not implemented")
}
func (UnimplementedDynamicConfigServiceServer) AddVolatileConfig(context.Context, *Ydb_DynamicConfig.AddVolatileConfigRequest) (*Ydb_DynamicConfig.AddVolatileConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVolatileConfig not implemented")
}
func (UnimplementedDynamicConfigServiceServer) RemoveVolatileConfig(context.Context, *Ydb_DynamicConfig.RemoveVolatileConfigRequest) (*Ydb_DynamicConfig.RemoveVolatileConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveVolatileConfig not implemented")
}
func (UnimplementedDynamicConfigServiceServer) GetNodeLabels(context.Context, *Ydb_DynamicConfig.GetNodeLabelsRequest) (*Ydb_DynamicConfig.GetNodeLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeLabels not implemented")
}
func (UnimplementedDynamicConfigServiceServer) ResolveConfig(context.Context, *Ydb_DynamicConfig.ResolveConfigRequest) (*Ydb_DynamicConfig.ResolveConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveConfig not implemented")
}
func (UnimplementedDynamicConfigServiceServer) ResolveAllConfig(context.Context, *Ydb_DynamicConfig.ResolveAllConfigRequest) (*Ydb_DynamicConfig.ResolveAllConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveAllConfig not implemented")
}
func (UnimplementedDynamicConfigServiceServer) mustEmbedUnimplementedDynamicConfigServiceServer() {}

// UnsafeDynamicConfigServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DynamicConfigServiceServer will
// result in compilation errors.
type UnsafeDynamicConfigServiceServer interface {
	mustEmbedUnimplementedDynamicConfigServiceServer()
}

func RegisterDynamicConfigServiceServer(s grpc.ServiceRegistrar, srv DynamicConfigServiceServer) {
	s.RegisterService(&DynamicConfigService_ServiceDesc, srv)
}

func _DynamicConfigService_SetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_DynamicConfig.SetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DynamicConfigServiceServer).SetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DynamicConfigService_SetConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DynamicConfigServiceServer).SetConfig(ctx, req.(*Ydb_DynamicConfig.SetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DynamicConfigService_ReplaceConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_DynamicConfig.ReplaceConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DynamicConfigServiceServer).ReplaceConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DynamicConfigService_ReplaceConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DynamicConfigServiceServer).ReplaceConfig(ctx, req.(*Ydb_DynamicConfig.ReplaceConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DynamicConfigService_GetMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_DynamicConfig.GetMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DynamicConfigServiceServer).GetMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DynamicConfigService_GetMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DynamicConfigServiceServer).GetMetadata(ctx, req.(*Ydb_DynamicConfig.GetMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DynamicConfigService_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_DynamicConfig.GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DynamicConfigServiceServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DynamicConfigService_GetConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DynamicConfigServiceServer).GetConfig(ctx, req.(*Ydb_DynamicConfig.GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DynamicConfigService_DropConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_DynamicConfig.DropConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DynamicConfigServiceServer).DropConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DynamicConfigService_DropConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DynamicConfigServiceServer).DropConfig(ctx, req.(*Ydb_DynamicConfig.DropConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DynamicConfigService_AddVolatileConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_DynamicConfig.AddVolatileConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DynamicConfigServiceServer).AddVolatileConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DynamicConfigService_AddVolatileConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DynamicConfigServiceServer).AddVolatileConfig(ctx, req.(*Ydb_DynamicConfig.AddVolatileConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DynamicConfigService_RemoveVolatileConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_DynamicConfig.RemoveVolatileConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DynamicConfigServiceServer).RemoveVolatileConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DynamicConfigService_RemoveVolatileConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DynamicConfigServiceServer).RemoveVolatileConfig(ctx, req.(*Ydb_DynamicConfig.RemoveVolatileConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DynamicConfigService_GetNodeLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_DynamicConfig.GetNodeLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DynamicConfigServiceServer).GetNodeLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DynamicConfigService_GetNodeLabels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DynamicConfigServiceServer).GetNodeLabels(ctx, req.(*Ydb_DynamicConfig.GetNodeLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DynamicConfigService_ResolveConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_DynamicConfig.ResolveConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DynamicConfigServiceServer).ResolveConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DynamicConfigService_ResolveConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DynamicConfigServiceServer).ResolveConfig(ctx, req.(*Ydb_DynamicConfig.ResolveConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DynamicConfigService_ResolveAllConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_DynamicConfig.ResolveAllConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DynamicConfigServiceServer).ResolveAllConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DynamicConfigService_ResolveAllConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DynamicConfigServiceServer).ResolveAllConfig(ctx, req.(*Ydb_DynamicConfig.ResolveAllConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DynamicConfigService_ServiceDesc is the grpc.ServiceDesc for DynamicConfigService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DynamicConfigService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Ydb.DynamicConfig.V1.DynamicConfigService",
	HandlerType: (*DynamicConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetConfig",
			Handler:    _DynamicConfigService_SetConfig_Handler,
		},
		{
			MethodName: "ReplaceConfig",
			Handler:    _DynamicConfigService_ReplaceConfig_Handler,
		},
		{
			MethodName: "GetMetadata",
			Handler:    _DynamicConfigService_GetMetadata_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _DynamicConfigService_GetConfig_Handler,
		},
		{
			MethodName: "DropConfig",
			Handler:    _DynamicConfigService_DropConfig_Handler,
		},
		{
			MethodName: "AddVolatileConfig",
			Handler:    _DynamicConfigService_AddVolatileConfig_Handler,
		},
		{
			MethodName: "RemoveVolatileConfig",
			Handler:    _DynamicConfigService_RemoveVolatileConfig_Handler,
		},
		{
			MethodName: "GetNodeLabels",
			Handler:    _DynamicConfigService_GetNodeLabels_Handler,
		},
		{
			MethodName: "ResolveConfig",
			Handler:    _DynamicConfigService_ResolveConfig_Handler,
		},
		{
			MethodName: "ResolveAllConfig",
			Handler:    _DynamicConfigService_ResolveAllConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "draft/ydb_dynamic_config_v1.proto",
}
