package main

import (
	"context"
	"fmt"
	chat_api "github.com/kimcodec/microservices/chat_server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

type server struct {
	chat_api.UnimplementedChatApiV1Server
}

func (s *server) Create(ctx context.Context, req *chat_api.CreateRequest) (*chat_api.CreateResponse, error) {
	log.Printf("Usernames: %v", req.GetUsernames())
	return nil, nil
}
func (s *server) Delete(ctx context.Context, req *chat_api.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Id: %d", req.GetId())
	return nil, nil
}

func (s *server) SendMessage(ctx context.Context, req *chat_api.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Text: %s, From: %s, Time: %v", req.GetText(), req.GetFrom(), req.GetTimestamp().AsTime())
	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}

	s := grpc.NewServer()
	reflection.Register(s)
	chat_api.RegisterChatApiV1Server(s, &server{})

	log.Printf("server listening at %d port", 50051)

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to server: ", err.Error())
	}
}
