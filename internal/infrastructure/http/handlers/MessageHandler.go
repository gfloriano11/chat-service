package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	application "chat-service/internal/application/useCases/message"
	"chat-service/internal/infrastructure/http/model/request"
)

type MessageHandler struct {
	sendMessageUseCase application.SendMessageUseCase
	findAllMessagesUseCase application.GetMessageUseCase
}

func NewMessageHandler(sendMessage application.SendMessageUseCase, findAllMessages application.GetMessageUseCase) MessageHandler {
	return MessageHandler{
		sendMessageUseCase: sendMessage,
		findAllMessagesUseCase: findAllMessages,
	}
}

func (handler MessageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("route on!")
	log.Println(r.URL.Path)
	url, method := getRequestInfo(r)

	log.Println(url)

	if method == http.MethodGet && url == "" {
		log.Println("Trying to find chat messages")
		w.WriteHeader(http.StatusOK)
	}

	if method == http.MethodPost && url == "" {
		log.Println("Trying to send message")
		var newMessageRequest request.NewMessageRequest

		err := json.NewDecoder(r.Body).Decode(&newMessageRequest)
		if err != nil {
			http.Error(w, "invalid body", http.StatusBadRequest)
			return
		}

		handler.sendMessageUseCase.Execute(newMessageRequest.ToNewMessageInput())
		w.WriteHeader(http.StatusCreated)
	}
}

func getRequestInfo(r *http.Request) (string, string) {
	return strings.TrimPrefix(r.URL.Path, "/messages"), r.Method
}