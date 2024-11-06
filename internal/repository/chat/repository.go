package chat

import (
	"github.com/NikolosHGW/chat-server/internal/client/db"
)

type repo struct {
	db db.Client
}

// NewRepo - конструктор для чат репо.
func NewRepo(db db.Client) *repo {
	return &repo{db: db}
}

type PublicRepo = repo
