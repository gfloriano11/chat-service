package websocket

import (
	"chat-service/internal/infrastructure/security/auth"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[int]*websocket.Conn)
var mu sync.Mutex

func WebsocketHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := wsUpgrader.Upgrade(w, r, nil)

		if err != nil {
			log.Println("error upgrading:", err)
			return
		}
			
		userId, ok := auth.GetUserIdFromContext(r.Context())
		if !ok {
			log.Println("userId not found in context")
			conn.Close()
			return
		}

		mu.Lock()
		clients[userId] = conn
		mu.Unlock()

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("disconnected:", userId)
				delete(clients, userId)
				conn.Close()
				break
			}

			log.Println("received:", string(msg))
		}
	}
}