// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.1
// source: ydb_export_v1.proto

package Ydb_Export_V1

import (
	Ydb_Export "github.com/EgorkaZ/ydb-go-genproto/protos/Ydb_Export"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_ydb_export_v1_proto protoreflect.FileDescriptor

var file_ydb_export_v1_proto_rawDesc = []byte{
	0x0a, 0x13, 0x79, 0x64, 0x62, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x76, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x59, 0x64, 0x62, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x56, 0x31, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62,
	0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa9, 0x01,
	0x0a, 0x0d, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4b, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x6f, 0x59, 0x74, 0x12, 0x1d, 0x2e,
	0x59, 0x64, 0x62, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x54, 0x6f, 0x59, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x59,
	0x64, 0x62, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x54, 0x6f, 0x59, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0a,
	0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x6f, 0x53, 0x33, 0x12, 0x1d, 0x2e, 0x59, 0x64, 0x62,
	0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x6f,
	0x53, 0x33, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x59, 0x64, 0x62, 0x2e,
	0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x6f, 0x53,
	0x33, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4c, 0x0a, 0x18, 0x74, 0x65, 0x63,
	0x68, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x76, 0x31, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x45, 0x67, 0x6f, 0x72, 0x6b, 0x61, 0x5a, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x67, 0x6f,
	0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x59, 0x64, 0x62, 0x5f, 0x45, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_ydb_export_v1_proto_goTypes = []interface{}{
	(*Ydb_Export.ExportToYtRequest)(nil),  // 0: Ydb.Export.ExportToYtRequest
	(*Ydb_Export.ExportToS3Request)(nil),  // 1: Ydb.Export.ExportToS3Request
	(*Ydb_Export.ExportToYtResponse)(nil), // 2: Ydb.Export.ExportToYtResponse
	(*Ydb_Export.ExportToS3Response)(nil), // 3: Ydb.Export.ExportToS3Response
}
var file_ydb_export_v1_proto_depIdxs = []int32{
	0, // 0: Ydb.Export.V1.ExportService.ExportToYt:input_type -> Ydb.Export.ExportToYtRequest
	1, // 1: Ydb.Export.V1.ExportService.ExportToS3:input_type -> Ydb.Export.ExportToS3Request
	2, // 2: Ydb.Export.V1.ExportService.ExportToYt:output_type -> Ydb.Export.ExportToYtResponse
	3, // 3: Ydb.Export.V1.ExportService.ExportToS3:output_type -> Ydb.Export.ExportToS3Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ydb_export_v1_proto_init() }
func file_ydb_export_v1_proto_init() {
	if File_ydb_export_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ydb_export_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ydb_export_v1_proto_goTypes,
		DependencyIndexes: file_ydb_export_v1_proto_depIdxs,
	}.Build()
	File_ydb_export_v1_proto = out.File
	file_ydb_export_v1_proto_rawDesc = nil
	file_ydb_export_v1_proto_goTypes = nil
	file_ydb_export_v1_proto_depIdxs = nil
}
