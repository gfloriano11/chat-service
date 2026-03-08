package entity

import "time"

type Message struct {
	Id        int      `gorm:"primaryKey"`
	Content   string
	UserId    int
	ChatId    int
	CreatedAt time.Time
}