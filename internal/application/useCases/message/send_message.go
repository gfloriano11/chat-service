package message

import (
	"errors"
	"time"

	inputs "chat-service/internal/application/Inputs"
	message "chat-service/internal/domain/message"
)

type SendMessageUseCase struct {
	Repository message.MessageRepository
}

type SendMessageInput struct {
	Content string
	UserId int
}

func NewSendMessageUseCase(repository message.MessageRepository) SendMessageUseCase {
	return SendMessageUseCase{
		Repository: repository,
	}
}

func (useCase SendMessageUseCase) Execute(sendMessageInput inputs.NewMessageInput) (message.Message, error) {
	if sendMessageInput.UserId == 0 || sendMessageInput.Content == "" {
		return message.Message{}, errors.New("It was impossible to send your message")
	}

	newMessage := &message.Message{
		CreatedBy: sendMessageInput.UserId,
		CreatedAt: time.Now().UTC(),
		Content: sendMessageInput.Content,
		ChatId: 1,
	}

	err := useCase.Repository.Save(newMessage)
	if err != nil {
		return message.Message{}, err
	}

	return *newMessage, nil
}