package main

import (
	"context"
	"fmt"
	pb "go-playground/diver/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	totalNum int32
}

func (s *Server) AddNum(ctx context.Context, addingNum *pb.AddNumParams) (*pb.TotalNum, error) {
	log.Printf("Add Number")
	s.totalNum += addingNum.Number
	total := &pb.TotalNum{Total: s.totalNum}
	return total, nil
}

func (s *Server) GetTotalNum(ctx context.Context, _ *pb.GetTotalNumParams) (*pb.TotalNum, error) {
	log.Printf("Return total number")
	total := &pb.TotalNum{Total: s.totalNum}
	return total, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9999))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := Server{}
	grpcServer := grpc.NewServer()

	pb.RegisterAddNumServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to serve: %s", err)
	} else {
		log.Println("Server started successfully")
	}
}
