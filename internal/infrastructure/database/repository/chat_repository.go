package repository

import (
	"chat-service/internal/domain/chat"
	"chat-service/internal/infrastructure/database/entity"
	"time"

	"gorm.io/gorm"
)

type ChatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) *ChatRepository {
	return &ChatRepository{
		db: db,
	}
}

func (repository ChatRepository) Save(chat *chat.Chat) error {

	entity := entity.Chat{
		Title: chat.Title,
		CreatedAt: time.Now().UTC(),
		CreatedBy: chat.CreatedBy,
		SecondUserId: chat.SecondUserId,
		UpdatedAt: time.Now().UTC(),
		UpdatedBy: chat.CreatedBy,
	}
	
	err := repository.db.Create(&entity).Error

	if err != nil {
		return err
	}

	chat.Id = entity.Id 

	return nil
}