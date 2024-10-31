package domain

import "time"

// Message - модель для бизнес-слоя, то есть для сервисного слоя.
type Message struct {
	ID         int64
	ChatID     int64
	FromUserID int64
	Text       string
	Timestamp  time.Time
}
