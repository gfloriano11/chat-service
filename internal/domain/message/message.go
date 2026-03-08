package message

import "time"

type Message struct {
	Id        int
	Content   string
	ChatId    int
	CreatedAt time.Time
	CreatedBy int
	UpdatedAt time.Time
	updatedBy int
	DeletedAt time.Time
	DeletedBy int
}