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
// 	protoc-gen-go v1.34.1
// 	protoc        v4.25.3
// source: google/cloud/networkservices/v1/common.proto

package networkservicespb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Possible criteria values that define logic of how matching is made.
type EndpointMatcher_MetadataLabelMatcher_MetadataLabelMatchCriteria int32

const (
	// Default value. Should not be used.
	EndpointMatcher_MetadataLabelMatcher_METADATA_LABEL_MATCH_CRITERIA_UNSPECIFIED EndpointMatcher_MetadataLabelMatcher_MetadataLabelMatchCriteria = 0
	// At least one of the Labels specified in the matcher should match the
	// metadata presented by xDS client.
	EndpointMatcher_MetadataLabelMatcher_MATCH_ANY EndpointMatcher_MetadataLabelMatcher_MetadataLabelMatchCriteria = 1
	// The metadata presented by the xDS client should contain all of the
	// labels specified here.
	EndpointMatcher_MetadataLabelMatcher_MATCH_ALL EndpointMatcher_MetadataLabelMatcher_MetadataLabelMatchCriteria = 2
)

// Enum value maps for EndpointMatcher_MetadataLabelMatcher_MetadataLabelMatchCriteria.
var (
	EndpointMatcher_MetadataLabelMatcher_MetadataLabelMatchCriteria_name = map[int32]string{
		0: "METADATA_LABEL_MATCH_CRITERIA_UNSPECIFIED",
		1: "MATCH_ANY",
		2: "MATCH_ALL",
	}
	EndpointMatcher_MetadataLabelMatcher_MetadataLabelMatchCriteria_value = map[string]int32{
		"METADATA_LABEL_MATCH_CRITERIA_UNSPECIFIED": 0,
		"MATCH_ANY": 1,
		"MATCH_ALL": 2,
	}
)

func (x EndpointMatcher_MetadataLabelMatcher_MetadataLabelMatchCriteria) Enum() *EndpointMatcher_MetadataLabelMatcher_MetadataLabelMatchCriteria {
	p := new(EndpointMatcher_MetadataLabelMatcher_MetadataLabelMatchCriteria)
	*p = x
	return p
}

func (x EndpointMatcher_MetadataLabelMatcher_MetadataLabelMatchCriteria) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EndpointMatcher_MetadataLabelMatcher_MetadataLabelMatchCriteria) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_networkservices_v1_common_proto_enumTypes[0].Descriptor()
}

func (EndpointMatcher_MetadataLabelMatcher_MetadataLabelMatchCriteria) Type() protoreflect.EnumType {
	return &file_google_cloud_networkservices_v1_common_proto_enumTypes[0]
}

func (x EndpointMatcher_MetadataLabelMatcher_MetadataLabelMatchCriteria) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EndpointMatcher_MetadataLabelMatcher_MetadataLabelMatchCriteria.Descriptor instead.
func (EndpointMatcher_MetadataLabelMatcher_MetadataLabelMatchCriteria) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_networkservices_v1_common_proto_rawDescGZIP(), []int{2, 0, 0}
}

// Represents the metadata of the long-running operation.
type OperationMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The time the operation was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The time the operation finished running.
	EndTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Output only. Server-defined resource path for the target of the operation.
	Target string `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	// Output only. Name of the verb executed by the operation.
	Verb string `protobuf:"bytes,4,opt,name=verb,proto3" json:"verb,omitempty"`
	// Output only. Human-readable status of the operation, if any.
	StatusMessage string `protobuf:"bytes,5,opt,name=status_message,json=statusMessage,proto3" json:"status_message,omitempty"`
	// Output only. Identifies whether the user has requested cancellation
	// of the operation. Operations that have successfully been cancelled
	// have [Operation.error][] value with a
	// [google.rpc.Status.code][google.rpc.Status.code] of 1, corresponding to
	// `Code.CANCELLED`.
	RequestedCancellation bool `protobuf:"varint,6,opt,name=requested_cancellation,json=requestedCancellation,proto3" json:"requested_cancellation,omitempty"`
	// Output only. API version used to start the operation.
	ApiVersion string `protobuf:"bytes,7,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
}

func (x *OperationMetadata) Reset() {
	*x = OperationMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_networkservices_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationMetadata) ProtoMessage() {}

func (x *OperationMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_networkservices_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationMetadata.ProtoReflect.Descriptor instead.
func (*OperationMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_networkservices_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *OperationMetadata) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *OperationMetadata) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *OperationMetadata) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *OperationMetadata) GetVerb() string {
	if x != nil {
		return x.Verb
	}
	return ""
}

func (x *OperationMetadata) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

