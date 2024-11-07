package app

import (
	"context"
	"log"

	apiChat "github.com/NikolosHGW/chat-server/internal/api/chat"
	"github.com/NikolosHGW/chat-server/internal/infrastructure/config"
	repoChat "github.com/NikolosHGW/chat-server/internal/repository/chat"
	serviceChat "github.com/NikolosHGW/chat-server/internal/service/chat"
	"github.com/NikolosHGW/platform-common/pkg/closer"
	"github.com/NikolosHGW/platform-common/pkg/db"
	"github.com/NikolosHGW/platform-common/pkg/db/pg"
	"github.com/NikolosHGW/platform-common/pkg/db/transaction"
)

type serviceProvider struct {
	pgConfig   *config.PublicPG
	grpcConfig *config.PublicGRPC

	dbClient  db.Client
	txManager db.TxManager

	chatRepo *repoChat.PublicRepo

	chatService *serviceChat.PublicService

	chatServer *apiChat.PublicServerImplementation
}

// NewServiceProvider - конструктор сервис провайдера.
func NewServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (sp *serviceProvider) PGConfig() *config.PublicPG {
	if sp.pgConfig == nil {
		pgConfig, err := config.NewPG()
		if err != nil {
			log.Fatalf("не удалось инициализировать постгрес конфигурацию: %s", err.Error())
		}

		sp.pgConfig = pgConfig
	}

	return sp.pgConfig
}

func (sp *serviceProvider) GRPCConfig() *config.PublicGRPC {
	if sp.grpcConfig == nil {
		grpcConfig, err := config.NewGRPC()
		if err != nil {
			log.Fatalf("не удалось инициализировать grpc конфигурацию: %s", err.Error())
		}

		sp.grpcConfig = grpcConfig
	}

	return sp.grpcConfig
}

func (sp *serviceProvider) DBClient(ctx context.Context) db.Client {
	if sp.dbClient == nil {
		db, err := pg.New(ctx, sp.PGConfig().GetDatabaseDSN())
		if err != nil {
			log.Fatalf("не удалось настроить соединение с бд постгрес: %s", err.Error())
		}

		closer.Add(db.Close)

		sp.dbClient = db
	}

	return sp.dbClient
}

func (sp *serviceProvider) ChatRepo(ctx context.Context) *repoChat.PublicRepo {
	if sp.chatRepo == nil {
		sp.chatRepo = repoChat.NewRepo(sp.DBClient(ctx))
	}

	return sp.chatRepo
}

func (sp *serviceProvider) ChatService(ctx context.Context) *serviceChat.PublicService {
	if sp.chatService == nil {
		sp.chatService = serviceChat.NewService(sp.ChatRepo(ctx), sp.TxManager(ctx))
	}

	return sp.chatService
}

func (sp *serviceProvider) ChatServer(ctx context.Context) *apiChat.PublicServerImplementation {
	if sp.chatServer == nil {
		sp.chatServer = apiChat.NewImplementation(sp.ChatService(ctx))
	}

	return sp.chatServer
}

func (sp *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if sp.txManager == nil {
		sp.txManager = transaction.NewTransactionManager(sp.DBClient(ctx).DB())
	}

	return sp.txManager
}
