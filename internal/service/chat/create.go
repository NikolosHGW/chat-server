package chat

import (
	"context"
	"fmt"
)

func (s *service) Create(ctx context.Context, userIDs []int64) (int64, error) {
	// chatID, err := s.chatRepo.CreateChat(ctx)
	// if err != nil {
	// 	return 0, fmt.Errorf("layer service: ошибка при создании нового чата: %w", err)
	// }

	// err = s.chatRepo.AddUsersToChat(ctx, chatID, userIDs)
	// if err != nil {
	// 	return 0, fmt.Errorf("layer service: ошибка при добавлении пользователей в новый чат: %w", err)
	// }

	var chatID int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		chatID, errTx = s.chatRepo.CreateChat(ctx)
		if errTx != nil {
			return fmt.Errorf("layer service: ошибка при создании нового чата: %w", errTx)
		}

		errTx = s.chatRepo.AddUsersToChat(ctx, chatID, userIDs)
		if errTx != nil {
			return fmt.Errorf("layer service: ошибка при добавлении пользователей в новый чат: %w", errTx)
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return chatID, nil
}
