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
// source: google/cloud/aiplatform/v1beta1/model_garden_service.proto

package aiplatformpb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// View enumeration of PublisherModel.
type PublisherModelView int32

const (
	// The default / unset value. The API will default to the BASIC view.
	PublisherModelView_PUBLISHER_MODEL_VIEW_UNSPECIFIED PublisherModelView = 0
	// Include basic metadata about the publisher model, but not the full
	// contents.
	PublisherModelView_PUBLISHER_MODEL_VIEW_BASIC PublisherModelView = 1
	// Include everything.
	PublisherModelView_PUBLISHER_MODEL_VIEW_FULL PublisherModelView = 2
	// Include: VersionId, ModelVersionExternalName, and SupportedActions.
	PublisherModelView_PUBLISHER_MODEL_VERSION_VIEW_BASIC PublisherModelView = 3
)

// Enum value maps for PublisherModelView.
var (
	PublisherModelView_name = map[int32]string{
		0: "PUBLISHER_MODEL_VIEW_UNSPECIFIED",
		1: "PUBLISHER_MODEL_VIEW_BASIC",
		2: "PUBLISHER_MODEL_VIEW_FULL",
		3: "PUBLISHER_MODEL_VERSION_VIEW_BASIC",
	}
	PublisherModelView_value = map[string]int32{
		"PUBLISHER_MODEL_VIEW_UNSPECIFIED":   0,
		"PUBLISHER_MODEL_VIEW_BASIC":         1,
		"PUBLISHER_MODEL_VIEW_FULL":          2,
		"PUBLISHER_MODEL_VERSION_VIEW_BASIC": 3,
	}
)

func (x PublisherModelView) Enum() *PublisherModelView {
	p := new(PublisherModelView)
	*p = x
	return p
}

func (x PublisherModelView) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PublisherModelView) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_enumTypes[0].Descriptor()
}

func (PublisherModelView) Type() protoreflect.EnumType {
	return &file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_enumTypes[0]
}

func (x PublisherModelView) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PublisherModelView.Descriptor instead.
func (PublisherModelView) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_rawDescGZIP(), []int{0}
}

