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
// source: google/maps/routing/v2/route_label.proto

package routingpb

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

// Labels for the [`Route`][google.maps.routing.v2.Route] that are useful to
// identify specific properties of the route to compare against others.
type RouteLabel int32

const (
	// Default - not used.
	RouteLabel_ROUTE_LABEL_UNSPECIFIED RouteLabel = 0
	// The default "best" route returned for the route computation.
	RouteLabel_DEFAULT_ROUTE RouteLabel = 1
	// An alternative to the default "best" route. Routes like this will be
	// returned when
	// [`compute_alternative_routes`][google.maps.routing.v2.ComputeRoutesRequest.compute_alternative_routes]
	// is specified.
	RouteLabel_DEFAULT_ROUTE_ALTERNATE RouteLabel = 2
	// Fuel efficient route. Routes labeled with this value are determined to be
	// optimized for Eco parameters such as fuel consumption.
	RouteLabel_FUEL_EFFICIENT RouteLabel = 3
)

// Enum value maps for RouteLabel.
var (
	RouteLabel_name = map[int32]string{
		0: "ROUTE_LABEL_UNSPECIFIED",
		1: "DEFAULT_ROUTE",
		2: "DEFAULT_ROUTE_ALTERNATE",
		3: "FUEL_EFFICIENT",
	}
	RouteLabel_value = map[string]int32{
		"ROUTE_LABEL_UNSPECIFIED": 0,
		"DEFAULT_ROUTE":           1,
		"DEFAULT_ROUTE_ALTERNATE": 2,
		"FUEL_EFFICIENT":          3,
	}
)

func (x RouteLabel) Enum() *RouteLabel {
	p := new(RouteLabel)
	*p = x
	return p
}

func (x RouteLabel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RouteLabel) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_routing_v2_route_label_proto_enumTypes[0].Descriptor()
}

func (RouteLabel) Type() protoreflect.EnumType {
	return &file_google_maps_routing_v2_route_label_proto_enumTypes[0]
}

func (x RouteLabel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RouteLabel.Descriptor instead.
func (RouteLabel) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_routing_v2_route_label_proto_rawDescGZIP(), []int{0}
}

var File_google_maps_routing_v2_route_label_proto protoreflect.FileDescriptor

var file_google_maps_routing_v2_route_label_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x32, 0x2a, 0x6d, 0x0a, 0x0a, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x12, 0x1b, 0x0a, 0x17, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x5f, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a,
	0x0d, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x10, 0x01,
	0x12, 0x1b, 0x0a, 0x17, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x52, 0x4f, 0x55, 0x54,
	0x45, 0x5f, 0x41, 0x4c, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x12, 0x0a,
	0x0e, 0x46, 0x55, 0x45, 0x4c, 0x5f, 0x45, 0x46, 0x46, 0x49, 0x43, 0x49, 0x45, 0x4e, 0x54, 0x10,
	0x03, 0x42, 0xc4, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32,
	0x42, 0x0f, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x70, 0x62, 0x3b, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0xf8,
	0x01, 0x01, 0xa2, 0x02, 0x05, 0x47, 0x4d, 0x52, 0x56, 0x32, 0xaa, 0x02, 0x16, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x56, 0x32, 0xca, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x4d, 0x61, 0x70,
	0x73, 0x5c, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x32, 0xea, 0x02, 0x19, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x4d, 0x61, 0x70, 0x73, 0x3a, 0x3a, 0x52, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_routing_v2_route_label_proto_rawDescOnce sync.Once
	file_google_maps_routing_v2_route_label_proto_rawDescData = file_google_maps_routing_v2_route_label_proto_rawDesc
)

func file_google_maps_routing_v2_route_label_proto_rawDescGZIP() []byte {
	file_google_maps_routing_v2_route_label_proto_rawDescOnce.Do(func() {
		file_google_maps_routing_v2_route_label_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_routing_v2_route_label_proto_rawDescData)
	})
	return file_google_maps_routing_v2_route_label_proto_rawDescData
}

var file_google_maps_routing_v2_route_label_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_maps_routing_v2_route_label_proto_goTypes = []any{
	(RouteLabel)(0), // 0: google.maps.routing.v2.RouteLabel
}
var file_google_maps_routing_v2_route_label_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_maps_routing_v2_route_label_proto_init() }
func file_google_maps_routing_v2_route_label_proto_init() {
	if File_google_maps_routing_v2_route_label_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_maps_routing_v2_route_label_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_routing_v2_route_label_proto_goTypes,
		DependencyIndexes: file_google_maps_routing_v2_route_label_proto_depIdxs,
		EnumInfos:         file_google_maps_routing_v2_route_label_proto_enumTypes,
	}.Build()
	File_google_maps_routing_v2_route_label_proto = out.File
	file_google_maps_routing_v2_route_label_proto_rawDesc = nil
	file_google_maps_routing_v2_route_label_proto_goTypes = nil
	file_google_maps_routing_v2_route_label_proto_depIdxs = nil
}
