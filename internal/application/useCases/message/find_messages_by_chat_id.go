package message

import "chat-service/internal/domain/message"

type FindMessagesByChatId struct {
	Repository message.MessageRepository
}

func NewFindMessagesByChatId(repository message.MessageRepository) FindMessagesByChatId {
	return FindMessagesByChatId{
		Repository: repository,
	}
}

func (useCase FindMessagesByChatId) Execute(chatId int) ([]message.Message, error) {
	messages, err := useCase.Repository.FindMessagesByChatId(chatId)

	if err != nil {
		return messages, err
	}

	return messages, nil
}