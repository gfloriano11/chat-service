package request

import (
	inputs "chat-service/internal/application/Inputs"
)

type NewChatRequest struct {
	Title        string `json:"title"`
	CreatedBy    int    `json:"userId"`
	SecondUserId int    `json:"secondUserid"` 
}

func (newChatRequest NewChatRequest) ToNewChatInput() inputs.NewChatInput {
	return inputs.NewChatInput{
		Title: newChatRequest.Title,
		UserId: newChatRequest.CreatedBy,
		SecondUserId: newChatRequest.SecondUserId,
	}
}