package chat

import "time"

type Chat struct {
	Id        int
	Title 	  string
	CreatedBy int
	CreatedAt time.Time
	SecondUserId int
	isDeleted bool
	DeletedAt time.Time
	UpdatedAt time.Time
	UpdatedBy int
}