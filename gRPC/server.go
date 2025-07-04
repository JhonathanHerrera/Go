package main

import (
    "log"
    "net"

    "google.golang.org/grpc"
    "github.com/JhonathanHerrera/grpc/chat"
)

func main() {
    lis, err := net.Listen("tcp", ":9000")
    if err != nil {
        log.Fatalf("failed to listen on port 9000: %v", err)
    }

    s := chat.Server{}
    grpcServer := grpc.NewServer()
    chat.RegisterChatServiceServer(grpcServer, &s)

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
