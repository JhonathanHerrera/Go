package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "gRPC/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Connect to the server
	conn, err := grpc.Dial("localhost:50051", 
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := pb.NewGreeterClient(conn)

	// Test 1: Simple unary call
	log.Println("=== Test 1: Simple greeting ===")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Alice"})
	if err != nil {
		log.Fatalf("SayHello failed: %v", err)
	}
	log.Printf("Response: %s\n", response.Message)

	// Test 2: Server streaming
	log.Println("\n=== Test 2: Streaming greetings ===")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel2()

	stream, err := client.SayHelloMany(ctx2, &pb.HelloRequest{Name: "Bob"})
	if err != nil {
		log.Fatalf("SayHelloMany failed: %v", err)
	}

	// Receive streamed messages
	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			// Stream finished
			log.Println("Stream ended")
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive: %v", err)
		}
		log.Printf("Received: %s", reply.Message)
	}

	log.Println("\n=== All tests completed ===")
}