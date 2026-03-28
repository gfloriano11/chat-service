package router

import (
	"chat-service/internal/infrastructure/security/auth"
	"chat-service/internal/infrastructure/websocket"
	"chat-service/internal/module"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(modules *module.Modules, jwtService auth.JwtService) http.Handler {

	router := chi.NewRouter()

	router.Mount("/users", NewUserRouter(modules.UserModule))

	router.Route("/messages", func(r chi.Router) {
		r.Use(jwtService.AuthMiddleware())
		r.Mount("/", NewMessageRouter(modules.MessageModule))
	})

	router.Route("/chats", func(r chi.Router) {
		r.Use(jwtService.AuthMiddleware())
		r.Mount("/", NewChatRouter(modules.ChatModule))
	})

	router.Route("/ws", func(r chi.Router) {
		r.Use(jwtService.AuthMiddleware())
		r.Mount("/", websocket.WebsocketHandler())
	})

	return router
}