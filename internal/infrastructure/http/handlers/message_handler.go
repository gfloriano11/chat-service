package handlers

import (
	application "chat-service/internal/application/useCases/message"
	"chat-service/internal/infrastructure/http/model/request"
	response "chat-service/internal/infrastructure/http/model/response/message"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type MessageHandler struct {
	sendMessageUseCase application.SendMessageUseCase
	findMessagesByChatIdUseCase application.FindMessagesByChatId
}

func NewMessageHandler(sendMessage application.SendMessageUseCase, findAllMessages application.FindMessagesByChatId) MessageHandler {
	return MessageHandler{
		sendMessageUseCase: sendMessage,
		findMessagesByChatIdUseCase: findAllMessages,
	}
}

func (handler MessageHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
	
	chatId, err := strconv.Atoi(chi.URLParam(r, "chatId"))

	if err != nil {
		http.Error(w, "invalid body", http.StatusInternalServerError)
	}

	messages, err := handler.findMessagesByChatIdUseCase.Execute(chatId)

	if err != nil {
		http.Error(w, "Error trying to get messages", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.NewMessagesResponse(messages))
}

func (handler MessageHandler) SendMessage(w http.ResponseWriter, r *http.Request) {

	chatId, err := strconv.Atoi(chi.URLParam(r, "chatId"))

	if err != nil {
		http.Error(w, "Chat Id must be a valid id", http.StatusInternalServerError)
	}

	var newMessageRequest request.NewMessageRequest

	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(&newMessageRequest)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	message, err := handler.sendMessageUseCase.Execute(newMessageRequest.ToNewMessageInput(chatId))
	
	if err != nil {
		switch err {

		case application.ErrUserNotInChat:
			http.Error(w, err.Error(), http.StatusUnauthorized)

		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response.NewMessageResponse(message))
}