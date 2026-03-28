package user

import (
	"chat-service/internal/domain/user"
)

type GetMe struct {
	Repository user.UserRepository
}

func NewGetMeUseCase(repository user.UserRepository) GetMe {
	return GetMe{
		Repository: repository,
	}
}

func (useCase GetMe) Execute(id int) (*user.User, error) {
	foundUser, err := useCase.Repository.FindUserById(id)

	if err != nil {
		return nil, ErrSearchUser
	}

	if foundUser == nil {
		return nil, ErrUserNotFound
	}

	return foundUser, nil
} 