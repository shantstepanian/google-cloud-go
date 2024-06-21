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
// source: google/cloud/securitycenter/v2/log_entry.proto

package securitycenterpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// An individual entry in a log.
type LogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to LogEntry:
	//
	//	*LogEntry_CloudLoggingEntry
	LogEntry isLogEntry_LogEntry `protobuf_oneof:"log_entry"`
}

func (x *LogEntry) Reset() {
	*x = LogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v2_log_entry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntry) ProtoMessage() {}

func (x *LogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v2_log_entry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntry.ProtoReflect.Descriptor instead.
func (*LogEntry) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v2_log_entry_proto_rawDescGZIP(), []int{0}
}

func (m *LogEntry) GetLogEntry() isLogEntry_LogEntry {
	if m != nil {
		return m.LogEntry
	}
	return nil
}

func (x *LogEntry) GetCloudLoggingEntry() *CloudLoggingEntry {
	if x, ok := x.GetLogEntry().(*LogEntry_CloudLoggingEntry); ok {
		return x.CloudLoggingEntry
	}
	return nil
}

type isLogEntry_LogEntry interface {
	isLogEntry_LogEntry()
}

type LogEntry_CloudLoggingEntry struct {
	// An individual entry in a log stored in Cloud Logging.
	CloudLoggingEntry *CloudLoggingEntry `protobuf:"bytes,1,opt,name=cloud_logging_entry,json=cloudLoggingEntry,proto3,oneof"`
}

func (*LogEntry_CloudLoggingEntry) isLogEntry_LogEntry() {}

// Metadata taken from a [Cloud Logging
// LogEntry](https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry)
type CloudLoggingEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique identifier for the log entry.
	InsertId string `protobuf:"bytes,1,opt,name=insert_id,json=insertId,proto3" json:"insert_id,omitempty"`
	// The type of the log (part of `log_name`. `log_name` is the resource name of
	// the log to which this log entry belongs). For example:
	// `cloudresourcemanager.googleapis.com/activity` Note that this field is not
	// URL-encoded, unlike in `LogEntry`.
	LogId string `protobuf:"bytes,2,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
	// The organization, folder, or project of the monitored resource that
	// produced this log entry.
	ResourceContainer string `protobuf:"bytes,3,opt,name=resource_container,json=resourceContainer,proto3" json:"resource_container,omitempty"`
	// The time the event described by the log entry occurred.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *CloudLoggingEntry) Reset() {
	*x = CloudLoggingEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v2_log_entry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudLoggingEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudLoggingEntry) ProtoMessage() {}

func (x *CloudLoggingEntry) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v2_log_entry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudLoggingEntry.ProtoReflect.Descriptor instead.
func (*CloudLoggingEntry) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v2_log_entry_proto_rawDescGZIP(), []int{1}
}

func (x *CloudLoggingEntry) GetInsertId() string {
	if x != nil {
		return x.InsertId
	}
	return ""
}

func (x *CloudLoggingEntry) GetLogId() string {
	if x != nil {
		return x.LogId
	}
	return ""
}

func (x *CloudLoggingEntry) GetResourceContainer() string {
	if x != nil {
		return x.ResourceContainer
	}
	return ""
}

func (x *CloudLoggingEntry) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_google_cloud_securitycenter_v2_log_entry_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_v2_log_entry_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x32,
	0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x7c, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x63, 0x0a,
	0x13, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x48, 0x00, 0x52,
	0x11, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x42, 0x0b, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x22,
	0xb0, 0x01, 0x0a, 0x11, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74,
	0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0xe7, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x42, 0x0d, 0x4c, 0x6f, 0x67, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x32, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_securitycenter_v2_log_entry_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_v2_log_entry_proto_rawDescData = file_google_cloud_securitycenter_v2_log_entry_proto_rawDesc
)

func file_google_cloud_securitycenter_v2_log_entry_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_v2_log_entry_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_v2_log_entry_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_v2_log_entry_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_v2_log_entry_proto_rawDescData
}

var file_google_cloud_securitycenter_v2_log_entry_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_securitycenter_v2_log_entry_proto_goTypes = []any{
	(*LogEntry)(nil),              // 0: google.cloud.securitycenter.v2.LogEntry
	(*CloudLoggingEntry)(nil),     // 1: google.cloud.securitycenter.v2.CloudLoggingEntry
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_google_cloud_securitycenter_v2_log_entry_proto_depIdxs = []int32{
	1, // 0: google.cloud.securitycenter.v2.LogEntry.cloud_logging_entry:type_name -> google.cloud.securitycenter.v2.CloudLoggingEntry
	2, // 1: google.cloud.securitycenter.v2.CloudLoggingEntry.timestamp:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_securitycenter_v2_log_entry_proto_init() }
func file_google_cloud_securitycenter_v2_log_entry_proto_init() {
	if File_google_cloud_securitycenter_v2_log_entry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_securitycenter_v2_log_entry_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*LogEntry); i {
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
		file_google_cloud_securitycenter_v2_log_entry_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CloudLoggingEntry); i {
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
	file_google_cloud_securitycenter_v2_log_entry_proto_msgTypes[0].OneofWrappers = []any{
		(*LogEntry_CloudLoggingEntry)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_securitycenter_v2_log_entry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_v2_log_entry_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_v2_log_entry_proto_depIdxs,
		MessageInfos:      file_google_cloud_securitycenter_v2_log_entry_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_v2_log_entry_proto = out.File
	file_google_cloud_securitycenter_v2_log_entry_proto_rawDesc = nil
	file_google_cloud_securitycenter_v2_log_entry_proto_goTypes = nil
	file_google_cloud_securitycenter_v2_log_entry_proto_depIdxs = nil
}
