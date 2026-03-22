package user

import (
	inputs "chat-service/internal/application/Inputs"
	"chat-service/internal/domain/user"
	"chat-service/internal/infrastructure/security"
	"chat-service/internal/infrastructure/security/auth"
	"errors"
)

type Login struct {
	Repository user.UserRepository
	JwtService auth.JwtService
	EmailService security.EmailService
	PasswordService security.PasswordService
}

func NewLoginUseCase(
	repository user.UserRepository,
	jwt auth.JwtService,
	email security.EmailService,
	password security.PasswordService,
) Login {
	return Login{
		Repository: repository,
		JwtService: jwt,
		EmailService: email,
		PasswordService: password,
	} 
}

func (useCase Login) Execute(input inputs.NewLoginInput) (string, error) {

	if input.Email == "" || input.Password == "" {
		return "", errors.New("E-mail or password can't be empty!")
	}

	isEmailValid := useCase.EmailService.IsValid(input.Email)

	if !isEmailValid {
		return "", errors.New("Invalid e-mail.")
	}

	foundUser, err := useCase.Repository.FindUserByEmail(input.Email)
	
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	
	isPasswordValid := useCase.PasswordService.Check(input.Password, foundUser.Password)

	if !isPasswordValid {
		return "", errors.New("E-mail or password is wrong.")
	}

	token, err := useCase.JwtService.Generate(foundUser.Id)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	return token, nil
}