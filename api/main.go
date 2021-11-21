package main

import (
	"context"
	"log"
	"net"

	pb "github.com/leslesnoa/grpc-gateway/pb"
	"google.golang.org/grpc"
)

type EchoService struct {
}

func (s *EchoService) Echo(ctx context.Context, message *pb.EchoRequest) (*pb.EchoResponse, error) {
	// log.Println(message)
	log.Printf("Received: %v", message.Message)
	// time.Sleep(3 * time.Second)
	return &pb.EchoResponse{Message: "Hello " + message.Message}, nil
}

func main() {
	addr := ":9090"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEchoServiceServer(s, &EchoService{})

	log.Printf("gRPC server listening on " + addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
