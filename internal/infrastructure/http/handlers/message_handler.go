package handlers

import (
	application "chat-service/internal/application/useCases/message"
	"chat-service/internal/infrastructure/http/model/request"
	response "chat-service/internal/infrastructure/http/model/response/messages"
	"encoding/json"
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

func (handler MessageHandler) GetMessages(w http.ResponseWriter, r *http.Request) ([]response.MessageResponse, error) {
	
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

func (handler MessageHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
	var newMessageRequest request.NewMessageRequest

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&newMessageRequest)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	err = handler.sendMessageUseCase.Execute(newMessageRequest.ToNewMessageInput())
	if err != nil {
		http.Error(w, "error creating message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response.MessageResponse{})
}