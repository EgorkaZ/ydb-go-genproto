// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: draft/ydb_maintenance_v1.proto

package Ydb_Maintenance_V1

import (
	Ydb_Maintenance "github.com/ydb-platform/ydb-go-genproto/draft/protos/Ydb_Maintenance"
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

var File_draft_ydb_maintenance_v1_proto protoreflect.FileDescriptor

var file_draft_ydb_maintenance_v1_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x64, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x79, 0x64, 0x62, 0x5f, 0x6d, 0x61, 0x69, 0x6e,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x59, 0x64, 0x62, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63,
	0x65, 0x2e, 0x56, 0x31, 0x1a, 0x22, 0x64, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x84, 0x09, 0x0a, 0x12, 0x4d, 0x61, 0x69,
	0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x67, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x12, 0x28, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x59, 0x64, 0x62, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x70, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x2d, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x72, 0x0a, 0x16, 0x52, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x12, 0x2e, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x4d, 0x61,
	0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x74,
	0x0a, 0x19, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x2a, 0x2e, 0x59, 0x64,
	0x62, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x4d, 0x61,
	0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69,
	0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x73, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x69, 0x6e,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x2c, 0x2e, 0x59,
	0x64, 0x62, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x59, 0x64, 0x62,
	0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x72, 0x0a, 0x13, 0x44, 0x72, 0x6f,
	0x70, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x2b, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x44, 0x72, 0x6f, 0x70, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e,
	0x59, 0x64, 0x62, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7e, 0x0a,
	0x19, 0x50, 0x72, 0x6f, 0x6c, 0x6f, 0x6e, 0x67, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x31, 0x2e, 0x59, 0x64, 0x62,
	0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f,
	0x6c, 0x6f, 0x6e, 0x67, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e,
	0x59, 0x64, 0x62, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a,
	0x0d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x12, 0x25,
	0x2e, 0x59, 0x64, 0x62, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x4d, 0x61, 0x69, 0x6e,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x10,
	0x50, 0x72, 0x6f, 0x6c, 0x6f, 0x6e, 0x67, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x74,
	0x12, 0x28, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6c, 0x6f, 0x6e, 0x67, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x59, 0x64, 0x62,
	0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x7c, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x2f, 0x2e, 0x59,
	0x64, 0x62, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e,
	0x59, 0x64, 0x62, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x5e, 0x0a, 0x1a, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x64, 0x72, 0x61, 0x66,
	0x74, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5a, 0x40, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x59, 0x64, 0x62,
	0x5f, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_draft_ydb_maintenance_v1_proto_goTypes = []interface{}{
	(*Ydb_Maintenance.ListClusterNodesRequest)(nil),          // 0: Ydb.Maintenance.ListClusterNodesRequest
	(*Ydb_Maintenance.CreateMaintenanceTaskRequest)(nil),     // 1: Ydb.Maintenance.CreateMaintenanceTaskRequest
	(*Ydb_Maintenance.RefreshMaintenanceTaskRequest)(nil),    // 2: Ydb.Maintenance.RefreshMaintenanceTaskRequest
	(*Ydb_Maintenance.GetMaintenanceTaskRequest)(nil),        // 3: Ydb.Maintenance.GetMaintenanceTaskRequest
	(*Ydb_Maintenance.ListMaintenanceTasksRequest)(nil),      // 4: Ydb.Maintenance.ListMaintenanceTasksRequest
	(*Ydb_Maintenance.DropMaintenanceTaskRequest)(nil),       // 5: Ydb.Maintenance.DropMaintenanceTaskRequest
	(*Ydb_Maintenance.ProlongateMaintenanceTaskRequest)(nil), // 6: Ydb.Maintenance.ProlongateMaintenanceTaskRequest
	(*Ydb_Maintenance.ReleasePermitRequest)(nil),             // 7: Ydb.Maintenance.ReleasePermitRequest
	(*Ydb_Maintenance.ProlongatePermitRequest)(nil),          // 8: Ydb.Maintenance.ProlongatePermitRequest
	(*Ydb_Maintenance.GetReadableActionReasonRequest)(nil),   // 9: Ydb.Maintenance.GetReadableActionReasonRequest
	(*Ydb_Maintenance.ListClusterNodesResponse)(nil),         // 10: Ydb.Maintenance.ListClusterNodesResponse
	(*Ydb_Maintenance.MaintenanceTaskResponse)(nil),          // 11: Ydb.Maintenance.MaintenanceTaskResponse
	(*Ydb_Maintenance.GetMaintenanceTaskResponse)(nil),       // 12: Ydb.Maintenance.GetMaintenanceTaskResponse
	(*Ydb_Maintenance.ListMaintenanceTasksResponse)(nil),     // 13: Ydb.Maintenance.ListMaintenanceTasksResponse
	(*Ydb_Maintenance.ManageMaintenanceTaskResponse)(nil),    // 14: Ydb.Maintenance.ManageMaintenanceTaskResponse
	(*Ydb_Maintenance.ManagePermitResponse)(nil),             // 15: Ydb.Maintenance.ManagePermitResponse
	(*Ydb_Maintenance.GetReadableActionReasonResponse)(nil),  // 16: Ydb.Maintenance.GetReadableActionReasonResponse
}
var file_draft_ydb_maintenance_v1_proto_depIdxs = []int32{
	0,  // 0: Ydb.Maintenance.V1.MaintenanceService.ListClusterNodes:input_type -> Ydb.Maintenance.ListClusterNodesRequest
	1,  // 1: Ydb.Maintenance.V1.MaintenanceService.CreateMaintenanceTask:input_type -> Ydb.Maintenance.CreateMaintenanceTaskRequest
	2,  // 2: Ydb.Maintenance.V1.MaintenanceService.RefreshMaintenanceTask:input_type -> Ydb.Maintenance.RefreshMaintenanceTaskRequest
	3,  // 3: Ydb.Maintenance.V1.MaintenanceService.GetMaintenanceTaskDetails:input_type -> Ydb.Maintenance.GetMaintenanceTaskRequest
	4,  // 4: Ydb.Maintenance.V1.MaintenanceService.ListMaintenanceTasks:input_type -> Ydb.Maintenance.ListMaintenanceTasksRequest
	5,  // 5: Ydb.Maintenance.V1.MaintenanceService.DropMaintenanceTask:input_type -> Ydb.Maintenance.DropMaintenanceTaskRequest
	6,  // 6: Ydb.Maintenance.V1.MaintenanceService.ProlongateMaintenanceTask:input_type -> Ydb.Maintenance.ProlongateMaintenanceTaskRequest
	7,  // 7: Ydb.Maintenance.V1.MaintenanceService.ReleasePermit:input_type -> Ydb.Maintenance.ReleasePermitRequest
	8,  // 8: Ydb.Maintenance.V1.MaintenanceService.ProlongatePermit:input_type -> Ydb.Maintenance.ProlongatePermitRequest
	9,  // 9: Ydb.Maintenance.V1.MaintenanceService.GetReadableActionReason:input_type -> Ydb.Maintenance.GetReadableActionReasonRequest
	10, // 10: Ydb.Maintenance.V1.MaintenanceService.ListClusterNodes:output_type -> Ydb.Maintenance.ListClusterNodesResponse
	11, // 11: Ydb.Maintenance.V1.MaintenanceService.CreateMaintenanceTask:output_type -> Ydb.Maintenance.MaintenanceTaskResponse
	11, // 12: Ydb.Maintenance.V1.MaintenanceService.RefreshMaintenanceTask:output_type -> Ydb.Maintenance.MaintenanceTaskResponse
	12, // 13: Ydb.Maintenance.V1.MaintenanceService.GetMaintenanceTaskDetails:output_type -> Ydb.Maintenance.GetMaintenanceTaskResponse
	13, // 14: Ydb.Maintenance.V1.MaintenanceService.ListMaintenanceTasks:output_type -> Ydb.Maintenance.ListMaintenanceTasksResponse
	14, // 15: Ydb.Maintenance.V1.MaintenanceService.DropMaintenanceTask:output_type -> Ydb.Maintenance.ManageMaintenanceTaskResponse
	14, // 16: Ydb.Maintenance.V1.MaintenanceService.ProlongateMaintenanceTask:output_type -> Ydb.Maintenance.ManageMaintenanceTaskResponse
	15, // 17: Ydb.Maintenance.V1.MaintenanceService.ReleasePermit:output_type -> Ydb.Maintenance.ManagePermitResponse
	15, // 18: Ydb.Maintenance.V1.MaintenanceService.ProlongatePermit:output_type -> Ydb.Maintenance.ManagePermitResponse
	16, // 19: Ydb.Maintenance.V1.MaintenanceService.GetReadableActionReason:output_type -> Ydb.Maintenance.GetReadableActionReasonResponse
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_draft_ydb_maintenance_v1_proto_init() }
func file_draft_ydb_maintenance_v1_proto_init() {
	if File_draft_ydb_maintenance_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_draft_ydb_maintenance_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_draft_ydb_maintenance_v1_proto_goTypes,
		DependencyIndexes: file_draft_ydb_maintenance_v1_proto_depIdxs,
	}.Build()
	File_draft_ydb_maintenance_v1_proto = out.File
	file_draft_ydb_maintenance_v1_proto_rawDesc = nil
	file_draft_ydb_maintenance_v1_proto_goTypes = nil
	file_draft_ydb_maintenance_v1_proto_depIdxs = nil
}
