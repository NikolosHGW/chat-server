package chat

import (
	"context"
	"fmt"
)

func (s *service) Create(ctx context.Context, userIDs []int64) (int64, error) {
	chatID, err := s.chatRepo.CreateChat(ctx)
	if err != nil {
		return 0, fmt.Errorf("layer service: ошибка при создании нового чата: %w", err)
	}

	err = s.chatRepo.AddUsersToChat(ctx, chatID, userIDs)
	if err != nil {
		return 0, fmt.Errorf("layer service: ошибка при добавлении пользователей в новый чат: %w", err)
	}

	return chatID, nil
}
