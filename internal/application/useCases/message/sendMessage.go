package message

import (
	"chat-service/internal/domain"
	"errors"
	"time"
)

type SendMessageUseCase struct {
	Repository domain.MessageRepository
}

type SendMessageInput struct {
	Content string
	UserId int
}

func (useCase SendMessageUseCase) Execute(sendMessageInput SendMessageInput) error {
	userId := sendMessageInput.UserId
	content := sendMessageInput.Content
	id := 1
	createdAt := time.Now()

	if (userId == 0 || content == "") {
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