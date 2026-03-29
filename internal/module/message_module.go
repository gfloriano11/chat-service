package module

import (
	application "chat-service/internal/application/useCases/message"
	chatApplication "chat-service/internal/application/useCases/chat"
	repository "chat-service/internal/infrastructure/database/repository"

	"gorm.io/gorm"
)

type MessageModule struct {
	SendMessage 					application.SendMessageUseCase
	FindMessagesByChatId 	application.FindMessagesByChatId
	FindChatById					chatApplication.FindChatById
}

func NewMessageModule(db *gorm.DB) MessageModule {
	messageRepository := repository.NewMessageRepository(db)
	chatRepository := repository.NewChatRepository(db)
	findMessagesByChatId := application.NewFindMessagesByChatId(messageRepository)
	sendMessageUseCase := application.NewSendMessageUseCase(messageRepository, chatRepository)
	findChatByIdUseCase := chatApplication.NewFindChatByIdUsecase(chatRepository)

	return MessageModule{
		sendMessageUseCase,
		findMessagesByChatId,
		findChatByIdUseCase,
	}
}