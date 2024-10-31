package chat

import (
	"context"
	"fmt"

	chatpb "github.com/NikolosHGW/chat-server/pkg/chat/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *implementation) Delete(ctx context.Context, req *chatpb.DeleteRequest) (*emptypb.Empty, error) {
	err := i.chatService.Delete(ctx, req.Id)
	if err != nil {
		return &emptypb.Empty{}, fmt.Errorf("layer api: не получилось удалить чат: %w", err)
	}

	return &emptypb.Empty{}, nil
}
