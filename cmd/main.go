package main

import (
	"fmt"
	"log"
	"net"

	apiChat "github.com/NikolosHGW/chat-server/internal/api/chat"
	"github.com/NikolosHGW/chat-server/internal/infrastructure/config"
	repoChat "github.com/NikolosHGW/chat-server/internal/repository/chat"
	serviceChat "github.com/NikolosHGW/chat-server/internal/service/chat"
	chatpb "github.com/NikolosHGW/chat-server/pkg/chat/v1"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 3200

func main() {
	cfg := config.NewConfig()

	db, err := sqlx.Connect("postgres", cfg.GetDatabaseDSN())
	if err != nil {
		log.Fatalln(err)
	}

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	reflection.Register(s)

	chatRepo := repoChat.NewRepo(db)
	chatService := serviceChat.NewService(chatRepo)
	chatAPI := apiChat.NewImplementation(chatService)

	chatpb.RegisterChatV1Server(s, chatAPI)

	fmt.Println("Сервер gRPC начал работу")
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
