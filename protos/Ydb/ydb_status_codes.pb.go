// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.1
// source: protos/ydb_status_codes.proto

package Ydb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// reserved range [400000, 400999]
type StatusIds_StatusCode int32

const (
	StatusIds_STATUS_CODE_UNSPECIFIED StatusIds_StatusCode = 0
	StatusIds_SUCCESS                 StatusIds_StatusCode = 400000
	StatusIds_BAD_REQUEST             StatusIds_StatusCode = 400010
	StatusIds_UNAUTHORIZED            StatusIds_StatusCode = 400020
	StatusIds_INTERNAL_ERROR          StatusIds_StatusCode = 400030
	StatusIds_ABORTED                 StatusIds_StatusCode = 400040
	StatusIds_UNAVAILABLE             StatusIds_StatusCode = 400050
	StatusIds_OVERLOADED              StatusIds_StatusCode = 400060
	StatusIds_SCHEME_ERROR            StatusIds_StatusCode = 400070
	StatusIds_GENERIC_ERROR           StatusIds_StatusCode = 400080
	StatusIds_TIMEOUT                 StatusIds_StatusCode = 400090
	StatusIds_BAD_SESSION             StatusIds_StatusCode = 400100
	StatusIds_PRECONDITION_FAILED     StatusIds_StatusCode = 400120
	StatusIds_ALREADY_EXISTS          StatusIds_StatusCode = 400130
	StatusIds_NOT_FOUND               StatusIds_StatusCode = 400140
	StatusIds_SESSION_EXPIRED         StatusIds_StatusCode = 400150
	StatusIds_CANCELLED               StatusIds_StatusCode = 400160
	StatusIds_UNDETERMINED            StatusIds_StatusCode = 400170
	StatusIds_UNSUPPORTED             StatusIds_StatusCode = 400180
	StatusIds_SESSION_BUSY            StatusIds_StatusCode = 400190
	StatusIds_EXTERNAL_ERROR          StatusIds_StatusCode = 400200
)

// Enum value maps for StatusIds_StatusCode.
var (
	StatusIds_StatusCode_name = map[int32]string{
		0:      "STATUS_CODE_UNSPECIFIED",
		400000: "SUCCESS",
		400010: "BAD_REQUEST",
		400020: "UNAUTHORIZED",
		400030: "INTERNAL_ERROR",
		400040: "ABORTED",
		400050: "UNAVAILABLE",
		400060: "OVERLOADED",
		400070: "SCHEME_ERROR",
		400080: "GENERIC_ERROR",
		400090: "TIMEOUT",
		400100: "BAD_SESSION",
		400120: "PRECONDITION_FAILED",
		400130: "ALREADY_EXISTS",
		400140: "NOT_FOUND",
		400150: "SESSION_EXPIRED",
		400160: "CANCELLED",
		400170: "UNDETERMINED",
		400180: "UNSUPPORTED",
		400190: "SESSION_BUSY",
		400200: "EXTERNAL_ERROR",
	}
	StatusIds_StatusCode_value = map[string]int32{
		"STATUS_CODE_UNSPECIFIED": 0,
		"SUCCESS":                 400000,
		"BAD_REQUEST":             400010,
		"UNAUTHORIZED":            400020,
		"INTERNAL_ERROR":          400030,
		"ABORTED":                 400040,
		"UNAVAILABLE":             400050,
		"OVERLOADED":              400060,
		"SCHEME_ERROR":            400070,
		"GENERIC_ERROR":           400080,
		"TIMEOUT":                 400090,
		"BAD_SESSION":             400100,
		"PRECONDITION_FAILED":     400120,
		"ALREADY_EXISTS":          400130,
		"NOT_FOUND":               400140,
		"SESSION_EXPIRED":         400150,
		"CANCELLED":               400160,
		"UNDETERMINED":            400170,
		"UNSUPPORTED":             400180,
		"SESSION_BUSY":            400190,
		"EXTERNAL_ERROR":          400200,
	}
)

func (x StatusIds_StatusCode) Enum() *StatusIds_StatusCode {
	p := new(StatusIds_StatusCode)
	*p = x
	return p
}

func (x StatusIds_StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusIds_StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_ydb_status_codes_proto_enumTypes[0].Descriptor()
}

func (StatusIds_StatusCode) Type() protoreflect.EnumType {
	return &file_protos_ydb_status_codes_proto_enumTypes[0]
}

func (x StatusIds_StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusIds_StatusCode.Descriptor instead.
func (StatusIds_StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_protos_ydb_status_codes_proto_rawDescGZIP(), []int{0, 0}
}

type StatusIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatusIds) Reset() {
	*x = StatusIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_status_codes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusIds) ProtoMessage() {}

func (x *StatusIds) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_status_codes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusIds.ProtoReflect.Descriptor instead.
func (*StatusIds) Descriptor() ([]byte, []int) {
	return file_protos_ydb_status_codes_proto_rawDescGZIP(), []int{0}
}

