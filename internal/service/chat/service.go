package chat

import (
	"context"

	"github.com/NikolosHGW/chat-server/internal/infrastructure/db/dto"
	"github.com/NikolosHGW/platform-common/pkg/db"
)

type repoChat interface {
	CreateChat(context.Context) (int64, error)
	AddUsersToChat(ctx context.Context, chatID int64, userIDs []int64) error
	DeleteChat(ctx context.Context, chatID int64) error
	CreateMessage(context.Context, dto.MessageDTO) error
}

type service struct {
	chatRepo  repoChat
	txManager db.TxManager
}

// NewService - конструктор сервиса чата.
func NewService(chatRepo repoChat, txManager db.TxManager) *service {
	return &service{chatRepo: chatRepo, txManager: txManager}
}

// PublicService - алиас для структуры чат-сервиса.
// Сделано в экспериментальных целях, чтоб пробросить
// сервис провайдер чёткую структуру без менющихся интерфейсов.
type PublicService = service
