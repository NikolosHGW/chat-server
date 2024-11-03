package chat

import (
	"context"
	"fmt"

	"github.com/NikolosHGW/chat-server/internal/infrastructure/db/dto"
)

func (r *repo) CreateMessage(ctx context.Context, messageDTO dto.MessageDTO) error {
	query := `
		INSERT INTO messages
			(chat_id, from_user_id, text, timestamp)
		VALUES
			(:chat_id, :from_user_id, :text, :timestamp)
		`
	_, err := r.db.NamedExecContext(ctx, query, &messageDTO)
	if err != nil {
		return fmt.Errorf("layer repository: ошибка при создании чата: %w", err)
	}

	return nil
}