func (x *OperationMetadata) GetRequestedCancellation() bool {
	if x != nil {
		return x.RequestedCancellation
	}
	return false
}

func (x *OperationMetadata) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

// Specification of a port-based selector.
type TrafficPortSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. A list of ports. Can be port numbers or port range
	// (example, [80-90] specifies all ports from 80 to 90, including
	// 80 and 90) or named ports or * to specify all ports. If the
	// list is empty, all ports are selected.
	Ports []string `protobuf:"bytes,1,rep,name=ports,proto3" json:"ports,omitempty"`
}

func (x *TrafficPortSelector) Reset() {
	*x = TrafficPortSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_networkservices_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrafficPortSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrafficPortSelector) ProtoMessage() {}

func (x *TrafficPortSelector) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_networkservices_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrafficPortSelector.ProtoReflect.Descriptor instead.
func (*TrafficPortSelector) Descriptor() ([]byte, []int) {
	return file_google_cloud_networkservices_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *TrafficPortSelector) GetPorts() []string {
	if x != nil {
		return x.Ports
	}
	return nil
}

// A definition of a matcher that selects endpoints to which the policies
// should be applied.
type EndpointMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies type of the matcher used for this endpoint matcher.
	//
	// Types that are assignable to MatcherType:
	//
	//	*EndpointMatcher_MetadataLabelMatcher_
	MatcherType isEndpointMatcher_MatcherType `protobuf_oneof:"matcher_type"`
}

func (x *EndpointMatcher) Reset() {
	*x = EndpointMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_networkservices_v1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointMatcher) ProtoMessage() {}

func (x *EndpointMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_networkservices_v1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointMatcher.ProtoReflect.Descriptor instead.
func (*EndpointMatcher) Descriptor() ([]byte, []int) {
	return file_google_cloud_networkservices_v1_common_proto_rawDescGZIP(), []int{2}
}

func (m *EndpointMatcher) GetMatcherType() isEndpointMatcher_MatcherType {
	if m != nil {
		return m.MatcherType
	}
	return nil
}

func (x *EndpointMatcher) GetMetadataLabelMatcher() *EndpointMatcher_MetadataLabelMatcher {
	if x, ok := x.GetMatcherType().(*EndpointMatcher_MetadataLabelMatcher_); ok {
		return x.MetadataLabelMatcher
	}
	return nil
}

type isEndpointMatcher_MatcherType interface {
	isEndpointMatcher_MatcherType()
}

type EndpointMatcher_MetadataLabelMatcher_ struct {
	// The matcher is based on node metadata presented by xDS clients.
	MetadataLabelMatcher *EndpointMatcher_MetadataLabelMatcher `protobuf:"bytes,1,opt,name=metadata_label_matcher,json=metadataLabelMatcher,proto3,oneof"`
}

func (*EndpointMatcher_MetadataLabelMatcher_) isEndpointMatcher_MatcherType() {}

// The matcher that is based on node metadata presented by xDS clients.
type EndpointMatcher_MetadataLabelMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies how matching should be done.
	//
	// Supported values are:
	// MATCH_ANY: At least one of the Labels specified in the
	//
	//	matcher should match the metadata presented by xDS client.
	//
	// MATCH_ALL: The metadata presented by the xDS client should
	//
	//	contain all of the labels specified here.
	//
	// The selection is determined based on the best match. For
	// example, suppose there are three EndpointPolicy
	// resources P1, P2 and P3 and if P1 has a the matcher as
	// MATCH_ANY <A:1, B:1>, P2 has MATCH_ALL <A:1,B:1>, and P3 has
	// MATCH_ALL <A:1,B:1,C:1>.
	//
	// If a client with label <A:1> connects, the config from P1
	// will be selected.
	//
	// If a client with label <A:1,B:1> connects, the config from P2
	// will be selected.
	//
	// If a client with label <A:1,B:1,C:1> connects, the config
	// from P3 will be selected.
	//
	// If there is more than one best match, (for example, if a
	// config P4 with selector <A:1,D:1> exists and if a client with
	// label <A:1,B:1,D:1> connects), an error will be thrown.
	MetadataLabelMatchCriteria EndpointMatcher_MetadataLabelMatcher_MetadataLabelMatchCriteria `protobuf:"varint,1,opt,name=metadata_label_match_criteria,json=metadataLabelMatchCriteria,proto3,enum=google.cloud.networkservices.v1.EndpointMatcher_MetadataLabelMatcher_MetadataLabelMatchCriteria" json:"metadata_label_match_criteria,omitempty"`
	// The list of label value pairs that must match labels in the
	// provided metadata based on filterMatchCriteria This list can
	// have at most 64 entries. The list can be empty if the match
	// criteria is MATCH_ANY, to specify a wildcard match (i.e this
	// matches any client).
	MetadataLabels []*EndpointMatcher_MetadataLabelMatcher_MetadataLabels `protobuf:"bytes,2,rep,name=metadata_labels,json=metadataLabels,proto3" json:"metadata_labels,omitempty"`
}

