package chat

import (
	"context"
	"fmt"

	"github.com/NikolosHGW/platform-common/pkg/db"
)

func (r *repo) CreateChat(ctx context.Context) (int64, error) {
	var id int64

	query := db.Query{
		Name:     repositoryName + ".add_user_to_chat",
		QueryRaw: `INSERT INTO chats DEFAULT VALUES RETURNING id`,
	}

	err := r.db.DB().QueryRowxContext(ctx, query).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("layer repository: ошибка при создании чата: %w", err)
	}

	return id, nil
}
