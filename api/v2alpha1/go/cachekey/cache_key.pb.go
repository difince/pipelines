// Copyright 2021 The Kubeflow Authors
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
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: cache_key.proto

package cachekey

import (
	_struct "github.com/golang/protobuf/ptypes/struct"
	pipelinespec "github.com/kubeflow/pipelines/api/v2alpha1/go/pipelinespec"
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

type CacheKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InputArtifactNames   map[string]*ArtifactNameList             `protobuf:"bytes,1,rep,name=inputArtifactNames,proto3" json:"inputArtifactNames,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	InputParameters      map[string]*pipelinespec.Value           `protobuf:"bytes,2,rep,name=inputParameters,proto3" json:"inputParameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	OutputArtifactsSpec  map[string]*pipelinespec.RuntimeArtifact `protobuf:"bytes,3,rep,name=outputArtifactsSpec,proto3" json:"outputArtifactsSpec,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	OutputParametersSpec map[string]string                        `protobuf:"bytes,4,rep,name=outputParametersSpec,proto3" json:"outputParametersSpec,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ContainerSpec        *ContainerSpec                           `protobuf:"bytes,5,opt,name=containerSpec,proto3" json:"containerSpec,omitempty"`
	InputParameterValues map[string]*_struct.Value                `protobuf:"bytes,6,rep,name=input_parameter_values,json=inputParameterValues,proto3" json:"input_parameter_values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CacheKey) Reset() {
	*x = CacheKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_key_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheKey) ProtoMessage() {}

func (x *CacheKey) ProtoReflect() protoreflect.Message {
	mi := &file_cache_key_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheKey.ProtoReflect.Descriptor instead.
func (*CacheKey) Descriptor() ([]byte, []int) {
	return file_cache_key_proto_rawDescGZIP(), []int{0}
}

func (x *CacheKey) GetInputArtifactNames() map[string]*ArtifactNameList {
	if x != nil {
		return x.InputArtifactNames
	}
	return nil
}

func (x *CacheKey) GetInputParameters() map[string]*pipelinespec.Value {
	if x != nil {
		return x.InputParameters
	}
	return nil
}

func (x *CacheKey) GetOutputArtifactsSpec() map[string]*pipelinespec.RuntimeArtifact {
	if x != nil {
		return x.OutputArtifactsSpec
	}
	return nil
}

func (x *CacheKey) GetOutputParametersSpec() map[string]string {
	if x != nil {
		return x.OutputParametersSpec
	}
	return nil
}

func (x *CacheKey) GetContainerSpec() *ContainerSpec {
	if x != nil {
		return x.ContainerSpec
	}
	return nil
}

func (x *CacheKey) GetInputParameterValues() map[string]*_struct.Value {
	if x != nil {
		return x.InputParameterValues
	}
	return nil
}

type ContainerSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image   string   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	CmdArgs []string `protobuf:"bytes,2,rep,name=cmdArgs,proto3" json:"cmdArgs,omitempty"`
}

func (x *ContainerSpec) Reset() {
	*x = ContainerSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_key_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerSpec) ProtoMessage() {}

func (x *ContainerSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cache_key_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerSpec.ProtoReflect.Descriptor instead.
func (*ContainerSpec) Descriptor() ([]byte, []int) {
	return file_cache_key_proto_rawDescGZIP(), []int{1}
}

func (x *ContainerSpec) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *ContainerSpec) GetCmdArgs() []string {
	if x != nil {
		return x.CmdArgs
	}
	return nil
}

type ArtifactNameList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArtifactNames []string `protobuf:"bytes,1,rep,name=artifactNames,proto3" json:"artifactNames,omitempty"`
}

func (x *ArtifactNameList) Reset() {
	*x = ArtifactNameList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_key_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtifactNameList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtifactNameList) ProtoMessage() {}

