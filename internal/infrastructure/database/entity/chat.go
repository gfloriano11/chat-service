package entity

import "time"

type Chat struct {
	Id 						 int        `gorm:"primaryKey"`
	Title          string     `gorm:"type:varchar(255)"`
	CreatedAt      time.Time  `gorm:"type:timestamp"`
	CreatedBy      int
	SecondUserId   int
	UpdatedAt      time.Time  `gorm:"type:timestamp"`
	UpdatedBy 	   int
	IsDeleted 	 	 bool
	DeletedAt 		 *time.Time  `gorm:"type:timestamp"`
	DeletedBy 		 *int
}