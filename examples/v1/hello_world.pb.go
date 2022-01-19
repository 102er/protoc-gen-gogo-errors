// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: examples/v1/hello_world.proto

package v1

import (
	fmt "fmt"
	_ "github.com/102er/protoc-gen-gogo-errors/errors"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ErrorReason int32

const (
	ErrorReason_USER_NOT_FOUND   ErrorReason = 0
	ErrorReason_USER_NOT_FOUND_t ErrorReason = 1
)

var ErrorReason_name = map[int32]string{
	0: "USER_NOT_FOUND",
	1: "USER_NOT_FOUND_t",
}

var ErrorReason_value = map[string]int32{
	"USER_NOT_FOUND":   0,
	"USER_NOT_FOUND_t": 1,
}

func (x ErrorReason) String() string {
	return proto.EnumName(ErrorReason_name, int32(x))
}

func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c3336ab888c6aac6, []int{0}
}

func init() {
	proto.RegisterEnum("errors.ErrorReason", ErrorReason_name, ErrorReason_value)
}

func init() { proto.RegisterFile("examples/v1/hello_world.proto", fileDescriptor_c3336ab888c6aac6) }

var fileDescriptor_c3336ab888c6aac6 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x2f, 0x33, 0xd4, 0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0x8f, 0x2f, 0xcf,
	0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0x2d, 0x2a, 0xca, 0x2f,
	0x2a, 0x96, 0x12, 0x86, 0xd0, 0xfa, 0x10, 0x0a, 0x22, 0xa9, 0x55, 0xc9, 0xc5, 0xed, 0x0a, 0xe2,
	0x07, 0xa5, 0x26, 0x16, 0xe7, 0xe7, 0x09, 0xe9, 0x70, 0xf1, 0x85, 0x06, 0xbb, 0x06, 0xc5, 0xfb,
	0xf9, 0x87, 0xc4, 0xbb, 0xf9, 0x87, 0xfa, 0xb9, 0x08, 0x30, 0x48, 0x49, 0x6c, 0x70, 0x9d, 0xc2,
	0xbc, 0xc3, 0xf5, 0x3c, 0xe7, 0x21, 0x57, 0x56, 0x13, 0x03, 0x93, 0xd4, 0xbc, 0x53, 0x60, 0xba,
	0x2a, 0x43, 0xc8, 0x98, 0x4b, 0x00, 0x55, 0x75, 0x7c, 0x89, 0x00, 0xa3, 0x94, 0xec, 0x06, 0xd7,
	0xc9, 0xcc, 0x3b, 0x5c, 0xbf, 0x1c, 0x64, 0x3a, 0xe4, 0xca, 0x6e, 0x62, 0x60, 0x62, 0x60, 0x08,
	0xd2, 0x02, 0x61, 0x55, 0x65, 0x48, 0xb1, 0x2f, 0x70, 0xfd, 0xc2, 0xbc, 0xc2, 0x95, 0xc1, 0xc9,
	0x22, 0xca, 0x2c, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0xdf, 0xd0, 0xc0,
	0x28, 0xb5, 0x48, 0x1f, 0xec, 0xa8, 0x64, 0xdd, 0xf4, 0xd4, 0x3c, 0xdd, 0xf4, 0xfc, 0xf4, 0x7c,
	0x5d, 0x98, 0x93, 0x11, 0x1e, 0xb4, 0x2e, 0x33, 0x4c, 0x62, 0x03, 0x2b, 0x33, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x27, 0x36, 0x14, 0x9e, 0xf9, 0x00, 0x00, 0x00,
}
