package messageHttp

import (
	"net/http"

	handlers "chat-service/internal/infrastructure/http/handlers"
	"chat-service/internal/module"
)

func NewRouter(messageModule module.MessageModule) *http.ServeMux {
	mux := http.NewServeMux()

	// mux.Handle("/health", handlers.HealthHandler{})

	messageHandler := handlers.NewMessageHandler(messageModule.SendMessage, messageModule.FindMessagesByChatId)
	mux.Handle("/messages", messageHandler)
	return mux
}