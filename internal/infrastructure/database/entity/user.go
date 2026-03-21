package entity

import "time"

type User struct {
	Id 				int					`gorm:"primaryKey"`
	Fullname 	string
	Email 		string			`gorm:"type:varchar(255)"`
	Password 	string			`gorm:"type:varchar(255)"`
	Username 	string 		 	`gorm:"type:varchar(255)"`
	CreatedAt time.Time  	`gorm:"type:timestamp"`
	UpdatedAt time.Time  	`gorm:"type:timestamp"`
	UpdatedBy *int
	IsDeleted bool
	DeletedAt *time.Time 	`gorm:"type:timestamp"`
	DeletedBy *int
} 