func (x *ArtifactNameList) ProtoReflect() protoreflect.Message {
	mi := &file_cache_key_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArtifactNameList.ProtoReflect.Descriptor instead.
func (*ArtifactNameList) Descriptor() ([]byte, []int) {
	return file_cache_key_proto_rawDescGZIP(), []int{2}
}

func (x *ArtifactNameList) GetArtifactNames() []string {
	if x != nil {
		return x.ArtifactNames
	}
	return nil
}

var File_cache_key_proto protoreflect.FileDescriptor

var file_cache_key_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x6d, 0x6c, 0x5f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x86, 0x08, 0x0a, 0x08, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4b, 0x65, 0x79, 0x12,
	0x5e, 0x0a, 0x12, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6d, 0x6c,
	0x5f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x4b, 0x65, 0x79, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x12, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12,
	0x55, 0x0a, 0x0f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6d, 0x6c, 0x5f, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4b, 0x65, 0x79,
	0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x61, 0x0a, 0x13, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x53, 0x70, 0x65, 0x63, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6d, 0x6c, 0x5f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x73, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4b, 0x65, 0x79, 0x2e, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x53, 0x70, 0x65, 0x63, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x13, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x64, 0x0a, 0x14, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x53, 0x70, 0x65,
	0x63, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6d, 0x6c, 0x5f, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4b, 0x65, 0x79, 0x2e,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x53, 0x70, 0x65, 0x63, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x14, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x41, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x6c, 0x5f, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53,
	0x70, 0x65, 0x63, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x66, 0x0a, 0x16, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6d, 0x6c, 0x5f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x73, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4b, 0x65, 0x79, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x14, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x65, 0x0a, 0x17, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x6c, 0x5f, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x57, 0x0a, 0x14, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6c, 0x5f,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x65, 0x0a, 0x18, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x53, 0x70, 0x65,
	0x63, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x33, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x6c, 0x5f, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x41, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x47, 0x0a, 0x19, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x53, 0x70, 0x65, 0x63, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x5f, 0x0a, 0x19, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3f, 0x0a, 0x0d, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6d, 0x64, 0x41, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6d, 0x64, 0x41, 0x72, 0x67, 0x73, 0x22, 0x38, 0x0a, 0x10,
	0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x0d, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x6b, 0x65, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cache_key_proto_rawDescOnce sync.Once
	file_cache_key_proto_rawDescData = file_cache_key_proto_rawDesc
)

func file_cache_key_proto_rawDescGZIP() []byte {
	file_cache_key_proto_rawDescOnce.Do(func() {
		file_cache_key_proto_rawDescData = protoimpl.X.CompressGZIP(file_cache_key_proto_rawDescData)
	})
	return file_cache_key_proto_rawDescData
}

var file_cache_key_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cache_key_proto_goTypes = []interface{}{
	(*CacheKey)(nil),                     // 0: ml_pipelines.CacheKey
	(*ContainerSpec)(nil),                // 1: ml_pipelines.ContainerSpec
	(*ArtifactNameList)(nil),             // 2: ml_pipelines.ArtifactNameList
	nil,                                  // 3: ml_pipelines.CacheKey.InputArtifactNamesEntry
	nil,                                  // 4: ml_pipelines.CacheKey.InputParametersEntry
	nil,                                  // 5: ml_pipelines.CacheKey.OutputArtifactsSpecEntry
	nil,                                  // 6: ml_pipelines.CacheKey.OutputParametersSpecEntry
	nil,                                  // 7: ml_pipelines.CacheKey.InputParameterValuesEntry
	(*pipelinespec.Value)(nil),           // 8: ml_pipelines.Value
	(*pipelinespec.RuntimeArtifact)(nil), // 9: ml_pipelines.RuntimeArtifact
	(*_struct.Value)(nil),                // 10: google.protobuf.Value
}
var file_cache_key_proto_depIdxs = []int32{
	3,  // 0: ml_pipelines.CacheKey.inputArtifactNames:type_name -> ml_pipelines.CacheKey.InputArtifactNamesEntry
	4,  // 1: ml_pipelines.CacheKey.inputParameters:type_name -> ml_pipelines.CacheKey.InputParametersEntry
	5,  // 2: ml_pipelines.CacheKey.outputArtifactsSpec:type_name -> ml_pipelines.CacheKey.OutputArtifactsSpecEntry
	6,  // 3: ml_pipelines.CacheKey.outputParametersSpec:type_name -> ml_pipelines.CacheKey.OutputParametersSpecEntry
	1,  // 4: ml_pipelines.CacheKey.containerSpec:type_name -> ml_pipelines.ContainerSpec
	7,  // 5: ml_pipelines.CacheKey.input_parameter_values:type_name -> ml_pipelines.CacheKey.InputParameterValuesEntry
	2,  // 6: ml_pipelines.CacheKey.InputArtifactNamesEntry.value:type_name -> ml_pipelines.ArtifactNameList
	8,  // 7: ml_pipelines.CacheKey.InputParametersEntry.value:type_name -> ml_pipelines.Value
	9,  // 8: ml_pipelines.CacheKey.OutputArtifactsSpecEntry.value:type_name -> ml_pipelines.RuntimeArtifact
	10, // 9: ml_pipelines.CacheKey.InputParameterValuesEntry.value:type_name -> google.protobuf.Value
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_cache_key_proto_init() }
func file_cache_key_proto_init() {
	if File_cache_key_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cache_key_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CacheKey); i {
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
		file_cache_key_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerSpec); i {
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
		file_cache_key_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArtifactNameList); i {
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
			RawDescriptor: file_cache_key_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cache_key_proto_goTypes,
		DependencyIndexes: file_cache_key_proto_depIdxs,
		MessageInfos:      file_cache_key_proto_msgTypes,
	}.Build()
	File_cache_key_proto = out.File
	file_cache_key_proto_rawDesc = nil
	file_cache_key_proto_goTypes = nil
	file_cache_key_proto_depIdxs = nil
}
