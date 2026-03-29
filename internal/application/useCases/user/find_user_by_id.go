package user

import "chat-service/internal/domain/user"

type FindUserById struct {
	Repository user.UserRepository
}

func NewFindUserByIdUseCase(repository user.UserRepository) FindUserById {
	return FindUserById{
		Repository: repository,
	}
}

func (useCase FindUserById) Execute(id int) (*user.User, error) {
	foundUser, err := useCase.Repository.FindUserById(id)

	if err != nil {
		return nil, ErrSearchUser
	}

	if foundUser == nil {
		return nil, ErrUserNotFound
	}

	return foundUser, nil
}