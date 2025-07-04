package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    "github.com/JhonathanHerrera/grpc/chat"
)

func main() {
    var conn *grpc.ClientConn
    conn, err := grpc.Dial(":9000", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("could not connect: %v", err)
    }
    defer conn.Close()

    c := chat.NewChatServiceClient(conn)

    message := chat.Message{
        Body: "Hello from the client!",
    }

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    response, err := c.SayHello(ctx, &message)
    if err != nil {
        log.Fatalf("Error when calling SayHello: %v", err)
    }

    log.Printf("Response from server: %s", response.Body)
}
