package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/NikolosHGW/chat-server/internal/api/chat"
	"github.com/NikolosHGW/chat-server/internal/api/chat/mocks"
	chatpb "github.com/NikolosHGW/chat-server/pkg/chat/v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestDelete(t *testing.T) {
	type chatServiceMockFunc func(mc *minimock.Controller) chat.Service

	type args struct {
		ctx context.Context
		req *chatpb.DeleteRequest
	}

	var (
		ctx        = context.Background()
		chatID     = gofakeit.Int64()
		req        = &chatpb.DeleteRequest{Id: chatID}
		mc         = minimock.NewController(t)
		serviceErr = fmt.Errorf("service error")
	)

	tests := []struct {
		name                string
		args                args
		want                *emptypb.Empty
		err                 error
		chatServiceMockFunc chatServiceMockFunc
	}{
		{
			name: "успешное удаление чата",
			args: args{ctx: ctx, req: req},
			want: &emptypb.Empty{},
			err:  nil,
			chatServiceMockFunc: func(mc *minimock.Controller) chat.Service {
				mock := mocks.NewServiceMock(mc)
				mock.DeleteMock.Expect(ctx, chatID).Return(nil)

				return mock
			},
		},
		{
			name: "ошибка сервиса при удалении чата",
			args: args{ctx: ctx, req: req},
			want: nil,
			err:  serviceErr,
			chatServiceMockFunc: func(mc *minimock.Controller) chat.Service {
				mock := mocks.NewServiceMock(mc)
				mock.DeleteMock.Expect(ctx, chatID).Return(serviceErr)

				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chatService := tt.chatServiceMockFunc(mc)

			handler := chat.NewImplementation(chatService)

			empty, err := handler.Delete(tt.args.ctx, tt.args.req)

			assert.Equal(t, tt.want, empty)
			if err != nil {
				assert.ErrorContains(t, err, "layer api: не получилось удалить чат:")
				assert.ErrorIs(t, err, tt.err)
			} else {
				assert.Equal(t, tt.err, err)
			}
		})
	}
}
