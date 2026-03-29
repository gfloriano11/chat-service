package chat

import (
	"chat-service/internal/domain/chat"
)

type FindChatById struct {
	Repository chat.ChatRepository
}

func NewFindChatByIdUsecase(repository chat.ChatRepository) FindChatById {
	return FindChatById{
		Repository: repository,
	}
}

func (useCase FindChatById) Execute(id int) (*chat.Chat, error) {
	foundChat, err := useCase.Repository.FindChatById(id)

	if err != nil {
		return nil, ErrGeneric
	}

	if foundChat == nil {
		return &chat.Chat{}, nil
	}

	return foundChat, nil
}