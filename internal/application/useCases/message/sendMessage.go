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

func (useCase SendMessageUseCase) Execute(sendMessageInput inputs.NewMessageInput) error {
	userId := sendMessageInput.UserId
	content := sendMessageInput.Content
	id := 1 // ID FIXO DE MSG
	createdAt := time.Now()

	if userId == 0 || content == "" {
		return errors.New("It was impossible to send your message")
	}

	message := &domain.Message{
			Id: id,
			UserId: userId,
			Content: content,
			CreatedAt: createdAt,
	}

	err := useCase.Repository.Save(message)
	if err != nil {
		return err
	}

	return nil
}