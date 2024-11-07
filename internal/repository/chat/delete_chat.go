package chat

import (
	"context"
	"fmt"

	"github.com/NikolosHGW/platform-common/pkg/db"
)

func (r *repo) DeleteChat(ctx context.Context, chatID int64) error {
	query := db.Query{
		Name:     repositoryName + ".add_user_to_chat",
		QueryRaw: `DELETE FROM chats WHERE id = $1`,
	}

	_, err := r.db.DB().ExecContext(ctx, query, chatID)
	if err != nil {
		return fmt.Errorf("layer repository: ошибка при удалении чата: %w", err)
	}

	return nil
}
