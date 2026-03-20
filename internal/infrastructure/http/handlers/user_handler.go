package handlers

import (
	application "chat-service/internal/application/useCases/user"
	response "chat-service/internal/infrastructure/http/model/response/user"
	"chat-service/internal/infrastructure/http/model/request"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	CreateUserUseCase application.CreateUserUseCase
}

func NewUserHandler(createUser application.CreateUserUseCase) UserHandler {
	return UserHandler{
		CreateUserUseCase: createUser,
	}
}

func (handler UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUserRequest request.CreateUserRequest

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&newUserRequest)

	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	user, err := handler.CreateUserUseCase.Execute(newUserRequest.ToCreateUserInput())

	if err != nil {
		http.Error(w, "Error while trying to create user!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response.NewUserResponse(user))
}