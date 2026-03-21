package router

import (
	"chat-service/internal/infrastructure/http/handlers"
	"chat-service/internal/module"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewMessageRouter(messageModule module.MessageModule) http.Handler {
	r := chi.NewRouter()

	messageHandler := handlers.NewMessageHandler(
		messageModule.SendMessage,
		messageModule.FindMessagesByChatId,
	)

	r.Get("/{chatId}", messageHandler.GetMessages)
	r.Post("/{chatId}", messageHandler.SendMessage)

	return r
}