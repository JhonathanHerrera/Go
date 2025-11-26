package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "gRPC/proto"

	"google.golang.org/grpc"
)

// server implements the GreeterServer interface
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements the SayHello RPC method
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received request from: %s", req.Name)
	
	return &pb.HelloReply{
		Message: fmt.Sprintf("Hello, %s! Welcome to gRPC.", req.Name),
	}, nil
}

// SayHelloMany implements server streaming
func (s *server) SayHelloMany(req *pb.HelloRequest, stream pb.Greeter_SayHelloManyServer) error {
	log.Printf("Streaming greetings to: %s", req.Name)
	
	greetings := []string{
		"Hello",
		"Hi there",
		"Greetings",
		"Welcome",
		"Good to see you",
	}
	
	for i, greeting := range greetings {
		msg := fmt.Sprintf("%s, %s! (message %d/5)", greeting, req.Name, i+1)
		
		if err := stream.Send(&pb.HelloReply{Message: msg}); err != nil {
			return err
		}
		
		time.Sleep(500 * time.Millisecond) // Simulate work
	}
	
	return nil
}

func main() {
	// Listen on TCP port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	
	// Register our service with the server
	pb.RegisterGreeterServer(grpcServer, &server{})
	
	log.Println("gRPC server listening on :50051")
	
	// Start serving
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}