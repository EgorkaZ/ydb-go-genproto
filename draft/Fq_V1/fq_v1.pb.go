// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.1
// source: draft/fq_v1.proto

package Fq_V1

import (
	Fq "github.com/ydb-platform/ydb-go-genproto/draft/protos/Fq"
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

var File_draft_fq_v1_proto protoreflect.FileDescriptor

var file_draft_fq_v1_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x66, 0x71, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x56, 0x31, 0x1a, 0x15, 0x64, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd6, 0x0f,
	0x0a, 0x15, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x22, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x46, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x56, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x22,
	0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0d, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x24, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0b, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x22, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x46, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56,
	0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x22, 0x2e,
	0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x23, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x46, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x24, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4d, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x1f, 0x2e, 0x46, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x46,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56,
	0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x22, 0x2e,
	0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x46, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a,
	0x0f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x26, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x6b, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65,
	0x0a, 0x10, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x46, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x46, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x28, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x0e,
	0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25,
	0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x24,
	0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x0c, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x23, 0x2e, 0x46, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x2e, 0x46, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0d, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x24, 0x2e, 0x46, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x24, 0x2e, 0x46, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4d, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x62,
	0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x67, 0x6f,
	0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x72, 0x61, 0x66, 0x74, 0x2f,
	0x46, 0x71, 0x5f, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_draft_fq_v1_proto_goTypes = []interface{}{
	(*Fq.CreateQueryRequest)(nil),         // 0: FederatedQuery.CreateQueryRequest
	(*Fq.ListQueriesRequest)(nil),         // 1: FederatedQuery.ListQueriesRequest
	(*Fq.DescribeQueryRequest)(nil),       // 2: FederatedQuery.DescribeQueryRequest
	(*Fq.GetQueryStatusRequest)(nil),      // 3: FederatedQuery.GetQueryStatusRequest
	(*Fq.ModifyQueryRequest)(nil),         // 4: FederatedQuery.ModifyQueryRequest
	(*Fq.DeleteQueryRequest)(nil),         // 5: FederatedQuery.DeleteQueryRequest
	(*Fq.ControlQueryRequest)(nil),        // 6: FederatedQuery.ControlQueryRequest
	(*Fq.GetResultDataRequest)(nil),       // 7: FederatedQuery.GetResultDataRequest
	(*Fq.ListJobsRequest)(nil),            // 8: FederatedQuery.ListJobsRequest
	(*Fq.DescribeJobRequest)(nil),         // 9: FederatedQuery.DescribeJobRequest
	(*Fq.CreateConnectionRequest)(nil),    // 10: FederatedQuery.CreateConnectionRequest
	(*Fq.ListConnectionsRequest)(nil),     // 11: FederatedQuery.ListConnectionsRequest
	(*Fq.DescribeConnectionRequest)(nil),  // 12: FederatedQuery.DescribeConnectionRequest
	(*Fq.ModifyConnectionRequest)(nil),    // 13: FederatedQuery.ModifyConnectionRequest
	(*Fq.DeleteConnectionRequest)(nil),    // 14: FederatedQuery.DeleteConnectionRequest
	(*Fq.TestConnectionRequest)(nil),      // 15: FederatedQuery.TestConnectionRequest
	(*Fq.CreateBindingRequest)(nil),       // 16: FederatedQuery.CreateBindingRequest
	(*Fq.ListBindingsRequest)(nil),        // 17: FederatedQuery.ListBindingsRequest
	(*Fq.DescribeBindingRequest)(nil),     // 18: FederatedQuery.DescribeBindingRequest
	(*Fq.ModifyBindingRequest)(nil),       // 19: FederatedQuery.ModifyBindingRequest
	(*Fq.DeleteBindingRequest)(nil),       // 20: FederatedQuery.DeleteBindingRequest
	(*Fq.CreateQueryResponse)(nil),        // 21: FederatedQuery.CreateQueryResponse
	(*Fq.ListQueriesResponse)(nil),        // 22: FederatedQuery.ListQueriesResponse
	(*Fq.DescribeQueryResponse)(nil),      // 23: FederatedQuery.DescribeQueryResponse
	(*Fq.GetQueryStatusResponse)(nil),     // 24: FederatedQuery.GetQueryStatusResponse
	(*Fq.ModifyQueryResponse)(nil),        // 25: FederatedQuery.ModifyQueryResponse
	(*Fq.DeleteQueryResponse)(nil),        // 26: FederatedQuery.DeleteQueryResponse
	(*Fq.ControlQueryResponse)(nil),       // 27: FederatedQuery.ControlQueryResponse
	(*Fq.GetResultDataResponse)(nil),      // 28: FederatedQuery.GetResultDataResponse
	(*Fq.ListJobsResponse)(nil),           // 29: FederatedQuery.ListJobsResponse
	(*Fq.DescribeJobResponse)(nil),        // 30: FederatedQuery.DescribeJobResponse
	(*Fq.CreateConnectionResponse)(nil),   // 31: FederatedQuery.CreateConnectionResponse
	(*Fq.ListConnectionsResponse)(nil),    // 32: FederatedQuery.ListConnectionsResponse
	(*Fq.DescribeConnectionResponse)(nil), // 33: FederatedQuery.DescribeConnectionResponse
	(*Fq.ModifyConnectionResponse)(nil),   // 34: FederatedQuery.ModifyConnectionResponse
	(*Fq.DeleteConnectionResponse)(nil),   // 35: FederatedQuery.DeleteConnectionResponse
	(*Fq.TestConnectionResponse)(nil),     // 36: FederatedQuery.TestConnectionResponse
	(*Fq.CreateBindingResponse)(nil),      // 37: FederatedQuery.CreateBindingResponse
	(*Fq.ListBindingsResponse)(nil),       // 38: FederatedQuery.ListBindingsResponse
	(*Fq.DescribeBindingResponse)(nil),    // 39: FederatedQuery.DescribeBindingResponse
	(*Fq.ModifyBindingResponse)(nil),      // 40: FederatedQuery.ModifyBindingResponse
	(*Fq.DeleteBindingResponse)(nil),      // 41: FederatedQuery.DeleteBindingResponse
}
var file_draft_fq_v1_proto_depIdxs = []int32{
	0,  // 0: FederatedQuery.V1.FederatedQueryService.CreateQuery:input_type -> FederatedQuery.CreateQueryRequest
	1,  // 1: FederatedQuery.V1.FederatedQueryService.ListQueries:input_type -> FederatedQuery.ListQueriesRequest
	2,  // 2: FederatedQuery.V1.FederatedQueryService.DescribeQuery:input_type -> FederatedQuery.DescribeQueryRequest
	3,  // 3: FederatedQuery.V1.FederatedQueryService.GetQueryStatus:input_type -> FederatedQuery.GetQueryStatusRequest
	4,  // 4: FederatedQuery.V1.FederatedQueryService.ModifyQuery:input_type -> FederatedQuery.ModifyQueryRequest
	5,  // 5: FederatedQuery.V1.FederatedQueryService.DeleteQuery:input_type -> FederatedQuery.DeleteQueryRequest
	6,  // 6: FederatedQuery.V1.FederatedQueryService.ControlQuery:input_type -> FederatedQuery.ControlQueryRequest
	7,  // 7: FederatedQuery.V1.FederatedQueryService.GetResultData:input_type -> FederatedQuery.GetResultDataRequest
	8,  // 8: FederatedQuery.V1.FederatedQueryService.ListJobs:input_type -> FederatedQuery.ListJobsRequest
	9,  // 9: FederatedQuery.V1.FederatedQueryService.DescribeJob:input_type -> FederatedQuery.DescribeJobRequest
	10, // 10: FederatedQuery.V1.FederatedQueryService.CreateConnection:input_type -> FederatedQuery.CreateConnectionRequest
	11, // 11: FederatedQuery.V1.FederatedQueryService.ListConnections:input_type -> FederatedQuery.ListConnectionsRequest
	12, // 12: FederatedQuery.V1.FederatedQueryService.DescribeConnection:input_type -> FederatedQuery.DescribeConnectionRequest
	13, // 13: FederatedQuery.V1.FederatedQueryService.ModifyConnection:input_type -> FederatedQuery.ModifyConnectionRequest
	14, // 14: FederatedQuery.V1.FederatedQueryService.DeleteConnection:input_type -> FederatedQuery.DeleteConnectionRequest
	15, // 15: FederatedQuery.V1.FederatedQueryService.TestConnection:input_type -> FederatedQuery.TestConnectionRequest
	16, // 16: FederatedQuery.V1.FederatedQueryService.CreateBinding:input_type -> FederatedQuery.CreateBindingRequest
	17, // 17: FederatedQuery.V1.FederatedQueryService.ListBindings:input_type -> FederatedQuery.ListBindingsRequest
	18, // 18: FederatedQuery.V1.FederatedQueryService.DescribeBinding:input_type -> FederatedQuery.DescribeBindingRequest
	19, // 19: FederatedQuery.V1.FederatedQueryService.ModifyBinding:input_type -> FederatedQuery.ModifyBindingRequest
	20, // 20: FederatedQuery.V1.FederatedQueryService.DeleteBinding:input_type -> FederatedQuery.DeleteBindingRequest
	21, // 21: FederatedQuery.V1.FederatedQueryService.CreateQuery:output_type -> FederatedQuery.CreateQueryResponse
	22, // 22: FederatedQuery.V1.FederatedQueryService.ListQueries:output_type -> FederatedQuery.ListQueriesResponse
	23, // 23: FederatedQuery.V1.FederatedQueryService.DescribeQuery:output_type -> FederatedQuery.DescribeQueryResponse
	24, // 24: FederatedQuery.V1.FederatedQueryService.GetQueryStatus:output_type -> FederatedQuery.GetQueryStatusResponse
	25, // 25: FederatedQuery.V1.FederatedQueryService.ModifyQuery:output_type -> FederatedQuery.ModifyQueryResponse
	26, // 26: FederatedQuery.V1.FederatedQueryService.DeleteQuery:output_type -> FederatedQuery.DeleteQueryResponse
	27, // 27: FederatedQuery.V1.FederatedQueryService.ControlQuery:output_type -> FederatedQuery.ControlQueryResponse
	28, // 28: FederatedQuery.V1.FederatedQueryService.GetResultData:output_type -> FederatedQuery.GetResultDataResponse
	29, // 29: FederatedQuery.V1.FederatedQueryService.ListJobs:output_type -> FederatedQuery.ListJobsResponse
	30, // 30: FederatedQuery.V1.FederatedQueryService.DescribeJob:output_type -> FederatedQuery.DescribeJobResponse
	31, // 31: FederatedQuery.V1.FederatedQueryService.CreateConnection:output_type -> FederatedQuery.CreateConnectionResponse
	32, // 32: FederatedQuery.V1.FederatedQueryService.ListConnections:output_type -> FederatedQuery.ListConnectionsResponse
	33, // 33: FederatedQuery.V1.FederatedQueryService.DescribeConnection:output_type -> FederatedQuery.DescribeConnectionResponse
	34, // 34: FederatedQuery.V1.FederatedQueryService.ModifyConnection:output_type -> FederatedQuery.ModifyConnectionResponse
	35, // 35: FederatedQuery.V1.FederatedQueryService.DeleteConnection:output_type -> FederatedQuery.DeleteConnectionResponse
	36, // 36: FederatedQuery.V1.FederatedQueryService.TestConnection:output_type -> FederatedQuery.TestConnectionResponse
	37, // 37: FederatedQuery.V1.FederatedQueryService.CreateBinding:output_type -> FederatedQuery.CreateBindingResponse
	38, // 38: FederatedQuery.V1.FederatedQueryService.ListBindings:output_type -> FederatedQuery.ListBindingsResponse
	39, // 39: FederatedQuery.V1.FederatedQueryService.DescribeBinding:output_type -> FederatedQuery.DescribeBindingResponse
	40, // 40: FederatedQuery.V1.FederatedQueryService.ModifyBinding:output_type -> FederatedQuery.ModifyBindingResponse
	41, // 41: FederatedQuery.V1.FederatedQueryService.DeleteBinding:output_type -> FederatedQuery.DeleteBindingResponse
	21, // [21:42] is the sub-list for method output_type
	0,  // [0:21] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_draft_fq_v1_proto_init() }
func file_draft_fq_v1_proto_init() {
	if File_draft_fq_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_draft_fq_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_draft_fq_v1_proto_goTypes,
		DependencyIndexes: file_draft_fq_v1_proto_depIdxs,
	}.Build()
	File_draft_fq_v1_proto = out.File
	file_draft_fq_v1_proto_rawDesc = nil
	file_draft_fq_v1_proto_goTypes = nil
	file_draft_fq_v1_proto_depIdxs = nil
}
