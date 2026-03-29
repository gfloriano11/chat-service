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

func (repository ChatRepository) IsUserParticipant(id, userId int) bool {
	var entity entity.Chat
	var count int64
	repository.db.Where("id = ? AND (created_by = ? OR second_user_id = ?)", id, userId, userId).Find(&entity).Count(&count)
	return count > 0
}

func (repository ChatRepository) FindChatsByUserId(id int) (*[]chat.ChatListItem, error) {
	var chats []chat.ChatListItem

	err := repository.db.Raw(`
		SELECT 
			c.id AS chat_id,
			u.id AS user_id,
			u.fullname AS user_name,
			m.content AS last_message,
			m.created_at AS last_message_at
		FROM chats c

		JOIN users u ON u.id = 
			CASE 
				WHEN c.created_by = ? THEN c.second_user_id
				ELSE c.created_by
			END

		LEFT JOIN messages m ON m.id = (
			SELECT id 
			FROM messages 
			WHERE chat_id = c.id 
			ORDER BY created_at DESC 
			LIMIT 1
		)

		WHERE c.created_by = ? OR c.second_user_id = ?
		ORDER BY m.created_at DESC
	`, id, id, id).
	Scan(&chats).Error

	if err != nil {
		return nil, err
	}

	return &chats, nil
}

func (repository ChatRepository) FindChatById(id int) (*chat.Chat, error) {
	var chat chat.Chat

	err := repository.db.Where("id = ?", id).First(&chat).Error

	if err != nil {
		return nil, err
	}

	return &chat, nil
}