package chat

import (
	"context"
	"fmt"

	"github.com/NikolosHGW/chat-server/internal/domain"
)

func (s *service) SendMessage(ctx context.Context, message domain.Message) error {
	err := s.chatRepo.CreateMessage(ctx, convertToMessageDTO(message))
	if err != nil {
		return fmt.Errorf("layer service: ошибка при создании сообщения в бд: %w", err)
	}

	return nil
}
