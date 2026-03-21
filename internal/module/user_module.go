package module

import (
	application "chat-service/internal/application/useCases/user"
	userRepository "chat-service/internal/infrastructure/database/repository"

	"gorm.io/gorm"
)

type UserModule struct {
	CreateUser				application.CreateUserUseCase
	FindUserByEmail 	application.FindUserByEmail
}

func NewUserModule(db *gorm.DB) UserModule {
	repository := userRepository.NewUserRepository(db)
	createUserUseCase := application.NewCreateUserUseCase(repository)
	findUserByEmailUseCase := application.NewFindUserByEmail(repository)

	return UserModule{
		CreateUser: createUserUseCase,
		FindUserByEmail: findUserByEmailUseCase,
	}
}