// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.1
// source: draft/ydb_dynamic_config_v1.proto

package Ydb_DynamicConfig_V1

import (
	Ydb_DynamicConfig "github.com/EgorkaZ/ydb-go-genproto/draft/protos/Ydb_DynamicConfig"
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

var File_draft_ydb_dynamic_config_v1_proto protoreflect.FileDescriptor

var file_draft_ydb_dynamic_config_v1_proto_rawDesc = []byte{
	0x0a, 0x21, 0x64, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x79, 0x64, 0x62, 0x5f, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x59, 0x64, 0x62, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x56, 0x31, 0x1a, 0x25, 0x64, 0x72, 0x61, 0x66, 0x74,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62, 0x5f, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0x81, 0x08, 0x0a, 0x14, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x09, 0x53, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x23, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x44, 0x79, 0x6e,
	0x61, 0x6d, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x59, 0x64,
	0x62, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x62, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x27, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x59, 0x64,
	0x62, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d,
	0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x59, 0x64,
	0x62, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x23, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x44, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x0a, 0x44,
	0x72, 0x6f, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x24, 0x2e, 0x59, 0x64, 0x62, 0x2e,
	0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x44, 0x72,
	0x6f, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x44, 0x72, 0x6f, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6e, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x56, 0x6f, 0x6c,
	0x61, 0x74, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2b, 0x2e, 0x59, 0x64,
	0x62, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x41, 0x64, 0x64, 0x56, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x44,
	0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x64, 0x64,
	0x56, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x77, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x56, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2e,
	0x2e, 0x59, 0x64, 0x62, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x56, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6c,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f,
	0x2e, 0x59, 0x64, 0x62, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x56, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6c,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x62, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x12, 0x27, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x59, 0x64, 0x62, 0x2e,
	0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x65,
	0x74, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x27, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d,
	0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e,
	0x59, 0x64, 0x62, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2a, 0x2e, 0x59, 0x64,
	0x62, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x44, 0x79,
	0x6e, 0x61, 0x6d, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x69, 0x0a, 0x25, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x79, 0x64, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x64, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x64, 0x79, 0x6e,
	0x61, 0x6d, 0x69, 0x63, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x5a, 0x3d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x67, 0x6f, 0x72, 0x6b, 0x61,
	0x5a, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x64, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x59, 0x64, 0x62, 0x5f, 0x44, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x56, 0x31, 0xf8, 0x01, 0x01, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_draft_ydb_dynamic_config_v1_proto_goTypes = []interface{}{
	(*Ydb_DynamicConfig.SetConfigRequest)(nil),             // 0: Ydb.DynamicConfig.SetConfigRequest
	(*Ydb_DynamicConfig.ReplaceConfigRequest)(nil),         // 1: Ydb.DynamicConfig.ReplaceConfigRequest
	(*Ydb_DynamicConfig.GetMetadataRequest)(nil),           // 2: Ydb.DynamicConfig.GetMetadataRequest
	(*Ydb_DynamicConfig.GetConfigRequest)(nil),             // 3: Ydb.DynamicConfig.GetConfigRequest
	(*Ydb_DynamicConfig.DropConfigRequest)(nil),            // 4: Ydb.DynamicConfig.DropConfigRequest
	(*Ydb_DynamicConfig.AddVolatileConfigRequest)(nil),     // 5: Ydb.DynamicConfig.AddVolatileConfigRequest
	(*Ydb_DynamicConfig.RemoveVolatileConfigRequest)(nil),  // 6: Ydb.DynamicConfig.RemoveVolatileConfigRequest
	(*Ydb_DynamicConfig.GetNodeLabelsRequest)(nil),         // 7: Ydb.DynamicConfig.GetNodeLabelsRequest
	(*Ydb_DynamicConfig.ResolveConfigRequest)(nil),         // 8: Ydb.DynamicConfig.ResolveConfigRequest
	(*Ydb_DynamicConfig.ResolveAllConfigRequest)(nil),      // 9: Ydb.DynamicConfig.ResolveAllConfigRequest
	(*Ydb_DynamicConfig.SetConfigResponse)(nil),            // 10: Ydb.DynamicConfig.SetConfigResponse
	(*Ydb_DynamicConfig.ReplaceConfigResponse)(nil),        // 11: Ydb.DynamicConfig.ReplaceConfigResponse
	(*Ydb_DynamicConfig.GetMetadataResponse)(nil),          // 12: Ydb.DynamicConfig.GetMetadataResponse
	(*Ydb_DynamicConfig.GetConfigResponse)(nil),            // 13: Ydb.DynamicConfig.GetConfigResponse
	(*Ydb_DynamicConfig.DropConfigResponse)(nil),           // 14: Ydb.DynamicConfig.DropConfigResponse
	(*Ydb_DynamicConfig.AddVolatileConfigResponse)(nil),    // 15: Ydb.DynamicConfig.AddVolatileConfigResponse
	(*Ydb_DynamicConfig.RemoveVolatileConfigResponse)(nil), // 16: Ydb.DynamicConfig.RemoveVolatileConfigResponse
	(*Ydb_DynamicConfig.GetNodeLabelsResponse)(nil),        // 17: Ydb.DynamicConfig.GetNodeLabelsResponse
	(*Ydb_DynamicConfig.ResolveConfigResponse)(nil),        // 18: Ydb.DynamicConfig.ResolveConfigResponse
	(*Ydb_DynamicConfig.ResolveAllConfigResponse)(nil),     // 19: Ydb.DynamicConfig.ResolveAllConfigResponse
}
var file_draft_ydb_dynamic_config_v1_proto_depIdxs = []int32{
	0,  // 0: Ydb.DynamicConfig.V1.DynamicConfigService.SetConfig:input_type -> Ydb.DynamicConfig.SetConfigRequest
	1,  // 1: Ydb.DynamicConfig.V1.DynamicConfigService.ReplaceConfig:input_type -> Ydb.DynamicConfig.ReplaceConfigRequest
	2,  // 2: Ydb.DynamicConfig.V1.DynamicConfigService.GetMetadata:input_type -> Ydb.DynamicConfig.GetMetadataRequest
	3,  // 3: Ydb.DynamicConfig.V1.DynamicConfigService.GetConfig:input_type -> Ydb.DynamicConfig.GetConfigRequest
	4,  // 4: Ydb.DynamicConfig.V1.DynamicConfigService.DropConfig:input_type -> Ydb.DynamicConfig.DropConfigRequest
	5,  // 5: Ydb.DynamicConfig.V1.DynamicConfigService.AddVolatileConfig:input_type -> Ydb.DynamicConfig.AddVolatileConfigRequest
	6,  // 6: Ydb.DynamicConfig.V1.DynamicConfigService.RemoveVolatileConfig:input_type -> Ydb.DynamicConfig.RemoveVolatileConfigRequest
	7,  // 7: Ydb.DynamicConfig.V1.DynamicConfigService.GetNodeLabels:input_type -> Ydb.DynamicConfig.GetNodeLabelsRequest
	8,  // 8: Ydb.DynamicConfig.V1.DynamicConfigService.ResolveConfig:input_type -> Ydb.DynamicConfig.ResolveConfigRequest
	9,  // 9: Ydb.DynamicConfig.V1.DynamicConfigService.ResolveAllConfig:input_type -> Ydb.DynamicConfig.ResolveAllConfigRequest
	10, // 10: Ydb.DynamicConfig.V1.DynamicConfigService.SetConfig:output_type -> Ydb.DynamicConfig.SetConfigResponse
	11, // 11: Ydb.DynamicConfig.V1.DynamicConfigService.ReplaceConfig:output_type -> Ydb.DynamicConfig.ReplaceConfigResponse
	12, // 12: Ydb.DynamicConfig.V1.DynamicConfigService.GetMetadata:output_type -> Ydb.DynamicConfig.GetMetadataResponse
	13, // 13: Ydb.DynamicConfig.V1.DynamicConfigService.GetConfig:output_type -> Ydb.DynamicConfig.GetConfigResponse
	14, // 14: Ydb.DynamicConfig.V1.DynamicConfigService.DropConfig:output_type -> Ydb.DynamicConfig.DropConfigResponse
	15, // 15: Ydb.DynamicConfig.V1.DynamicConfigService.AddVolatileConfig:output_type -> Ydb.DynamicConfig.AddVolatileConfigResponse
	16, // 16: Ydb.DynamicConfig.V1.DynamicConfigService.RemoveVolatileConfig:output_type -> Ydb.DynamicConfig.RemoveVolatileConfigResponse
	17, // 17: Ydb.DynamicConfig.V1.DynamicConfigService.GetNodeLabels:output_type -> Ydb.DynamicConfig.GetNodeLabelsResponse
	18, // 18: Ydb.DynamicConfig.V1.DynamicConfigService.ResolveConfig:output_type -> Ydb.DynamicConfig.ResolveConfigResponse
	19, // 19: Ydb.DynamicConfig.V1.DynamicConfigService.ResolveAllConfig:output_type -> Ydb.DynamicConfig.ResolveAllConfigResponse
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_draft_ydb_dynamic_config_v1_proto_init() }
func file_draft_ydb_dynamic_config_v1_proto_init() {
	if File_draft_ydb_dynamic_config_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_draft_ydb_dynamic_config_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_draft_ydb_dynamic_config_v1_proto_goTypes,
		DependencyIndexes: file_draft_ydb_dynamic_config_v1_proto_depIdxs,
	}.Build()
	File_draft_ydb_dynamic_config_v1_proto = out.File
	file_draft_ydb_dynamic_config_v1_proto_rawDesc = nil
	file_draft_ydb_dynamic_config_v1_proto_goTypes = nil
	file_draft_ydb_dynamic_config_v1_proto_depIdxs = nil
}
