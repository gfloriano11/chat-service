package chat

import (
	"chat-service/internal/domain/chat"
)

type FindChatsByUserIdUseCase struct {
	Repository chat.ChatRepository
}

func NewFindChatsByUserIdUseCase(repository chat.ChatRepository) FindChatsByUserIdUseCase {
	return FindChatsByUserIdUseCase{
		Repository: repository,
	}
}

func (useCase FindChatsByUserIdUseCase) Execute(userId int) (*[]chat.ChatListItem, error) {
	chats, err := useCase.Repository.FindChatsByUserId(userId)

	if err != nil {
		return nil, ErrGeneric
	}

	if chats == nil {
		return &[]chat.ChatListItem{}, nil
	}

	return chats, nil
}