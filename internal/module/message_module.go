package module

import (
	application "chat-service/internal/application/useCases/message"
	repository "chat-service/internal/infrastructure/database/repository"

	"gorm.io/gorm"
)

type MessageModule struct {
	SendMessage application.SendMessageUseCase
	FindMessagesByChatId application.FindMessagesByChatId
}

func NewMessageModule(db *gorm.DB) MessageModule {
	messageRepository := repository.NewMessageRepository(db)
	chatRepository := repository.NewChatRepository(db)
	findMessagesByChatId := application.NewFindMessagesByChatId(messageRepository)
	sendMessageUseCase := application.NewSendMessageUseCase(messageRepository, chatRepository)

	return MessageModule{
		sendMessageUseCase,
		findMessagesByChatId,
	}
}