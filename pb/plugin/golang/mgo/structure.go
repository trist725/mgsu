package mgo

import (
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/trist725/mgsu/pb/plugin/golang"
	"github.com/trist725/mgsu/pb/plugin/golang/generator"
)

type field struct {
	*golang.Field

	Msg                      string
	MsgRepeatedMessageGoType string
}

func newField(g *generator.Generator, d *generator.Descriptor, fdp *descriptor.FieldDescriptorProto, commentIndex int) *field {
	f := &field{
		Field: golang.NewField(g, d, fdp, commentIndex),
	}
	return f
}

type message struct {
	*golang.Message

	Fields []*field

	Msg string
	Rpc string

	Slice bool
}

func newMessage(g *generator.Generator, d *generator.Descriptor) *message {
	m := &message{
		Message: golang.NewMessage(g, d),
	}
	return m
}

type file struct {
	*golang.File
	Messages []*message
}

func newFile(fd *generator.FileDescriptor) *file {
	f := &file{
		File: golang.NewFile(fd),
	}
	return f
}
