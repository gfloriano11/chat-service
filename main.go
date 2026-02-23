package main

import (
	"log"
	"net/http"

	messageHttp "chat-service/internal/infrastructure/http"
)

func main() {
	router := messageHttp.NewRouter()
	log.Println("server running at http://localhost:8080");
	http.ListenAndServe(":8080", router)
}
