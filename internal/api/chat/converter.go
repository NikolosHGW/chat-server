package chat

import (
	"github.com/NikolosHGW/chat-server/internal/domain"
	chatpb "github.com/NikolosHGW/chat-server/pkg/chat/v1"
)

func convertToMessageDomain(req *chatpb.SendMessageRequest) domain.Message {
	return domain.Message{
		ChatID:     req.ChatId,
		FromUserID: req.FromUserId,
		Text:       req.Text,
		Timestamp:  req.Timestamp.AsTime(),
	}
}
