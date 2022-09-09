package service

import (
	"google.golang.org/grpc"
)

type IRPCServerImpl interface {
	Serve()
	GetAddr() string
}

type IRPCClientImpl interface {
	Close()
	Dial(...grpc.DialOption)
	SetAddr(string)
}
