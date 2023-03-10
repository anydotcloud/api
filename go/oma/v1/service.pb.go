// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: oma/v1/service.proto

package omav1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_oma_v1_service_proto protoreflect.FileDescriptor

var file_oma_v1_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6f, 0x6d, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6f, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x19,
	0x6f, 0x6d, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xfe, 0x02, 0x0a, 0x0a, 0x4f, 0x4d,
	0x41, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x6f, 0x6d,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6f, 0x6d,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x20, 0x2e, 0x6f, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6f, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x5d, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x6f, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6f, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x12, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x21, 0x2e, 0x6f, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6f, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x7f, 0x0a, 0x0a, 0x63, 0x6f,
	0x6d, 0x2e, 0x6f, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x79, 0x64, 0x6f, 0x74, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x6f, 0x6d, 0x61, 0x2f, 0x76, 0x31, 0x3b, 0x6f,
	0x6d, 0x61, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4f, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x4f, 0x6d, 0x61,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x06, 0x4f, 0x6d, 0x61, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x12, 0x4f,
	0x6d, 0x61, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x07, 0x4f, 0x6d, 0x61, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_oma_v1_service_proto_goTypes = []interface{}{
	(*GetOrganizationRequest)(nil),     // 0: oma.v1.GetOrganizationRequest
	(*ListOrganizationsRequest)(nil),   // 1: oma.v1.ListOrganizationsRequest
	(*CreateOrganizationRequest)(nil),  // 2: oma.v1.CreateOrganizationRequest
	(*DeleteOrganizationRequest)(nil),  // 3: oma.v1.DeleteOrganizationRequest
	(*GetOrganizationResponse)(nil),    // 4: oma.v1.GetOrganizationResponse
	(*ListOrganizationsResponse)(nil),  // 5: oma.v1.ListOrganizationsResponse
	(*CreateOrganizationResponse)(nil), // 6: oma.v1.CreateOrganizationResponse
	(*DeleteOrganizationResponse)(nil), // 7: oma.v1.DeleteOrganizationResponse
}
var file_oma_v1_service_proto_depIdxs = []int32{
	0, // 0: oma.v1.OMAService.GetOrganization:input_type -> oma.v1.GetOrganizationRequest
	1, // 1: oma.v1.OMAService.ListOrganizations:input_type -> oma.v1.ListOrganizationsRequest
	2, // 2: oma.v1.OMAService.CreateOrganization:input_type -> oma.v1.CreateOrganizationRequest
	3, // 3: oma.v1.OMAService.DeleteOrganization:input_type -> oma.v1.DeleteOrganizationRequest
	4, // 4: oma.v1.OMAService.GetOrganization:output_type -> oma.v1.GetOrganizationResponse
	5, // 5: oma.v1.OMAService.ListOrganizations:output_type -> oma.v1.ListOrganizationsResponse
	6, // 6: oma.v1.OMAService.CreateOrganization:output_type -> oma.v1.CreateOrganizationResponse
	7, // 7: oma.v1.OMAService.DeleteOrganization:output_type -> oma.v1.DeleteOrganizationResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_oma_v1_service_proto_init() }
func file_oma_v1_service_proto_init() {
	if File_oma_v1_service_proto != nil {
		return
	}
	file_oma_v1_organization_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_oma_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_oma_v1_service_proto_goTypes,
		DependencyIndexes: file_oma_v1_service_proto_depIdxs,
	}.Build()
	File_oma_v1_service_proto = out.File
	file_oma_v1_service_proto_rawDesc = nil
	file_oma_v1_service_proto_goTypes = nil
	file_oma_v1_service_proto_depIdxs = nil
}
