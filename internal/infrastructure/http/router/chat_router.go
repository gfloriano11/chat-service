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
		chatModule.FindChatsByUserId,
	)

	r.Get("/{userId}", chatHandler.GetChatsByUserId)
	r.Post("/", chatHandler.CreateChat)

	return r
}