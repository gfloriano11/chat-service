package module

import (
	application "chat-service/internal/application/useCases/message"
	memoryRepository "chat-service/internal/infrastructure/repository/memory"
)

type MessageModule struct {
	SendMessage application.SendMessageUseCase
	FindMessagesByChatId application.FindMessagesByChatId
}

func NewMessageModule() MessageModule {
	repository := memoryRepository.NewMessageRepositoryMemory()
	findMessagesByChatId := application.NewFindMessagesByChatId(repository)
	sendMessageUseCase := application.NewSendMessageUseCase(repository)

	return MessageModule{
		sendMessageUseCase,
		findMessagesByChatId,
	}
}