package messageHttp

import (
	"net/http"

	handlers "chat-service/internal/infrastructure/http/handlers"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// mux.Handle("/health", handlers.HealthHandler{})

	mux.Handle("/messages", handlers.MessageHandler{})
	return mux
}