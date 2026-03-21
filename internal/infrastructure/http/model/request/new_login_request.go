package request

import inputs "chat-service/internal/application/Inputs"

type NewLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (newLoginRequest NewLoginRequest) ToLoginInput() inputs.NewLoginInput {
	return inputs.NewLoginInput{
		Email: newLoginRequest.Email,
		Password: newLoginRequest.Password,
	}
}