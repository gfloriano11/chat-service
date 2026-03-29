package handlers

import (
	application "chat-service/internal/application/useCases/user"
	"chat-service/internal/infrastructure/http/model/request"
	response "chat-service/internal/infrastructure/http/model/response/user"
	"chat-service/internal/infrastructure/security/auth"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	CreateUserUseCase								application.CreateUserUseCase
	LoginUseCase 										application.Login
	GetMeUseCase 										application.GetMe
	FindUsersNotInChatWithMeUseCase application.FindUsersNotInChatWithMe
}

func NewUserHandler(
	createUser 								application.CreateUserUseCase, 
	login 										application.Login,
	getMe 										application.GetMe,
	findUsersNotInChatWithMe 	application.FindUsersNotInChatWithMe,
) UserHandler {
	return UserHandler{
		CreateUserUseCase: createUser,
		LoginUseCase: login,
		GetMeUseCase: getMe,
		FindUsersNotInChatWithMeUseCase: findUsersNotInChatWithMe,
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

	http.SetCookie(w, handler.CreateUserUseCase.JwtService.NewAuthCookie(output.Token))
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

	http.SetCookie(w, handler.LoginUseCase.JwtService.NewAuthCookie(token))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}

func (handler UserHandler) GetMe(w http.ResponseWriter, r *http.Request) {
	id, ok := auth.GetUserIdFromContext(r.Context())
	
	if !ok {
		http.Error(w, "User not found", http.StatusForbidden)
		return
	}

	user, err := handler.GetMeUseCase.Execute(id)

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// http.SetCookie(w, handler.LoginUseCase.JwtService.NewAuthCookie(token))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.NewUserResponse(*user))
}

func (handler UserHandler) FindUsersNotInChatWithMe(w http.ResponseWriter, r *http.Request) {
	id, ok := auth.GetUserIdFromContext(r.Context())

	if !ok {
		http.Error(w, "User not found", http.StatusForbidden)
		return
	}

	users, err := handler.FindUsersNotInChatWithMeUseCase.Execute(id)

	if err != nil {
		http.Error(w, "Users not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.NewUsersResponse(*users))
}