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
// source: google/cloud/datacatalog/v1/common.proto

package datacatalogpb

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

// This enum lists all the systems that Data Catalog integrates with.
type IntegratedSystem int32

const (
	// Default unknown system.
	IntegratedSystem_INTEGRATED_SYSTEM_UNSPECIFIED IntegratedSystem = 0
	// BigQuery.
	IntegratedSystem_BIGQUERY IntegratedSystem = 1
	// Cloud Pub/Sub.
	IntegratedSystem_CLOUD_PUBSUB IntegratedSystem = 2
	// Dataproc Metastore.
	IntegratedSystem_DATAPROC_METASTORE IntegratedSystem = 3
	// Dataplex.
	IntegratedSystem_DATAPLEX IntegratedSystem = 4
	// Cloud Spanner
	IntegratedSystem_CLOUD_SPANNER IntegratedSystem = 6
	// Cloud Bigtable
	IntegratedSystem_CLOUD_BIGTABLE IntegratedSystem = 7
	// Cloud Sql
	IntegratedSystem_CLOUD_SQL IntegratedSystem = 8
	// Looker
	IntegratedSystem_LOOKER IntegratedSystem = 9
	// Vertex AI
	IntegratedSystem_VERTEX_AI IntegratedSystem = 10
)

// Enum value maps for IntegratedSystem.
var (
	IntegratedSystem_name = map[int32]string{
		0:  "INTEGRATED_SYSTEM_UNSPECIFIED",
		1:  "BIGQUERY",
		2:  "CLOUD_PUBSUB",
		3:  "DATAPROC_METASTORE",
		4:  "DATAPLEX",
		6:  "CLOUD_SPANNER",
		7:  "CLOUD_BIGTABLE",
		8:  "CLOUD_SQL",
		9:  "LOOKER",
		10: "VERTEX_AI",
	}
	IntegratedSystem_value = map[string]int32{
		"INTEGRATED_SYSTEM_UNSPECIFIED": 0,
		"BIGQUERY":                      1,
		"CLOUD_PUBSUB":                  2,
		"DATAPROC_METASTORE":            3,
		"DATAPLEX":                      4,
		"CLOUD_SPANNER":                 6,
		"CLOUD_BIGTABLE":                7,
		"CLOUD_SQL":                     8,
		"LOOKER":                        9,
		"VERTEX_AI":                     10,
	}
)

func (x IntegratedSystem) Enum() *IntegratedSystem {
	p := new(IntegratedSystem)
	*p = x
	return p
}

func (x IntegratedSystem) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IntegratedSystem) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_datacatalog_v1_common_proto_enumTypes[0].Descriptor()
}

func (IntegratedSystem) Type() protoreflect.EnumType {
	return &file_google_cloud_datacatalog_v1_common_proto_enumTypes[0]
}

func (x IntegratedSystem) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IntegratedSystem.Descriptor instead.
func (IntegratedSystem) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_datacatalog_v1_common_proto_rawDescGZIP(), []int{0}
}

// This enum describes all the systems that manage
// Taxonomy and PolicyTag resources in DataCatalog.
type ManagingSystem int32

const (
	// Default value
	ManagingSystem_MANAGING_SYSTEM_UNSPECIFIED ManagingSystem = 0
	// Dataplex.
	ManagingSystem_MANAGING_SYSTEM_DATAPLEX ManagingSystem = 1
	// Other
	ManagingSystem_MANAGING_SYSTEM_OTHER ManagingSystem = 2
)

// Enum value maps for ManagingSystem.
var (
	ManagingSystem_name = map[int32]string{
		0: "MANAGING_SYSTEM_UNSPECIFIED",
		1: "MANAGING_SYSTEM_DATAPLEX",
		2: "MANAGING_SYSTEM_OTHER",
	}
	ManagingSystem_value = map[string]int32{
		"MANAGING_SYSTEM_UNSPECIFIED": 0,
		"MANAGING_SYSTEM_DATAPLEX":    1,
		"MANAGING_SYSTEM_OTHER":       2,
	}
)

func (x ManagingSystem) Enum() *ManagingSystem {
	p := new(ManagingSystem)
	*p = x
	return p
}

func (x ManagingSystem) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ManagingSystem) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_datacatalog_v1_common_proto_enumTypes[1].Descriptor()
}

func (ManagingSystem) Type() protoreflect.EnumType {
	return &file_google_cloud_datacatalog_v1_common_proto_enumTypes[1]
}

func (x ManagingSystem) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ManagingSystem.Descriptor instead.
func (ManagingSystem) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_datacatalog_v1_common_proto_rawDescGZIP(), []int{1}
}

// Entry metadata relevant only to the user and private to them.
type PersonalDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// True if the entry is starred by the user; false otherwise.
	Starred bool `protobuf:"varint,1,opt,name=starred,proto3" json:"starred,omitempty"`
	// Set if the entry is starred; unset otherwise.
	StarTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=star_time,json=starTime,proto3" json:"star_time,omitempty"`
}

