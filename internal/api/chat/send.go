package chat

import (
	"context"
	"fmt"

	chatpb "github.com/NikolosHGW/chat-server/pkg/chat/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *implementation) SendMessage(ctx context.Context, req *chatpb.SendMessageRequest) (*emptypb.Empty, error) {
	err := i.chatService.SendMessage(ctx, convertToMessageDomain(req))
	if err != nil {
		return &emptypb.Empty{}, fmt.Errorf("layer api: не удалось отправить сообщение: %w", err)
	}

	return &emptypb.Empty{}, nil
}
