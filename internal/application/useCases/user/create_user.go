package user

import (
	inputs "chat-service/internal/application/Inputs"
	"chat-service/internal/domain/user"
	"chat-service/internal/infrastructure/database/repository"
)

type CreateUserUseCase struct {
	Repository repository.UserRepository
}

func NewCreateUserUseCase(repository repository.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		Repository: repository,
	}
}

func (useCase CreateUserUseCase) CreateUser(createUserInput inputs.CreateUserInput) (user.User, error) {
	
}