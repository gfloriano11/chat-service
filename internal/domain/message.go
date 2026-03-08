package domain

import "time"

type Message struct {
	Id        int
	UserId    int
	Content   string
	CreatedAt time.Time
	ChatId    int
}