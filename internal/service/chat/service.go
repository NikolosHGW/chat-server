package chat

import (
	"context"

	"github.com/NikolosHGW/chat-server/internal/infrastructure/db/dto"
)

type repoChat interface {
	CreateChat(context.Context) (int64, error)
	AddUsersToChat(ctx context.Context, chatID int64, userIDs []int64) error
	DeleteChat(ctx context.Context, chatID int64) error
	CreateMessage(context.Context, dto.MessageDTO) error
}

type service struct {
	chatRepo repoChat
}

// NewService - конструктор сервиса чата.
func NewService(chatRepo repoChat) *service {
	return &service{chatRepo: chatRepo}
}
