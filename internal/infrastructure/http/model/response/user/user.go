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

func NewUserResponse(user user.User) UserResponse {
	userResponse := UserResponse{
		user.Id,
		user.Username,
		user.Fullname,
		user.Email,
		user.CreatedAt,
	}

	return userResponse
}