// Request message for
// [ModelGardenService.GetPublisherModel][google.cloud.aiplatform.v1beta1.ModelGardenService.GetPublisherModel]
type GetPublisherModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the PublisherModel resource.
	// Format:
	// `publishers/{publisher}/models/{publisher_model}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. The IETF BCP-47 language code representing the language in which
	// the publisher model's text information should be written in (see go/bcp47).
	LanguageCode string `protobuf:"bytes,2,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// Optional. PublisherModel view specifying which fields to read.
	View PublisherModelView `protobuf:"varint,3,opt,name=view,proto3,enum=google.cloud.aiplatform.v1beta1.PublisherModelView" json:"view,omitempty"`
}

func (x *GetPublisherModelRequest) Reset() {
	*x = GetPublisherModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPublisherModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPublisherModelRequest) ProtoMessage() {}

func (x *GetPublisherModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPublisherModelRequest.ProtoReflect.Descriptor instead.
func (*GetPublisherModelRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetPublisherModelRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetPublisherModelRequest) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

func (x *GetPublisherModelRequest) GetView() PublisherModelView {
	if x != nil {
		return x.View
	}
	return PublisherModelView_PUBLISHER_MODEL_VIEW_UNSPECIFIED
}

// Request message for
// [ModelGardenService.ListPublisherModels][google.cloud.aiplatform.v1beta1.ModelGardenService.ListPublisherModels].
type ListPublisherModelsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the Publisher from which to list the PublisherModels.
	// Format: `publishers/{publisher}`
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. The standard list filter.
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	// Optional. The standard list page size.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. The standard list page token.
	// Typically obtained via
	// [ListPublisherModelsResponse.next_page_token][google.cloud.aiplatform.v1beta1.ListPublisherModelsResponse.next_page_token]
	// of the previous
	// [ModelGardenService.ListPublisherModels][google.cloud.aiplatform.v1beta1.ModelGardenService.ListPublisherModels]
	// call.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Optional. PublisherModel view specifying which fields to read.
	View PublisherModelView `protobuf:"varint,5,opt,name=view,proto3,enum=google.cloud.aiplatform.v1beta1.PublisherModelView" json:"view,omitempty"`
	// Optional. A comma-separated list of fields to order by, sorted in ascending
	// order. Use "desc" after a field name for descending.
	OrderBy string `protobuf:"bytes,6,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	// Optional. The IETF BCP-47 language code representing the language in which
	// the publisher models' text information should be written in (see go/bcp47).
	// If not set, by default English (en).
	LanguageCode string `protobuf:"bytes,7,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
}

func (x *ListPublisherModelsRequest) Reset() {
	*x = ListPublisherModelsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPublisherModelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPublisherModelsRequest) ProtoMessage() {}

func (x *ListPublisherModelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPublisherModelsRequest.ProtoReflect.Descriptor instead.
func (*ListPublisherModelsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListPublisherModelsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListPublisherModelsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *ListPublisherModelsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListPublisherModelsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListPublisherModelsRequest) GetView() PublisherModelView {
	if x != nil {
		return x.View
	}
	return PublisherModelView_PUBLISHER_MODEL_VIEW_UNSPECIFIED
}

func (x *ListPublisherModelsRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *ListPublisherModelsRequest) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

// Response message for
// [ModelGardenService.ListPublisherModels][google.cloud.aiplatform.v1beta1.ModelGardenService.ListPublisherModels].
type ListPublisherModelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of PublisherModels in the requested page.
	PublisherModels []*PublisherModel `protobuf:"bytes,1,rep,name=publisher_models,json=publisherModels,proto3" json:"publisher_models,omitempty"`
	// A token to retrieve next page of results.
	// Pass to [ListPublisherModels.page_token][] to obtain that page.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListPublisherModelsResponse) Reset() {
	*x = ListPublisherModelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPublisherModelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPublisherModelsResponse) ProtoMessage() {}

func (x *ListPublisherModelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPublisherModelsResponse.ProtoReflect.Descriptor instead.
func (*ListPublisherModelsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListPublisherModelsResponse) GetPublisherModels() []*PublisherModel {
	if x != nil {
		return x.PublisherModels
	}
	return nil
}

func (x *ListPublisherModelsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_google_cloud_aiplatform_v1beta1_model_garden_service_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x30, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2a, 0x0a, 0x28, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0d, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x4c, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x56, 0x69, 0x65, 0x77, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x04, 0x76, 0x69,
	0x65, 0x77, 0x22, 0xb4, 0x02, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x01, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03,
	0xe0, 0x41, 0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x4c, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x56, 0x69, 0x65, 0x77, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x04, 0x76, 0x69, 0x65, 0x77, 0x12,
	0x1e, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12,
	0x28, 0x0a, 0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0c, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x22, 0xa1, 0x01, 0x0a, 0x1b, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x10, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x52, 0x0f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x2a, 0xa1, 0x01,
	0x0a, 0x12, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x56, 0x69, 0x65, 0x77, 0x12, 0x24, 0x0a, 0x20, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x45,
	0x52, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x4c, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x55,
	0x42, 0x4c, 0x49, 0x53, 0x48, 0x45, 0x52, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x4c, 0x5f, 0x56, 0x49,
	0x45, 0x57, 0x5f, 0x42, 0x41, 0x53, 0x49, 0x43, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x55,
	0x42, 0x4c, 0x49, 0x53, 0x48, 0x45, 0x52, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x4c, 0x5f, 0x56, 0x49,
	0x45, 0x57, 0x5f, 0x46, 0x55, 0x4c, 0x4c, 0x10, 0x02, 0x12, 0x26, 0x0a, 0x22, 0x50, 0x55, 0x42,
	0x4c, 0x49, 0x53, 0x48, 0x45, 0x52, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x4c, 0x5f, 0x56, 0x45, 0x52,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x42, 0x41, 0x53, 0x49, 0x43, 0x10,
	0x03, 0x32, 0xe6, 0x03, 0x0a, 0x12, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x47, 0x61, 0x72, 0x64, 0x65,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xb5, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x39,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x34, 0xda, 0x41, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12, 0x25, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x2a, 0x7d,
	0x12, 0xc8, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x36, 0xda, 0x41, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x27, 0x12, 0x25, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72,
	0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x4d, 0xca, 0x41, 0x19,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0xee, 0x01, 0x0a, 0x23, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x42, 0x17, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70,
	0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x70, 0x62, 0x3b, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x70, 0x62, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0x42,
	0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_rawDescData = file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_rawDesc
)

func file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_rawDescData
}

var file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_goTypes = []any{
	(PublisherModelView)(0),             // 0: google.cloud.aiplatform.v1beta1.PublisherModelView
	(*GetPublisherModelRequest)(nil),    // 1: google.cloud.aiplatform.v1beta1.GetPublisherModelRequest
	(*ListPublisherModelsRequest)(nil),  // 2: google.cloud.aiplatform.v1beta1.ListPublisherModelsRequest
	(*ListPublisherModelsResponse)(nil), // 3: google.cloud.aiplatform.v1beta1.ListPublisherModelsResponse
	(*PublisherModel)(nil),              // 4: google.cloud.aiplatform.v1beta1.PublisherModel
}
var file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_depIdxs = []int32{
	0, // 0: google.cloud.aiplatform.v1beta1.GetPublisherModelRequest.view:type_name -> google.cloud.aiplatform.v1beta1.PublisherModelView
	0, // 1: google.cloud.aiplatform.v1beta1.ListPublisherModelsRequest.view:type_name -> google.cloud.aiplatform.v1beta1.PublisherModelView
	4, // 2: google.cloud.aiplatform.v1beta1.ListPublisherModelsResponse.publisher_models:type_name -> google.cloud.aiplatform.v1beta1.PublisherModel
	1, // 3: google.cloud.aiplatform.v1beta1.ModelGardenService.GetPublisherModel:input_type -> google.cloud.aiplatform.v1beta1.GetPublisherModelRequest
	2, // 4: google.cloud.aiplatform.v1beta1.ModelGardenService.ListPublisherModels:input_type -> google.cloud.aiplatform.v1beta1.ListPublisherModelsRequest
	4, // 5: google.cloud.aiplatform.v1beta1.ModelGardenService.GetPublisherModel:output_type -> google.cloud.aiplatform.v1beta1.PublisherModel
	3, // 6: google.cloud.aiplatform.v1beta1.ModelGardenService.ListPublisherModels:output_type -> google.cloud.aiplatform.v1beta1.ListPublisherModelsResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_init() }
func file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_init() {
	if File_google_cloud_aiplatform_v1beta1_model_garden_service_proto != nil {
		return
	}
	file_google_cloud_aiplatform_v1beta1_publisher_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetPublisherModelRequest); i {
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
		file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListPublisherModelsRequest); i {
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
		file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListPublisherModelsResponse); i {
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
			RawDescriptor: file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_depIdxs,
		EnumInfos:         file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_enumTypes,
		MessageInfos:      file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1beta1_model_garden_service_proto = out.File
	file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_goTypes = nil
	file_google_cloud_aiplatform_v1beta1_model_garden_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ModelGardenServiceClient is the client API for ModelGardenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ModelGardenServiceClient interface {
	// Gets a Model Garden publisher model.
	GetPublisherModel(ctx context.Context, in *GetPublisherModelRequest, opts ...grpc.CallOption) (*PublisherModel, error)
	// Lists publisher models in Model Garden.
	ListPublisherModels(ctx context.Context, in *ListPublisherModelsRequest, opts ...grpc.CallOption) (*ListPublisherModelsResponse, error)
}

type modelGardenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewModelGardenServiceClient(cc grpc.ClientConnInterface) ModelGardenServiceClient {
	return &modelGardenServiceClient{cc}
}

func (c *modelGardenServiceClient) GetPublisherModel(ctx context.Context, in *GetPublisherModelRequest, opts ...grpc.CallOption) (*PublisherModel, error) {
	out := new(PublisherModel)
	err := c.cc.Invoke(ctx, "/google.cloud.aiplatform.v1beta1.ModelGardenService/GetPublisherModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelGardenServiceClient) ListPublisherModels(ctx context.Context, in *ListPublisherModelsRequest, opts ...grpc.CallOption) (*ListPublisherModelsResponse, error) {
	out := new(ListPublisherModelsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.aiplatform.v1beta1.ModelGardenService/ListPublisherModels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelGardenServiceServer is the server API for ModelGardenService service.
type ModelGardenServiceServer interface {
	// Gets a Model Garden publisher model.
	GetPublisherModel(context.Context, *GetPublisherModelRequest) (*PublisherModel, error)
	// Lists publisher models in Model Garden.
	ListPublisherModels(context.Context, *ListPublisherModelsRequest) (*ListPublisherModelsResponse, error)
}

// UnimplementedModelGardenServiceServer can be embedded to have forward compatible implementations.
type UnimplementedModelGardenServiceServer struct {
}

func (*UnimplementedModelGardenServiceServer) GetPublisherModel(context.Context, *GetPublisherModelRequest) (*PublisherModel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublisherModel not implemented")
}
func (*UnimplementedModelGardenServiceServer) ListPublisherModels(context.Context, *ListPublisherModelsRequest) (*ListPublisherModelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPublisherModels not implemented")
}

func RegisterModelGardenServiceServer(s *grpc.Server, srv ModelGardenServiceServer) {
	s.RegisterService(&_ModelGardenService_serviceDesc, srv)
}

func _ModelGardenService_GetPublisherModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublisherModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelGardenServiceServer).GetPublisherModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.aiplatform.v1beta1.ModelGardenService/GetPublisherModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelGardenServiceServer).GetPublisherModel(ctx, req.(*GetPublisherModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelGardenService_ListPublisherModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPublisherModelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelGardenServiceServer).ListPublisherModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.aiplatform.v1beta1.ModelGardenService/ListPublisherModels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelGardenServiceServer).ListPublisherModels(ctx, req.(*ListPublisherModelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ModelGardenService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.aiplatform.v1beta1.ModelGardenService",
	HandlerType: (*ModelGardenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPublisherModel",
			Handler:    _ModelGardenService_GetPublisherModel_Handler,
		},
		{
			MethodName: "ListPublisherModels",
			Handler:    _ModelGardenService_ListPublisherModels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/aiplatform/v1beta1/model_garden_service.proto",
}
