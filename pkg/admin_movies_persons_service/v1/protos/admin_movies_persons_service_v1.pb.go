// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: admin_movies_persons_service_v1.proto

package protos

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_admin_movies_persons_service_v1_proto protoreflect.FileDescriptor

var file_admin_movies_persons_service_v1_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x5f, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x76,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x76, 0x31, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0x8f, 0x0e, 0x0a, 0x16, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x31, 0x12, 0x79, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x2f, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x73, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x84, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x73, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x97, 0x01,
	0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x42, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x73, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2f, 0x7b, 0x4e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0xb3, 0x01, 0x0a, 0x14, 0x49, 0x73, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x49, 0x44, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73,
	0x12, 0x39, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x5f,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x49, 0x73, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x49, 0x44, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x73, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x49, 0x44, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12,
	0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x7b, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x49, 0x44, 0x7d, 0x2f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x96, 0x01,
	0x0a, 0x0e, 0x49, 0x73, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73,
	0x12, 0x33, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x5f,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x49, 0x73, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x73, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x9a, 0x01, 0x0a, 0x0f, 0x49, 0x73, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x34, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x73, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x35, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x5f,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x49, 0x73, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12,
	0x12, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x78, 0x69,
	0x73, 0x74, 0x73, 0x12, 0xd1, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x37, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x6a, 0x92, 0x41, 0x46,
	0x4a, 0x44, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x3d, 0x0a, 0x1e, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x20,
	0x6e, 0x6f, 0x74, 0x20, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1b, 0x0a, 0x19, 0x1a, 0x17, 0x23,
	0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x7b, 0x49, 0x44, 0x7d, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0xbe, 0x01, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x63, 0x92, 0x41, 0x46, 0x4a, 0x44, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12,
	0x3d, 0x0a, 0x1e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e,
	0x20, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x12, 0x1b, 0x0a, 0x19, 0x1a, 0x17, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x2f, 0x7b, 0x49, 0x44, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0xfd, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65,
	0x22, 0x85, 0x01, 0x92, 0x41, 0x6d, 0x4a, 0x6b, 0x0a, 0x03, 0x34, 0x30, 0x39, 0x12, 0x64, 0x0a,
	0x45, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x20, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x20, 0x28, 0x77,
	0x68, 0x65, 0x6e, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x20, 0x65,
	0x78, 0x65, 0x70, 0x74, 0x20, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x20, 0x73, 0x61, 0x6d, 0x65, 0x29,
	0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x19, 0x1a, 0x17, 0x23, 0x2f, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0xd7, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x32, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x5f, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x63, 0x65, 0x22, 0x5d, 0x92, 0x41, 0x47, 0x4a, 0x45, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12,
	0x3e, 0x0a, 0x1f, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e,
	0x20, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x12, 0x1b, 0x0a, 0x19, 0x1a, 0x17, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x2a, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x73, 0x42, 0xc8, 0x02, 0x5a, 0x26, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x92, 0x41, 0x9c,
	0x02, 0x12, 0x64, 0x0a, 0x1c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x20, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x73, 0x20, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x22, 0x3f, 0x0a, 0x07, 0x46, 0x61, 0x6c, 0x6f, 0x6b, 0x75, 0x74, 0x12, 0x1a, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x46, 0x61, 0x6c, 0x6f, 0x6b, 0x75, 0x74, 0x1a, 0x18, 0x74, 0x69, 0x6d, 0x75, 0x72, 0x2e,
	0x73, 0x69, 0x6e, 0x65, 0x6c, 0x6e, 0x69, 0x6b, 0x40, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x72, 0x75, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x01, 0x01, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x52, 0x50,
	0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x49, 0x0a, 0x2a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65,
	0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78, 0x69,
	0x73, 0x74, 0x2e, 0x12, 0x1b, 0x0a, 0x19, 0x1a, 0x17, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x3b, 0x0a, 0x03, 0x35, 0x30, 0x30, 0x12, 0x34, 0x0a, 0x15, 0x53, 0x6f, 0x6d, 0x65, 0x74,
	0x68, 0x69, 0x6e, 0x67, 0x20, 0x77, 0x65, 0x6e, 0x74, 0x20, 0x77, 0x72, 0x6f, 0x6e, 0x67, 0x2e,
	0x12, 0x1b, 0x0a, 0x19, 0x1a, 0x17, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_admin_movies_persons_service_v1_proto_goTypes = []interface{}{
	(*GetPersonsRequest)(nil),            // 0: admin_movies_persons_service.GetPersonsRequest
	(*SearchPersonRequest)(nil),          // 1: admin_movies_persons_service.SearchPersonRequest
	(*SearchPersonByNameRequest)(nil),    // 2: admin_movies_persons_service.SearchPersonByNameRequest
	(*IsPersonWithIDExistsRequest)(nil),  // 3: admin_movies_persons_service.IsPersonWithIDExistsRequest
	(*IsPersonExistsRequest)(nil),        // 4: admin_movies_persons_service.IsPersonExistsRequest
	(*IsPersonsExistsRequest)(nil),       // 5: admin_movies_persons_service.IsPersonsExistsRequest
	(*UpdatePersonFieldsRequest)(nil),    // 6: admin_movies_persons_service.UpdatePersonFieldsRequest
	(*UpdatePersonRequest)(nil),          // 7: admin_movies_persons_service.UpdatePersonRequest
	(*CreatePersonRequest)(nil),          // 8: admin_movies_persons_service.CreatePersonRequest
	(*DeletePersonsRequest)(nil),         // 9: admin_movies_persons_service.DeletePersonsRequest
	(*Persons)(nil),                      // 10: admin_movies_persons_service.Persons
	(*IsPersonWithIDExistsResponse)(nil), // 11: admin_movies_persons_service.IsPersonWithIDExistsResponse
	(*IsPersonExistsResponse)(nil),       // 12: admin_movies_persons_service.IsPersonExistsResponse
	(*IsPersonsExistsResponse)(nil),      // 13: admin_movies_persons_service.IsPersonsExistsResponse
	(*emptypb.Empty)(nil),                // 14: google.protobuf.Empty
	(*CreatePersonResponce)(nil),         // 15: admin_movies_persons_service.CreatePersonResponce
	(*DeletePersonsResponce)(nil),        // 16: admin_movies_persons_service.DeletePersonsResponce
}
var file_admin_movies_persons_service_v1_proto_depIdxs = []int32{
	0,  // 0: admin_movies_persons_service.moviesPersonsServiceV1.GetPersons:input_type -> admin_movies_persons_service.GetPersonsRequest
	1,  // 1: admin_movies_persons_service.moviesPersonsServiceV1.SearchPerson:input_type -> admin_movies_persons_service.SearchPersonRequest
	2,  // 2: admin_movies_persons_service.moviesPersonsServiceV1.SearchPersonByName:input_type -> admin_movies_persons_service.SearchPersonByNameRequest
	3,  // 3: admin_movies_persons_service.moviesPersonsServiceV1.IsPersonWithIDExists:input_type -> admin_movies_persons_service.IsPersonWithIDExistsRequest
	4,  // 4: admin_movies_persons_service.moviesPersonsServiceV1.IsPersonExists:input_type -> admin_movies_persons_service.IsPersonExistsRequest
	5,  // 5: admin_movies_persons_service.moviesPersonsServiceV1.IsPersonsExists:input_type -> admin_movies_persons_service.IsPersonsExistsRequest
	6,  // 6: admin_movies_persons_service.moviesPersonsServiceV1.UpdatePersonFields:input_type -> admin_movies_persons_service.UpdatePersonFieldsRequest
	7,  // 7: admin_movies_persons_service.moviesPersonsServiceV1.UpdatePerson:input_type -> admin_movies_persons_service.UpdatePersonRequest
	8,  // 8: admin_movies_persons_service.moviesPersonsServiceV1.CreatePerson:input_type -> admin_movies_persons_service.CreatePersonRequest
	9,  // 9: admin_movies_persons_service.moviesPersonsServiceV1.DeletePersons:input_type -> admin_movies_persons_service.DeletePersonsRequest
	10, // 10: admin_movies_persons_service.moviesPersonsServiceV1.GetPersons:output_type -> admin_movies_persons_service.Persons
	10, // 11: admin_movies_persons_service.moviesPersonsServiceV1.SearchPerson:output_type -> admin_movies_persons_service.Persons
	10, // 12: admin_movies_persons_service.moviesPersonsServiceV1.SearchPersonByName:output_type -> admin_movies_persons_service.Persons
	11, // 13: admin_movies_persons_service.moviesPersonsServiceV1.IsPersonWithIDExists:output_type -> admin_movies_persons_service.IsPersonWithIDExistsResponse
	12, // 14: admin_movies_persons_service.moviesPersonsServiceV1.IsPersonExists:output_type -> admin_movies_persons_service.IsPersonExistsResponse
	13, // 15: admin_movies_persons_service.moviesPersonsServiceV1.IsPersonsExists:output_type -> admin_movies_persons_service.IsPersonsExistsResponse
	14, // 16: admin_movies_persons_service.moviesPersonsServiceV1.UpdatePersonFields:output_type -> google.protobuf.Empty
	14, // 17: admin_movies_persons_service.moviesPersonsServiceV1.UpdatePerson:output_type -> google.protobuf.Empty
	15, // 18: admin_movies_persons_service.moviesPersonsServiceV1.CreatePerson:output_type -> admin_movies_persons_service.CreatePersonResponce
	16, // 19: admin_movies_persons_service.moviesPersonsServiceV1.DeletePersons:output_type -> admin_movies_persons_service.DeletePersonsResponce
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_admin_movies_persons_service_v1_proto_init() }
func file_admin_movies_persons_service_v1_proto_init() {
	if File_admin_movies_persons_service_v1_proto != nil {
		return
	}
	file_admin_movies_persons_service_v1_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_admin_movies_persons_service_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_movies_persons_service_v1_proto_goTypes,
		DependencyIndexes: file_admin_movies_persons_service_v1_proto_depIdxs,
	}.Build()
	File_admin_movies_persons_service_v1_proto = out.File
	file_admin_movies_persons_service_v1_proto_rawDesc = nil
	file_admin_movies_persons_service_v1_proto_goTypes = nil
	file_admin_movies_persons_service_v1_proto_depIdxs = nil
}
