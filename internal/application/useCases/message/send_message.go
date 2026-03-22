package message

import (
	"errors"
	"time"

	inputs "chat-service/internal/application/Inputs"
	"chat-service/internal/domain/chat"
	message "chat-service/internal/domain/message"
)

type SendMessageUseCase struct {
	MessageRepository message.MessageRepository
	ChatRepository chat.ChatRepository
}

type SendMessageInput struct {
	Content string
	UserId int
}

func NewSendMessageUseCase(repository message.MessageRepository) SendMessageUseCase {
	return SendMessageUseCase{
		MessageRepository: repository,
	}
}

func (useCase SendMessageUseCase) Execute(sendMessageInput inputs.NewMessageInput) (message.Message, error) {

	if sendMessageInput.Content == "" {
		return message.Message{}, errors.New("It was impossible to send your message")
	}

	isUserInChat := useCase.ChatRepository.IsUserParticipant(sendMessageInput.ChatId, sendMessageInput.UserId)

	if !isUserInChat {
		return message.Message{}, ErrUserNotInChat
	}

	newMessage := &message.Message{
		CreatedBy: sendMessageInput.UserId,
		CreatedAt: time.Now().UTC(),
		Content: sendMessageInput.Content,
		ChatId: sendMessageInput.ChatId,
	}

	err := useCase.MessageRepository.Save(newMessage)
	if err != nil {
		return message.Message{}, err
	}

	return *newMessage, nil
}