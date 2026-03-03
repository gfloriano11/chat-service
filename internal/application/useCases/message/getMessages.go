package message

import "chat-service/internal/domain"

type FindMessagesByChatId struct {
	Repository domain.MessageRepository
}

func NewFindMessagesByChatId(repository domain.MessageRepository) FindMessagesByChatId {
	return FindMessagesByChatId{
		Repository: repository,
	}
}

func (useCase FindMessagesByChatId) Execute(chatId int) error {
	

	return nil
}