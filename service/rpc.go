package service

import (
	context "context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

// GreeterServiceImpl is used to implement GreeterServer.
type GreeterServiceImpl struct {
	UnimplementedGreeterServer
}

func (s *GreeterServiceImpl) Init() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", Conf.GRPCPort))
	if err != nil {
		panic(err)
	}
	svc := grpc.NewServer()
	RegisterGreeterServer(svc, &GreeterServiceImpl{})
	log.Printf("GreeterServer listening at %v", lis.Addr())
	if err := svc.Serve(lis); err != nil {
		panic(err)
	}
}

// SayHello implements GreeterServer
func (s *GreeterServiceImpl) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *GreeterServiceImpl) SayHelloAgain(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &HelloReply{Message: "Hello Again" + in.GetName()}, nil
}
