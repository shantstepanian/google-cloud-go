// Copyright 2023 Google LLC
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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: google/cloud/deploy/v1/jobrun_notification_payload.proto

package deploypb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Payload proto for "clouddeploy.googleapis.com/jobrun_notification"
// Platform Log event that describes the failure to send JobRun resource update
// Pub/Sub notification.
type JobRunNotificationEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Debug message for when a notification fails to send.
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// The name of the `JobRun`.
	JobRun string `protobuf:"bytes,2,opt,name=job_run,json=jobRun,proto3" json:"job_run,omitempty"`
	// Unique identifier of the `DeliveryPipeline`.
	PipelineUid string `protobuf:"bytes,3,opt,name=pipeline_uid,json=pipelineUid,proto3" json:"pipeline_uid,omitempty"`
	// Unique identifier of the `Release`.
	ReleaseUid string `protobuf:"bytes,4,opt,name=release_uid,json=releaseUid,proto3" json:"release_uid,omitempty"`
	// Unique identifier of the `Rollout`.
	RolloutUid string `protobuf:"bytes,5,opt,name=rollout_uid,json=rolloutUid,proto3" json:"rollout_uid,omitempty"`
	// ID of the `Target`.
	TargetId string `protobuf:"bytes,6,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	// Type of this notification, e.g. for a Pub/Sub failure.
	Type Type `protobuf:"varint,7,opt,name=type,proto3,enum=google.cloud.deploy.v1.Type" json:"type,omitempty"`
}

func (x *JobRunNotificationEvent) Reset() {
	*x = JobRunNotificationEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_deploy_v1_jobrun_notification_payload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobRunNotificationEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobRunNotificationEvent) ProtoMessage() {}

func (x *JobRunNotificationEvent) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_deploy_v1_jobrun_notification_payload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobRunNotificationEvent.ProtoReflect.Descriptor instead.
func (*JobRunNotificationEvent) Descriptor() ([]byte, []int) {
	return file_google_cloud_deploy_v1_jobrun_notification_payload_proto_rawDescGZIP(), []int{0}
}

func (x *JobRunNotificationEvent) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *JobRunNotificationEvent) GetJobRun() string {
	if x != nil {
		return x.JobRun
	}
	return ""
}

func (x *JobRunNotificationEvent) GetPipelineUid() string {
	if x != nil {
		return x.PipelineUid
	}
	return ""
}

func (x *JobRunNotificationEvent) GetReleaseUid() string {
	if x != nil {
		return x.ReleaseUid
	}
	return ""
}

func (x *JobRunNotificationEvent) GetRolloutUid() string {
	if x != nil {
		return x.RolloutUid
	}
	return ""
}

func (x *JobRunNotificationEvent) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *JobRunNotificationEvent) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_TYPE_UNSPECIFIED
}

var File_google_cloud_deploy_v1_jobrun_notification_payload_proto protoreflect.FileDescriptor

var file_google_cloud_deploy_v1_jobrun_notification_payload_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x72, 0x75, 0x6e, 0x5f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e,
	0x76, 0x31, 0x1a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x02, 0x0a, 0x17, 0x4a,
	0x6f, 0x62, 0x52, 0x75, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x6a, 0x6f, 0x62, 0x5f, 0x72, 0x75, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6a, 0x6f, 0x62, 0x52, 0x75, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x55, 0x69, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x72, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x55, 0x69, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x72, 0x0a,
	0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x1e, 0x4a, 0x6f, 0x62,
	0x52, 0x75, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x70, 0x62, 0x3b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_deploy_v1_jobrun_notification_payload_proto_rawDescOnce sync.Once
	file_google_cloud_deploy_v1_jobrun_notification_payload_proto_rawDescData = file_google_cloud_deploy_v1_jobrun_notification_payload_proto_rawDesc
)

func file_google_cloud_deploy_v1_jobrun_notification_payload_proto_rawDescGZIP() []byte {
	file_google_cloud_deploy_v1_jobrun_notification_payload_proto_rawDescOnce.Do(func() {
		file_google_cloud_deploy_v1_jobrun_notification_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_deploy_v1_jobrun_notification_payload_proto_rawDescData)
	})
	return file_google_cloud_deploy_v1_jobrun_notification_payload_proto_rawDescData
}

var file_google_cloud_deploy_v1_jobrun_notification_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_deploy_v1_jobrun_notification_payload_proto_goTypes = []interface{}{
	(*JobRunNotificationEvent)(nil), // 0: google.cloud.deploy.v1.JobRunNotificationEvent
	(Type)(0),                       // 1: google.cloud.deploy.v1.Type
}
var file_google_cloud_deploy_v1_jobrun_notification_payload_proto_depIdxs = []int32{
	1, // 0: google.cloud.deploy.v1.JobRunNotificationEvent.type:type_name -> google.cloud.deploy.v1.Type
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_deploy_v1_jobrun_notification_payload_proto_init() }
func file_google_cloud_deploy_v1_jobrun_notification_payload_proto_init() {
	if File_google_cloud_deploy_v1_jobrun_notification_payload_proto != nil {
		return
	}
	file_google_cloud_deploy_v1_log_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_deploy_v1_jobrun_notification_payload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobRunNotificationEvent); i {
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
			RawDescriptor: file_google_cloud_deploy_v1_jobrun_notification_payload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_deploy_v1_jobrun_notification_payload_proto_goTypes,
		DependencyIndexes: file_google_cloud_deploy_v1_jobrun_notification_payload_proto_depIdxs,
		MessageInfos:      file_google_cloud_deploy_v1_jobrun_notification_payload_proto_msgTypes,
	}.Build()
	File_google_cloud_deploy_v1_jobrun_notification_payload_proto = out.File
	file_google_cloud_deploy_v1_jobrun_notification_payload_proto_rawDesc = nil
	file_google_cloud_deploy_v1_jobrun_notification_payload_proto_goTypes = nil
	file_google_cloud_deploy_v1_jobrun_notification_payload_proto_depIdxs = nil
}
