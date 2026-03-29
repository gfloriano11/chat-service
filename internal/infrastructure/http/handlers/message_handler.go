package handlers

import (
	chatApplication "chat-service/internal/application/useCases/chat"
	application "chat-service/internal/application/useCases/message"
	"chat-service/internal/infrastructure/http/model/request"
	response "chat-service/internal/infrastructure/http/model/response/message"
	"chat-service/internal/infrastructure/security/auth"
	"chat-service/internal/infrastructure/websocket"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type MessageHandler struct {
	sendMessageUseCase application.SendMessageUseCase
	findMessagesByChatIdUseCase application.FindMessagesByChatId
	findChatByIdUseCase 				chatApplication.FindChatById
}

func NewMessageHandler(
	sendMessage application.SendMessageUseCase, 
	findAllMessages application.FindMessagesByChatId,
	findChatById 		chatApplication.FindChatById,
) MessageHandler {
	return MessageHandler{
		sendMessageUseCase: sendMessage,
		findMessagesByChatIdUseCase: findAllMessages,
		findChatByIdUseCase: findChatById,
	}
}

func (handler MessageHandler) GetMessagesByChatId(w http.ResponseWriter, r *http.Request) {
	
	chatId, err := strconv.Atoi(chi.URLParam(r, "chatId"))

	if err != nil {
		http.Error(w, "invalid body", http.StatusInternalServerError)
	}

	messages, err := handler.findMessagesByChatIdUseCase.Execute(chatId)

	if err != nil {
		http.Error(w, "Error trying to get messages", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.NewMessagesResponse(messages))
}

func (handler MessageHandler) SendMessage(w http.ResponseWriter, r *http.Request) {

	chatId, err := strconv.Atoi(chi.URLParam(r, "chatId"))
	userId, _ := auth.GetUserIdFromContext(r.Context())

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

	message, err := handler.sendMessageUseCase.Execute(newMessageRequest.ToNewMessageInput(chatId, userId))
	
	if err != nil {
		switch err {

		case application.ErrUserNotInChat:
			http.Error(w, err.Error(), http.StatusUnauthorized)

		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	
	chat, err := handler.findChatByIdUseCase.Execute(chatId)
	
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	
	var userToSendEvent int
	
	if message.CreatedBy == chat.CreatedBy {
		userToSendEvent = chat.SecondUserId
	} else {
		userToSendEvent = chat.CreatedBy
	}

	event := websocket.WsEvent{
		Type:    "NEW_MESSAGE",
		Payload: response.NewFullMessageResponse(message),
	}

	websocket.SendToUser(userToSendEvent, event)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response.NewMessageResponse(message))
}