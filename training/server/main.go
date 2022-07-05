package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"webinar/training/proto/generated"
)

func main() {
	fmt.Println("let's go!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := my_chat.UnimplementedChatServiceServer{}

	grpcServer := grpc.NewServer()

	my_chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
