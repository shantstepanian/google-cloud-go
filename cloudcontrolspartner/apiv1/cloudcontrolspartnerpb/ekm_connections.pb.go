// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: google/cloud/cloudcontrolspartner/v1/ekm_connections.proto

package cloudcontrolspartnerpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The EKM connection state.
type EkmConnection_ConnectionState int32

const (
	// Unspecified EKM connection state
	EkmConnection_CONNECTION_STATE_UNSPECIFIED EkmConnection_ConnectionState = 0
	// Available EKM connection state
	EkmConnection_AVAILABLE EkmConnection_ConnectionState = 1
	// Not available EKM connection state
	EkmConnection_NOT_AVAILABLE EkmConnection_ConnectionState = 2
	// Error EKM connection state
	EkmConnection_ERROR EkmConnection_ConnectionState = 3
	// Permission denied EKM connection state
	EkmConnection_PERMISSION_DENIED EkmConnection_ConnectionState = 4
)

// Enum value maps for EkmConnection_ConnectionState.
var (
	EkmConnection_ConnectionState_name = map[int32]string{
		0: "CONNECTION_STATE_UNSPECIFIED",
		1: "AVAILABLE",
		2: "NOT_AVAILABLE",
		3: "ERROR",
		4: "PERMISSION_DENIED",
	}
	EkmConnection_ConnectionState_value = map[string]int32{
		"CONNECTION_STATE_UNSPECIFIED": 0,
		"AVAILABLE":                    1,
		"NOT_AVAILABLE":                2,
		"ERROR":                        3,
		"PERMISSION_DENIED":            4,
	}
)

func (x EkmConnection_ConnectionState) Enum() *EkmConnection_ConnectionState {
	p := new(EkmConnection_ConnectionState)
	*p = x
	return p
}

func (x EkmConnection_ConnectionState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EkmConnection_ConnectionState) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_enumTypes[0].Descriptor()
}

func (EkmConnection_ConnectionState) Type() protoreflect.EnumType {
	return &file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_enumTypes[0]
}

func (x EkmConnection_ConnectionState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EkmConnection_ConnectionState.Descriptor instead.
func (EkmConnection_ConnectionState) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_rawDescGZIP(), []int{2, 0}
}

// The EKM connections associated with a workload
type EkmConnections struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier. Format:
	// `organizations/{organization}/locations/{location}/customers/{customer}/workloads/{workload}/ekmConnections`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The EKM connections associated with the workload
	EkmConnections []*EkmConnection `protobuf:"bytes,2,rep,name=ekm_connections,json=ekmConnections,proto3" json:"ekm_connections,omitempty"`
}

func (x *EkmConnections) Reset() {
	*x = EkmConnections{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EkmConnections) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EkmConnections) ProtoMessage() {}

func (x *EkmConnections) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EkmConnections.ProtoReflect.Descriptor instead.
func (*EkmConnections) Descriptor() ([]byte, []int) {
	return file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_rawDescGZIP(), []int{0}
}

func (x *EkmConnections) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EkmConnections) GetEkmConnections() []*EkmConnection {
	if x != nil {
		return x.EkmConnections
	}
	return nil
}

// Request for getting the EKM connections associated with a workload
type GetEkmConnectionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Format:
	// `organizations/{organization}/locations/{location}/customers/{customer}/workloads/{workload}/ekmConnections`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetEkmConnectionsRequest) Reset() {
	*x = GetEkmConnectionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEkmConnectionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEkmConnectionsRequest) ProtoMessage() {}

func (x *GetEkmConnectionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEkmConnectionsRequest.ProtoReflect.Descriptor instead.
func (*GetEkmConnectionsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_rawDescGZIP(), []int{1}
}

func (x *GetEkmConnectionsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Details about the EKM connection
type EkmConnection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name of the EKM connection in the format:
	// projects/{project}/locations/{location}/ekmConnections/{ekm_connection}
	ConnectionName string `protobuf:"bytes,1,opt,name=connection_name,json=connectionName,proto3" json:"connection_name,omitempty"`
	// Output only. The connection state
	ConnectionState EkmConnection_ConnectionState `protobuf:"varint,2,opt,name=connection_state,json=connectionState,proto3,enum=google.cloud.cloudcontrolspartner.v1.EkmConnection_ConnectionState" json:"connection_state,omitempty"`
	// The connection error that occurred if any
	ConnectionError *EkmConnection_ConnectionError `protobuf:"bytes,3,opt,name=connection_error,json=connectionError,proto3" json:"connection_error,omitempty"`
}

func (x *EkmConnection) Reset() {
	*x = EkmConnection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EkmConnection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EkmConnection) ProtoMessage() {}

func (x *EkmConnection) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EkmConnection.ProtoReflect.Descriptor instead.
func (*EkmConnection) Descriptor() ([]byte, []int) {
	return file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_rawDescGZIP(), []int{2}
}

func (x *EkmConnection) GetConnectionName() string {
	if x != nil {
		return x.ConnectionName
	}
	return ""
}

func (x *EkmConnection) GetConnectionState() EkmConnection_ConnectionState {
	if x != nil {
		return x.ConnectionState
	}
	return EkmConnection_CONNECTION_STATE_UNSPECIFIED
}

func (x *EkmConnection) GetConnectionError() *EkmConnection_ConnectionError {
	if x != nil {
		return x.ConnectionError
	}
	return nil
}

// Information around the error that occurred if the connection state is
// anything other than available or unspecified
type EkmConnection_ConnectionError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The error domain for the error
	ErrorDomain string `protobuf:"bytes,1,opt,name=error_domain,json=errorDomain,proto3" json:"error_domain,omitempty"`
	// The error message for the error
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *EkmConnection_ConnectionError) Reset() {
	*x = EkmConnection_ConnectionError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EkmConnection_ConnectionError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EkmConnection_ConnectionError) ProtoMessage() {}

