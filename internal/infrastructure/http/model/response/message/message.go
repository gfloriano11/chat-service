package message

import (
	"chat-service/internal/domain/message"
	"time"
)

type MessageResponse struct {
	Id        int       `json:"id"`
	Content   string    `json:"content"`
	CreatedBy int       `json:"createdBy"`
	SentAt    time.Time `json:"sentAt"`
}

func NewMessageResponse(message message.Message) MessageResponse {
	messageResponse := MessageResponse{
		message.Id,
		message.Content,
		message.CreatedBy,
		message.CreatedAt,
	}

	return messageResponse
}

func NewMessagesResponse(messages []message.Message) []MessageResponse {
	messagesResponse := []MessageResponse{}

	for _, message := range messages {
		messageResponse := NewMessageResponse(message)
		messagesResponse = append(messagesResponse, messageResponse)
	}

	return messagesResponse
}