var File_protos_ydb_status_codes_proto protoreflect.FileDescriptor

var file_protos_ydb_status_codes_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x59, 0x64, 0x62, 0x22, 0xbd, 0x03, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49,
	0x64, 0x73, 0x22, 0xaf, 0x03, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d,
	0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x80, 0xb5, 0x18, 0x12, 0x11, 0x0a,
	0x0b, 0x42, 0x41, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x8a, 0xb5, 0x18,
	0x12, 0x12, 0x0a, 0x0c, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44,
	0x10, 0x94, 0xb5, 0x18, 0x12, 0x14, 0x0a, 0x0e, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x9e, 0xb5, 0x18, 0x12, 0x0d, 0x0a, 0x07, 0x41, 0x42,
	0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0xa8, 0xb5, 0x18, 0x12, 0x11, 0x0a, 0x0b, 0x55, 0x4e, 0x41,
	0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0xb2, 0xb5, 0x18, 0x12, 0x10, 0x0a, 0x0a,
	0x4f, 0x56, 0x45, 0x52, 0x4c, 0x4f, 0x41, 0x44, 0x45, 0x44, 0x10, 0xbc, 0xb5, 0x18, 0x12, 0x12,
	0x0a, 0x0c, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xc6,
	0xb5, 0x18, 0x12, 0x13, 0x0a, 0x0d, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x49, 0x43, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0xd0, 0xb5, 0x18, 0x12, 0x0d, 0x0a, 0x07, 0x54, 0x49, 0x4d, 0x45, 0x4f,
	0x55, 0x54, 0x10, 0xda, 0xb5, 0x18, 0x12, 0x11, 0x0a, 0x0b, 0x42, 0x41, 0x44, 0x5f, 0x53, 0x45,
	0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0xe4, 0xb5, 0x18, 0x12, 0x19, 0x0a, 0x13, 0x50, 0x52, 0x45,
	0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44,
	0x10, 0xf8, 0xb5, 0x18, 0x12, 0x14, 0x0a, 0x0e, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f,
	0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x82, 0xb6, 0x18, 0x12, 0x0f, 0x0a, 0x09, 0x4e, 0x4f,
	0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x8c, 0xb6, 0x18, 0x12, 0x15, 0x0a, 0x0f, 0x53,
	0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x96,
	0xb6, 0x18, 0x12, 0x0f, 0x0a, 0x09, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10,
	0xa0, 0xb6, 0x18, 0x12, 0x12, 0x0a, 0x0c, 0x55, 0x4e, 0x44, 0x45, 0x54, 0x45, 0x52, 0x4d, 0x49,
	0x4e, 0x45, 0x44, 0x10, 0xaa, 0xb6, 0x18, 0x12, 0x11, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x55, 0x50,
	0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0xb4, 0xb6, 0x18, 0x12, 0x12, 0x0a, 0x0c, 0x53, 0x45,
	0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x42, 0x55, 0x53, 0x59, 0x10, 0xbe, 0xb6, 0x18, 0x12, 0x14,
	0x0a, 0x0e, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0xc8, 0xb6, 0x18, 0x42, 0x52, 0x0a, 0x0e, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x79, 0x64, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x11, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x67, 0x6f, 0x72, 0x6b, 0x61, 0x5a, 0x2f, 0x79, 0x64,
	0x62, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x59, 0x64, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_ydb_status_codes_proto_rawDescOnce sync.Once
	file_protos_ydb_status_codes_proto_rawDescData = file_protos_ydb_status_codes_proto_rawDesc
)

func file_protos_ydb_status_codes_proto_rawDescGZIP() []byte {
	file_protos_ydb_status_codes_proto_rawDescOnce.Do(func() {
		file_protos_ydb_status_codes_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_ydb_status_codes_proto_rawDescData)
	})
	return file_protos_ydb_status_codes_proto_rawDescData
}

var file_protos_ydb_status_codes_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_ydb_status_codes_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_ydb_status_codes_proto_goTypes = []interface{}{
	(StatusIds_StatusCode)(0), // 0: Ydb.StatusIds.StatusCode
	(*StatusIds)(nil),         // 1: Ydb.StatusIds
}
var file_protos_ydb_status_codes_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_ydb_status_codes_proto_init() }
func file_protos_ydb_status_codes_proto_init() {
	if File_protos_ydb_status_codes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_ydb_status_codes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusIds); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_ydb_status_codes_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_ydb_status_codes_proto_goTypes,
		DependencyIndexes: file_protos_ydb_status_codes_proto_depIdxs,
		EnumInfos:         file_protos_ydb_status_codes_proto_enumTypes,
		MessageInfos:      file_protos_ydb_status_codes_proto_msgTypes,
	}.Build()
	File_protos_ydb_status_codes_proto = out.File
	file_protos_ydb_status_codes_proto_rawDesc = nil
	file_protos_ydb_status_codes_proto_goTypes = nil
	file_protos_ydb_status_codes_proto_depIdxs = nil
}
