package router

import (
	"chat-service/internal/infrastructure/http/handlers"
	"chat-service/internal/module"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewUserRouter(userModule module.UserModule) http.Handler {
	r := chi.NewRouter()

	userHandler := handlers.NewUserHandler(
		userModule.CreateUser,
		userModule.FindUserByEmail,
	)
	
	r.Post("/register", userHandler.CreateUser)
	r.Post("/login", userHandler.Login)

	return r
}