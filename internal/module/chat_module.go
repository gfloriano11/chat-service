package module

import (
	"gorm.io/gorm"
	application "chat-service/internal/application/useCases/chat"
	chatRepository "chat-service/internal/infrastructure/database/repository"
)

type ChatModule struct{
	CreateChat application.CreateChatUseCase
}

func NewChatModule(db *gorm.DB) ChatModule {
	repository := chatRepository.NewChatRepository(db)
	createChatUseCase := application.NewCreateChatUseCase(repository)

	return ChatModule{
		CreateChat: createChatUseCase,
	}
}