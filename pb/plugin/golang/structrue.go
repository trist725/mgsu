package golang

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/trist725/mgsu/pb/plugin/golang/generator"

	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

const (
	EnterReplaceString = "\n//"
)

var (
	enumEnableZeroRegexp = regexp.MustCompile(`\s*@enable_0\s*`)
)

type Field struct {
	*descriptor.FieldDescriptorProto

	Name                string
	JsonName            string
	Comment             string
	DescriptorProtoType string

	IsMap   bool
	IsOneof bool

	GoType       string
	GoTypeToName string

	MapType         *generator.GoMapDescriptor
	KeyField        *descriptor.FieldDescriptorProto
	KeyType         string
	KeyTypeToName   string
	ValueField      *descriptor.FieldDescriptorProto
	ValueType       string
	ValueTypeToName string
}

func NewField(g *generator.Generator, d *generator.Descriptor, fdp *descriptor.FieldDescriptorProto, commentIndex int) *Field {
	field := &Field{
		FieldDescriptorProto: fdp,
		Name:                 g.GetFieldName(d, fdp),
		JsonName:             fdp.GetJsonName(),
		Comment:              g.Comments(fmt.Sprintf("%s,%d,%d", d.Path(), 2, commentIndex)),
		DescriptorProtoType:  fieldDescriptorProtoTypes[*fdp.Type],
		IsMap:                g.IsMap(fdp),
		IsOneof:              fdp.OneofIndex != nil,
	}

	if field.Comment != "" {
		field.Comment = strings.ReplaceAll(field.Comment, "\n", EnterReplaceString)
	}

	field.GoType, _ = g.GoType(d, fdp)
	field.GoTypeToName = generator.GoTypeToName(field.GoType)

	if field.IsMap {
		desc := g.ObjectNamed(fdp.GetTypeName())
		if d, ok := desc.(*generator.Descriptor); ok && d.GetOptions().GetMapEntry() {
			field.MapType = g.GoMapType(d, fdp)
			field.GoType = field.MapType.GoType
			field.GoTypeToName = generator.GoTypeToName(field.GoType)
			field.KeyField = field.MapType.KeyField
			field.KeyType, _ = g.GoType(d, field.KeyField)
			field.KeyTypeToName = generator.GoTypeToName(field.KeyType)
			field.ValueField = field.MapType.ValueField
			field.ValueType, _ = g.GoType(d, field.ValueField)
			field.ValueTypeToName = generator.GoTypeToName(field.ValueType)
		}
	}

	return field
}

type MessageMap struct {
	Name string
	Key  string
}

func NewMessageMap(name string, key string) *MessageMap {
	return &MessageMap{
		Name: name,
		Key:  key,
	}
}

type Message struct {
	*generator.Descriptor

	ID      string
	Name    string
	Type    string
	Comment string
	Fields  []*Field
	Maps    []*MessageMap
	Oneofs  []string
}

func NewMessage(g *generator.Generator, d *generator.Descriptor) *Message {
	message := &Message{
		Descriptor: d,
		Comment:    g.Comments(d.Path()),
	}

	message.Name = generator.CamelCaseSlice(d.TypeName())

	if message.Comment != "" {
		message.Comment = strings.ReplaceAll(message.Comment, "\n", EnterReplaceString)
	}

	return message
}

type EnumValue struct {
	*descriptor.EnumValueDescriptorProto

	Name     string
	TypeName string
	Value    int32
	Comment  string
}

func NewEnumValue(g *generator.Generator, enumName string, evdp *descriptor.EnumValueDescriptorProto, commentIndex1 int, commentIndex2 int) *EnumValue {
	enumValue := &EnumValue{
		EnumValueDescriptorProto: evdp,
		Name:                     evdp.GetName(),
		TypeName:                 enumName,
		Value:                    evdp.GetNumber(),
		Comment:                  g.Comments(fmt.Sprintf("5,%d,2,%d", commentIndex1, commentIndex2)),
	}

	if enumValue.Comment != "" {
		enumValue.Comment = strings.ReplaceAll(enumValue.Comment, "\n", EnterReplaceString)
	}

	return enumValue
}

type Enum struct {
	*generator.EnumDescriptor

	Name    string
	Comment string
	Values  []*EnumValue
	Enable0 bool
}

func NewEnum(g *generator.Generator, ed *generator.EnumDescriptor, commonIndex int) *Enum {
	enum := &Enum{
		EnumDescriptor: ed,
		Name:           generator.CamelCaseSlice(ed.TypeName()),
		Comment:        g.Comments(fmt.Sprintf("5,%d", commonIndex)),
	}

	if matches := enumEnableZeroRegexp.FindStringSubmatch(enum.Comment); len(matches) > 0 {
		enum.Enable0 = true
	}

	if enum.Comment != "" {
		enum.Comment = strings.ReplaceAll(enum.Comment, "\n", EnterReplaceString)
	}

	return enum
}

type File struct {
	*generator.FileDescriptor

	PackageName string
	Messages    []*Message
	Enums       []*Enum
}

func NewFile(fd *generator.FileDescriptor) *File {
	f := &File{
		FileDescriptor: fd,
		PackageName:    fd.GetPackage(),
	}
	return f
}
