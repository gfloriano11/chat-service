package handlers

import (
	application "chat-service/internal/application/useCases/chat"
	"chat-service/internal/infrastructure/http/model/request"
	response "chat-service/internal/infrastructure/http/model/response/chat"
	"chat-service/internal/infrastructure/security/auth"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ChatHandler struct {
	CreateChatUseCase 				application.CreateChatUseCase
	FindChatsByUserIdUseCase 	application.FindChatsByUserIdUseCase
}

func NewChatHandler(
	createChat 				application.CreateChatUseCase,
	findChatsByUserId application.FindChatsByUserIdUseCase,
) ChatHandler {
	return ChatHandler{
		CreateChatUseCase: createChat,
		FindChatsByUserIdUseCase: findChatsByUserId,
	}
}

func (handler ChatHandler) GetChatsByUserId(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(chi.URLParam(r, "userId"))

	if err != nil {
		http.Error(w, "invalid body", http.StatusInternalServerError)
	}

	chats, err := handler.FindChatsByUserIdUseCase.Execute(userId)

	if err != nil {

		switch err {
			default:
				http.Error(w, "Something went wrong.", http.StatusInternalServerError)
				return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.NewChatListResponse(*chats))
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response.NewChatResponse(chat))
}