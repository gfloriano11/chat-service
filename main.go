package main

import (
	"log"
	"net/http"

	messageHttp "chat-service/internal/infrastructure/http/router"
	"chat-service/internal/module"
)

func main() {
	messageModule := module.NewMessageModule()
	router := messageHttp.NewRouter(messageModule)
	log.Println("server running at http://localhost:8080");
	http.ListenAndServe(":8080", router)
}