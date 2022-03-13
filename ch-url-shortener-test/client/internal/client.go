package main

import (
	"context"
	pb "github.com/channel-io/grpc-study/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewRegisterAPIClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Ping(ctx, &pb.Empty{})

	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}
	log.Printf("Client: ping -> Server: %s", r.Value)
}
