package request

import (
	inputs "chat-service/internal/application/Inputs"
)

type NewMessageRequest struct {
	Content string `json:"content"`
}

func (newMessageRequest NewMessageRequest) ToNewMessageInput(chatId, userId int) inputs.NewMessageInput {
	return inputs.NewMessageInput{
		Content: newMessageRequest.Content,
		UserId: userId,
		ChatId: chatId,
	}
}