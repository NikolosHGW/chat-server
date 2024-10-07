package main

import (
	"context"
	"fmt"
	"log"
	"net"

	chatpb "github.com/NikolosHGW/chat-server/pkg/chat/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type userServer struct {
	chatpb.ChatV1Server
}

const grpcPort = 3200

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	chatpb.RegisterChatV1Server(s, &userServer{})

	fmt.Println("Сервер gRPC начал работу")
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}

func (s *userServer) Create(ctx context.Context, req *chatpb.CreateRequest) (*chatpb.CreateResponse, error) {
	_, cancel := context.WithCancel(ctx)
	defer cancel()
	fmt.Println(req.Usernames)

	return &chatpb.CreateResponse{
		Id: 1,
	}, nil
}

func (s *userServer) Delete(ctx context.Context, req *chatpb.DeleteRequest) (*emptypb.Empty, error) {
	_, cancel := context.WithCancel(ctx)
	defer cancel()
	fmt.Println(req.Id)

	return &emptypb.Empty{}, nil
}

func (s *userServer) SendMessage(ctx context.Context, req *chatpb.SendMessageRequest) (*emptypb.Empty, error) {
	_, cancel := context.WithCancel(ctx)
	defer cancel()
	fmt.Println(req.From, req.Text, req.Timestamp)

	return &emptypb.Empty{}, nil
}
