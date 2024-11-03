package chat

import "github.com/jmoiron/sqlx"

type repo struct {
	db *sqlx.DB
}

// NewRepo - конструктор для чат репо.
func NewRepo(db *sqlx.DB) *repo {
	return &repo{db: db}
}