func (x *EndpointMatcher_MetadataLabelMatcher) Reset() {
	*x = EndpointMatcher_MetadataLabelMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_networkservices_v1_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointMatcher_MetadataLabelMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointMatcher_MetadataLabelMatcher) ProtoMessage() {}

func (x *EndpointMatcher_MetadataLabelMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_networkservices_v1_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointMatcher_MetadataLabelMatcher.ProtoReflect.Descriptor instead.
func (*EndpointMatcher_MetadataLabelMatcher) Descriptor() ([]byte, []int) {
	return file_google_cloud_networkservices_v1_common_proto_rawDescGZIP(), []int{2, 0}
}

func (x *EndpointMatcher_MetadataLabelMatcher) GetMetadataLabelMatchCriteria() EndpointMatcher_MetadataLabelMatcher_MetadataLabelMatchCriteria {
	if x != nil {
		return x.MetadataLabelMatchCriteria
	}
	return EndpointMatcher_MetadataLabelMatcher_METADATA_LABEL_MATCH_CRITERIA_UNSPECIFIED
}

func (x *EndpointMatcher_MetadataLabelMatcher) GetMetadataLabels() []*EndpointMatcher_MetadataLabelMatcher_MetadataLabels {
	if x != nil {
		return x.MetadataLabels
	}
	return nil
}

// Defines a name-pair value for a single label.
type EndpointMatcher_MetadataLabelMatcher_MetadataLabels struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Label name presented as key in xDS Node Metadata.
	LabelName string `protobuf:"bytes,1,opt,name=label_name,json=labelName,proto3" json:"label_name,omitempty"`
	// Required. Label value presented as value corresponding to the above
	// key, in xDS Node Metadata.
	LabelValue string `protobuf:"bytes,2,opt,name=label_value,json=labelValue,proto3" json:"label_value,omitempty"`
}

func (x *EndpointMatcher_MetadataLabelMatcher_MetadataLabels) Reset() {
	*x = EndpointMatcher_MetadataLabelMatcher_MetadataLabels{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_networkservices_v1_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointMatcher_MetadataLabelMatcher_MetadataLabels) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointMatcher_MetadataLabelMatcher_MetadataLabels) ProtoMessage() {}

func (x *EndpointMatcher_MetadataLabelMatcher_MetadataLabels) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_networkservices_v1_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointMatcher_MetadataLabelMatcher_MetadataLabels.ProtoReflect.Descriptor instead.
func (*EndpointMatcher_MetadataLabelMatcher_MetadataLabels) Descriptor() ([]byte, []int) {
	return file_google_cloud_networkservices_v1_common_proto_rawDescGZIP(), []int{2, 0, 0}
}

func (x *EndpointMatcher_MetadataLabelMatcher_MetadataLabels) GetLabelName() string {
	if x != nil {
		return x.LabelName
	}
	return ""
}

func (x *EndpointMatcher_MetadataLabelMatcher_MetadataLabels) GetLabelValue() string {
	if x != nil {
		return x.LabelValue
	}
	return ""
}

var File_google_cloud_networkservices_v1_common_proto protoreflect.FileDescriptor

var file_google_cloud_networkservices_v1_common_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd5, 0x02, 0x0a, 0x11, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x07, 0x65, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x76, 0x65, 0x72, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x76, 0x65, 0x72, 0x62, 0x12, 0x2a, 0x0a, 0x0e, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x16, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x15, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x61,
	0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x30, 0x0a, 0x13, 0x54, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x19, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x01, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x22, 0xa5, 0x05, 0x0a, 0x0f,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12,
	0x7d, 0x0a, 0x16, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x45, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x48, 0x00, 0x52, 0x14, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x1a, 0x82,
	0x04, 0x0a, 0x14, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0xa3, 0x01, 0x0a, 0x1d, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x5f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x60, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x52, 0x1a, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x12, 0x7d, 0x0a,
	0x0f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x54, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x0e, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x5a, 0x0a, 0x0e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x22,
	0x0a, 0x0a, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0b, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x69, 0x0a, 0x1a, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x12, 0x2d, 0x0a, 0x29, 0x4d, 0x45, 0x54, 0x41, 0x44, 0x41,
	0x54, 0x41, 0x5f, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x43,
	0x52, 0x49, 0x54, 0x45, 0x52, 0x49, 0x41, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x41,
	0x4e, 0x59, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x41, 0x4c,
	0x4c, 0x10, 0x02, 0x42, 0x0e, 0x0a, 0x0c, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x42, 0xec, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x70, 0x62, 0x3b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x70, 0x62, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1f, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x22,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_networkservices_v1_common_proto_rawDescOnce sync.Once
	file_google_cloud_networkservices_v1_common_proto_rawDescData = file_google_cloud_networkservices_v1_common_proto_rawDesc
)

