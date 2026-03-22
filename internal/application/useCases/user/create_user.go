package user

import (
	inputs "chat-service/internal/application/Inputs"
	"chat-service/internal/domain/user"
	"chat-service/internal/infrastructure/security"
	"chat-service/internal/infrastructure/security/auth"
)

type CreateUserUseCase struct {
	Repository user.UserRepository
	JwtService auth.JwtService
	EmailService security.EmailService
	PasswordService security.PasswordService
}

type CreateUserOutput struct {
	User 	user.User
	Token string
}

func NewCreateUserUseCase(
	repository user.UserRepository,
	jwt auth.JwtService,
	email security.EmailService,
	password security.PasswordService,	
) CreateUserUseCase {
	return CreateUserUseCase{
		Repository: repository,
		JwtService: jwt,
		EmailService: email,
		PasswordService: password,
	}
}

func (useCase CreateUserUseCase) Execute(createUserInput inputs.CreateUserInput) (CreateUserOutput, error) {

	if createUserInput.Email == "" || createUserInput.Fullname == "" || createUserInput.Username == "" || createUserInput.Password == "" {
		return CreateUserOutput{}, ErrEmptyFields
	}

	isEmailValid := useCase.EmailService.IsValid(createUserInput.Email)
	
	if !isEmailValid {
		return CreateUserOutput{}, ErrInvalidEmail
	}

	hashedPassword, err := useCase.PasswordService.Hash(createUserInput.Password)

	if err != nil {
		return CreateUserOutput{}, ErrGeneric
	}

	userToCreate := &user.User{
		Email: createUserInput.Email,
		Username: createUserInput.Username,
		Fullname: createUserInput.Fullname,
		Password: hashedPassword,
	}

	createdUser, err := useCase.Repository.Save(userToCreate)

	if err != nil {

	}

	token, err := useCase.JwtService.Generate(createdUser.Id)

	if err != nil {
		return CreateUserOutput{}, ErrGeneric
	}

	output := CreateUserOutput{
		User: createdUser,
		Token: token,
	}

	return output, nil
}