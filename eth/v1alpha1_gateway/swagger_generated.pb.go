// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: eth/v1alpha1/swagger_generated.proto

package eth

import (
	reflect "reflect"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_eth_v1alpha1_swagger_generated_proto protoreflect.FileDescriptor

var file_eth_v1alpha1_swagger_generated_proto_rawDesc = []byte{
	0x0a, 0x24, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73,
	0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x2c, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65,
	0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0xc8, 0x1e, 0x5a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d,
	0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75,
	0x6d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x3b, 0x65, 0x74, 0x68, 0x92, 0x41, 0x8c, 0x1e, 0x12, 0x92, 0x1c, 0x0a, 0x16, 0x45,
	0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x20, 0x53, 0x65, 0x72, 0x65, 0x6e, 0x69, 0x74, 0x79,
	0x20, 0x41, 0x50, 0x49, 0x73, 0x12, 0xa4, 0x1b, 0x5b, 0x21, 0x5b, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x20, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f,
	0x2f, 0x62, 0x61, 0x64, 0x67, 0x65, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x6b, 0x69, 0x74, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x36, 0x32, 0x62, 0x65, 0x30, 0x38, 0x30, 0x39, 0x39, 0x65, 0x39,
	0x65, 0x32, 0x32, 0x38, 0x62, 0x31, 0x36, 0x35, 0x63, 0x32, 0x64, 0x62, 0x61, 0x36, 0x39, 0x63,
	0x36, 0x33, 0x37, 0x65, 0x62, 0x39, 0x63, 0x61, 0x37, 0x61, 0x31, 0x63, 0x61, 0x39, 0x35, 0x65,
	0x66, 0x64, 0x35, 0x34, 0x62, 0x39, 0x66, 0x2e, 0x73, 0x76, 0x67, 0x3f, 0x62, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x3d, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x29, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x6b, 0x69, 0x74, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x2d, 0x6c, 0x61, 0x62, 0x73,
	0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x29, 0x5b,
	0x21, 0x5b, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x37, 0x32, 0x38, 0x38, 0x33, 0x32, 0x32, 0x2f, 0x33, 0x34, 0x34,
	0x37, 0x31, 0x39, 0x36, 0x37, 0x2d, 0x31, 0x64, 0x66, 0x37, 0x38, 0x30, 0x38, 0x61, 0x2d, 0x65,
	0x66, 0x62, 0x62, 0x2d, 0x31, 0x31, 0x65, 0x37, 0x2d, 0x39, 0x30, 0x38, 0x38, 0x2d, 0x65, 0x64,
	0x30, 0x62, 0x30, 0x34, 0x31, 0x35, 0x31, 0x32, 0x39, 0x31, 0x2e, 0x70, 0x6e, 0x67, 0x29, 0x5d,
	0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64,
	0x2e, 0x67, 0x67, 0x2f, 0x4b, 0x53, 0x41, 0x37, 0x72, 0x50, 0x72, 0x29, 0x5b, 0x21, 0x5b, 0x47,
	0x69, 0x74, 0x74, 0x65, 0x72, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x62,
	0x61, 0x64, 0x67, 0x65, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x69, 0x6d, 0x2f,
	0x4a, 0x6f, 0x69, 0x6e, 0x25, 0x32, 0x30, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x73, 0x76, 0x67, 0x29,
	0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x74, 0x65, 0x72,
	0x2e, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62,
	0x73, 0x2f, 0x67, 0x65, 0x74, 0x68, 0x2d, 0x73, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x3f,
	0x75, 0x74, 0x6d, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3d, 0x62, 0x61, 0x64, 0x67, 0x65,
	0x26, 0x75, 0x74, 0x6d, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x3d, 0x62, 0x61, 0x64, 0x67,
	0x65, 0x26, 0x75, 0x74, 0x6d, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x3d, 0x70,
	0x72, 0x2d, 0x62, 0x61, 0x64, 0x67, 0x65, 0x29, 0x5b, 0x21, 0x5b, 0x45, 0x54, 0x48, 0x32, 0x2e,
	0x30, 0x5f, 0x53, 0x70, 0x65, 0x63, 0x5f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x30,
	0x2e, 0x31, 0x32, 0x2e, 0x31, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x69,
	0x6d, 0x67, 0x2e, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x62, 0x61,
	0x64, 0x67, 0x65, 0x2f, 0x45, 0x54, 0x48, 0x32, 0x2e, 0x30, 0x25, 0x32, 0x30, 0x53, 0x70, 0x65,
	0x63, 0x25, 0x32, 0x30, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2d, 0x76, 0x30, 0x2e, 0x31,
	0x32, 0x2e, 0x31, 0x2d, 0x62, 0x6c, 0x75, 0x65, 0x2e, 0x73, 0x76, 0x67, 0x29, 0x5d, 0x28, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2f, 0x65, 0x74, 0x68, 0x32, 0x2e,
	0x30, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x2f, 0x76, 0x30, 0x2e,
	0x31, 0x32, 0x2e, 0x31, 0x29, 0x0a, 0x0a, 0x54, 0x68, 0x69, 0x73, 0x20, 0x73, 0x77, 0x61, 0x67,
	0x67, 0x65, 0x72, 0x20, 0x73, 0x69, 0x74, 0x65, 0x20, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75,
	0x6d, 0x20, 0x53, 0x65, 0x72, 0x65, 0x6e, 0x69, 0x74, 0x79, 0x20, 0x4a, 0x53, 0x4f, 0x4e, 0x20,
	0x41, 0x50, 0x49, 0x2e, 0x0a, 0x0a, 0x23, 0x23, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x20, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x0a, 0x0a, 0x41, 0x74, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x63, 0x6f, 0x72, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x20, 0x53, 0x65, 0x72, 0x65, 0x6e, 0x69, 0x74, 0x79, 0x20, 0x6c, 0x69, 0x65, 0x73,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x22, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x20, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x22, 0x2c, 0x20, 0x61, 0x20, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x2d, 0x6f, 0x66, 0x2d,
	0x73, 0x74, 0x61, 0x6b, 0x65, 0x20, 0x62, 0x61, 0x73, 0x65, 0x64, 0x20, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x20, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x20,
	0x61, 0x20, 0x73, 0x65, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x73, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x22, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73, 0x22,
	0x20, 0x61, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2c, 0x20, 0x77, 0x68, 0x69, 0x63, 0x68, 0x20, 0x61, 0x72, 0x65, 0x20, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x66, 0x75, 0x6c, 0x2c, 0x20, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x20, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20,
	0x66, 0x6f, 0x72, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20,
	0x61, 0x6e, 0x64, 0x20, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x20, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x20, 0x54, 0x68,
	0x65, 0x20, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x20, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x20, 0x72,
	0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x73, 0x20, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x73, 0x75, 0x73, 0x20, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x20, 0x61, 0x63, 0x72, 0x6f,
	0x73, 0x73, 0x20, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x62, 0x61, 0x63, 0x6b, 0x62, 0x6f, 0x6e,
	0x65, 0x20, 0x6f, 0x66, 0x20, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x20, 0x53, 0x65,
	0x72, 0x65, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x20, 0x54, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x20, 0x62, 0x65, 0x6c, 0x6f, 0x77, 0x20, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x20, 0x66, 0x6f, 0x72, 0x20, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x20, 0x6f,
	0x66, 0x20, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x68, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x20, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x6f,
	0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2c, 0x20, 0x63, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2c, 0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x2c, 0x20,
	0x61, 0x6e, 0x64, 0x20, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x20, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x0a, 0x0a, 0x7c, 0x20, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x20, 0x7c, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20,
	0x7c, 0x20, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x7c, 0x20, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x7c, 0x0a, 0x7c, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d,
	0x2d, 0x2d, 0x2d, 0x2d, 0x7c, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x7c, 0x2d,
	0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x7c, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d,
	0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x7c, 0x0a, 0x7c, 0x20, 0x65, 0x74, 0x68, 0x20, 0x7c, 0x20,
	0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x20, 0x7c, 0x20, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x20, 0x7c, 0x20, 0x54, 0x68, 0x69, 0x73, 0x20, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x69, 0x73, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x74, 0x6f,
	0x20, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x20, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63,
	0x61, 0x6c, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x72, 0x65, 0x6c, 0x65, 0x76, 0x61, 0x6e, 0x74,
	0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x20, 0x32, 0x2e, 0x30, 0x20, 0x70, 0x68, 0x61, 0x73, 0x65, 0x20, 0x30, 0x20, 0x62, 0x65, 0x61,
	0x63, 0x6f, 0x6e, 0x20, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2c, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x75,
	0x64, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x68, 0x65, 0x20, 0x6d, 0x6f, 0x73, 0x74, 0x20, 0x72, 0x65,
	0x63, 0x65, 0x6e, 0x74, 0x20, 0x68, 0x65, 0x61, 0x64, 0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2c,
	0x20, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x20, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x20, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x73, 0x2c, 0x20, 0x74, 0x68, 0x65, 0x20, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x20, 0x73, 0x74, 0x61, 0x74, 0x65, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x6d,
	0x6f, 0x72, 0x65, 0x2e, 0x20, 0x7c, 0x0a, 0x7c, 0x20, 0x65, 0x74, 0x68, 0x20, 0x7c, 0x20, 0x4e,
	0x6f, 0x64, 0x65, 0x20, 0x7c, 0x20, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x20, 0x7c,
	0x20, 0x54, 0x68, 0x65, 0x20, 0x4e, 0x6f, 0x64, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x20, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x20, 0x6e, 0x6f, 0x64, 0x65, 0x20, 0x69, 0x74,
	0x73, 0x65, 0x6c, 0x66, 0x2c, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x69, 0x6e, 0x67, 0x20,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x20, 0x61, 0x73, 0x20, 0x77, 0x65, 0x6c, 0x6c, 0x20, 0x61, 0x73, 0x20, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x20, 0x73, 0x79, 0x6e, 0x63, 0x20, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x61, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x20, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x6c, 0x79, 0x20, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x20, 0x6f,
	0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x0a, 0x7c, 0x20, 0x65, 0x74,
	0x68, 0x20, 0x7c, 0x20, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x20, 0x7c, 0x20,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x20, 0x7c, 0x20, 0x54, 0x68, 0x69, 0x73, 0x20,
	0x41, 0x50, 0x49, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x73, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x61, 0x20, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x20, 0x6e, 0x65, 0x65, 0x64, 0x73, 0x20, 0x74,
	0x6f, 0x20, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x20, 0x74, 0x68, 0x72, 0x6f, 0x75,
	0x67, 0x68, 0x6f, 0x75, 0x74, 0x20, 0x69, 0x74, 0x73, 0x20, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79,
	0x63, 0x6c, 0x65, 0x2c, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x72,
	0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x64, 0x20, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x74, 0x68, 0x65, 0x20, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2c, 0x20, 0x69, 0x74, 0x73, 0x20, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x20, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2c, 0x20, 0x61, 0x73, 0x20, 0x77, 0x65, 0x6c, 0x6c, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x70, 0x65,
	0x6e, 0x61, 0x6c, 0x74, 0x69, 0x65, 0x73, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x68, 0x61, 0x76,
	0x65, 0x20, 0x62, 0x65, 0x65, 0x6e, 0x20, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x20, 0x74,
	0x6f, 0x20, 0x69, 0x74, 0x2e, 0x0a, 0x0a, 0x23, 0x23, 0x23, 0x20, 0x4a, 0x53, 0x4f, 0x4e, 0x20,
	0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x0a, 0x0a, 0x54, 0x68, 0x65, 0x20, 0x6d, 0x61, 0x6a,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x20, 0x6f, 0x66, 0x20, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x20, 0x70,
	0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x20, 0x74, 0x79, 0x70, 0x65, 0x73, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x20, 0x61, 0x72, 0x65, 0x20,
	0x65, 0x69, 0x74, 0x68, 0x65, 0x72, 0x20, 0x60, 0x62, 0x79, 0x74, 0x65, 0x73, 0x60, 0x20, 0x6f,
	0x72, 0x20, 0x60, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x60, 0x2e, 0x20, 0x54, 0x68, 0x65, 0x20,
	0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x20, 0x4a, 0x53, 0x4f, 0x4e, 0x20, 0x6d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x68, 0x6f, 0x73, 0x65,
	0x20, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x20, 0x61, 0x72, 0x65, 0x20, 0x61, 0x20, 0x42, 0x61,
	0x73, 0x65, 0x36, 0x34, 0x20, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x20, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x60, 0x62, 0x79, 0x74, 0x65, 0x73, 0x60, 0x2c,
	0x20, 0x6f, 0x72, 0x20, 0x61, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x72, 0x65, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6f, 0x66, 0x20, 0x60,
	0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x60, 0x2e, 0x20, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x20, 0x4a,
	0x61, 0x76, 0x61, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x20, 0x6c, 0x6f, 0x73, 0x65, 0x73, 0x20,
	0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x20, 0x6f, 0x76, 0x65, 0x72, 0x20, 0x5b, 0x4d, 0x41, 0x58, 0x5f, 0x53,
	0x41, 0x46, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x47, 0x45, 0x52, 0x5d, 0x28, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x6d,
	0x6f, 0x7a, 0x69, 0x6c, 0x6c, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x65, 0x6e, 0x2d, 0x55, 0x53,
	0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2f, 0x57, 0x65, 0x62, 0x2f, 0x4a, 0x61, 0x76, 0x61, 0x53, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x2f, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x47,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x2f, 0x4d, 0x41, 0x58, 0x5f, 0x53, 0x41, 0x46, 0x45, 0x5f, 0x49, 0x4e,
	0x54, 0x45, 0x47, 0x45, 0x52, 0x29, 0x2c, 0x20, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x20, 0x6d,
	0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20, 0x61, 0x20, 0x4a, 0x53, 0x4f, 0x4e, 0x20, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x20, 0x69, 0x6e, 0x20, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x20, 0x74, 0x6f,
	0x20, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x65, 0x20, 0x74, 0x68, 0x6f,
	0x73, 0x65, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e, 0x20, 0x49, 0x66, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x6e, 0x6f,
	0x74, 0x20, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x69, 0x73,
	0x20, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x20, 0x73, 0x65, 0x74, 0x20, 0x74, 0x6f, 0x20, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x27, 0x73, 0x20, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x2c, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x20, 0x77, 0x69, 0x6c,
	0x6c, 0x20, 0x62, 0x65, 0x20, 0x6f, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x20, 0x66, 0x72, 0x6f,
	0x6d, 0x20, 0x74, 0x68, 0x65, 0x20, 0x4a, 0x53, 0x4f, 0x4e, 0x20, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x20, 0x65, 0x6e, 0x74, 0x69, 0x72, 0x65, 0x6c, 0x79, 0x2e, 0x0a, 0x0a, 0x46,
	0x6f, 0x72, 0x20, 0x6d, 0x6f, 0x72, 0x65, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x20,
	0x6f, 0x6e, 0x20, 0x4a, 0x53, 0x4f, 0x4e, 0x20, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x20,
	0x66, 0x6f, 0x72, 0x20, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x20, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2c,
	0x20, 0x76, 0x69, 0x65, 0x77, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x6c, 0x65, 0x76, 0x61,
	0x6e, 0x74, 0x20, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x5b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x20, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x20, 0x67, 0x75, 0x69, 0x64, 0x65, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2d, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33, 0x23, 0x6a, 0x73, 0x6f, 0x6e, 0x29, 0x2e, 0x0a, 0x0a, 0x23, 0x23, 0x23,
	0x20, 0x48, 0x65, 0x6c, 0x70, 0x66, 0x75, 0x6c, 0x20, 0x67, 0x52, 0x50, 0x43, 0x20, 0x54, 0x6f,
	0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x0a, 0x0a, 0x2d, 0x20, 0x5b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x27, 0x73,
	0x20, 0x41, 0x50, 0x49, 0x20, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x20, 0x47, 0x75, 0x69, 0x64, 0x65,
	0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x2f, 0x29, 0x0a, 0x2d, 0x20, 0x5b, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x20, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x33, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2d, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x29, 0x0a, 0x2d, 0x20, 0x5b, 0x41, 0x77, 0x65, 0x73, 0x6f,
	0x6d, 0x65, 0x20, 0x67, 0x52, 0x50, 0x43, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x61, 0x77, 0x65, 0x73, 0x6f,
	0x6d, 0x65, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x29, 0x0a, 0x2d, 0x20, 0x5b, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x20, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x20, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x73, 0x3a, 0x20, 0x47, 0x6f, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2d, 0x62,
	0x75, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x75,
	0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x29, 0x0a, 0x2d, 0x20, 0x5b, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x67, 0x52, 0x50, 0x43, 0x20, 0x74, 0x6f, 0x20, 0x4a,
	0x53, 0x4f, 0x4e, 0x2f, 0x48, 0x54, 0x54, 0x50, 0x20, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x20, 0x45,
	0x6e, 0x76, 0x6f, 0x79, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x62, 0x6c,
	0x6f, 0x67, 0x2e, 0x6a, 0x64, 0x72, 0x69, 0x76, 0x65, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x32,
	0x30, 0x31, 0x38, 0x2f, 0x31, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x69,
	0x6e, 0x67, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x74, 0x6f, 0x2d, 0x68, 0x74, 0x74, 0x70, 0x2d,
	0x6a, 0x73, 0x6f, 0x6e, 0x2d, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x2d, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x29, 0x0a, 0x0a, 0x0a, 0x23, 0x23, 0x20, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x0a, 0x57, 0x65, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x70, 0x75, 0x74,
	0x20, 0x61, 0x6c, 0x6c, 0x20, 0x6f, 0x66, 0x20, 0x6f, 0x75, 0x72, 0x20, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x67, 0x75, 0x69, 0x64, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x73, 0x20, 0x69, 0x6e, 0x74, 0x6f, 0x20, 0x5b, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x49,
	0x42, 0x55, 0x54, 0x49, 0x4e, 0x47, 0x2e, 0x6d, 0x64, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72,
	0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x79, 0x73,
	0x6d, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x43, 0x4f,
	0x4e, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x49, 0x4e, 0x47, 0x2e, 0x6d, 0x64, 0x29, 0x21, 0x20,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x20, 0x69, 0x74, 0x20, 0x6f, 0x75, 0x74, 0x20, 0x74, 0x6f, 0x20,
	0x67, 0x65, 0x74, 0x20, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x2e, 0x2a, 0x47, 0x0a, 0x19,
	0x4c, 0x69, 0x63, 0x65, 0x6e, 0x63, 0x65, 0x64, 0x20, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x20, 0x41,
	0x70, 0x61, 0x63, 0x68, 0x65, 0x20, 0x32, 0x2e, 0x30, 0x12, 0x2a, 0x68, 0x74, 0x74, 0x70, 0x3a,
	0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53,
	0x45, 0x2d, 0x32, 0x2e, 0x30, 0x32, 0x08, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a,
	0x0f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x79, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x6e, 0x65, 0x74,
	0x2a, 0x02, 0x02, 0x04, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x32, 0x19, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x77, 0x65, 0x62, 0x2d, 0x74, 0x65, 0x78,
	0x74, 0x32, 0x19, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x77, 0x65, 0x62, 0x2d, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x19,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x77, 0x65, 0x62, 0x2d, 0x74, 0x65, 0x78, 0x74, 0x3a, 0x19, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x77, 0x65, 0x62, 0x2d,
	0x6a, 0x73, 0x6f, 0x6e, 0x72, 0x50, 0x0a, 0x24, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x20, 0x32, 0x2e, 0x30, 0x20, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x20, 0x6f, 0x6e, 0x20, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x12, 0x28, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2f, 0x65, 0x74, 0x68, 0x32, 0x2e, 0x30,
	0x2d, 0x73, 0x70, 0x65, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_eth_v1alpha1_swagger_generated_proto_goTypes = []interface{}{}
var file_eth_v1alpha1_swagger_generated_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eth_v1alpha1_swagger_generated_proto_init() }
func file_eth_v1alpha1_swagger_generated_proto_init() {
	if File_eth_v1alpha1_swagger_generated_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eth_v1alpha1_swagger_generated_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eth_v1alpha1_swagger_generated_proto_goTypes,
		DependencyIndexes: file_eth_v1alpha1_swagger_generated_proto_depIdxs,
	}.Build()
	File_eth_v1alpha1_swagger_generated_proto = out.File
	file_eth_v1alpha1_swagger_generated_proto_rawDesc = nil
	file_eth_v1alpha1_swagger_generated_proto_goTypes = nil
	file_eth_v1alpha1_swagger_generated_proto_depIdxs = nil
}
