package user

import "time"

type User struct {
	Id        int
	Fullname  string
	Email     string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
	UpdatedBy int
	IsDeleted bool
	DeletedAt *time.Time
	DeletedBy *int
}