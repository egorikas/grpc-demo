package my_chat

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"time"
)

type MyBeatifulEndpoint struct {
	UnimplementedChatServiceServer
}

func (MyBeatifulEndpoint) Hi(context.Context, *ChatMessage) (*ChatMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hi not implemented")
}
func (MyBeatifulEndpoint) HiStream(_ *ChatMessage, server ChatService_HiStreamServer) error {
	for i := 1; i < 10; i++ {
		err := server.Send(&ChatMessage{
			Text: fmt.Sprint(i),
		})
		time.Sleep(1 * time.Second)
		if err != io.EOF && err != nil {
			panic(err)
		}
	}
	return status.Errorf(codes.Unimplemented, "method HiStream not implemented")
}
