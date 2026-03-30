package router

import (
	"chat-service/internal/infrastructure/security/auth"
	"chat-service/internal/infrastructure/websocket"
	"chat-service/internal/module"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func NewRouter(modules *module.Modules, jwtService auth.JwtService) http.Handler {

	router := chi.NewRouter()

	godotenv.Load()
	environment := os.Getenv("environment")
	var allowedRoutes []string

	if environment == "" {
		allowedRoutes = append(allowedRoutes, os.Getenv("FRONT_DEV_URL"), os.Getenv("FRONT_PROD_URL"))
	} else {
		allowedRoutes = append(allowedRoutes, os.Getenv("FINAL_PROD_URL"))
	}

	log.Println(allowedRoutes)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   allowedRoutes,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	router.Mount("/users", NewUserRouter(modules.UserModule, jwtService))

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