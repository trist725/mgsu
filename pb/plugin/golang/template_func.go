package golang

import (
	"strings"
	"text/template"

	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

var TemplateFunc = template.FuncMap{
	"CamelCase":           generator.CamelCase,
	"CamelCaseSlice":      generator.CamelCaseSlice,
	"DescriptorProtoType": DescriptorProtoType,
	"SplitPackageName":    SplitPackageName,
	"SplitType":           SplitType,
}

var (
	fieldDescriptorProtoTypes = map[descriptor.FieldDescriptorProto_Type]string{
		descriptor.FieldDescriptorProto_TYPE_DOUBLE:   "double",
		descriptor.FieldDescriptorProto_TYPE_FLOAT:    "float",
		descriptor.FieldDescriptorProto_TYPE_INT64:    "int64",
		descriptor.FieldDescriptorProto_TYPE_UINT64:   "uint64",
		descriptor.FieldDescriptorProto_TYPE_INT32:    "int32",
		descriptor.FieldDescriptorProto_TYPE_FIXED64:  "fixed64",
		descriptor.FieldDescriptorProto_TYPE_FIXED32:  "fixed32",
		descriptor.FieldDescriptorProto_TYPE_BOOL:     "bool",
		descriptor.FieldDescriptorProto_TYPE_STRING:   "string",
		descriptor.FieldDescriptorProto_TYPE_GROUP:    "group",
		descriptor.FieldDescriptorProto_TYPE_MESSAGE:  "message",
		descriptor.FieldDescriptorProto_TYPE_BYTES:    "bytes",
		descriptor.FieldDescriptorProto_TYPE_UINT32:   "uint32",
		descriptor.FieldDescriptorProto_TYPE_ENUM:     "enum",
		descriptor.FieldDescriptorProto_TYPE_SFIXED32: "sfixed32",
		descriptor.FieldDescriptorProto_TYPE_SFIXED64: "sfixed64",
		descriptor.FieldDescriptorProto_TYPE_SINT32:   "sint32",
		descriptor.FieldDescriptorProto_TYPE_SINT64:   "sint64",
	}
)

func DescriptorProtoType(t descriptor.FieldDescriptorProto_Type) string {
	if s, ok := fieldDescriptorProtoTypes[t]; ok {
		return s
	}
	return ""
}

func SplitPackageName(t string) (packageName string) {
	words := strings.Split(t, ".")
	if len(words) == 2 {
		packageName = words[0]
	}
	return
}

func SplitType(t string) (tYpE string) {
	words := strings.Split(t, ".")
	if len(words) == 2 {
		tYpE = words[1]
	}
	return
}

func IsGoBaseType(t string) bool {
	switch t {
	case "int", "uint", "int32", "uint32", "int64", "uint64", "int8", "uint8", "int16", "uint16",
		"string", "bool",
		"double", "float32", "float64":
		return true
	}
	return false
}
