package handlers

import (
	application "chat-service/internal/application/useCases/chat"
	"chat-service/internal/infrastructure/http/model/request"
	response "chat-service/internal/infrastructure/http/model/response/chat"
	"chat-service/internal/infrastructure/security/auth"
	"encoding/json"
	"net/http"
)

type ChatHandler struct {
	CreateChatUseCase application.CreateChatUseCase
}

func NewChatHandler(createChat application.CreateChatUseCase) ChatHandler {
	return ChatHandler{
		CreateChatUseCase: createChat,
	}
}

func (handler ChatHandler) GetChat(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Connected!"))
}

func (handler ChatHandler) CreateChat(w http.ResponseWriter, r *http.Request) {

	userId, _ := auth.GetUserIdFromContext(r.Context())
	var newChatRequest request.NewChatRequest

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&newChatRequest)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	chat, err := handler.CreateChatUseCase.Execute(newChatRequest.ToNewChatInput(userId))

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response.NewChatResponse(chat))
}