package message

import (
	"chat-service/internal/domain"
	"time"
)

type MessageResponse struct {
	Id        int
	Content   string
	CreatedBy int
	SentAt    time.Time
}

func NewMessageResponse(message domain.Message) MessageResponse {
	messageResponse := MessageResponse{
		message.Id,
		message.Content,
		message.UserId,
		message.CreatedAt,
	}

	return messageResponse
}

func NewMessagesResponse(messages []domain.Message) []MessageResponse {
	messagesResponse := []MessageResponse{}

	for _, message := range messages {
		messageResponse := NewMessageResponse(message)
		messagesResponse = append(messagesResponse, messageResponse)
	}

	return messagesResponse
}