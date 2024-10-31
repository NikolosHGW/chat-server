package chat

import (
	"context"
	"fmt"

	chatpb "github.com/NikolosHGW/chat-server/pkg/chat/v1"
)

func (i *implementation) Create(ctx context.Context, req *chatpb.CreateRequest) (*chatpb.CreateResponse, error) {
	chatID, err := i.chatService.Create(ctx, req.UserIds)
	if err != nil {
		return nil, fmt.Errorf("layer api: не получилось создать чат: %w", err)
	}

	return &chatpb.CreateResponse{
		Id: chatID,
	}, nil
}
