package messages

import (
	"chat-service/internal/domain"
	"time"
)

type MessageResponse struct {
	Id int
	Content string
	CreatedBy int
	SentAt time.Time
}

func NewMessagesResponse(messages []domain.Message) []MessageResponse {
	messagesResponse := []MessageResponse{}

	for _, message := range messages {
		messageResponse := MessageResponse{
			message.Id,
			message.Content,
			message.UserId,
			message.CreatedAt,
		}

		messagesResponse = append(messagesResponse, messageResponse)
	}

	return messagesResponse
}