// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.1
// source: ydb_rate_limiter_v1.proto

package Ydb_RateLimiter_V1

import (
	Ydb_RateLimiter "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_RateLimiter"
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

var File_ydb_rate_limiter_v1_proto protoreflect.FileDescriptor

var file_ydb_rate_limiter_v1_proto_rawDesc = []byte{
	0x0a, 0x19, 0x79, 0x64, 0x62, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x59, 0x64, 0x62,
	0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0x1a,
	0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62, 0x5f, 0x72, 0x61, 0x74, 0x65,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe3,
	0x04, 0x0a, 0x12, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x26, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x52, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65,
	0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x0d, 0x41, 0x6c, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x25, 0x2e, 0x59, 0x64, 0x62, 0x2e,
	0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x6c, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x65, 0x72, 0x2e, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x0c, 0x44, 0x72, 0x6f, 0x70,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x24, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x52,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x72, 0x6f, 0x70, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x59, 0x64, 0x62, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72,
	0x2e, 0x44, 0x72, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x25, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x52, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x59, 0x64, 0x62, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x10, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x28, 0x2e, 0x59, 0x64, 0x62, 0x2e,
	0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64,
	0x0a, 0x0f, 0x41, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x27, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x59, 0x64, 0x62,
	0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x0a, 0x1e, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x79, 0x64, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x59, 0x64, 0x62, 0x5f, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x65, 0x72, 0x5f, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_ydb_rate_limiter_v1_proto_goTypes = []interface{}{
	(*Ydb_RateLimiter.CreateResourceRequest)(nil),    // 0: Ydb.RateLimiter.CreateResourceRequest
	(*Ydb_RateLimiter.AlterResourceRequest)(nil),     // 1: Ydb.RateLimiter.AlterResourceRequest
	(*Ydb_RateLimiter.DropResourceRequest)(nil),      // 2: Ydb.RateLimiter.DropResourceRequest
	(*Ydb_RateLimiter.ListResourcesRequest)(nil),     // 3: Ydb.RateLimiter.ListResourcesRequest
	(*Ydb_RateLimiter.DescribeResourceRequest)(nil),  // 4: Ydb.RateLimiter.DescribeResourceRequest
	(*Ydb_RateLimiter.AcquireResourceRequest)(nil),   // 5: Ydb.RateLimiter.AcquireResourceRequest
	(*Ydb_RateLimiter.CreateResourceResponse)(nil),   // 6: Ydb.RateLimiter.CreateResourceResponse
	(*Ydb_RateLimiter.AlterResourceResponse)(nil),    // 7: Ydb.RateLimiter.AlterResourceResponse
	(*Ydb_RateLimiter.DropResourceResponse)(nil),     // 8: Ydb.RateLimiter.DropResourceResponse
	(*Ydb_RateLimiter.ListResourcesResponse)(nil),    // 9: Ydb.RateLimiter.ListResourcesResponse
	(*Ydb_RateLimiter.DescribeResourceResponse)(nil), // 10: Ydb.RateLimiter.DescribeResourceResponse
	(*Ydb_RateLimiter.AcquireResourceResponse)(nil),  // 11: Ydb.RateLimiter.AcquireResourceResponse
}
var file_ydb_rate_limiter_v1_proto_depIdxs = []int32{
	0,  // 0: Ydb.RateLimiter.V1.RateLimiterService.CreateResource:input_type -> Ydb.RateLimiter.CreateResourceRequest
	1,  // 1: Ydb.RateLimiter.V1.RateLimiterService.AlterResource:input_type -> Ydb.RateLimiter.AlterResourceRequest
	2,  // 2: Ydb.RateLimiter.V1.RateLimiterService.DropResource:input_type -> Ydb.RateLimiter.DropResourceRequest
	3,  // 3: Ydb.RateLimiter.V1.RateLimiterService.ListResources:input_type -> Ydb.RateLimiter.ListResourcesRequest
	4,  // 4: Ydb.RateLimiter.V1.RateLimiterService.DescribeResource:input_type -> Ydb.RateLimiter.DescribeResourceRequest
	5,  // 5: Ydb.RateLimiter.V1.RateLimiterService.AcquireResource:input_type -> Ydb.RateLimiter.AcquireResourceRequest
	6,  // 6: Ydb.RateLimiter.V1.RateLimiterService.CreateResource:output_type -> Ydb.RateLimiter.CreateResourceResponse
	7,  // 7: Ydb.RateLimiter.V1.RateLimiterService.AlterResource:output_type -> Ydb.RateLimiter.AlterResourceResponse
	8,  // 8: Ydb.RateLimiter.V1.RateLimiterService.DropResource:output_type -> Ydb.RateLimiter.DropResourceResponse
	9,  // 9: Ydb.RateLimiter.V1.RateLimiterService.ListResources:output_type -> Ydb.RateLimiter.ListResourcesResponse
	10, // 10: Ydb.RateLimiter.V1.RateLimiterService.DescribeResource:output_type -> Ydb.RateLimiter.DescribeResourceResponse
	11, // 11: Ydb.RateLimiter.V1.RateLimiterService.AcquireResource:output_type -> Ydb.RateLimiter.AcquireResourceResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_ydb_rate_limiter_v1_proto_init() }
func file_ydb_rate_limiter_v1_proto_init() {
	if File_ydb_rate_limiter_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ydb_rate_limiter_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ydb_rate_limiter_v1_proto_goTypes,
		DependencyIndexes: file_ydb_rate_limiter_v1_proto_depIdxs,
	}.Build()
	File_ydb_rate_limiter_v1_proto = out.File
	file_ydb_rate_limiter_v1_proto_rawDesc = nil
	file_ydb_rate_limiter_v1_proto_goTypes = nil
	file_ydb_rate_limiter_v1_proto_depIdxs = nil
}
