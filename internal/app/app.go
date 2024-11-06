package app

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/NikolosHGW/chat-server/internal/closer"
	"github.com/NikolosHGW/chat-server/internal/infrastructure/config"
	chatpb "github.com/NikolosHGW/chat-server/pkg/chat/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type app struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

func NewApp(ctx context.Context) (*app, error) {
	a := &app{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, fmt.Errorf("не удалось инициализировать зависимости приложения: %w", err)
	}

	return a, nil
}

func (a *app) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
		log.Println("Приложение завершено корректно.")
	}()

	return a.runGRPCServer()
}

func (a *app) initDeps(ctx context.Context) error {
	err := a.initConfigs()
	if err != nil {
		return err
	}

	a.initServiceProvider()
	a.initGRPCServer(ctx)

	return nil
}

func (a *app) initConfigs() error {
	err := config.Load()
	if err != nil {
		return fmt.Errorf("ошибка инициализации конфигов: %w", err)
	}

	return nil
}

func (a *app) initServiceProvider() {
	a.serviceProvider = NewServiceProvider()
}

func (a *app) initGRPCServer(ctx context.Context) {
	a.grpcServer = grpc.NewServer()
	reflection.Register(a.grpcServer)

	chatpb.RegisterChatV1Server(a.grpcServer, a.serviceProvider.ChatServer(ctx))
}

func (a *app) runGRPCServer() error {
	listen, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().GetRunAddress())
	if err != nil {
		return fmt.Errorf("не удалось прослушать TCP: %w", err)
	}

	closer.Add(
		func() error {
			log.Printf(
				"Отключаем сервер по адресу %s...",
				a.serviceProvider.GRPCConfig().GetRunAddress(),
			)
			a.grpcServer.GracefulStop()

			return nil
		},
	)

	log.Printf("Запуск gRPC сервера на адресе %s", a.serviceProvider.GRPCConfig().GetRunAddress())

	if err := a.grpcServer.Serve(listen); err != nil {
		return fmt.Errorf("ошибка при запуске сервера: %w", err)
	}

	return nil
}
