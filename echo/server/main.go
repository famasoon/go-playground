package main

import (
	"context"
	"log"
	"net"

	pb "go-playground/echo/proto"

	"google.golang.org/grpc"
)

type echoService struct{}

func (s *echoService) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{
		Message: req.GetMessage(),
	}, nil
}

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	srv := grpc.NewServer()
	pb.RegisterEchoServiceServer(srv, &echoService{})

	log.Printf("Start server on port%s\n", port)
	if err := srv.Serve(lis); err != nil {
		log.Printf("Failed to serve: %v\n", err)
	}

}
