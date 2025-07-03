package main

import (
	grpcdef "app/go/bp"
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:8080", "the address to connect to")
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := grpcdef.NewGreetingServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := c.Hello(ctx, &grpcdef.HelloRequest{
		Name: "AAA",
	})
	if err != nil {
		log.Fatalf("Error when calling SyaHello: %s", err)
	}
	log.Printf("Responce from server: %s", resp.Message)
}
