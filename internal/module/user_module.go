package module

import (
	application "chat-service/internal/application/useCases/user"
	userRepository "chat-service/internal/infrastructure/database/repository"
	"chat-service/internal/infrastructure/security"
	"chat-service/internal/infrastructure/security/auth"

	"gorm.io/gorm"
)

type UserModule struct {
	CreateUser	application.CreateUserUseCase
	Login 			application.Login
}

func NewUserModule(db *gorm.DB) UserModule {
	repository := userRepository.NewUserRepository(db)
	createUserUseCase := application.NewCreateUserUseCase(repository)
	loginUsecase := application.NewLoginUseCase(
		repository, 
		auth.NewJwtService(), 
		security.NewEmailService(), 
		security.NewPasswordService(),
	)

	return UserModule{
		CreateUser: createUserUseCase,
		Login: loginUsecase,
	}
}