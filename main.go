package main

import (
	"log"
	"net/http"

	"chat-service/internal/infrastructure/database"
	messageHttp "chat-service/internal/infrastructure/http/router"
	"chat-service/internal/module"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	messageModule := module.NewMessageModule(db)
	router := messageHttp.NewRouter(messageModule)
	log.Println("server running at http://localhost:8080");
	http.ListenAndServe(":8080", router)
}