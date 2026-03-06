package router

import (
	"chat-service/internal/module"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(messageModule module.MessageModule) http.Handler {

	router := chi.NewRouter()

	router.Mount("/messages", NewMessageRouter(messageModule))

	return router
}