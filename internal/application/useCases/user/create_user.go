package user

import (
	inputs "chat-service/internal/application/Inputs"
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

func (useCase CreateUserUseCase) CreateUser(createUserInput inputs.CreateUserInput) (user.User, error) {

	if createUserInput.Email == "" || createUserInput.Fullname == "" || createUserInput.Username == "" || createUserInput.Password == "" {
		return user.User{}, errors.New("Any of the fields can be empty")
	}

	user := &user.User{
		Email: createUserInput.Email,
		Username: createUserInput.Username,
		Fullname: createUserInput.Fullname,
	}

	useCase.Repository.Save(user)

	return *user, nil
}