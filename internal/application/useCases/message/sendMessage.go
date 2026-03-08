package message

import (
	"errors"
	"time"

	inputs "chat-service/internal/application/Inputs"
	domain "chat-service/internal/domain"
)

type SendMessageUseCase struct {
	Repository domain.MessageRepository
}

type SendMessageInput struct {
	Content string
	UserId int
}

func NewSendMessageUseCase(repository domain.MessageRepository) SendMessageUseCase {
	return SendMessageUseCase{
		Repository: repository,
	}
}

func (useCase SendMessageUseCase) Execute(sendMessageInput inputs.NewMessageInput) (domain.Message, error) {
	if sendMessageInput.UserId == 0 || sendMessageInput.Content == "" {
		return domain.Message{}, errors.New("It was impossible to send your message")
	}

	message := &domain.Message{
		CreatedBy: sendMessageInput.UserId,
		CreatedAt: time.Now().UTC(),
		Content: sendMessageInput.Content,
		ChatId: 1,
	}

	err := useCase.Repository.Save(message)
	if err != nil {
		return domain.Message{}, err
	}

	return *message, nil
}