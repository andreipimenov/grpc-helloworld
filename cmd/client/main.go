package main

import (
	"context"
	"flag"
	"log"

	"github.com/andreipimenov/grpc-helloworld/pb"
	"google.golang.org/grpc"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:8080", "address of the server")
	flag.Parse()

	name := flag.Arg(0)
	if name == "" {
		log.Fatal("Name should be specified as first cmd argument")
	}

	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial %s: %v", *addr, err)
	}
	if conn != nil {
		defer conn.Close()
	}

	c := pb.NewHelloServiceClient(conn)
	res, err := c.Hello(context.Background(), &pb.HelloRequest{
		Name: name,
	})
	if err != nil {
		log.Fatalf("Failed to call RPC method Hello: %v", err)
	}
	log.Printf("Response: %v", res.Message)

}
