package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/andreipimenov/grpc-helloworld/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type helloService struct{}

func (h helloService) Hello(context.Context, *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{}, status.Error(codes.Unimplemented, "Not implemented")
}

func main() {
	port := flag.Int("p", 8080, "port")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	srv := helloService{}
	pb.RegisterHelloServiceServer(s, srv)
	log.Printf("Listen on %v\n", fmt.Sprintf(":%d", *port))
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
