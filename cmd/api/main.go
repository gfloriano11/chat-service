package main

import (
	"log"
	"net/http"
	"os"

	"chat-service/internal/infrastructure/database"
	messageHttp "chat-service/internal/infrastructure/http/router"
	"chat-service/internal/infrastructure/security/auth"
	"chat-service/internal/module"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	environment := os.Getenv("ENV")
	if environment == "" {
		environment = "development"
	}
	log.Printf("Running in %s mode", environment)
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	modules := module.CreateModules(db)
	router := messageHttp.NewRouter(modules, auth.NewJwtService())
	log.Println("server running at http://localhost:8080");
	http.ListenAndServe(":8080", router)
}