// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: ydb_cms_v1.proto

package Ydb_Cms_V1

import (
	context "context"
	Ydb_Cms "github.com/EgorkaZ/ydb-go-genproto/protos/Ydb_Cms"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CmsService_CreateDatabase_FullMethodName          = "/Ydb.Cms.V1.CmsService/CreateDatabase"
	CmsService_GetDatabaseStatus_FullMethodName       = "/Ydb.Cms.V1.CmsService/GetDatabaseStatus"
	CmsService_AlterDatabase_FullMethodName           = "/Ydb.Cms.V1.CmsService/AlterDatabase"
	CmsService_ListDatabases_FullMethodName           = "/Ydb.Cms.V1.CmsService/ListDatabases"
	CmsService_RemoveDatabase_FullMethodName          = "/Ydb.Cms.V1.CmsService/RemoveDatabase"
	CmsService_DescribeDatabaseOptions_FullMethodName = "/Ydb.Cms.V1.CmsService/DescribeDatabaseOptions"
)

// CmsServiceClient is the client API for CmsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CmsServiceClient interface {
	// Create a new database.
	CreateDatabase(ctx context.Context, in *Ydb_Cms.CreateDatabaseRequest, opts ...grpc.CallOption) (*Ydb_Cms.CreateDatabaseResponse, error)
	// Get current database's status.
	GetDatabaseStatus(ctx context.Context, in *Ydb_Cms.GetDatabaseStatusRequest, opts ...grpc.CallOption) (*Ydb_Cms.GetDatabaseStatusResponse, error)
	// Alter database resources.
	AlterDatabase(ctx context.Context, in *Ydb_Cms.AlterDatabaseRequest, opts ...grpc.CallOption) (*Ydb_Cms.AlterDatabaseResponse, error)
	// List all databases.
	ListDatabases(ctx context.Context, in *Ydb_Cms.ListDatabasesRequest, opts ...grpc.CallOption) (*Ydb_Cms.ListDatabasesResponse, error)
	// Remove database.
	RemoveDatabase(ctx context.Context, in *Ydb_Cms.RemoveDatabaseRequest, opts ...grpc.CallOption) (*Ydb_Cms.RemoveDatabaseResponse, error)
	// Describe supported database options.
	DescribeDatabaseOptions(ctx context.Context, in *Ydb_Cms.DescribeDatabaseOptionsRequest, opts ...grpc.CallOption) (*Ydb_Cms.DescribeDatabaseOptionsResponse, error)
}

type cmsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCmsServiceClient(cc grpc.ClientConnInterface) CmsServiceClient {
	return &cmsServiceClient{cc}
}

