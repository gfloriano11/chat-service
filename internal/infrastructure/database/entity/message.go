package entity

import "time"

type Message struct {
	Id        int       `gorm:"primaryKey"`
	Content   string
	ChatId    int
	CreatedAt time.Time `gorm:"type:timestamp"`
	CreatedBy int
	UpdatedAt time.Time `gorm:"type:timestamp"`
	UpdatedBy int
	IsDeleted bool
	DeletedAt *time.Time `gorm:"type:timestamp"`
	DeletedBy *int
}