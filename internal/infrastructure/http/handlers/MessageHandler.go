package handlers

import (
	"net/http"

	application "chat-service/application/useCases/message"
)

type MessageHandler struct {
	SendMessage application.SendMessageUseCase
}

func (handler MessageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}