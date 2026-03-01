package request

import (
	inputs "chat-service/internal/application/Inputs"
)

type NewMessageRequest struct {
	Content string `json:"content"`
	UserId  int    `json:"userId"`
}

func (newMessageRequest NewMessageRequest) ToNewMessageInput() inputs.NewMessageInput {
	return inputs.NewMessageInput{
		Content: newMessageRequest.Content,
		UserId: newMessageRequest.UserId,
	}
}