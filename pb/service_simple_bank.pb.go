// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: service_simple_bank.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_simple_bank_proto protoreflect.FileDescriptor

var file_service_simple_bank_proto_rawDesc = string([]byte{
	0x0a, 0x19, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x5f, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x15, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x70,
	0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xd2, 0x0b, 0x0a, 0x0a, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x61, 0x6e, 0x6b,
	0x12, 0xb7, 0x03, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xf9,
	0x02, 0x92, 0x41, 0xe1, 0x02, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x20, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61,
	0x20, 0x6e, 0x65, 0x77, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x1a, 0x5a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x73, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77,
	0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x20, 0x75, 0x73, 0x65,
	0x72, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x8d, 0x01,
	0x0a, 0x03, 0x32, 0x30, 0x31, 0x12, 0x85, 0x01, 0x0a, 0x19, 0x55, 0x73, 0x65, 0x72, 0x20, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75,
	0x6c, 0x6c, 0x79, 0x22, 0x68, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x54, 0x7b, 0x22, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x75, 0x73, 0x65, 0x72, 0x31, 0x32, 0x33, 0x22, 0x2c, 0x20,
	0x22, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x6a, 0x6f, 0x68,
	0x6e, 0x5f, 0x64, 0x6f, 0x65, 0x22, 0x2c, 0x20, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x22, 0x3a, 0x20, 0x22, 0x32, 0x30, 0x32, 0x35, 0x2d, 0x30, 0x32, 0x2d, 0x32,
	0x38, 0x54, 0x31, 0x32, 0x3a, 0x30, 0x30, 0x3a, 0x30, 0x30, 0x5a, 0x22, 0x7d, 0x4a, 0x47, 0x0a,
	0x03, 0x34, 0x30, 0x39, 0x12, 0x40, 0x0a, 0x3e, 0x43, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74,
	0x20, 0x2d, 0x20, 0x55, 0x73, 0x65, 0x72, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x73, 0x61, 0x6d, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x6f,
	0x72, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x20, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x20,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22,
	0x09, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x82, 0x03, 0x0a, 0x0a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xc4, 0x02, 0x92, 0x41, 0xac, 0x02, 0x0a,
	0x0f, 0x55, 0x73, 0x65, 0x72, 0x20, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x44, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x20,
	0x61, 0x6e, 0x64, 0x20, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x4a, 0x8d, 0x01, 0x0a, 0x03, 0x32, 0x30, 0x31, 0x12, 0x85, 0x01, 0x0a, 0x19,
	0x55, 0x73, 0x65, 0x72, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x20, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x22, 0x68, 0x0a, 0x10, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x54, 0x7b,
	0x22, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x75, 0x73, 0x65, 0x72,
	0x31, 0x32, 0x33, 0x22, 0x2c, 0x20, 0x22, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x3a, 0x20, 0x22, 0x6a, 0x6f, 0x68, 0x6e, 0x5f, 0x64, 0x6f, 0x65, 0x22, 0x2c, 0x20, 0x22, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x3a, 0x20, 0x22, 0x32, 0x30, 0x32,
	0x35, 0x2d, 0x30, 0x32, 0x2d, 0x32, 0x38, 0x54, 0x31, 0x32, 0x3a, 0x30, 0x30, 0x3a, 0x30, 0x30,
	0x5a, 0x22, 0x7d, 0x4a, 0x26, 0x0a, 0x03, 0x34, 0x30, 0x39, 0x12, 0x1f, 0x0a, 0x1d, 0x43, 0x6f,
	0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x20, 0x2d, 0x20, 0x55, 0x73, 0x65, 0x72, 0x20, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x20, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0e, 0x3a, 0x01, 0x2a, 0x32, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12,
	0xc2, 0x03, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x2e,
	0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x87, 0x03, 0x92, 0x41, 0xef,
	0x02, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x13, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x20,
	0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x47, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x73, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x77, 0x69, 0x74,
	0x68, 0x20, 0x74, 0x68, 0x65, 0x69, 0x72, 0x20, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x20,
	0x61, 0x6e, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4a,
	0xcf, 0x01, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0xc7, 0x01, 0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x22, 0xb2, 0x01, 0x0a,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f,
	0x6e, 0x12, 0x9d, 0x01, 0x7b, 0x22, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x3a, 0x20, 0x22, 0x65, 0x79, 0x4a, 0x68, 0x62, 0x47, 0x63, 0x69, 0x4f, 0x69,
	0x4a, 0x49, 0x55, 0x7a, 0x49, 0x31, 0x4e, 0x69, 0x49, 0x73, 0x49, 0x6e, 0x52, 0x35, 0x63, 0x43,
	0x49, 0x36, 0x49, 0x6b, 0x70, 0x58, 0x56, 0x43, 0x4a, 0x39, 0x2e, 0x2e, 0x2e, 0x22, 0x2c, 0x20,
	0x22, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3a,
	0x20, 0x22, 0x65, 0x79, 0x4a, 0x68, 0x62, 0x47, 0x63, 0x69, 0x4f, 0x69, 0x4a, 0x49, 0x55, 0x7a,
	0x49, 0x31, 0x4e, 0x69, 0x49, 0x73, 0x49, 0x6e, 0x52, 0x35, 0x63, 0x43, 0x49, 0x36, 0x49, 0x6b,
	0x70, 0x58, 0x56, 0x43, 0x4a, 0x39, 0x2e, 0x2e, 0x2e, 0x22, 0x2c, 0x20, 0x22, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x22, 0x3a, 0x20, 0x22, 0x32, 0x30, 0x32, 0x35, 0x2d,
	0x30, 0x33, 0x2d, 0x30, 0x31, 0x54, 0x31, 0x32, 0x3a, 0x30, 0x30, 0x3a, 0x30, 0x30, 0x5a, 0x22,
	0x7d, 0x4a, 0x2b, 0x0a, 0x03, 0x34, 0x30, 0x31, 0x12, 0x24, 0x0a, 0x22, 0x55, 0x6e, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x2d, 0x20, 0x49, 0x6e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x20, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x62, 0x00,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x1a, 0xbf, 0x01, 0x92, 0x41, 0xbb, 0x01, 0x0a, 0x0a, 0x53, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x12, 0x44, 0x42, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67,
	0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x75, 0x73,
	0x65, 0x72, 0x20, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x61, 0x6e,
	0x64, 0x20, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x20, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x67, 0x0a,
	0x27, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x20, 0x6d, 0x6f, 0x72, 0x65, 0x20, 0x61, 0x62, 0x6f, 0x75,
	0x74, 0x20, 0x74, 0x68, 0x65, 0x20, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x61, 0x6e, 0x6b,
	0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x61, 0x64,
	0x69, 0x74, 0x79, 0x61, 0x61, 0x32, 0x36, 0x30, 0x36, 0x2f, 0x42, 0x61, 0x6e, 0x6b, 0x2d, 0x41,
	0x50, 0x49, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x52, 0x45, 0x41,
	0x44, 0x4d, 0x45, 0x2e, 0x6d, 0x64, 0x42, 0xcf, 0x08, 0x92, 0x41, 0xa6, 0x08, 0x12, 0xfc, 0x01,
	0x0a, 0x0f, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x20, 0x42, 0x61, 0x6e, 0x6b, 0x20, 0x41, 0x50,
	0x49, 0x12, 0x24, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x20, 0x73, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x20, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x20, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4c, 0x0a, 0x11, 0x41, 0x61, 0x64, 0x69, 0x74,
	0x79, 0x61, 0x61, 0x20, 0x4e, 0x61, 0x67, 0x61, 0x72, 0x6a, 0x61, 0x6e, 0x12, 0x1f, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x41, 0x61, 0x64, 0x69, 0x74, 0x79, 0x61, 0x61, 0x32, 0x36, 0x30, 0x36, 0x1a, 0x16, 0x61,
	0x61, 0x64, 0x69, 0x74, 0x79, 0x61, 0x61, 0x32, 0x36, 0x30, 0x36, 0x40, 0x67, 0x6d, 0x61, 0x69,
	0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x2a, 0x49, 0x0a, 0x0b, 0x4d, 0x49, 0x54, 0x20, 0x4c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x61, 0x64, 0x69, 0x74, 0x79, 0x61,
	0x61, 0x32, 0x36, 0x30, 0x36, 0x2f, 0x42, 0x61, 0x6e, 0x6b, 0x2d, 0x41, 0x50, 0x49, 0x2f, 0x62,
	0x6c, 0x6f, 0x62, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45,
	0x32, 0x03, 0x31, 0x2e, 0x31, 0x3a, 0x25, 0x0a, 0x0c, 0x78, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x15, 0x1a, 0x13, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2a, 0x02, 0x01, 0x02,
	0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x6a, 0x73, 0x6f, 0x6e, 0x52, 0x3e, 0x0a, 0x03, 0x34, 0x30, 0x30, 0x12, 0x37, 0x0a, 0x35, 0x42,
	0x61, 0x64, 0x20, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x2d, 0x20, 0x54, 0x68, 0x65,
	0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x73, 0x20, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x52, 0x4e, 0x0a, 0x03, 0x34, 0x30, 0x31, 0x12, 0x47, 0x0a, 0x45, 0x55,
	0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x2d, 0x20, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x66, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x20, 0x6f, 0x72, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x64, 0x6f, 0x65, 0x73,
	0x6e, 0x27, 0x74, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x50, 0x0a, 0x03, 0x34, 0x30, 0x33, 0x12, 0x49, 0x0a, 0x47, 0x46,
	0x6f, 0x72, 0x62, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x20, 0x2d, 0x20, 0x54, 0x68, 0x65, 0x20, 0x75,
	0x73, 0x65, 0x72, 0x20, 0x69, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x20, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x39, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x32, 0x0a,
	0x30, 0x4e, 0x6f, 0x74, 0x20, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x20, 0x2d, 0x20, 0x54, 0x68, 0x65,
	0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x6e, 0x27, 0x74, 0x20, 0x65, 0x78, 0x69, 0x73,
	0x74, 0x52, 0x43, 0x0a, 0x03, 0x35, 0x30, 0x30, 0x12, 0x3c, 0x0a, 0x3a, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x20, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x20, 0x2d, 0x20, 0x53, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x20, 0x77, 0x65,
	0x6e, 0x74, 0x20, 0x77, 0x72, 0x6f, 0x6e, 0x67, 0x20, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5a, 0xa6, 0x02, 0x0a, 0x4f, 0x0a, 0x0a, 0x42, 0x65, 0x61,
	0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x41, 0x08, 0x02, 0x12, 0x2c, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2c, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x27,
	0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x3a, 0x20, 0x27, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x0a, 0xd2, 0x01, 0x0a, 0x06, 0x4f,
	0x41, 0x75, 0x74, 0x68, 0x32, 0x12, 0xc7, 0x01, 0x08, 0x03, 0x28, 0x04, 0x32, 0x30, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x3a, 0x2c,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x62, 0x61, 0x6e,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x61, 0x0a, 0x25,
	0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1c, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x20,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x20, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x0a, 0x1a, 0x0a, 0x04, 0x72, 0x65, 0x61, 0x64, 0x12, 0x12, 0x47,
	0x72, 0x61, 0x6e, 0x74, 0x73, 0x20, 0x72, 0x65, 0x61, 0x64, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x0a, 0x1c, 0x0a, 0x05, 0x77, 0x72, 0x69, 0x74, 0x65, 0x12, 0x13, 0x47, 0x72, 0x61, 0x6e,
	0x74, 0x73, 0x20, 0x77, 0x72, 0x69, 0x74, 0x65, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x62,
	0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12,
	0x00, 0x72, 0x60, 0x0a, 0x2f, 0x46, 0x69, 0x6e, 0x64, 0x20, 0x6d, 0x6f, 0x72, 0x65, 0x20, 0x69,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x20, 0x42, 0x61, 0x6e, 0x6b,
	0x20, 0x41, 0x50, 0x49, 0x12, 0x2d, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x61, 0x64, 0x69, 0x74, 0x79, 0x61,
	0x61, 0x32, 0x36, 0x30, 0x36, 0x2f, 0x42, 0x61, 0x6e, 0x6b, 0x2d, 0x41, 0x50, 0x49, 0x2f, 0x77,
	0x69, 0x6b, 0x69, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x41, 0x61, 0x64, 0x69, 0x74, 0x79, 0x61, 0x61, 0x32, 0x36, 0x30, 0x36, 0x2f, 0x42, 0x61, 0x6e,
	0x6b, 0x2d, 0x41, 0x50, 0x49, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_service_simple_bank_proto_goTypes = []any{
	(*CreateUserRequest)(nil),  // 0: pb.CreateUserRequest
	(*UpdateUserRequest)(nil),  // 1: pb.UpdateUserRequest
	(*LoginUserRequest)(nil),   // 2: pb.LoginUserRequest
	(*CreateUserResponse)(nil), // 3: pb.CreateUserResponse
	(*UpdateUserResponse)(nil), // 4: pb.UpdateUserResponse
	(*LoginUserResponse)(nil),  // 5: pb.LoginUserResponse
}
var file_service_simple_bank_proto_depIdxs = []int32{
	0, // 0: pb.SimpleBank.CreateUser:input_type -> pb.CreateUserRequest
	1, // 1: pb.SimpleBank.UpdateUser:input_type -> pb.UpdateUserRequest
	2, // 2: pb.SimpleBank.LoginUser:input_type -> pb.LoginUserRequest
	3, // 3: pb.SimpleBank.CreateUser:output_type -> pb.CreateUserResponse
	4, // 4: pb.SimpleBank.UpdateUser:output_type -> pb.UpdateUserResponse
	5, // 5: pb.SimpleBank.LoginUser:output_type -> pb.LoginUserResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_simple_bank_proto_init() }
func file_service_simple_bank_proto_init() {
	if File_service_simple_bank_proto != nil {
		return
	}
	file_rpc_create_user_proto_init()
	file_rpc_login_user_proto_init()
	file_rpc_update_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_service_simple_bank_proto_rawDesc), len(file_service_simple_bank_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_simple_bank_proto_goTypes,
		DependencyIndexes: file_service_simple_bank_proto_depIdxs,
	}.Build()
	File_service_simple_bank_proto = out.File
	file_service_simple_bank_proto_goTypes = nil
	file_service_simple_bank_proto_depIdxs = nil
}
