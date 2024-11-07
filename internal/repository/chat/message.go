package chat

import (
	"context"
	"fmt"

	"github.com/NikolosHGW/chat-server/internal/infrastructure/db/dto"
	"github.com/NikolosHGW/platform-common/pkg/db"
)

func (r *repo) CreateMessage(ctx context.Context, messageDTO dto.MessageDTO) error {
	query := db.Query{
		Name: repositoryName + ".add_user_to_chat",
		QueryRaw: `
			INSERT INTO messages
				(chat_id, from_user_id, text, timestamp)
			VALUES
				(:chat_id, :from_user_id, :text, :timestamp)
		`,
	}

	_, err := r.db.DB().NamedExecContext(ctx, query, &messageDTO)
	if err != nil {
		return fmt.Errorf("layer repository: ошибка при создании чата: %w", err)
	}

	return nil
}
