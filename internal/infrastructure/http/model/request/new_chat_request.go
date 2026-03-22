package request

import (
	inputs "chat-service/internal/application/Inputs"
)

type NewChatRequest struct {
	Title        string `json:"title"`
	SecondUserId int    `json:"secondUserId"` 
}

func (newChatRequest NewChatRequest) ToNewChatInput(createdBy int) inputs.NewChatInput {
	return inputs.NewChatInput{
		Title: newChatRequest.Title,
		UserId: createdBy,
		SecondUserId: newChatRequest.SecondUserId,
	}
}