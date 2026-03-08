package chat

import (
	"chat-service/internal/domain/chat"
	"time"
)

type ChatResponse struct {
	Id           int
	Title        string
	SecondUserId int
	CreatedBy    int
	CreatedAt    time.Time
}

func NewChatResponse(chat chat.Chat) ChatResponse {
	return ChatResponse{
		Id: chat.Id,
		Title: chat.Title,
		CreatedBy: chat.CreatedBy,
		SecondUserId: chat.SecondUserId,
		CreatedAt: chat.CreatedAt,
	}
}