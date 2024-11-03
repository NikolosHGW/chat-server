package dto

import "time"

// MessageDTO - модель сообщения для слоя репозиторий.
type MessageDTO struct {
	Text       string    `db:"text"`
	ID         int64     `db:"id"`
	ChatID     int64     `db:"chat_id"`
	FromUserID int64     `db:"from_user_id"`
	Timestamp  time.Time `db:"timestamp"`
}
