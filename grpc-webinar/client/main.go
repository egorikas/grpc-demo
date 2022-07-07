package main

import (
	"context"
	my_chat "grpc-webinar/generated"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := my_chat.NewChatServiceClient(conn)

	client, err := c.HiStream(context.Background(), &my_chat.ChatMessage{Text: "Это сообщение вернется!"})
	if err != nil {
		panic(err)
	}

	for {
		msg, err := client.Recv()
		if err == io.EOF {
			return
		}

		println(msg.Text)
	}
}