func file_google_cloud_networkservices_v1_common_proto_rawDescGZIP() []byte {
	file_google_cloud_networkservices_v1_common_proto_rawDescOnce.Do(func() {
		file_google_cloud_networkservices_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_networkservices_v1_common_proto_rawDescData)
	})
	return file_google_cloud_networkservices_v1_common_proto_rawDescData
}

var file_google_cloud_networkservices_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_networkservices_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_networkservices_v1_common_proto_goTypes = []interface{}{
	(EndpointMatcher_MetadataLabelMatcher_MetadataLabelMatchCriteria)(0), // 0: google.cloud.networkservices.v1.EndpointMatcher.MetadataLabelMatcher.MetadataLabelMatchCriteria
	(*OperationMetadata)(nil),                                   // 1: google.cloud.networkservices.v1.OperationMetadata
	(*TrafficPortSelector)(nil),                                 // 2: google.cloud.networkservices.v1.TrafficPortSelector
	(*EndpointMatcher)(nil),                                     // 3: google.cloud.networkservices.v1.EndpointMatcher
	(*EndpointMatcher_MetadataLabelMatcher)(nil),                // 4: google.cloud.networkservices.v1.EndpointMatcher.MetadataLabelMatcher
	(*EndpointMatcher_MetadataLabelMatcher_MetadataLabels)(nil), // 5: google.cloud.networkservices.v1.EndpointMatcher.MetadataLabelMatcher.MetadataLabels
	(*timestamppb.Timestamp)(nil),                               // 6: google.protobuf.Timestamp
}
var file_google_cloud_networkservices_v1_common_proto_depIdxs = []int32{
	6, // 0: google.cloud.networkservices.v1.OperationMetadata.create_time:type_name -> google.protobuf.Timestamp
	6, // 1: google.cloud.networkservices.v1.OperationMetadata.end_time:type_name -> google.protobuf.Timestamp
	4, // 2: google.cloud.networkservices.v1.EndpointMatcher.metadata_label_matcher:type_name -> google.cloud.networkservices.v1.EndpointMatcher.MetadataLabelMatcher
	0, // 3: google.cloud.networkservices.v1.EndpointMatcher.MetadataLabelMatcher.metadata_label_match_criteria:type_name -> google.cloud.networkservices.v1.EndpointMatcher.MetadataLabelMatcher.MetadataLabelMatchCriteria
	5, // 4: google.cloud.networkservices.v1.EndpointMatcher.MetadataLabelMatcher.metadata_labels:type_name -> google.cloud.networkservices.v1.EndpointMatcher.MetadataLabelMatcher.MetadataLabels
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_networkservices_v1_common_proto_init() }
func file_google_cloud_networkservices_v1_common_proto_init() {
	if File_google_cloud_networkservices_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_networkservices_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationMetadata); i {
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
		file_google_cloud_networkservices_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrafficPortSelector); i {
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
		file_google_cloud_networkservices_v1_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointMatcher); i {
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
		file_google_cloud_networkservices_v1_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointMatcher_MetadataLabelMatcher); i {
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
		file_google_cloud_networkservices_v1_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointMatcher_MetadataLabelMatcher_MetadataLabels); i {
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
	file_google_cloud_networkservices_v1_common_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*EndpointMatcher_MetadataLabelMatcher_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_networkservices_v1_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_networkservices_v1_common_proto_goTypes,
		DependencyIndexes: file_google_cloud_networkservices_v1_common_proto_depIdxs,
		EnumInfos:         file_google_cloud_networkservices_v1_common_proto_enumTypes,
		MessageInfos:      file_google_cloud_networkservices_v1_common_proto_msgTypes,
	}.Build()
	File_google_cloud_networkservices_v1_common_proto = out.File
	file_google_cloud_networkservices_v1_common_proto_rawDesc = nil
	file_google_cloud_networkservices_v1_common_proto_goTypes = nil
	file_google_cloud_networkservices_v1_common_proto_depIdxs = nil
}
