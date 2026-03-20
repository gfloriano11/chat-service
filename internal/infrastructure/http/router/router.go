package router

import (
	"chat-service/internal/module"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(modules *module.Modules) http.Handler {

	router := chi.NewRouter()

	router.Mount("/messages", NewMessageRouter(modules.MessageModule))
	router.Mount("/chats", NewChatRouter(modules.ChatModule))
	router.Mount("/users", NewUserRouter(modules.UserModule))

	return router
}