package main

import (
	"context"
	"log"
	"net"

	pb "github.com/leslesnoa/grpc-gateway/pb"
	"google.golang.org/grpc"
)

// ---service----
type EchoService struct {
}

func (s *EchoService) Echo(ctx context.Context, message *pb.EchoRequest) (*pb.EchoResponse, error) {
	// log.Println(message)
	log.Printf("Received: %v", message.Message)
	// time.Sleep(3 * time.Second)
	return &pb.EchoResponse{Message: "Hello " + message.Message}, nil
}

// -------------

func main() {
	addr := ":9090"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// -------TLS認証処理を追加-------
	// cred, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// s := grpc.NewServer(grpc.Creds(cred))
	// -----------------------------
	s := grpc.NewServer()
	pb.RegisterEchoServiceServer(s, &EchoService{})

	log.Printf("gRPC server listening on " + addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// func (h HelloService) Run() int {
// 	s := grpc.NewServer()
// 	pb.RegisterHelloServiceServer(s, h)

// 	lis, err := net.Listen("tcp", ":5000")
// 	if err != nil {
// 		fmt.Printf("%+v\n", err)
// 		return 1
// 	}
// 	if err := s.Serve(lis); err != nil {
// 		fmt.Printf("%+v\n", err)
// 		return 1
// 	}
// 	return 0
// }

// func (h HelloService) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
// 	return &pb.HelloResponse{
// 		Message: "Hello, " + in.Name,
// 	}, nil
// }
