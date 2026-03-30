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
	environment := os.Getenv("environment")
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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("server running at port ", port);
	http.ListenAndServe(":8080"+port, router)
}