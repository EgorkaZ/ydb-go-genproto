// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: protos/ydb_discovery.proto

package Ydb_Discovery

import (
	Ydb_Operations "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Operations"
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

type ListEndpointsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database string   `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Service  []string `protobuf:"bytes,2,rep,name=service,proto3" json:"service,omitempty"`
}

func (x *ListEndpointsRequest) Reset() {
	*x = ListEndpointsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_discovery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEndpointsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEndpointsRequest) ProtoMessage() {}

func (x *ListEndpointsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_discovery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEndpointsRequest.ProtoReflect.Descriptor instead.
func (*ListEndpointsRequest) Descriptor() ([]byte, []int) {
	return file_protos_ydb_discovery_proto_rawDescGZIP(), []int{0}
}

func (x *ListEndpointsRequest) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *ListEndpointsRequest) GetService() []string {
	if x != nil {
		return x.Service
	}
	return nil
}

type EndpointInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is an address (usually fqdn) and port of this node's grpc endpoint
	Address    string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port       uint32   `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	LoadFactor float32  `protobuf:"fixed32,3,opt,name=load_factor,json=loadFactor,proto3" json:"load_factor,omitempty"`
	Ssl        bool     `protobuf:"varint,4,opt,name=ssl,proto3" json:"ssl,omitempty"`
	Service    []string `protobuf:"bytes,5,rep,name=service,proto3" json:"service,omitempty"`
	Location   string   `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	NodeId     uint32   `protobuf:"varint,7,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Optional ipv4 and/or ipv6 addresses of the endpoint, which clients may
	// use instead of a dns name in the address field.
	IpV4 []string `protobuf:"bytes,8,rep,name=ip_v4,json=ipV4,proto3" json:"ip_v4,omitempty"`
	IpV6 []string `protobuf:"bytes,9,rep,name=ip_v6,json=ipV6,proto3" json:"ip_v6,omitempty"`
	// Optional value for grpc.ssl_target_name_override option that must be
	// used when connecting to this endpoint. This may be specified when an ssl
	// endpoint is using certificate chain valid for a balancer hostname, and
	// not this specific node hostname.
	SslTargetNameOverride string `protobuf:"bytes,10,opt,name=ssl_target_name_override,json=sslTargetNameOverride,proto3" json:"ssl_target_name_override,omitempty"`
}

func (x *EndpointInfo) Reset() {
	*x = EndpointInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_discovery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointInfo) ProtoMessage() {}

func (x *EndpointInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_discovery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointInfo.ProtoReflect.Descriptor instead.
func (*EndpointInfo) Descriptor() ([]byte, []int) {
	return file_protos_ydb_discovery_proto_rawDescGZIP(), []int{1}
}

func (x *EndpointInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *EndpointInfo) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *EndpointInfo) GetLoadFactor() float32 {
	if x != nil {
		return x.LoadFactor
	}
	return 0
}

func (x *EndpointInfo) GetSsl() bool {
	if x != nil {
		return x.Ssl
	}
	return false
}

func (x *EndpointInfo) GetService() []string {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *EndpointInfo) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *EndpointInfo) GetNodeId() uint32 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *EndpointInfo) GetIpV4() []string {
	if x != nil {
		return x.IpV4
	}
	return nil
}

func (x *EndpointInfo) GetIpV6() []string {
	if x != nil {
		return x.IpV6
	}
	return nil
}

func (x *EndpointInfo) GetSslTargetNameOverride() string {
	if x != nil {
		return x.SslTargetNameOverride
	}
	return ""
}

type ListEndpointsResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoints    []*EndpointInfo `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	SelfLocation string          `protobuf:"bytes,2,opt,name=self_location,json=selfLocation,proto3" json:"self_location,omitempty"`
}

func (x *ListEndpointsResult) Reset() {
	*x = ListEndpointsResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_discovery_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEndpointsResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEndpointsResult) ProtoMessage() {}

func (x *ListEndpointsResult) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_discovery_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEndpointsResult.ProtoReflect.Descriptor instead.
func (*ListEndpointsResult) Descriptor() ([]byte, []int) {
	return file_protos_ydb_discovery_proto_rawDescGZIP(), []int{2}
}

func (x *ListEndpointsResult) GetEndpoints() []*EndpointInfo {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *ListEndpointsResult) GetSelfLocation() string {
	if x != nil {
		return x.SelfLocation
	}
	return ""
}

type ListEndpointsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation *Ydb_Operations.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
}

func (x *ListEndpointsResponse) Reset() {
	*x = ListEndpointsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_discovery_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEndpointsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEndpointsResponse) ProtoMessage() {}

func (x *ListEndpointsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_discovery_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEndpointsResponse.ProtoReflect.Descriptor instead.
func (*ListEndpointsResponse) Descriptor() ([]byte, []int) {
	return file_protos_ydb_discovery_proto_rawDescGZIP(), []int{3}
}

func (x *ListEndpointsResponse) GetOperation() *Ydb_Operations.Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

type WhoAmIRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Include user groups in response
	IncludeGroups bool `protobuf:"varint,1,opt,name=include_groups,json=includeGroups,proto3" json:"include_groups,omitempty"`
}

func (x *WhoAmIRequest) Reset() {
	*x = WhoAmIRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_discovery_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhoAmIRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhoAmIRequest) ProtoMessage() {}

func (x *WhoAmIRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_discovery_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhoAmIRequest.ProtoReflect.Descriptor instead.
func (*WhoAmIRequest) Descriptor() ([]byte, []int) {
	return file_protos_ydb_discovery_proto_rawDescGZIP(), []int{4}
}

func (x *WhoAmIRequest) GetIncludeGroups() bool {
	if x != nil {
		return x.IncludeGroups
	}
	return false
}

type WhoAmIResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User SID (Security ID)
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// List of group SIDs (Security IDs) for the user
	Groups []string `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *WhoAmIResult) Reset() {
	*x = WhoAmIResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_discovery_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhoAmIResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhoAmIResult) ProtoMessage() {}

