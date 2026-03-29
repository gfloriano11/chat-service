package module

import (
	application "chat-service/internal/application/useCases/user"
	userRepository "chat-service/internal/infrastructure/database/repository"
	"chat-service/internal/infrastructure/security"
	"chat-service/internal/infrastructure/security/auth"

	"gorm.io/gorm"
)

type UserModule struct {
	CreateUser								application.CreateUserUseCase
	Login 										application.Login
	GetMe 										application.GetMe
	FindUsersNotInChatWithMe 	application.FindUsersNotInChatWithMe
}

func NewUserModule(db *gorm.DB) UserModule {
	repository := userRepository.NewUserRepository(db)
	
	createUserUseCase := application.NewCreateUserUseCase(
		repository,
		auth.NewJwtService(), 
		security.NewEmailService(), 
		security.NewPasswordService(),
	)

	loginUsecase := application.NewLoginUseCase(
		repository, 
		auth.NewJwtService(), 
		security.NewEmailService(), 
		security.NewPasswordService(),
	)

	getMeUseCase := application.NewGetMeUseCase(repository)

	findUsersNotInChatWithMeUseCase := application.NewFindUsersNotInChatWithMeUseCase(repository)

	return UserModule{
		CreateUser: createUserUseCase,
		Login: loginUsecase,
		GetMe: getMeUseCase,
		FindUsersNotInChatWithMe: findUsersNotInChatWithMeUseCase,
	}
}