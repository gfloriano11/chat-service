package request

import (
	inputs "chat-service/internal/application/Inputs"
)

type CreateUserRequest struct {
	Email        	string `json:"email"`
	Username   		string `json:"username"`
	Fullname		 	string `json:"fullname"` 
	Password			string `json:"password"`
}

func (createUserRequest CreateUserRequest) ToCreateUserInput() inputs.CreateUserInput {
	return inputs.CreateUserInput{
		Email: createUserRequest.Email,
		Username: createUserRequest.Username,
		Fullname: createUserRequest.Fullname,
		Password: createUserRequest.Password,
	}
}