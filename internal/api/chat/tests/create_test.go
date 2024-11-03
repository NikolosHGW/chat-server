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
)

func TestCreate(t *testing.T) {
	type chatServiceMockFunc func(mc *minimock.Controller) chat.Service

	type args struct {
		ctx context.Context
		req *chatpb.CreateRequest
	}

	var (
		mc = minimock.NewController(t)

		userIDs = []int64{gofakeit.Int64(), gofakeit.Int64()}
		chatID  = gofakeit.Int64()
		ctx     = context.Background()
		req     = &chatpb.CreateRequest{
			UserIds: userIDs,
		}
		res        = &chatpb.CreateResponse{Id: chatID}
		serviceErr = fmt.Errorf("service error")
	)

	tests := []struct {
		name            string
		args            args
		want            *chatpb.CreateResponse
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "успешное создание чата",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) chat.Service {
				mock := mocks.NewServiceMock(mc)
				mock.CreateMock.Expect(ctx, userIDs).Return(chatID, nil)

				return mock
			},
		},
		{
			name: "ошибка из сервиса при создании чата",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			chatServiceMock: func(mc *minimock.Controller) chat.Service {
				mock := mocks.NewServiceMock(mc)
				mock.CreateMock.Expect(ctx, userIDs).Return(0, serviceErr)

				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chatServiceMock := tt.chatServiceMock(mc)

			chatHandler := chat.NewImplementation(chatServiceMock)

			res, err := chatHandler.Create(tt.args.ctx, tt.args.req)

			assert.Equal(t, tt.want, res)
			if tt.err != nil {
				assert.ErrorContains(t, err, "layer api: не получилось создать чат:")
				assert.ErrorIs(t, err, tt.err)
			} else {
				assert.Equal(t, tt.err, err)
			}
		})
	}
}
