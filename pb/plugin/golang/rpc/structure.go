package rpc

import (
	"github.com/trist725/mgsu/pb/plugin/golang"
	"github.com/trist725/mgsu/pb/plugin/golang/generator"
)

const (
	DefaultResponseTimeoutMillisecond = "10000"
)

type message struct {
	*golang.Message
	Response        string
	ResponseTimeout string
	CanIgnore       bool
}

func newMessage(g *generator.Generator, d *generator.Descriptor) *message {
	m := &message{
		Message:         golang.NewMessage(g, d),
		ResponseTimeout: DefaultResponseTimeoutMillisecond,
	}
	return m
}

type file struct {
	*golang.File
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
