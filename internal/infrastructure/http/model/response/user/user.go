package user

import (
	"chat-service/internal/domain/user"
	"time"
)

type UserResponse struct {
	Id        int
	Username  string
	Fullname  string
	Email     string
	CreatedAt time.Time
}

type UserTokenResponse struct {
	Id        int
	Username  string
	Fullname  string
	Email     string
	CreatedAt time.Time
	Token 		string
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