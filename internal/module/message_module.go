package module

import (
	application "chat-service/internal/application/useCases/message"
	memoryRepository "chat-service/internal/infrastructure/repository/memory"
)

type MessageModule struct {
	SendMessage application.SendMessageUseCase
	FindAllMessages application.GetMessageUseCase
}

func NewMessageModule() MessageModule {
	repository := memoryRepository.NewMessageRepositoryMemory()
	getMessageUseCase := application.NewGetMessageUseCase(repository)
	sendMessageUseCase := application.NewSendMessageUseCase(repository)

	return MessageModule{
		sendMessageUseCase,
		getMessageUseCase,
	}
}