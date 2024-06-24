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
// source: google/cloud/apphub/v1/workload.proto

package apphubpb

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

// Workload state.
type Workload_State int32

const (
	// Unspecified state.
	Workload_STATE_UNSPECIFIED Workload_State = 0
	// The Workload is being created.
	Workload_CREATING Workload_State = 1
	// The Workload is ready.
	Workload_ACTIVE Workload_State = 2
	// The Workload is being deleted.
	Workload_DELETING Workload_State = 3
	// The underlying compute resources have been deleted.
	Workload_DETACHED Workload_State = 4
)

// Enum value maps for Workload_State.
var (
	Workload_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "CREATING",
		2: "ACTIVE",
		3: "DELETING",
		4: "DETACHED",
	}
	Workload_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"CREATING":          1,
		"ACTIVE":            2,
		"DELETING":          3,
		"DETACHED":          4,
	}
)

func (x Workload_State) Enum() *Workload_State {
	p := new(Workload_State)
	*p = x
	return p
}

func (x Workload_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Workload_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_apphub_v1_workload_proto_enumTypes[0].Descriptor()
}

func (Workload_State) Type() protoreflect.EnumType {
	return &file_google_cloud_apphub_v1_workload_proto_enumTypes[0]
}

func (x Workload_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Workload_State.Descriptor instead.
func (Workload_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_apphub_v1_workload_proto_rawDescGZIP(), []int{0, 0}
}

// Workload is an App Hub data model that contains a discovered workload, which
// represents a binary deployment (such as managed instance groups (MIGs) and
// GKE deployments) that performs the smallest logical subset of business
// functionality.
type Workload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier. The resource name of the Workload. Format:
	// "projects/{host-project-id}/locations/{location}/applications/{application-id}/workloads/{workload-id}"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. User-defined name for the Workload.
	// Can have a maximum length of 63 characters.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Optional. User-defined description of a Workload.
	// Can have a maximum length of 2048 characters.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. Reference of an underlying compute resource represented by the
	// Workload. These are immutable.
	WorkloadReference *WorkloadReference `protobuf:"bytes,4,opt,name=workload_reference,json=workloadReference,proto3" json:"workload_reference,omitempty"`
	// Output only. Properties of an underlying compute resource represented by
	// the Workload. These are immutable.
	WorkloadProperties *WorkloadProperties `protobuf:"bytes,5,opt,name=workload_properties,json=workloadProperties,proto3" json:"workload_properties,omitempty"`
	// Required. Immutable. The resource name of the original discovered workload.
	DiscoveredWorkload string `protobuf:"bytes,6,opt,name=discovered_workload,json=discoveredWorkload,proto3" json:"discovered_workload,omitempty"`
	// Optional. Consumer provided attributes.
	Attributes *Attributes `protobuf:"bytes,7,opt,name=attributes,proto3" json:"attributes,omitempty"`
	// Output only. Create time.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Update time.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. A universally unique identifier (UUID) for the `Workload` in
	// the UUID4 format.
	Uid string `protobuf:"bytes,10,opt,name=uid,proto3" json:"uid,omitempty"`
	// Output only. Workload state.
	State Workload_State `protobuf:"varint,11,opt,name=state,proto3,enum=google.cloud.apphub.v1.Workload_State" json:"state,omitempty"`
}

func (x *Workload) Reset() {
	*x = Workload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apphub_v1_workload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Workload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Workload) ProtoMessage() {}

func (x *Workload) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apphub_v1_workload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Workload.ProtoReflect.Descriptor instead.
func (*Workload) Descriptor() ([]byte, []int) {
	return file_google_cloud_apphub_v1_workload_proto_rawDescGZIP(), []int{0}
}

func (x *Workload) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Workload) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Workload) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Workload) GetWorkloadReference() *WorkloadReference {
	if x != nil {
		return x.WorkloadReference
	}
	return nil
}

func (x *Workload) GetWorkloadProperties() *WorkloadProperties {
	if x != nil {
		return x.WorkloadProperties
	}
	return nil
}

func (x *Workload) GetDiscoveredWorkload() string {
	if x != nil {
		return x.DiscoveredWorkload
	}
	return ""
}

func (x *Workload) GetAttributes() *Attributes {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Workload) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Workload) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Workload) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Workload) GetState() Workload_State {
	if x != nil {
		return x.State
	}
	return Workload_STATE_UNSPECIFIED
}

// Reference of an underlying compute resource represented by the Workload.
type WorkloadReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The underlying compute resource uri.
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *WorkloadReference) Reset() {
	*x = WorkloadReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apphub_v1_workload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadReference) ProtoMessage() {}

func (x *WorkloadReference) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apphub_v1_workload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadReference.ProtoReflect.Descriptor instead.
func (*WorkloadReference) Descriptor() ([]byte, []int) {
	return file_google_cloud_apphub_v1_workload_proto_rawDescGZIP(), []int{1}
}

