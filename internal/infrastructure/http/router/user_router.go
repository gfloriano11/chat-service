package router

import (
	"chat-service/internal/infrastructure/http/handlers"
	"chat-service/internal/infrastructure/security/auth"
	"chat-service/internal/module"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewUserRouter(userModule module.UserModule, jwtService auth.JwtService) http.Handler {
	r := chi.NewRouter()

	userHandler := handlers.NewUserHandler(
		userModule.CreateUser,
		userModule.Login,
		userModule.GetMe,
	)
	
	r.Post("/register", userHandler.CreateUser)
	r.Post("/login", userHandler.Login)
	r.With(jwtService.AuthMiddleware()).Get("/me", userHandler.GetMe)

	return r
}