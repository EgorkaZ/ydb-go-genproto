// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package Ydb_Import_V1

import (
	context "context"
	Ydb_Import "github.com/YandexDatabase/ydb-go-genproto/protos/Ydb_Import"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ImportServiceClient is the client API for ImportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImportServiceClient interface {
	// Imports data from S3.
	// Method starts an asynchronous operation that can be cancelled while it is in progress.
	ImportFromS3(ctx context.Context, in *Ydb_Import.ImportFromS3Request, opts ...grpc.CallOption) (*Ydb_Import.ImportFromS3Response, error)
	// Writes data to a table.
	// Method accepts serialized data in the selected format and writes it non-transactionally.
	ImportData(ctx context.Context, in *Ydb_Import.ImportDataRequest, opts ...grpc.CallOption) (*Ydb_Import.ImportDataResponse, error)
}

type importServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImportServiceClient(cc grpc.ClientConnInterface) ImportServiceClient {
	return &importServiceClient{cc}
}

func (c *importServiceClient) ImportFromS3(ctx context.Context, in *Ydb_Import.ImportFromS3Request, opts ...grpc.CallOption) (*Ydb_Import.ImportFromS3Response, error) {
	out := new(Ydb_Import.ImportFromS3Response)
	err := c.cc.Invoke(ctx, "/Ydb.Import.V1.ImportService/ImportFromS3", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *importServiceClient) ImportData(ctx context.Context, in *Ydb_Import.ImportDataRequest, opts ...grpc.CallOption) (*Ydb_Import.ImportDataResponse, error) {
	out := new(Ydb_Import.ImportDataResponse)
	err := c.cc.Invoke(ctx, "/Ydb.Import.V1.ImportService/ImportData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImportServiceServer is the server API for ImportService service.
// All implementations must embed UnimplementedImportServiceServer
// for forward compatibility
type ImportServiceServer interface {
	// Imports data from S3.
	// Method starts an asynchronous operation that can be cancelled while it is in progress.
	ImportFromS3(context.Context, *Ydb_Import.ImportFromS3Request) (*Ydb_Import.ImportFromS3Response, error)
	// Writes data to a table.
	// Method accepts serialized data in the selected format and writes it non-transactionally.
	ImportData(context.Context, *Ydb_Import.ImportDataRequest) (*Ydb_Import.ImportDataResponse, error)
	mustEmbedUnimplementedImportServiceServer()
}

// UnimplementedImportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImportServiceServer struct {
}

func (UnimplementedImportServiceServer) ImportFromS3(context.Context, *Ydb_Import.ImportFromS3Request) (*Ydb_Import.ImportFromS3Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportFromS3 not implemented")
}
func (UnimplementedImportServiceServer) ImportData(context.Context, *Ydb_Import.ImportDataRequest) (*Ydb_Import.ImportDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportData not implemented")
}
func (UnimplementedImportServiceServer) mustEmbedUnimplementedImportServiceServer() {}

// UnsafeImportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImportServiceServer will
// result in compilation errors.
type UnsafeImportServiceServer interface {
	mustEmbedUnimplementedImportServiceServer()
}

func RegisterImportServiceServer(s grpc.ServiceRegistrar, srv ImportServiceServer) {
	s.RegisterService(&ImportService_ServiceDesc, srv)
}

func _ImportService_ImportFromS3_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Import.ImportFromS3Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImportServiceServer).ImportFromS3(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ydb.Import.V1.ImportService/ImportFromS3",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImportServiceServer).ImportFromS3(ctx, req.(*Ydb_Import.ImportFromS3Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImportService_ImportData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Import.ImportDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImportServiceServer).ImportData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ydb.Import.V1.ImportService/ImportData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImportServiceServer).ImportData(ctx, req.(*Ydb_Import.ImportDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImportService_ServiceDesc is the grpc.ServiceDesc for ImportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Ydb.Import.V1.ImportService",
	HandlerType: (*ImportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ImportFromS3",
			Handler:    _ImportService_ImportFromS3_Handler,
		},
		{
			MethodName: "ImportData",
			Handler:    _ImportService_ImportData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ydb_import_v1.proto",
}