func (x *PersonalDetails) Reset() {
	*x = PersonalDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datacatalog_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonalDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonalDetails) ProtoMessage() {}

func (x *PersonalDetails) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datacatalog_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonalDetails.ProtoReflect.Descriptor instead.
func (*PersonalDetails) Descriptor() ([]byte, []int) {
	return file_google_cloud_datacatalog_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *PersonalDetails) GetStarred() bool {
	if x != nil {
		return x.Starred
	}
	return false
}

func (x *PersonalDetails) GetStarTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StarTime
	}
	return nil
}

var File_google_cloud_datacatalog_v1_common_proto protoreflect.FileDescriptor

var file_google_cloud_datacatalog_v1_common_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x0f, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x61, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x74, 0x61, 0x72, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x74,
	0x61, 0x72, 0x72, 0x65, 0x64, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x2a, 0xcc,
	0x01, 0x0a, 0x10, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x65, 0x64, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x12, 0x21, 0x0a, 0x1d, 0x49, 0x4e, 0x54, 0x45, 0x47, 0x52, 0x41, 0x54, 0x45,
	0x44, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x49, 0x47, 0x51, 0x55, 0x45,
	0x52, 0x59, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x50, 0x55,
	0x42, 0x53, 0x55, 0x42, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x41, 0x54, 0x41, 0x50, 0x52,
	0x4f, 0x43, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x10, 0x03, 0x12, 0x0c,
	0x0a, 0x08, 0x44, 0x41, 0x54, 0x41, 0x50, 0x4c, 0x45, 0x58, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d,
	0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x53, 0x50, 0x41, 0x4e, 0x4e, 0x45, 0x52, 0x10, 0x06, 0x12,
	0x12, 0x0a, 0x0e, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x42, 0x49, 0x47, 0x54, 0x41, 0x42, 0x4c,
	0x45, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x53, 0x51, 0x4c,
	0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x4f, 0x4f, 0x4b, 0x45, 0x52, 0x10, 0x09, 0x12, 0x0d,
	0x0a, 0x09, 0x56, 0x45, 0x52, 0x54, 0x45, 0x58, 0x5f, 0x41, 0x49, 0x10, 0x0a, 0x2a, 0x6a, 0x0a,
	0x0e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12,
	0x1f, 0x0a, 0x1b, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x59, 0x53, 0x54,
	0x45, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x59, 0x53,
	0x54, 0x45, 0x4d, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x50, 0x4c, 0x45, 0x58, 0x10, 0x01, 0x12, 0x19,
	0x0a, 0x15, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45,
	0x4d, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x02, 0x42, 0xc6, 0x01, 0x0a, 0x1f, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a,
	0x41, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x70, 0x62, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x70, 0x62, 0xf8, 0x01, 0x01, 0xaa, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5c, 0x56,
	0x31, 0xea, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x44, 0x61, 0x74, 0x61, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_datacatalog_v1_common_proto_rawDescOnce sync.Once
	file_google_cloud_datacatalog_v1_common_proto_rawDescData = file_google_cloud_datacatalog_v1_common_proto_rawDesc
)

func file_google_cloud_datacatalog_v1_common_proto_rawDescGZIP() []byte {
	file_google_cloud_datacatalog_v1_common_proto_rawDescOnce.Do(func() {
		file_google_cloud_datacatalog_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_datacatalog_v1_common_proto_rawDescData)
	})
	return file_google_cloud_datacatalog_v1_common_proto_rawDescData
}

var file_google_cloud_datacatalog_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_datacatalog_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_datacatalog_v1_common_proto_goTypes = []any{
	(IntegratedSystem)(0),         // 0: google.cloud.datacatalog.v1.IntegratedSystem
	(ManagingSystem)(0),           // 1: google.cloud.datacatalog.v1.ManagingSystem
	(*PersonalDetails)(nil),       // 2: google.cloud.datacatalog.v1.PersonalDetails
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_google_cloud_datacatalog_v1_common_proto_depIdxs = []int32{
	3, // 0: google.cloud.datacatalog.v1.PersonalDetails.star_time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_datacatalog_v1_common_proto_init() }
func file_google_cloud_datacatalog_v1_common_proto_init() {
	if File_google_cloud_datacatalog_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_datacatalog_v1_common_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PersonalDetails); i {
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
			RawDescriptor: file_google_cloud_datacatalog_v1_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_datacatalog_v1_common_proto_goTypes,
		DependencyIndexes: file_google_cloud_datacatalog_v1_common_proto_depIdxs,
		EnumInfos:         file_google_cloud_datacatalog_v1_common_proto_enumTypes,
		MessageInfos:      file_google_cloud_datacatalog_v1_common_proto_msgTypes,
	}.Build()
	File_google_cloud_datacatalog_v1_common_proto = out.File
	file_google_cloud_datacatalog_v1_common_proto_rawDesc = nil
	file_google_cloud_datacatalog_v1_common_proto_goTypes = nil
	file_google_cloud_datacatalog_v1_common_proto_depIdxs = nil
}
