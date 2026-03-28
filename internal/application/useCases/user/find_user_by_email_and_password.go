package user

import (
	"chat-service/internal/domain/user"
	"chat-service/internal/infrastructure/security"
	"errors"
)

type FindUserByEmail struct {
	Repository user.UserRepository
}

func NewFindUserByEmail(repository user.UserRepository) FindUserByEmail {
	return FindUserByEmail{
		Repository: repository,
	}
}

func (useCase FindUserByEmail) Execute(email string) (user.User, error) {

	if email == "" {
		return user.User{}, errors.New("E-mail can't be empty!")
	}

	emailService := security.NewEmailService()
	isEmailValid := emailService.IsValid(email)

	if !isEmailValid {
		return user.User{}, errors.New("Invalid e-mail.")
	}
	
	foundUser, err := useCase.Repository.FindUserByEmail(email)
	
	if err != nil {
		return user.User{}, errors.New("It was an error while serch for users")
	}

	return *foundUser, nil
}