func (x *WorkloadReference) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

// Properties of an underlying compute resource represented by the Workload.
type WorkloadProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The service project identifier that the underlying cloud
	// resource resides in. Empty for non cloud resources.
	GcpProject string `protobuf:"bytes,1,opt,name=gcp_project,json=gcpProject,proto3" json:"gcp_project,omitempty"`
	// Output only. The location that the underlying compute resource resides in
	// (e.g us-west1).
	Location string `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	// Output only. The location that the underlying compute resource resides in
	// if it is zonal (e.g us-west1-a).
	Zone string `protobuf:"bytes,3,opt,name=zone,proto3" json:"zone,omitempty"`
}

func (x *WorkloadProperties) Reset() {
	*x = WorkloadProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apphub_v1_workload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadProperties) ProtoMessage() {}

func (x *WorkloadProperties) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apphub_v1_workload_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadProperties.ProtoReflect.Descriptor instead.
func (*WorkloadProperties) Descriptor() ([]byte, []int) {
	return file_google_cloud_apphub_v1_workload_proto_rawDescGZIP(), []int{2}
}

func (x *WorkloadProperties) GetGcpProject() string {
	if x != nil {
		return x.GcpProject
	}
	return ""
}

func (x *WorkloadProperties) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *WorkloadProperties) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

// DiscoveredWorkload is a binary deployment (such as managed instance groups
// (MIGs) and GKE deployments) that performs the smallest logical subset of
// business functionality. A discovered workload can be registered to an App Hub
// Workload.
type DiscoveredWorkload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier. The resource name of the discovered workload. Format:
	// "projects/{host-project-id}/locations/{location}/discoveredWorkloads/{uuid}"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Reference of an underlying compute resource represented by the
	// Workload. These are immutable.
	WorkloadReference *WorkloadReference `protobuf:"bytes,2,opt,name=workload_reference,json=workloadReference,proto3" json:"workload_reference,omitempty"`
	// Output only. Properties of an underlying compute resource represented by
	// the Workload. These are immutable.
	WorkloadProperties *WorkloadProperties `protobuf:"bytes,3,opt,name=workload_properties,json=workloadProperties,proto3" json:"workload_properties,omitempty"`
}

func (x *DiscoveredWorkload) Reset() {
	*x = DiscoveredWorkload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apphub_v1_workload_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoveredWorkload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoveredWorkload) ProtoMessage() {}

func (x *DiscoveredWorkload) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apphub_v1_workload_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoveredWorkload.ProtoReflect.Descriptor instead.
func (*DiscoveredWorkload) Descriptor() ([]byte, []int) {
	return file_google_cloud_apphub_v1_workload_proto_rawDescGZIP(), []int{3}
}

func (x *DiscoveredWorkload) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DiscoveredWorkload) GetWorkloadReference() *WorkloadReference {
	if x != nil {
		return x.WorkloadReference
	}
	return nil
}

func (x *DiscoveredWorkload) GetWorkloadProperties() *WorkloadProperties {
	if x != nil {
		return x.WorkloadProperties
	}
	return nil
}

var File_google_cloud_apphub_v1_workload_proto protoreflect.FileDescriptor

var file_google_cloud_apphub_v1_workload_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x70, 0x70, 0x68, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x70, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x70, 0x68, 0x75, 0x62, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb3, 0x07, 0x0a, 0x08, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x01, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5d, 0x0a, 0x12, 0x77, 0x6f, 0x72, 0x6b, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x70, 0x70, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x11, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x13, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x70, 0x70, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x12, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x64, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x2a,
	0x12, 0x28, 0x61, 0x70, 0x70, 0x68, 0x75, 0x62, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x12, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x47,
	0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x70, 0x70, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0a, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xe0, 0x41, 0x03, 0xe2, 0x8c, 0xcf,
	0xd7, 0x08, 0x02, 0x08, 0x01, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x41, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x70, 0x68, 0x75, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x54, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41,
	0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54,
	0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x54, 0x41, 0x43, 0x48, 0x45,
	0x44, 0x10, 0x04, 0x3a, 0x92, 0x01, 0xea, 0x41, 0x8e, 0x01, 0x0a, 0x1e, 0x61, 0x70, 0x70, 0x68,
	0x75, 0x62, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x57, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x7d, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x7b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2f, 0x7b, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x7d, 0x2a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x32, 0x08,
	0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x2a, 0x0a, 0x11, 0x57, 0x6f, 0x72, 0x6b,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x15, 0x0a,
	0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x03, 0x75, 0x72, 0x69, 0x22, 0x74, 0x0a, 0x12, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0b, 0x67, 0x63,
	0x70, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x67, 0x63, 0x70, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x1f, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x17, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x22, 0x9b, 0x03, 0x0a, 0x12, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x5d, 0x0a, 0x12, 0x77, 0x6f,
	0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x70, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x11, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x13, 0x77, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x70, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x12, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x3a, 0xaa, 0x01, 0xea, 0x41,
	0xa6, 0x01, 0x0a, 0x28, 0x61, 0x70, 0x70, 0x68, 0x75, 0x62, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x51, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x64,
	0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2f, 0x7b, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x7d, 0x2a,
	0x13, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x6c,
	0x6f, 0x61, 0x64, 0x73, 0x32, 0x12, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x64,
	0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0xaf, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70,
	0x70, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70,
	0x70, 0x68, 0x75, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x68, 0x75,
	0x62, 0x70, 0x62, 0x3b, 0x61, 0x70, 0x70, 0x68, 0x75, 0x62, 0x70, 0x62, 0xaa, 0x02, 0x16, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x70, 0x70, 0x48,
	0x75, 0x62, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x70, 0x70, 0x48, 0x75, 0x62, 0x5c, 0x56, 0x31, 0xea, 0x02,
	0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x41, 0x70, 0x70, 0x48, 0x75, 0x62, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_cloud_apphub_v1_workload_proto_rawDescOnce sync.Once
	file_google_cloud_apphub_v1_workload_proto_rawDescData = file_google_cloud_apphub_v1_workload_proto_rawDesc
)

func file_google_cloud_apphub_v1_workload_proto_rawDescGZIP() []byte {
	file_google_cloud_apphub_v1_workload_proto_rawDescOnce.Do(func() {
		file_google_cloud_apphub_v1_workload_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_apphub_v1_workload_proto_rawDescData)
	})
	return file_google_cloud_apphub_v1_workload_proto_rawDescData
}

var file_google_cloud_apphub_v1_workload_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_apphub_v1_workload_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_apphub_v1_workload_proto_goTypes = []any{
	(Workload_State)(0),           // 0: google.cloud.apphub.v1.Workload.State
	(*Workload)(nil),              // 1: google.cloud.apphub.v1.Workload
	(*WorkloadReference)(nil),     // 2: google.cloud.apphub.v1.WorkloadReference
	(*WorkloadProperties)(nil),    // 3: google.cloud.apphub.v1.WorkloadProperties
	(*DiscoveredWorkload)(nil),    // 4: google.cloud.apphub.v1.DiscoveredWorkload
	(*Attributes)(nil),            // 5: google.cloud.apphub.v1.Attributes
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_google_cloud_apphub_v1_workload_proto_depIdxs = []int32{
	2, // 0: google.cloud.apphub.v1.Workload.workload_reference:type_name -> google.cloud.apphub.v1.WorkloadReference
	3, // 1: google.cloud.apphub.v1.Workload.workload_properties:type_name -> google.cloud.apphub.v1.WorkloadProperties
	5, // 2: google.cloud.apphub.v1.Workload.attributes:type_name -> google.cloud.apphub.v1.Attributes
	6, // 3: google.cloud.apphub.v1.Workload.create_time:type_name -> google.protobuf.Timestamp
	6, // 4: google.cloud.apphub.v1.Workload.update_time:type_name -> google.protobuf.Timestamp
	0, // 5: google.cloud.apphub.v1.Workload.state:type_name -> google.cloud.apphub.v1.Workload.State
	2, // 6: google.cloud.apphub.v1.DiscoveredWorkload.workload_reference:type_name -> google.cloud.apphub.v1.WorkloadReference
	3, // 7: google.cloud.apphub.v1.DiscoveredWorkload.workload_properties:type_name -> google.cloud.apphub.v1.WorkloadProperties
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_google_cloud_apphub_v1_workload_proto_init() }
func file_google_cloud_apphub_v1_workload_proto_init() {
	if File_google_cloud_apphub_v1_workload_proto != nil {
		return
	}
	file_google_cloud_apphub_v1_attributes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_apphub_v1_workload_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Workload); i {
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
		file_google_cloud_apphub_v1_workload_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*WorkloadReference); i {
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
		file_google_cloud_apphub_v1_workload_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*WorkloadProperties); i {
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
		file_google_cloud_apphub_v1_workload_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DiscoveredWorkload); i {
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
			RawDescriptor: file_google_cloud_apphub_v1_workload_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_apphub_v1_workload_proto_goTypes,
		DependencyIndexes: file_google_cloud_apphub_v1_workload_proto_depIdxs,
		EnumInfos:         file_google_cloud_apphub_v1_workload_proto_enumTypes,
		MessageInfos:      file_google_cloud_apphub_v1_workload_proto_msgTypes,
	}.Build()
	File_google_cloud_apphub_v1_workload_proto = out.File
	file_google_cloud_apphub_v1_workload_proto_rawDesc = nil
	file_google_cloud_apphub_v1_workload_proto_goTypes = nil
	file_google_cloud_apphub_v1_workload_proto_depIdxs = nil
}
