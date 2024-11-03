package chat

import (
	"context"
	"fmt"
)

func (s *service) Delete(ctx context.Context, chatID int64) error {
	err := s.chatRepo.DeleteChat(ctx, chatID)
	if err != nil {
		return fmt.Errorf("layer service: ошибка при удалении чата: %w", err)
	}

	return nil
}