func (x *WhoAmIResult) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_discovery_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhoAmIResult.ProtoReflect.Descriptor instead.
func (*WhoAmIResult) Descriptor() ([]byte, []int) {
	return file_protos_ydb_discovery_proto_rawDescGZIP(), []int{5}
}

func (x *WhoAmIResult) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *WhoAmIResult) GetGroups() []string {
	if x != nil {
		return x.Groups
	}
	return nil
}

type WhoAmIResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation *Ydb_Operations.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
}

func (x *WhoAmIResponse) Reset() {
	*x = WhoAmIResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_discovery_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhoAmIResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhoAmIResponse) ProtoMessage() {}

func (x *WhoAmIResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_discovery_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhoAmIResponse.ProtoReflect.Descriptor instead.
func (*WhoAmIResponse) Descriptor() ([]byte, []int) {
	return file_protos_ydb_discovery_proto_rawDescGZIP(), []int{6}
}

func (x *WhoAmIResponse) GetOperation() *Ydb_Operations.Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

var File_protos_ydb_discovery_proto protoreflect.FileDescriptor

var file_protos_ydb_discovery_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62, 0x5f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x59, 0x64,
	0x62, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x1a, 0x1a, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xa1, 0x02, 0x0a, 0x0c, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x66, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x6c, 0x6f, 0x61, 0x64, 0x46,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x73, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x03, 0x73, 0x73, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a,
	0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x69, 0x70, 0x5f, 0x76, 0x34, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x69, 0x70, 0x56, 0x34, 0x12, 0x13, 0x0a, 0x05, 0x69,
	0x70, 0x5f, 0x76, 0x36, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x69, 0x70, 0x56, 0x36,
	0x12, 0x37, 0x0a, 0x18, 0x73, 0x73, 0x6c, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x15, 0x73, 0x73, 0x6c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x22, 0x75, 0x0a, 0x13, 0x4c, 0x69, 0x73,
	0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x39, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x73,
	0x65, 0x6c, 0x66, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x6c, 0x66, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x50, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x59,
	0x64, 0x62, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x36, 0x0a, 0x0d, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x3a, 0x0a, 0x0c, 0x57, 0x68,
	0x6f, 0x41, 0x6d, 0x49, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x49, 0x0a, 0x0e, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x59, 0x64,
	0x62, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x66, 0x0a, 0x12, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x42, 0x0f, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x59, 0x64, 0x62, 0x5f, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_protos_ydb_discovery_proto_rawDescOnce sync.Once
	file_protos_ydb_discovery_proto_rawDescData = file_protos_ydb_discovery_proto_rawDesc
)

func file_protos_ydb_discovery_proto_rawDescGZIP() []byte {
	file_protos_ydb_discovery_proto_rawDescOnce.Do(func() {
		file_protos_ydb_discovery_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_ydb_discovery_proto_rawDescData)
	})
	return file_protos_ydb_discovery_proto_rawDescData
}

var file_protos_ydb_discovery_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protos_ydb_discovery_proto_goTypes = []interface{}{
	(*ListEndpointsRequest)(nil),     // 0: Ydb.Discovery.ListEndpointsRequest
	(*EndpointInfo)(nil),             // 1: Ydb.Discovery.EndpointInfo
	(*ListEndpointsResult)(nil),      // 2: Ydb.Discovery.ListEndpointsResult
	(*ListEndpointsResponse)(nil),    // 3: Ydb.Discovery.ListEndpointsResponse
	(*WhoAmIRequest)(nil),            // 4: Ydb.Discovery.WhoAmIRequest
	(*WhoAmIResult)(nil),             // 5: Ydb.Discovery.WhoAmIResult
	(*WhoAmIResponse)(nil),           // 6: Ydb.Discovery.WhoAmIResponse
	(*Ydb_Operations.Operation)(nil), // 7: Ydb.Operations.Operation
}
var file_protos_ydb_discovery_proto_depIdxs = []int32{
	1, // 0: Ydb.Discovery.ListEndpointsResult.endpoints:type_name -> Ydb.Discovery.EndpointInfo
	7, // 1: Ydb.Discovery.ListEndpointsResponse.operation:type_name -> Ydb.Operations.Operation
	7, // 2: Ydb.Discovery.WhoAmIResponse.operation:type_name -> Ydb.Operations.Operation
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_ydb_discovery_proto_init() }
func file_protos_ydb_discovery_proto_init() {
	if File_protos_ydb_discovery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_ydb_discovery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEndpointsRequest); i {
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
		file_protos_ydb_discovery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointInfo); i {
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
		file_protos_ydb_discovery_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEndpointsResult); i {
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
		file_protos_ydb_discovery_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEndpointsResponse); i {
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
		file_protos_ydb_discovery_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhoAmIRequest); i {
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
		file_protos_ydb_discovery_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhoAmIResult); i {
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
		file_protos_ydb_discovery_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhoAmIResponse); i {
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
			RawDescriptor: file_protos_ydb_discovery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_ydb_discovery_proto_goTypes,
		DependencyIndexes: file_protos_ydb_discovery_proto_depIdxs,
		MessageInfos:      file_protos_ydb_discovery_proto_msgTypes,
	}.Build()
	File_protos_ydb_discovery_proto = out.File
	file_protos_ydb_discovery_proto_rawDesc = nil
	file_protos_ydb_discovery_proto_goTypes = nil
	file_protos_ydb_discovery_proto_depIdxs = nil
}
