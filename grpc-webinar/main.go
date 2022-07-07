package main

import (
	"google.golang.org/grpc"
	"grpc-webinar/generated"
	"net"
)

func main() {
	println("Let's go")

	l, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	uss := &my_chat.MyBeatifulEndpoint{}
	my_chat.RegisterChatServiceServer(grpcServer, uss)

	err = grpcServer.Serve(l)
	if err != nil {
		panic(err)
	}
}
