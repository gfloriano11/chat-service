package module

import (
	application "chat-service/internal/application/useCases/message"
	messageRepository "chat-service/internal/infrastructure/database/repository"

	"gorm.io/gorm"
)

type MessageModule struct {
	SendMessage application.SendMessageUseCase
	FindMessagesByChatId application.FindMessagesByChatId
}

func NewMessageModule(db *gorm.DB) MessageModule {
	repository := messageRepository.NewMessageRepository(db)
	findMessagesByChatId := application.NewFindMessagesByChatId(repository)
	sendMessageUseCase := application.NewSendMessageUseCase(repository)

	return MessageModule{
		sendMessageUseCase,
		findMessagesByChatId,
	}
}