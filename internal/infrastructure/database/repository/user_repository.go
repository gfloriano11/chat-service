package repository

import (
	"chat-service/internal/domain/user"
	"chat-service/internal/infrastructure/database/entity"
	"errors"
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
	User.CreatedAt = entity.CreatedAt
	User.UpdatedBy = &entity.Id
	User.UpdatedAt = entity.UpdatedAt

	return *User, nil
}

func (repository UserRepository) FindUserByEmail(email string) (*user.User, error) {
	var entity entity.User
	err := repository.db.Where("email = ?", email).First(&entity).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	user := user.User{
		Id: entity.Id,
		Fullname: entity.Fullname,
		Email: entity.Email,
		Username: entity.Username,
		Password: entity.Password,
		CreatedAt: entity.CreatedAt,
	}

	return &user, nil
}

func (repository UserRepository) FindUserById(id int) (*user.User, error) {
	var entity entity.User
	err := repository.db.Where("id = ?", id).First(&entity).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	user := user.User{
		Id: entity.Id,
		Fullname: entity.Fullname,
		Email: entity.Email,
		Username: entity.Username,
		Password: entity.Password,
		CreatedAt: entity.CreatedAt,
	}

	return &user, nil
}

func (repository UserRepository) FindUsersNotInChatsWithMeByUserId(id int) (*[]user.User, error) {
	var users []user.User

	err := repository.db.Table("users u").
		Joins(
			`LEFT JOIN chats c 
				ON (u.id = c.created_by AND c.second_user_id = ?) 
				OR (u.id = c.second_user_id AND c.created_by = ?)
			`, id, id).
		Where("u.id != ? AND c.id IS NULL", id).
		Find(&users).
		Error

	if err != nil {
			return nil, err
	}

	return &users, nil
}