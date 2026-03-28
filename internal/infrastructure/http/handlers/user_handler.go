package handlers

import (
	application "chat-service/internal/application/useCases/user"
	"chat-service/internal/infrastructure/http/model/request"
	response "chat-service/internal/infrastructure/http/model/response/user"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	CreateUserUseCase		application.CreateUserUseCase
	LoginUseCase 				application.Login
}

func NewUserHandler(createUser application.CreateUserUseCase, login application.Login) UserHandler {
	return UserHandler{
		CreateUserUseCase: createUser,
		LoginUseCase: login,
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

	output, err := handler.CreateUserUseCase.Execute(newUserRequest.ToCreateUserInput())

	if err != nil {
		switch err {

		case application.ErrEmptyFields, application.ErrInvalidEmail:
			http.Error(w, err.Error(), http.StatusBadRequest)

		case application.ErrEmailAlreadyExists:
			http.Error(w, err.Error(), http.StatusConflict)

		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response.NewUserTokenResponse(output.User, output.Token))
}

func (handler UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var newLoginRequest request.NewLoginRequest

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&newLoginRequest)

	if err != nil {
		http.Error(w, "Invalid login request!", http.StatusInternalServerError)
		return
	}

	token, err := handler.LoginUseCase.Execute(newLoginRequest.ToLoginInput())

	if err != nil {
		http.Error(w, "Invalid email or password.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}