package chat

import (
	"chat-service/internal/domain/chat"
	"time"
)

type ChatResponse struct {
	Id           int				`json:"id"`
	Title        string			`json:"name"`
	SecondUserId int				`json:"secondUserId"`
	CreatedBy    int				`json:"createdBy"`
	CreatedAt    time.Time	`json:"createdAt"`
}

type ChatListItemResponse struct {
	Id int 										`json:"id"`
	UserID int 								`json:"userId"`
	Fullname string 					`json:"name"`
	LastMessage *string 			`json:"lastMessage"`
	LastMessageAt *time.Time 	`json:"lastMessageSentAt"`
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

func NewChatListItemResponse(chat chat.ChatListItem) ChatListItemResponse {
	return ChatListItemResponse{
		Id: chat.ChatID,
		UserID: chat.UserID,
		Fullname: chat.UserName,
		LastMessage: chat.LastMessage,
		LastMessageAt: chat.LastMessageAt,
	}
} 

func NewChatListResponse(chats []chat.ChatListItem) []ChatListItemResponse {
	chatsResponse := []ChatListItemResponse{}

	for _, chat := range chats {
		chatsResponse = append(chatsResponse, NewChatListItemResponse(chat))
	}

	return chatsResponse
}