func (c *cmsServiceClient) CreateDatabase(ctx context.Context, in *Ydb_Cms.CreateDatabaseRequest, opts ...grpc.CallOption) (*Ydb_Cms.CreateDatabaseResponse, error) {
	out := new(Ydb_Cms.CreateDatabaseResponse)
	err := c.cc.Invoke(ctx, CmsService_CreateDatabase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsServiceClient) GetDatabaseStatus(ctx context.Context, in *Ydb_Cms.GetDatabaseStatusRequest, opts ...grpc.CallOption) (*Ydb_Cms.GetDatabaseStatusResponse, error) {
	out := new(Ydb_Cms.GetDatabaseStatusResponse)
	err := c.cc.Invoke(ctx, CmsService_GetDatabaseStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsServiceClient) AlterDatabase(ctx context.Context, in *Ydb_Cms.AlterDatabaseRequest, opts ...grpc.CallOption) (*Ydb_Cms.AlterDatabaseResponse, error) {
	out := new(Ydb_Cms.AlterDatabaseResponse)
	err := c.cc.Invoke(ctx, CmsService_AlterDatabase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsServiceClient) ListDatabases(ctx context.Context, in *Ydb_Cms.ListDatabasesRequest, opts ...grpc.CallOption) (*Ydb_Cms.ListDatabasesResponse, error) {
	out := new(Ydb_Cms.ListDatabasesResponse)
	err := c.cc.Invoke(ctx, CmsService_ListDatabases_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsServiceClient) RemoveDatabase(ctx context.Context, in *Ydb_Cms.RemoveDatabaseRequest, opts ...grpc.CallOption) (*Ydb_Cms.RemoveDatabaseResponse, error) {
	out := new(Ydb_Cms.RemoveDatabaseResponse)
	err := c.cc.Invoke(ctx, CmsService_RemoveDatabase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsServiceClient) DescribeDatabaseOptions(ctx context.Context, in *Ydb_Cms.DescribeDatabaseOptionsRequest, opts ...grpc.CallOption) (*Ydb_Cms.DescribeDatabaseOptionsResponse, error) {
	out := new(Ydb_Cms.DescribeDatabaseOptionsResponse)
	err := c.cc.Invoke(ctx, CmsService_DescribeDatabaseOptions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CmsServiceServer is the server API for CmsService service.
// All implementations must embed UnimplementedCmsServiceServer
// for forward compatibility
type CmsServiceServer interface {
	// Create a new database.
	CreateDatabase(context.Context, *Ydb_Cms.CreateDatabaseRequest) (*Ydb_Cms.CreateDatabaseResponse, error)
	// Get current database's status.
	GetDatabaseStatus(context.Context, *Ydb_Cms.GetDatabaseStatusRequest) (*Ydb_Cms.GetDatabaseStatusResponse, error)
	// Alter database resources.
	AlterDatabase(context.Context, *Ydb_Cms.AlterDatabaseRequest) (*Ydb_Cms.AlterDatabaseResponse, error)
	// List all databases.
	ListDatabases(context.Context, *Ydb_Cms.ListDatabasesRequest) (*Ydb_Cms.ListDatabasesResponse, error)
	// Remove database.
	RemoveDatabase(context.Context, *Ydb_Cms.RemoveDatabaseRequest) (*Ydb_Cms.RemoveDatabaseResponse, error)
	// Describe supported database options.
	DescribeDatabaseOptions(context.Context, *Ydb_Cms.DescribeDatabaseOptionsRequest) (*Ydb_Cms.DescribeDatabaseOptionsResponse, error)
	mustEmbedUnimplementedCmsServiceServer()
}

// UnimplementedCmsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCmsServiceServer struct {
}

func (UnimplementedCmsServiceServer) CreateDatabase(context.Context, *Ydb_Cms.CreateDatabaseRequest) (*Ydb_Cms.CreateDatabaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDatabase not implemented")
}
func (UnimplementedCmsServiceServer) GetDatabaseStatus(context.Context, *Ydb_Cms.GetDatabaseStatusRequest) (*Ydb_Cms.GetDatabaseStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatabaseStatus not implemented")
}
func (UnimplementedCmsServiceServer) AlterDatabase(context.Context, *Ydb_Cms.AlterDatabaseRequest) (*Ydb_Cms.AlterDatabaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlterDatabase not implemented")
}
func (UnimplementedCmsServiceServer) ListDatabases(context.Context, *Ydb_Cms.ListDatabasesRequest) (*Ydb_Cms.ListDatabasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDatabases not implemented")
}
func (UnimplementedCmsServiceServer) RemoveDatabase(context.Context, *Ydb_Cms.RemoveDatabaseRequest) (*Ydb_Cms.RemoveDatabaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDatabase not implemented")
}
func (UnimplementedCmsServiceServer) DescribeDatabaseOptions(context.Context, *Ydb_Cms.DescribeDatabaseOptionsRequest) (*Ydb_Cms.DescribeDatabaseOptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeDatabaseOptions not implemented")
}
func (UnimplementedCmsServiceServer) mustEmbedUnimplementedCmsServiceServer() {}

// UnsafeCmsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CmsServiceServer will
// result in compilation errors.
type UnsafeCmsServiceServer interface {
	mustEmbedUnimplementedCmsServiceServer()
}

func RegisterCmsServiceServer(s grpc.ServiceRegistrar, srv CmsServiceServer) {
	s.RegisterService(&CmsService_ServiceDesc, srv)
}

func _CmsService_CreateDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Cms.CreateDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServiceServer).CreateDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CmsService_CreateDatabase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServiceServer).CreateDatabase(ctx, req.(*Ydb_Cms.CreateDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmsService_GetDatabaseStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Cms.GetDatabaseStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServiceServer).GetDatabaseStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CmsService_GetDatabaseStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServiceServer).GetDatabaseStatus(ctx, req.(*Ydb_Cms.GetDatabaseStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmsService_AlterDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Cms.AlterDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServiceServer).AlterDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CmsService_AlterDatabase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServiceServer).AlterDatabase(ctx, req.(*Ydb_Cms.AlterDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmsService_ListDatabases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Cms.ListDatabasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServiceServer).ListDatabases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CmsService_ListDatabases_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServiceServer).ListDatabases(ctx, req.(*Ydb_Cms.ListDatabasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmsService_RemoveDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Cms.RemoveDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServiceServer).RemoveDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CmsService_RemoveDatabase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServiceServer).RemoveDatabase(ctx, req.(*Ydb_Cms.RemoveDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmsService_DescribeDatabaseOptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Cms.DescribeDatabaseOptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServiceServer).DescribeDatabaseOptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CmsService_DescribeDatabaseOptions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServiceServer).DescribeDatabaseOptions(ctx, req.(*Ydb_Cms.DescribeDatabaseOptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CmsService_ServiceDesc is the grpc.ServiceDesc for CmsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CmsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Ydb.Cms.V1.CmsService",
	HandlerType: (*CmsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDatabase",
			Handler:    _CmsService_CreateDatabase_Handler,
		},
		{
			MethodName: "GetDatabaseStatus",
			Handler:    _CmsService_GetDatabaseStatus_Handler,
		},
		{
			MethodName: "AlterDatabase",
			Handler:    _CmsService_AlterDatabase_Handler,
		},
		{
			MethodName: "ListDatabases",
			Handler:    _CmsService_ListDatabases_Handler,
		},
		{
			MethodName: "RemoveDatabase",
			Handler:    _CmsService_RemoveDatabase_Handler,
		},
		{
			MethodName: "DescribeDatabaseOptions",
			Handler:    _CmsService_DescribeDatabaseOptions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ydb_cms_v1.proto",
}
