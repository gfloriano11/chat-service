package domain

import "time"

type Message struct {
	ID        int
	UserID    int
	Content   string
	CreatedAt time.Time
}