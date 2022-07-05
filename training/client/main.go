package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"webinar/training/proto/generated"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := my_chat.NewChatServiceClient(conn)

	response, err := c.Hi(context.Background(), &my_chat.Message{Body: "Hello From Client!"})
	if err != nil {
		panic(err)
	}
	log.Printf("Response from server: %s", response.Body)
}
