package model

import "time"

// Chat структура чата
type Chat struct {
	ID        int64
	Info      ChatInfo
	State     int32
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

// ChatInfo структура информации о чате
type ChatInfo struct {
	Title   string
	UserIDs []int64
}
