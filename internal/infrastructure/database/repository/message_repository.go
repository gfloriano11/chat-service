package repository

import (
	"chat-service/internal/domain"
	"chat-service/internal/infrastructure/database/entity"
	"time"

	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{
		db: db,
	}
}

func (repository MessageRepository) Save(message *domain.Message) error {

	entity := entity.Message{
		Content: message.Content,
		CreatedBy: message.CreatedBy,
		ChatId: message.ChatId,
		CreatedAt: time.Now(),
	}

	err := repository.db.Create(&entity).Error

	if err != nil {
		return err
	}

	message.Id = int(entity.Id)

	return nil
}

func (repository MessageRepository) FindMessagesByChatId(id int) ([]domain.Message, error) {

	var entities []entity.Message
	err := repository.db.Find(&entities).Error

	if err != nil {
		return []domain.Message{}, nil
	}

	var messages []domain.Message

	for _, entity := range entities {
		messages = append(messages, domain.Message{
			Id:      entity.Id,
			Content: entity.Content,
			CreatedBy:  entity.CreatedBy,
			ChatId:  entity.ChatId,
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
		})
	}

	return messages, nil
}