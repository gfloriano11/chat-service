package router

import (
	"chat-service/internal/infrastructure/http/handlers"
	"chat-service/internal/module"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewChatRouter(chatModule module.ChatModule) http.Handler {
	r := chi.NewRouter()

	chatHandler := handlers.NewChatHandler(
		chatModule.CreateChat,
	)

	r.Get("/", chatHandler.GetChat)
	r.Post("/", chatHandler.CreateChat)

	return r
}