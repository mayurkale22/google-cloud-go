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
// source: google/cloud/discoveryengine/v1alpha/completion_service.proto

package discoveryenginepb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
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

// Request message for
// [CompletionService.CompleteQuery][google.cloud.discoveryengine.v1alpha.CompletionService.CompleteQuery]
// method.
type CompleteQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The parent data store resource name for which the completion is
	// performed, such as
	// `projects/*/locations/global/collections/default_collection/dataStores/default_data_store`.
	DataStore string `protobuf:"bytes,1,opt,name=data_store,json=dataStore,proto3" json:"data_store,omitempty"`
	// Required. The typeahead input used to fetch suggestions. Maximum length is
	// 128 characters.
	Query string `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	// Specifies the autocomplete data model. This overrides any model specified
	// in the Configuration > Autocomplete section of the Cloud console. Currently
	// supported values:
	//
	// * `document` - Using suggestions generated from user-imported documents.
	// * `search-history` - Using suggestions generated from the past history of
	// [SearchService.Search][google.cloud.discoveryengine.v1alpha.SearchService.Search]
	// API calls. Do not use it when there is no traffic for Search API.
	// * `user-event` - Using suggestions generated from user-imported search
	// events.
	// * `document-completable` - Using suggestions taken directly from
	// user-imported document fields marked as completable.
	//
	// Default values:
	//
	// * `document` is the default model for regular dataStores.
	// * `search-history` is the default model for site search dataStores.
	QueryModel string `protobuf:"bytes,3,opt,name=query_model,json=queryModel,proto3" json:"query_model,omitempty"`
	// A unique identifier for tracking visitors. For example, this could be
	// implemented with an HTTP cookie, which should be able to uniquely identify
	// a visitor on a single device. This unique identifier should not change if
	// the visitor logs in or out of the website.
	//
	// This field should NOT have a fixed value such as `unknown_visitor`.
	//
	// This should be the same identifier as
	// [UserEvent.user_pseudo_id][google.cloud.discoveryengine.v1alpha.UserEvent.user_pseudo_id]
	// and
	// [SearchRequest.user_pseudo_id][google.cloud.discoveryengine.v1alpha.SearchRequest.user_pseudo_id].
	//
	// The field must be a UTF-8 encoded string with a length limit of 128
	// characters. Otherwise, an `INVALID_ARGUMENT` error is returned.
	UserPseudoId string `protobuf:"bytes,4,opt,name=user_pseudo_id,json=userPseudoId,proto3" json:"user_pseudo_id,omitempty"`
	// Indicates if tail suggestions should be returned if there are no
	// suggestions that match the full query. Even if set to true, if there are
	// suggestions that match the full query, those are returned and no
	// tail suggestions are returned.
	IncludeTailSuggestions bool `protobuf:"varint,5,opt,name=include_tail_suggestions,json=includeTailSuggestions,proto3" json:"include_tail_suggestions,omitempty"`
}

func (x *CompleteQueryRequest) Reset() {
	*x = CompleteQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1alpha_completion_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteQueryRequest) ProtoMessage() {}

func (x *CompleteQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_completion_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteQueryRequest.ProtoReflect.Descriptor instead.
func (*CompleteQueryRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_completion_service_proto_rawDescGZIP(), []int{0}
}

func (x *CompleteQueryRequest) GetDataStore() string {
	if x != nil {
		return x.DataStore
	}
	return ""
}

func (x *CompleteQueryRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *CompleteQueryRequest) GetQueryModel() string {
	if x != nil {
		return x.QueryModel
	}
	return ""
}

func (x *CompleteQueryRequest) GetUserPseudoId() string {
	if x != nil {
		return x.UserPseudoId
	}
	return ""
}

func (x *CompleteQueryRequest) GetIncludeTailSuggestions() bool {
	if x != nil {
		return x.IncludeTailSuggestions
	}
	return false
}

// Response message for
// [CompletionService.CompleteQuery][google.cloud.discoveryengine.v1alpha.CompletionService.CompleteQuery]
// method.
type CompleteQueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Results of the matched query suggestions. The result list is ordered and
	// the first result is a top suggestion.
	QuerySuggestions []*CompleteQueryResponse_QuerySuggestion `protobuf:"bytes,1,rep,name=query_suggestions,json=querySuggestions,proto3" json:"query_suggestions,omitempty"`
	// True if the returned suggestions are all tail suggestions.
	//
	// For tail matching to be triggered, include_tail_suggestions in the request
	// must be true and there must be no suggestions that match the full query.
	TailMatchTriggered bool `protobuf:"varint,2,opt,name=tail_match_triggered,json=tailMatchTriggered,proto3" json:"tail_match_triggered,omitempty"`
}

func (x *CompleteQueryResponse) Reset() {
	*x = CompleteQueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1alpha_completion_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteQueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteQueryResponse) ProtoMessage() {}

func (x *CompleteQueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_completion_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteQueryResponse.ProtoReflect.Descriptor instead.
func (*CompleteQueryResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_completion_service_proto_rawDescGZIP(), []int{1}
}

func (x *CompleteQueryResponse) GetQuerySuggestions() []*CompleteQueryResponse_QuerySuggestion {
	if x != nil {
		return x.QuerySuggestions
	}
	return nil
}

func (x *CompleteQueryResponse) GetTailMatchTriggered() bool {
	if x != nil {
		return x.TailMatchTriggered
	}
	return false
}

// Suggestions as search queries.
type CompleteQueryResponse_QuerySuggestion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The suggestion for the query.
	Suggestion string `protobuf:"bytes,1,opt,name=suggestion,proto3" json:"suggestion,omitempty"`
	// The unique document field paths that serve as the source of this
	// suggestion if it was generated from completable fields.
	//
	// This field is only populated for the document-completable model.
	CompletableFieldPaths []string `protobuf:"bytes,2,rep,name=completable_field_paths,json=completableFieldPaths,proto3" json:"completable_field_paths,omitempty"`
}

func (x *CompleteQueryResponse_QuerySuggestion) Reset() {
	*x = CompleteQueryResponse_QuerySuggestion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1alpha_completion_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteQueryResponse_QuerySuggestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteQueryResponse_QuerySuggestion) ProtoMessage() {}

func (x *CompleteQueryResponse_QuerySuggestion) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_completion_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteQueryResponse_QuerySuggestion.ProtoReflect.Descriptor instead.
func (*CompleteQueryResponse_QuerySuggestion) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_completion_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CompleteQueryResponse_QuerySuggestion) GetSuggestion() string {
	if x != nil {
		return x.Suggestion
	}
	return ""
}

func (x *CompleteQueryResponse_QuerySuggestion) GetCompletableFieldPaths() []string {
	if x != nil {
		return x.CompletableFieldPaths
	}
	return nil
}

var File_google_cloud_discoveryengine_v1alpha_completion_service_proto protoreflect.FileDescriptor

var file_google_cloud_discoveryengine_v1alpha_completion_service_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x69,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x70, 0x75, 0x72, 0x67, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x83, 0x02, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4f, 0x0a, 0x0a, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0xe0,
	0x41, 0x02, 0xfa, 0x41, 0x2a, 0x0a, 0x28, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52,
	0x09, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x24, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70,
	0x73, 0x65, 0x75, 0x64, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x75, 0x73, 0x65, 0x72, 0x50, 0x73, 0x65, 0x75, 0x64, 0x6f, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x18,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x73, 0x75, 0x67,
	0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x54, 0x61, 0x69, 0x6c, 0x53, 0x75, 0x67, 0x67, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xae, 0x02, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x78, 0x0a, 0x11, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x75,
	0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x71, 0x75, 0x65, 0x72, 0x79, 0x53,
	0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x74, 0x61,
	0x69, 0x6c, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x74, 0x61, 0x69, 0x6c, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x65, 0x64, 0x1a, 0x69, 0x0a, 0x0f,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x36, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x15, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x50, 0x61, 0x74, 0x68, 0x73, 0x32, 0x9b, 0x0b, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xb4, 0x02,
	0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa9, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0xa2, 0x01, 0x5a, 0x57, 0x12, 0x55, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x2a, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x47, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0xfe, 0x03, 0x0a, 0x1f, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53,
	0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6e, 0x79, 0x4c, 0x69, 0x73,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x4c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x65, 0x6e, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xed, 0x02, 0xca, 0x41, 0x9c, 0x01, 0x0a, 0x4c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6e, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x65, 0x6e, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0xc6, 0x01, 0x3a,
	0x01, 0x2a, 0x5a, 0x5b, 0x3a, 0x01, 0x2a, 0x22, 0x56, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x73,
	0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6e, 0x79, 0x4c, 0x69, 0x73,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x3a, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x22,
	0x64, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x65, 0x6e, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x3a, 0x69,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0xf9, 0x03, 0x0a, 0x1e, 0x50, 0x75, 0x72, 0x67, 0x65, 0x53,
	0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6e, 0x79, 0x4c, 0x69, 0x73,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x4b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x50, 0x75, 0x72, 0x67, 0x65, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x65, 0x6e, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c,
	0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xea, 0x02, 0xca, 0x41, 0x9a, 0x01, 0x0a, 0x4b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x50, 0x75, 0x72, 0x67, 0x65, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x65, 0x6e, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x50,
	0x75, 0x72, 0x67, 0x65, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65,
	0x6e, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0xc5, 0x01, 0x3a, 0x01, 0x2a, 0x5a,
	0x5b, 0x3a, 0x01, 0x2a, 0x22, 0x56, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x2a, 0x2a, 0x7d, 0x2f, 0x73, 0x75, 0x67,
	0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6e, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x3a, 0x70, 0x75, 0x72, 0x67, 0x65, 0x22, 0x63, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x2a,
	0x7d, 0x2f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6e, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x3a, 0x70, 0x75, 0x72, 0x67,
	0x65, 0x1a, 0x52, 0xca, 0x41, 0x1e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77,
	0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0xa2, 0x02, 0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x42, 0x16, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x52, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0x3b, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62,
	0xa2, 0x02, 0x0f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x59, 0x45, 0x4e, 0x47, 0x49,
	0x4e, 0x45, 0xaa, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x56, 0x31, 0x41, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0xea, 0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_cloud_discoveryengine_v1alpha_completion_service_proto_rawDescOnce sync.Once
	file_google_cloud_discoveryengine_v1alpha_completion_service_proto_rawDescData = file_google_cloud_discoveryengine_v1alpha_completion_service_proto_rawDesc
)

func file_google_cloud_discoveryengine_v1alpha_completion_service_proto_rawDescGZIP() []byte {
	file_google_cloud_discoveryengine_v1alpha_completion_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_discoveryengine_v1alpha_completion_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_discoveryengine_v1alpha_completion_service_proto_rawDescData)
	})
	return file_google_cloud_discoveryengine_v1alpha_completion_service_proto_rawDescData
}

var file_google_cloud_discoveryengine_v1alpha_completion_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_discoveryengine_v1alpha_completion_service_proto_goTypes = []any{
	(*CompleteQueryRequest)(nil),                   // 0: google.cloud.discoveryengine.v1alpha.CompleteQueryRequest
	(*CompleteQueryResponse)(nil),                  // 1: google.cloud.discoveryengine.v1alpha.CompleteQueryResponse
	(*CompleteQueryResponse_QuerySuggestion)(nil),  // 2: google.cloud.discoveryengine.v1alpha.CompleteQueryResponse.QuerySuggestion
	(*ImportSuggestionDenyListEntriesRequest)(nil), // 3: google.cloud.discoveryengine.v1alpha.ImportSuggestionDenyListEntriesRequest
	(*PurgeSuggestionDenyListEntriesRequest)(nil),  // 4: google.cloud.discoveryengine.v1alpha.PurgeSuggestionDenyListEntriesRequest
	(*longrunningpb.Operation)(nil),                // 5: google.longrunning.Operation
}
var file_google_cloud_discoveryengine_v1alpha_completion_service_proto_depIdxs = []int32{
	2, // 0: google.cloud.discoveryengine.v1alpha.CompleteQueryResponse.query_suggestions:type_name -> google.cloud.discoveryengine.v1alpha.CompleteQueryResponse.QuerySuggestion
	0, // 1: google.cloud.discoveryengine.v1alpha.CompletionService.CompleteQuery:input_type -> google.cloud.discoveryengine.v1alpha.CompleteQueryRequest
	3, // 2: google.cloud.discoveryengine.v1alpha.CompletionService.ImportSuggestionDenyListEntries:input_type -> google.cloud.discoveryengine.v1alpha.ImportSuggestionDenyListEntriesRequest
	4, // 3: google.cloud.discoveryengine.v1alpha.CompletionService.PurgeSuggestionDenyListEntries:input_type -> google.cloud.discoveryengine.v1alpha.PurgeSuggestionDenyListEntriesRequest
	1, // 4: google.cloud.discoveryengine.v1alpha.CompletionService.CompleteQuery:output_type -> google.cloud.discoveryengine.v1alpha.CompleteQueryResponse
	5, // 5: google.cloud.discoveryengine.v1alpha.CompletionService.ImportSuggestionDenyListEntries:output_type -> google.longrunning.Operation
	5, // 6: google.cloud.discoveryengine.v1alpha.CompletionService.PurgeSuggestionDenyListEntries:output_type -> google.longrunning.Operation
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_discoveryengine_v1alpha_completion_service_proto_init() }
func file_google_cloud_discoveryengine_v1alpha_completion_service_proto_init() {
	if File_google_cloud_discoveryengine_v1alpha_completion_service_proto != nil {
		return
	}
	file_google_cloud_discoveryengine_v1alpha_import_config_proto_init()
	file_google_cloud_discoveryengine_v1alpha_purge_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_discoveryengine_v1alpha_completion_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CompleteQueryRequest); i {
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
		file_google_cloud_discoveryengine_v1alpha_completion_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CompleteQueryResponse); i {
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
		file_google_cloud_discoveryengine_v1alpha_completion_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CompleteQueryResponse_QuerySuggestion); i {
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
			RawDescriptor: file_google_cloud_discoveryengine_v1alpha_completion_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_discoveryengine_v1alpha_completion_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_discoveryengine_v1alpha_completion_service_proto_depIdxs,
		MessageInfos:      file_google_cloud_discoveryengine_v1alpha_completion_service_proto_msgTypes,
	}.Build()
	File_google_cloud_discoveryengine_v1alpha_completion_service_proto = out.File
	file_google_cloud_discoveryengine_v1alpha_completion_service_proto_rawDesc = nil
	file_google_cloud_discoveryengine_v1alpha_completion_service_proto_goTypes = nil
	file_google_cloud_discoveryengine_v1alpha_completion_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CompletionServiceClient is the client API for CompletionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CompletionServiceClient interface {
	// Completes the specified user input with keyword suggestions.
	CompleteQuery(ctx context.Context, in *CompleteQueryRequest, opts ...grpc.CallOption) (*CompleteQueryResponse, error)
	// Imports all
	// [SuggestionDenyListEntry][google.cloud.discoveryengine.v1alpha.SuggestionDenyListEntry]
	// for a DataStore.
	ImportSuggestionDenyListEntries(ctx context.Context, in *ImportSuggestionDenyListEntriesRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Permanently deletes all
	// [SuggestionDenyListEntry][google.cloud.discoveryengine.v1alpha.SuggestionDenyListEntry]
	// for a DataStore.
	PurgeSuggestionDenyListEntries(ctx context.Context, in *PurgeSuggestionDenyListEntriesRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
}

type completionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompletionServiceClient(cc grpc.ClientConnInterface) CompletionServiceClient {
	return &completionServiceClient{cc}
}

func (c *completionServiceClient) CompleteQuery(ctx context.Context, in *CompleteQueryRequest, opts ...grpc.CallOption) (*CompleteQueryResponse, error) {
	out := new(CompleteQueryResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.discoveryengine.v1alpha.CompletionService/CompleteQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *completionServiceClient) ImportSuggestionDenyListEntries(ctx context.Context, in *ImportSuggestionDenyListEntriesRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.discoveryengine.v1alpha.CompletionService/ImportSuggestionDenyListEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *completionServiceClient) PurgeSuggestionDenyListEntries(ctx context.Context, in *PurgeSuggestionDenyListEntriesRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.discoveryengine.v1alpha.CompletionService/PurgeSuggestionDenyListEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompletionServiceServer is the server API for CompletionService service.
type CompletionServiceServer interface {
	// Completes the specified user input with keyword suggestions.
	CompleteQuery(context.Context, *CompleteQueryRequest) (*CompleteQueryResponse, error)
	// Imports all
	// [SuggestionDenyListEntry][google.cloud.discoveryengine.v1alpha.SuggestionDenyListEntry]
	// for a DataStore.
	ImportSuggestionDenyListEntries(context.Context, *ImportSuggestionDenyListEntriesRequest) (*longrunningpb.Operation, error)
	// Permanently deletes all
	// [SuggestionDenyListEntry][google.cloud.discoveryengine.v1alpha.SuggestionDenyListEntry]
	// for a DataStore.
	PurgeSuggestionDenyListEntries(context.Context, *PurgeSuggestionDenyListEntriesRequest) (*longrunningpb.Operation, error)
}

// UnimplementedCompletionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCompletionServiceServer struct {
}

func (*UnimplementedCompletionServiceServer) CompleteQuery(context.Context, *CompleteQueryRequest) (*CompleteQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteQuery not implemented")
}
func (*UnimplementedCompletionServiceServer) ImportSuggestionDenyListEntries(context.Context, *ImportSuggestionDenyListEntriesRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportSuggestionDenyListEntries not implemented")
}
func (*UnimplementedCompletionServiceServer) PurgeSuggestionDenyListEntries(context.Context, *PurgeSuggestionDenyListEntriesRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurgeSuggestionDenyListEntries not implemented")
}

func RegisterCompletionServiceServer(s *grpc.Server, srv CompletionServiceServer) {
	s.RegisterService(&_CompletionService_serviceDesc, srv)
}

func _CompletionService_CompleteQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompletionServiceServer).CompleteQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.discoveryengine.v1alpha.CompletionService/CompleteQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompletionServiceServer).CompleteQuery(ctx, req.(*CompleteQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompletionService_ImportSuggestionDenyListEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportSuggestionDenyListEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompletionServiceServer).ImportSuggestionDenyListEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.discoveryengine.v1alpha.CompletionService/ImportSuggestionDenyListEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompletionServiceServer).ImportSuggestionDenyListEntries(ctx, req.(*ImportSuggestionDenyListEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompletionService_PurgeSuggestionDenyListEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurgeSuggestionDenyListEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompletionServiceServer).PurgeSuggestionDenyListEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.discoveryengine.v1alpha.CompletionService/PurgeSuggestionDenyListEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompletionServiceServer).PurgeSuggestionDenyListEntries(ctx, req.(*PurgeSuggestionDenyListEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CompletionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.discoveryengine.v1alpha.CompletionService",
	HandlerType: (*CompletionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CompleteQuery",
			Handler:    _CompletionService_CompleteQuery_Handler,
		},
		{
			MethodName: "ImportSuggestionDenyListEntries",
			Handler:    _CompletionService_ImportSuggestionDenyListEntries_Handler,
		},
		{
			MethodName: "PurgeSuggestionDenyListEntries",
			Handler:    _CompletionService_PurgeSuggestionDenyListEntries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/discoveryengine/v1alpha/completion_service.proto",
}