func (x *EkmConnection_ConnectionError) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EkmConnection_ConnectionError.ProtoReflect.Descriptor instead.
func (*EkmConnection_ConnectionError) Descriptor() ([]byte, []int) {
	return file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_rawDescGZIP(), []int{2, 0}
}

func (x *EkmConnection_ConnectionError) GetErrorDomain() string {
	if x != nil {
		return x.ErrorDomain
	}
	return ""
}

func (x *EkmConnection_ConnectionError) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto protoreflect.FileDescriptor

var file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6b, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae,
	0x02, 0x0a, 0x0e, 0x45, 0x6b, 0x6d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x5c, 0x0a, 0x0f, 0x65, 0x6b,
	0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6b, 0x6d, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x65, 0x6b, 0x6d, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0xa4, 0x01, 0xea, 0x41, 0xa0, 0x01, 0x0a,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x6b, 0x6d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x6a, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73,
	0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x7d, 0x2f, 0x77, 0x6f, 0x72, 0x6b,
	0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2f, 0x7b, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x7d,
	0x2f, 0x65, 0x6b, 0x6d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x6a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x45, 0x6b, 0x6d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4e, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3a, 0xe0, 0x41, 0x02, 0xfa, 0x41,
	0x34, 0x0a, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x6b, 0x6d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xf1, 0x03, 0x0a, 0x0d,
	0x45, 0x6b, 0x6d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a,
	0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x73, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6b, 0x6d, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x6e, 0x0a, 0x10, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6b, 0x6d,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x59, 0x0a, 0x0f, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x21,
	0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x77, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x4f, 0x4e,
	0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x41,
	0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f,
	0x54, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x09, 0x0a,
	0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x45, 0x52, 0x4d,
	0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x10, 0x04, 0x42,
	0x97, 0x02, 0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x45, 0x6b,
	0x6d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x5c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x70,
	0x62, 0xaa, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x50, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x73, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xea,
	0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x50, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_rawDescOnce sync.Once
	file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_rawDescData = file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_rawDesc
)

func file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_rawDescGZIP() []byte {
	file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_rawDescOnce.Do(func() {
		file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_rawDescData)
	})
	return file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_rawDescData
}

var file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_goTypes = []any{
	(EkmConnection_ConnectionState)(0),    // 0: google.cloud.cloudcontrolspartner.v1.EkmConnection.ConnectionState
	(*EkmConnections)(nil),                // 1: google.cloud.cloudcontrolspartner.v1.EkmConnections
	(*GetEkmConnectionsRequest)(nil),      // 2: google.cloud.cloudcontrolspartner.v1.GetEkmConnectionsRequest
	(*EkmConnection)(nil),                 // 3: google.cloud.cloudcontrolspartner.v1.EkmConnection
	(*EkmConnection_ConnectionError)(nil), // 4: google.cloud.cloudcontrolspartner.v1.EkmConnection.ConnectionError
}
var file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_depIdxs = []int32{
	3, // 0: google.cloud.cloudcontrolspartner.v1.EkmConnections.ekm_connections:type_name -> google.cloud.cloudcontrolspartner.v1.EkmConnection
	0, // 1: google.cloud.cloudcontrolspartner.v1.EkmConnection.connection_state:type_name -> google.cloud.cloudcontrolspartner.v1.EkmConnection.ConnectionState
	4, // 2: google.cloud.cloudcontrolspartner.v1.EkmConnection.connection_error:type_name -> google.cloud.cloudcontrolspartner.v1.EkmConnection.ConnectionError
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_init() }
func file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_init() {
	if File_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*EkmConnections); i {
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
		file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetEkmConnectionsRequest); i {
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
		file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*EkmConnection); i {
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
		file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*EkmConnection_ConnectionError); i {
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
			RawDescriptor: file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_goTypes,
		DependencyIndexes: file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_depIdxs,
		EnumInfos:         file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_enumTypes,
		MessageInfos:      file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_msgTypes,
	}.Build()
	File_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto = out.File
	file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_rawDesc = nil
	file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_goTypes = nil
	file_google_cloud_cloudcontrolspartner_v1_ekm_connections_proto_depIdxs = nil
}
