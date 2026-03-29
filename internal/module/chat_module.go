package module

import (
	"gorm.io/gorm"
	application "chat-service/internal/application/useCases/chat"
	chatRepository "chat-service/internal/infrastructure/database/repository"
)

type ChatModule struct{
	CreateChat 				application.CreateChatUseCase
	FindChatsByUserId application.FindChatsByUserIdUseCase
}

func NewChatModule(db *gorm.DB) ChatModule {
	repository := chatRepository.NewChatRepository(db)
	createChatUseCase := application.NewCreateChatUseCase(repository)
	findChatsByUserIdUseCase := application.NewFindChatsByUserIdUseCase(repository)

	return ChatModule{
		CreateChat: createChatUseCase,
		FindChatsByUserId: findChatsByUserIdUseCase,
	}
}