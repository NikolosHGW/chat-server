package chat

import (
	"context"

	"github.com/NikolosHGW/chat-server/internal/domain"
	chatpb "github.com/NikolosHGW/chat-server/pkg/chat/v1"
)

type implementation struct {
	chatpb.ChatV1Server

	chatService chatService
}

type chatService interface {
	Create(context.Context, []int64) (int64, error)
	Delete(context.Context, int64) error
	SendMessage(context.Context, *domain.Message) error
}

// NewImplementation - конструктор gRPC сервера.
func NewImplementation(chatService chatService) *implementation {
	return &implementation{chatService: chatService}
}
