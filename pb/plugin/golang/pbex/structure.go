package pbex

import (
	"github.com/trist725/mgsu/pb/plugin/golang"
	"github.com/trist725/mgsu/pb/plugin/golang/generator"
)

const (
	PluginName = "pbex-go"
)

type message struct {
	*golang.Message
	Response string
	Model    string
}

func newMessage(g *generator.Generator, d *generator.Descriptor) *message {
	m := &message{
		Message: golang.NewMessage(g, d),
	}
	return m
}

type file struct {
	*golang.File
	Protocol       string
	Messages       []*message
	messagesByName map[string]*message
}

func newFile(fd *generator.FileDescriptor) *file {
	f := &file{
		File:           golang.NewFile(fd),
		messagesByName: map[string]*message{},
	}
	return f
}

func (f file) ProtocolTagFunctionName() string {
	return generator.CamelCase(f.Protocol)
}
