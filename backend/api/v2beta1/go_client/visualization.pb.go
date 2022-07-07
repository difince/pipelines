// Copyright 2019 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: backend/api/v2beta1/visualization.proto

package go_client

import (
	context "context"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Type of visualization to be generated.
// This is required when creating the pipeline through CreateVisualization
// API.
type Visualization_Type int32

const (
	Visualization_ROC_CURVE Visualization_Type = 0
	Visualization_TFDV      Visualization_Type = 1
	Visualization_TFMA      Visualization_Type = 2
	Visualization_TABLE     Visualization_Type = 3
	Visualization_CUSTOM    Visualization_Type = 4
)

// Enum value maps for Visualization_Type.
var (
	Visualization_Type_name = map[int32]string{
		0: "ROC_CURVE",
		1: "TFDV",
		2: "TFMA",
		3: "TABLE",
		4: "CUSTOM",
	}
	Visualization_Type_value = map[string]int32{
		"ROC_CURVE": 0,
		"TFDV":      1,
		"TFMA":      2,
		"TABLE":     3,
		"CUSTOM":    4,
	}
)

func (x Visualization_Type) Enum() *Visualization_Type {
	p := new(Visualization_Type)
	*p = x
	return p
}

func (x Visualization_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Visualization_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_backend_api_v2beta1_visualization_proto_enumTypes[0].Descriptor()
}

func (Visualization_Type) Type() protoreflect.EnumType {
	return &file_backend_api_v2beta1_visualization_proto_enumTypes[0]
}

func (x Visualization_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Visualization_Type.Descriptor instead.
func (Visualization_Type) EnumDescriptor() ([]byte, []int) {
	return file_backend_api_v2beta1_visualization_proto_rawDescGZIP(), []int{1, 0}
}

// Create visualization by providing the type of visualization that is desired
// and input data paths. Input dat paths are assumed to be unique and are used
// for determining output path.
type CreateVisualizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Visualization *Visualization `protobuf:"bytes,1,opt,name=visualization,proto3" json:"visualization,omitempty"`
	Namespace     string         `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *CreateVisualizationRequest) Reset() {
	*x = CreateVisualizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_api_v2beta1_visualization_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVisualizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVisualizationRequest) ProtoMessage() {}

func (x *CreateVisualizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_api_v2beta1_visualization_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVisualizationRequest.ProtoReflect.Descriptor instead.
func (*CreateVisualizationRequest) Descriptor() ([]byte, []int) {
	return file_backend_api_v2beta1_visualization_proto_rawDescGZIP(), []int{0}
}

func (x *CreateVisualizationRequest) GetVisualization() *Visualization {
	if x != nil {
		return x.Visualization
	}
	return nil
}

func (x *CreateVisualizationRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type Visualization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Visualization_Type `protobuf:"varint,1,opt,name=type,proto3,enum=api.v2beta1.Visualization_Type" json:"type,omitempty"`
	// Path pattern of input data to be used during generation of visualizations.
	// This is required when creating the pipeline through CreateVisualization
	// API.
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	// Variables to be used during generation of a visualization.
	// This should be provided as a JSON string.
	// This is required when creating the pipeline through CreateVisualization
	// API.
	Arguments string `protobuf:"bytes,3,opt,name=arguments,proto3" json:"arguments,omitempty"`
	// Output. Generated visualization html.
	Html string `protobuf:"bytes,4,opt,name=html,proto3" json:"html,omitempty"`
	// In case any error happens when generating visualizations, only
	// visualization ID and the error message are returned. Client has the
	// flexibility of choosing how to handle the error.
	Error string `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Visualization) Reset() {
	*x = Visualization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_api_v2beta1_visualization_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Visualization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Visualization) ProtoMessage() {}

func (x *Visualization) ProtoReflect() protoreflect.Message {
	mi := &file_backend_api_v2beta1_visualization_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Visualization.ProtoReflect.Descriptor instead.
func (*Visualization) Descriptor() ([]byte, []int) {
	return file_backend_api_v2beta1_visualization_proto_rawDescGZIP(), []int{1}
}

func (x *Visualization) GetType() Visualization_Type {
	if x != nil {
		return x.Type
	}
	return Visualization_ROC_CURVE
}

func (x *Visualization) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Visualization) GetArguments() string {
	if x != nil {
		return x.Arguments
	}
	return ""
}

func (x *Visualization) GetHtml() string {
	if x != nil {
		return x.Html
	}
	return ""
}

func (x *Visualization) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_backend_api_v2beta1_visualization_proto protoreflect.FileDescriptor

var file_backend_api_v2beta1_visualization_proto_rawDesc = []byte{
	0x0a, 0x27, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x69, 0x73, 0x75,
	0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x40, 0x0a, 0x0d, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x22, 0xe6, 0x01, 0x0a, 0x0d, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x56,
	0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x74, 0x6d, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x74, 0x6d,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x40, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0d, 0x0a, 0x09, 0x52, 0x4f, 0x43, 0x5f, 0x43, 0x55, 0x52, 0x56, 0x45, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x54, 0x46, 0x44, 0x56, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x46, 0x4d, 0x41,
	0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x0a, 0x0a,
	0x06, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10, 0x04, 0x32, 0xb4, 0x01, 0x0a, 0x14, 0x56, 0x69,
	0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x9b, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x69, 0x73,
	0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56,
	0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x3f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x39, 0x22, 0x28, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76,
	0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x7d, 0x3a, 0x0d, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x95, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6b, 0x75, 0x62, 0x65, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x73, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x67, 0x6f, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x92, 0x41, 0x55, 0x52, 0x24, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x19,
	0x12, 0x17, 0x0a, 0x15, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5a, 0x1f, 0x0a, 0x1d, 0x0a, 0x06, 0x42,
	0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x13, 0x08, 0x02, 0x1a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06,
	0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_backend_api_v2beta1_visualization_proto_rawDescOnce sync.Once
	file_backend_api_v2beta1_visualization_proto_rawDescData = file_backend_api_v2beta1_visualization_proto_rawDesc
)

func file_backend_api_v2beta1_visualization_proto_rawDescGZIP() []byte {
	file_backend_api_v2beta1_visualization_proto_rawDescOnce.Do(func() {
		file_backend_api_v2beta1_visualization_proto_rawDescData = protoimpl.X.CompressGZIP(file_backend_api_v2beta1_visualization_proto_rawDescData)
	})
	return file_backend_api_v2beta1_visualization_proto_rawDescData
}

var file_backend_api_v2beta1_visualization_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_backend_api_v2beta1_visualization_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_backend_api_v2beta1_visualization_proto_goTypes = []interface{}{
	(Visualization_Type)(0),            // 0: api.v2beta1.Visualization.Type
	(*CreateVisualizationRequest)(nil), // 1: api.v2beta1.CreateVisualizationRequest
	(*Visualization)(nil),              // 2: api.v2beta1.Visualization
}
var file_backend_api_v2beta1_visualization_proto_depIdxs = []int32{
	2, // 0: api.v2beta1.CreateVisualizationRequest.visualization:type_name -> api.v2beta1.Visualization
	0, // 1: api.v2beta1.Visualization.type:type_name -> api.v2beta1.Visualization.Type
	1, // 2: api.v2beta1.VisualizationService.CreateVisualization:input_type -> api.v2beta1.CreateVisualizationRequest
	2, // 3: api.v2beta1.VisualizationService.CreateVisualization:output_type -> api.v2beta1.Visualization
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_backend_api_v2beta1_visualization_proto_init() }
func file_backend_api_v2beta1_visualization_proto_init() {
	if File_backend_api_v2beta1_visualization_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_backend_api_v2beta1_visualization_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVisualizationRequest); i {
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
		file_backend_api_v2beta1_visualization_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Visualization); i {
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
			RawDescriptor: file_backend_api_v2beta1_visualization_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_backend_api_v2beta1_visualization_proto_goTypes,
		DependencyIndexes: file_backend_api_v2beta1_visualization_proto_depIdxs,
		EnumInfos:         file_backend_api_v2beta1_visualization_proto_enumTypes,
		MessageInfos:      file_backend_api_v2beta1_visualization_proto_msgTypes,
	}.Build()
	File_backend_api_v2beta1_visualization_proto = out.File
	file_backend_api_v2beta1_visualization_proto_rawDesc = nil
	file_backend_api_v2beta1_visualization_proto_goTypes = nil
	file_backend_api_v2beta1_visualization_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// VisualizationServiceClient is the client API for VisualizationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VisualizationServiceClient interface {
	CreateVisualization(ctx context.Context, in *CreateVisualizationRequest, opts ...grpc.CallOption) (*Visualization, error)
}

type visualizationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVisualizationServiceClient(cc grpc.ClientConnInterface) VisualizationServiceClient {
	return &visualizationServiceClient{cc}
}

func (c *visualizationServiceClient) CreateVisualization(ctx context.Context, in *CreateVisualizationRequest, opts ...grpc.CallOption) (*Visualization, error) {
	out := new(Visualization)
	err := c.cc.Invoke(ctx, "/api.v2beta1.VisualizationService/CreateVisualization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VisualizationServiceServer is the server API for VisualizationService service.
type VisualizationServiceServer interface {
	CreateVisualization(context.Context, *CreateVisualizationRequest) (*Visualization, error)
}

// UnimplementedVisualizationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedVisualizationServiceServer struct {
}

func (*UnimplementedVisualizationServiceServer) CreateVisualization(context.Context, *CreateVisualizationRequest) (*Visualization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVisualization not implemented")
}

func RegisterVisualizationServiceServer(s *grpc.Server, srv VisualizationServiceServer) {
	s.RegisterService(&_VisualizationService_serviceDesc, srv)
}

func _VisualizationService_CreateVisualization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVisualizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisualizationServiceServer).CreateVisualization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v2beta1.VisualizationService/CreateVisualization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisualizationServiceServer).CreateVisualization(ctx, req.(*CreateVisualizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VisualizationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.v2beta1.VisualizationService",
	HandlerType: (*VisualizationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVisualization",
			Handler:    _VisualizationService_CreateVisualization_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend/api/v2beta1/visualization.proto",
}
