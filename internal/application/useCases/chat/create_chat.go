package chat

import (
	inputs "chat-service/internal/application/Inputs"
	"chat-service/internal/domain/chat"
	"errors"
	"time"
)

type CreateChatUseCase struct {
	Repository chat.ChatRepository
}

func NewCreateChatUseCase(repository chat.ChatRepository) CreateChatUseCase {
	return CreateChatUseCase{
		Repository: repository,
	}
}

func (useCase CreateChatUseCase) Execute(input inputs.NewChatInput) (chat.Chat, error) {

	if input.Title == "" {
		return chat.Chat{}, errors.New("It was not possible to create your chat!")
	}

	chat := &chat.Chat{
		Title: input.Title,
		CreatedBy: input.UserId,
		CreatedAt: time.Now().UTC(),
		SecondUserId: input.SecondUserId,
	}

	useCase.Repository.Save(chat)

	return *chat, nil
}