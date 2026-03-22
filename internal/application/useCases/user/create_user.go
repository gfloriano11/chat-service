package user

import (
	inputs "chat-service/internal/application/Inputs"
	"chat-service/internal/domain/user"
	"chat-service/internal/infrastructure/security"
	"chat-service/internal/infrastructure/security/auth"
	"errors"
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
		return CreateUserOutput{}, errors.New("Any of the fields can be empty")
	}

	isEmailValid := useCase.EmailService.IsValid(createUserInput.Email)
	
	if !isEmailValid {
		return CreateUserOutput{}, errors.New("Invalid e-mail")
	}

	hashedPassword, err := useCase.PasswordService.Hash(createUserInput.Password)

	if err != nil {
		return CreateUserOutput{}, errors.New("Error while trying to validate your password.")
	}

	createdUser := &user.User{
		Email: createUserInput.Email,
		Username: createUserInput.Username,
		Fullname: createUserInput.Fullname,
		Password: hashedPassword,
	}

	useCase.Repository.Save(createdUser)

	token, err := useCase.JwtService.Generate(createdUser.Id)

	if err != nil {
		return CreateUserOutput{}, errors.New("We can't create your account. Please, try again.")
	}

	output := CreateUserOutput{
		User: *createdUser,
		Token: token,
	}

	return output, nil
}