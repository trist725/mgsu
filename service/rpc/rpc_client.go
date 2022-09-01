package rpc

import (
	context "context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GreeterClientImpl struct {
	Addr    string
	Timeout time.Duration
}

func (gc *GreeterClientImpl) Dial() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(gc.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), gc.Timeout)
	defer cancel()
	r, err := c.SayHelloAgain(ctx, &HelloRequest{Name: "testName"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
