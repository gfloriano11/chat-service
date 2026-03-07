package handlers

import (
	application "chat-service/internal/application/useCases/message"
	response "chat-service/internal/infrastructure/http/model/response/messages"
	"net/http"
	"strconv"
)

type MessageHandler struct {
	sendMessageUseCase application.SendMessageUseCase
	findMessagesByChatId application.FindMessagesByChatId
}

func NewMessageHandler(sendMessage application.SendMessageUseCase, findAllMessages application.FindMessagesByChatId) MessageHandler {
	return MessageHandler{
		sendMessageUseCase: sendMessage,
		findMessagesByChatId: findAllMessages,
	}
}

func (handler MessageHandler) GetMessages(w http.ResponseWriter, r *http.Request) ([]response.MessageResponse, error){
	
	chatId, err := strconv.Atoi(r.URL.Query().Get("chatId"))

	if err != nil {
		return []response.MessageResponse{}, err
	}

	messages, err := handler.findMessagesByChatId.Execute(chatId)

	if err != nil {
		return []response.MessageResponse{}, err
	}

	return response.NewMessagesResponse(messages), nil
}