package model

import (
	"time"
)

// Message структура сообщения
type Message struct {
	ID        int64
	Info      MessageInfo
	State     int32
	Type      int32
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

// MessageInfo структура информации о сообщении
type MessageInfo struct {
	ChatID int64
	UserID int64
	Body   string
}
