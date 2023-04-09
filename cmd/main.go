package main

import (
	"context"
	"log"
	"net"

	pb "go-pkg/protobuf"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *protobuf.HelloRequest) (*pb.HelloResponse, error) {
	pb

	protobuf.HelloRequest
	api

	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("Server started on port %v", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
