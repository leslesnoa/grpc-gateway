package main

import (
	"context"
	"log"
	"os"

	pb "github.com/leslesnoa/grpc-gateway/pb"
	"google.golang.org/grpc"
)

func main() {
	addr := "localhost:5000"
	// --------TLS認証を追加--------
	// creds, err := credentials.NewClientTLSFromFile("server.crt", "")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(creds))
	// ---------------------------

	// --------Interceptor--------
	// conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(creds), grpc.WithUnaryInterceptor(unaryInterceptor))
	// conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithUnaryInterceptor(unaryInterceptor))
	// ---------------------------
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoServiceClient(conn)

	name := os.Args[1]

	ctx := context.Background()
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	cancel()
	// }()
	r, err := c.Echo(ctx, &pb.EchoRequest{Message: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}

// --------Interceptor--------
// func unaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
// 	log.Printf("before call: %s, request: %+v", method, req)
// 	err := invoker(ctx, method, req, reply, cc, opts...)
// 	log.Printf("after call: %s, response: %+v", method, reply)
// 	return err
// }

// --------Interceptor--------