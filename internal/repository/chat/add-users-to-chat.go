package chat

import (
	"context"
	"fmt"

	"github.com/NikolosHGW/platform-common/pkg/db"
)

const repositoryName = "chat_repository"

type chatUser struct {
	ChatID int64 `db:"chat_id"`
	UserID int64 `db:"user_id"`
}

func (r *repo) AddUsersToChat(ctx context.Context, chatID int64, userIDs []int64) error {
	chatUsers := make([]chatUser, len(userIDs))
	for i, userID := range userIDs {
		chatUsers[i] = chatUser{ChatID: chatID, UserID: userID}
	}

	query := db.Query{
		Name: repositoryName + ".add_user_to_chat",
		QueryRaw: `
			INSERT INTO chat_users
				(chat_id, user_id)
			VALUES
				(:chat_id, :user_id)
		`,
	}

	_, err := r.db.DB().NamedExecContext(ctx, query, chatUsers)
	if err != nil {
		return fmt.Errorf("layer repository: ошибка при добавлении пользователей в чат: %w", err)
	}

	return nil
}
