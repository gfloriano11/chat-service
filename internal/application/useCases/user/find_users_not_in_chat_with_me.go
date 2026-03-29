package user

import "chat-service/internal/domain/user"

type FindUsersNotInChatWithMe struct {
	Repository user.UserRepository
}

func NewFindUsersNotInChatWithMeUseCase(repository user.UserRepository) FindUsersNotInChatWithMe {
	return FindUsersNotInChatWithMe{
		Repository: repository,
	}
}

func (useCase FindUsersNotInChatWithMe) Execute(id int) (*[]user.User, error) {
	foundUsers, err := useCase.Repository.FindUsersNotInChatsWithMeByUserId(id)

	if err != nil {
		return nil, ErrSearchUser
	}

	if foundUsers == nil {
		return nil, ErrUserNotFound
	}

	return foundUsers, nil
}