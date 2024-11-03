package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/NikolosHGW/chat-server/internal/api/chat"
	"github.com/NikolosHGW/chat-server/internal/api/chat/mocks"
	"github.com/NikolosHGW/chat-server/internal/domain"
	chatpb "github.com/NikolosHGW/chat-server/pkg/chat/v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestSendMessage(t *testing.T) {
	type args struct {
		ctx context.Context
		req *chatpb.SendMessageRequest
	}

	var (
		ctx         = context.Background()
		chatID      = gofakeit.Int64()
		userID      = gofakeit.Int64()
		textMessage = gofakeit.SentenceSimple()
		req         = &chatpb.SendMessageRequest{
			ChatId:     chatID,
			FromUserId: userID,
			Text:       textMessage,
			Timestamp:  timestamppb.Now(),
		}
		domainMessage = &domain.Message{
			ChatID:     chatID,
			FromUserID: userID,
			Text:       textMessage,
			Timestamp:  req.Timestamp.AsTime(),
		}
		mc         = minimock.NewController(t)
		serviceErr = fmt.Errorf("service error")
	)

	tests := []struct {
		name                string
		args                args
		want                *emptypb.Empty
		err                 error
		chatServiceMockFunc func(mc *minimock.Controller) chat.Service
	}{
		{
			name: "успешная отправка сообщения в чат",
			args: args{ctx: ctx, req: req},
			want: &emptypb.Empty{},
			err:  nil,
			chatServiceMockFunc: func(mc *minimock.Controller) chat.Service {
				mock := mocks.NewServiceMock(mc)
				mock.SendMessageMock.Expect(ctx, domainMessage).Return(nil)

				return mock
			},
		},
		{
			name: "ошибка сервиса при отправки сообщения в чат",
			args: args{ctx: ctx, req: req},
			want: nil,
			err:  serviceErr,
			chatServiceMockFunc: func(mc *minimock.Controller) chat.Service {
				mock := mocks.NewServiceMock(mc)
				mock.SendMessageMock.Expect(ctx, domainMessage).Return(serviceErr)

				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chatService := tt.chatServiceMockFunc(mc)

			handler := chat.NewImplementation(chatService)

			empty, err := handler.SendMessage(tt.args.ctx, tt.args.req)

			assert.Equal(t, tt.want, empty)
			if tt.err != nil {
				assert.ErrorContains(t, err, "layer api: не удалось отправить сообщение:")
				assert.ErrorIs(t, err, tt.err)
			} else {
				assert.Equal(t, tt.err, err)
			}
		})
	}
}
