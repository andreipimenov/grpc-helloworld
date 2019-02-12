package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net"

	"github.com/andreipimenov/grpc-helloworld/pb"

	"google.golang.org/grpc"
)

type helloService struct{}

func (h helloService) Hello(context.Context, *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{}, errors.New("not implemented")
}

func main() {
	addr := flag.String("addr", ":8080", "address")
	flag.Parse()

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	srv := helloService{}
	pb.RegisterHelloServiceServer(s, srv)
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
