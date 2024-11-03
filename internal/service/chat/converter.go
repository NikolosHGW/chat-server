package chat

import (
	"github.com/NikolosHGW/chat-server/internal/domain"
	"github.com/NikolosHGW/chat-server/internal/infrastructure/db/dto"
)

func convertToMessageDTO(message domain.Message) dto.MessageDTO {
	return dto.MessageDTO{
		ID:         message.ID,
		ChatID:     message.ChatID,
		FromUserID: message.FromUserID,
		Text:       message.Text,
		Timestamp:  message.Timestamp,
	}
}
