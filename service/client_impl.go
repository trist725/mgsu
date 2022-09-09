package service

import (
	context "context"
	"log"
	"time"

	"google.golang.org/grpc"
)

type GreeterClientImpl struct {
	*BaseClient

	Cli     GreeterClient
	Timeout time.Duration
}

func (gc *GreeterClientImpl) Dial(opts ...grpc.DialOption) {
	gc.BaseClient.Dial(opts...)
	gc.Cli = NewGreeterClient(gc.BaseClient.Conn)
}

func (gc *GreeterClientImpl) Close() {
	gc.BaseClient.Close()
	gc.Cli = nil
}

func (gc *GreeterClientImpl) Do() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := gc.Cli.SayHello(ctx, &HelloRequest{Name: "Tristone"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return
	}
	log.Printf("Greeting: %s", r.GetMessage())
	r, err = gc.Cli.SayHelloAgain(ctx, &HelloRequest{Name: "Tristone"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
