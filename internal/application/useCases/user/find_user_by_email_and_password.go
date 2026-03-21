package user

import (
	inputs "chat-service/internal/application/Inputs"
	"chat-service/internal/domain/user"
	"chat-service/internal/infrastructure/security"
	"errors"
	"log"
)

type FindUserByEmail struct {
	Repository user.UserRepository
}

func NewFindUserByEmail(repository user.UserRepository) FindUserByEmail {
	return FindUserByEmail{
		Repository: repository,
	}
}

func (useCase FindUserByEmail) Execute(input inputs.NewLoginInput) (user.User, error) {

	if input.Email == "" || input.Password == "" {
		return user.User{}, errors.New("E-mail or password can't be empty!")
	}

	emailService := security.NewEmailService()
	isEmailValid := emailService.IsValid(input.Email)

	if !isEmailValid {
		return user.User{}, errors.New("Invalid e-mail.")
	}
	
	foundUser, err := useCase.Repository.FindUserByEmail(input.Email)
	
	if err != nil {
		return user.User{}, errors.New("It was an error while serch for users")
	}
	
	passwordService := security.NewPasswordService()
	log.Println("user pass: ", input.Password, "foundUserPass: ", foundUser.Password)
	isPasswordValid := passwordService.Check(input.Password, foundUser.Password)

	log.Println("is valid pass? ", isPasswordValid)

	if !isPasswordValid {
		return user.User{}, errors.New("E-mail or password is wrong.")
	}

	return foundUser, nil
}