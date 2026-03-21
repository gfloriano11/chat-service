package user

import (
	inputs "chat-service/internal/application/Inputs"
	services "chat-service/internal/infrastructure/security"
	"chat-service/internal/domain/user"
	"errors"
)

type CreateUserUseCase struct {
	Repository user.UserRepository
}

func NewCreateUserUseCase(repository user.UserRepository) CreateUserUseCase {
	return CreateUserUseCase{
		Repository: repository,
	}
}

func (useCase CreateUserUseCase) Execute(createUserInput inputs.CreateUserInput) (user.User, error) {

	if createUserInput.Email == "" || createUserInput.Fullname == "" || createUserInput.Username == "" || createUserInput.Password == "" {
		return user.User{}, errors.New("Any of the fields can be empty")
	}

	passwordService := services.NewPasswordService()
	hashedPassword, err := passwordService.Hash(createUserInput.Password)

	if err != nil {
		return user.User{}, errors.New("Error while trying to validate your password.")
	}

	user := &user.User{
		Email: createUserInput.Email,
		Username: createUserInput.Username,
		Fullname: createUserInput.Fullname,
		Password: hashedPassword,
	}

	useCase.Repository.Save(user)

	return *user, nil
}