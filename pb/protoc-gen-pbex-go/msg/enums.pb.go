// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: enums.proto

package msg

import (
	fmt "fmt"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

/// 协议版本枚举
type Version int32

const (
	Version_Version_ Version = 0
	/// 协议版本
	Version_Num Version = 1
)

var Version_name = map[int32]string{
	0: "Version_",
	1: "Num",
}

var Version_value = map[string]int32{
	"Version_": 0,
	"Num":      1,
}

func (x Version) String() string {
	return proto.EnumName(Version_name, int32(x))
}

func (Version) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_888b6bd9597961ff, []int{0}
}

func init() {
	proto.RegisterEnum("msg.Version", Version_name, Version_value)
}

func init() { proto.RegisterFile("enums.proto", fileDescriptor_888b6bd9597961ff) }

var fileDescriptor_888b6bd9597961ff = []byte{
	// 103 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xcd, 0x2b, 0xcd,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2d, 0x4e, 0xd7, 0x52, 0xe0, 0x62,
	0x0f, 0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0x13, 0xe2, 0xe1, 0xe2, 0x80, 0x32, 0xe3, 0x05, 0x18,
	0x84, 0xd8, 0xb9, 0x98, 0xfd, 0x4a, 0x73, 0x05, 0x18, 0x9d, 0x24, 0x4e, 0x3c, 0x92, 0x63, 0xbc,
	0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63,
	0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x89, 0x0d, 0x6c, 0x8e, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xe6,
	0x42, 0x22, 0x9e, 0x56, 0x00, 0x00, 0x00,
}