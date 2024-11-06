package chat

import (
	"context"

	"github.com/NikolosHGW/chat-server/internal/domain"
	chatpb "github.com/NikolosHGW/chat-server/pkg/chat/v1"
)

type implementation struct {
	chatpb.ChatV1Server

	chatService Service
}

// Service интерфейс для инъекции сервиса чата
type Service interface {
	Create(context.Context, []int64) (int64, error)
	Delete(context.Context, int64) error
	SendMessage(context.Context, domain.Message) error
}

// NewImplementation - конструктор gRPC сервера.
func NewImplementation(chatService Service) *implementation {
	return &implementation{chatService: chatService}
}

// PublicServerImplementation - алиас для структуры gRPC имплементации чат-сервера.
// Сделано в экспериментальных целях, чтоб пробросить
// сервис провайдер чёткую структуру без менющихся интерфейсов.
type PublicServerImplementation = implementation
