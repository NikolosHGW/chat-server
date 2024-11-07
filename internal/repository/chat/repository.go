package chat

import (
	"github.com/NikolosHGW/platform-common/pkg/db"
)

type repo struct {
	db db.Client
}

// NewRepo - конструктор для чат репо.
func NewRepo(db db.Client) *repo {
	return &repo{db: db}
}

// PublicRepo - алиас для структуры чат-репозитория.
// Сделано в экспериментальных целях, чтоб пробросить
// сервис провайдер чёткую структуру без менющихся интерфейсов.
type PublicRepo = repo
