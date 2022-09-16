package service

import (
	"github.com/trist725/myleaf/log"
	"google.golang.org/grpc"
)

type BaseClient struct {
	Addr string
	Conn *grpc.ClientConn
}

func NewBaseClient(addr string) *BaseClient {
	return &BaseClient{Addr: addr}
}

func (bc *BaseClient) SetAddr(addr string) {
	bc.Addr = addr
}

func (bc *BaseClient) Dial(opts ...grpc.DialOption) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(bc.Addr, opts...)
	if err != nil {
		log.Error("did not connect: %v", err)
		return
	}
	bc.Conn = conn
}

func (bc *BaseClient) Close() {
	if bc.Conn != nil {
		bc.Conn.Close()
	}
}
