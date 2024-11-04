package chat

import (
	"context"
	"fmt"
)

type chatUser struct {
	ChatID int64 `db:"chat_id"`
	UserID int64 `db:"user_id"`
}

func (r *repo) AddUsersToChat(ctx context.Context, chatID int64, userIDs []int64) error {
	chatUsers := make([]chatUser, len(userIDs))
	for i, userID := range userIDs {
		chatUsers[i] = chatUser{ChatID: chatID, UserID: userID}
	}

	query := `
		INSERT INTO chat_users
			(chat_id, user_id)
        VALUES
			(:chat_id, :user_id)
	`

	_, err := r.db.NamedExecContext(ctx, query, chatUsers)
	if err != nil {
		return fmt.Errorf("layer repository: ошибка при добавлении пользователей в чат: %w", err)
	}

	return nil
}
