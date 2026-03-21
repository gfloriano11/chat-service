package repository

import (
	"chat-service/internal/domain/user"
	"chat-service/internal/infrastructure/database/entity"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repository UserRepository) Save(User *user.User) (user.User, error) {

	entity := entity.User{
		Email: User.Email,
		Username: User.Username,
		Fullname: User.Fullname,
		Password: User.Password,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UpdatedBy: nil,
		IsDeleted: false,
		DeletedAt: nil,
		DeletedBy: nil,
	}

	err := repository.db.Create(&entity).Error

	if err != nil {
		return user.User{}, err
	}

	User.Id = entity.Id
	User.UpdatedBy = entity.UpdatedBy

	return *User, nil
}