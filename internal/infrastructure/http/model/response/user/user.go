package user

import (
	"chat-service/internal/domain/user"
	"time"
)

type UserResponse struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Fullname  string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserTokenResponse struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	Token 		string		`json:"token"`
}

func NewUserResponse(user user.User) UserResponse {
	return UserResponse{
		user.Id,
		user.Username,
		user.Fullname,
		user.Email,
		user.CreatedAt,
	}
}

func NewUserTokenResponse(user user.User, token string) UserTokenResponse {
	return UserTokenResponse{
		user.Id,
		user.Username,
		user.Fullname,
		user.Email,
		user.CreatedAt,
		token,
	}
}

func NewUsersResponse(users []user.User) []UserResponse {
	var usersResponse []UserResponse
	for _, user := range users {
		usersResponse = append(usersResponse, NewUserResponse(user))	
	}

	return usersResponse
}