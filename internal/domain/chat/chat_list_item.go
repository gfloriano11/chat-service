package chat

import (
	"time"
)

type ChatListItem struct {
	ChatID        int
	UserID        int
	UserName      string
	LastMessage   *string
	LastMessageAt *time.Time
}