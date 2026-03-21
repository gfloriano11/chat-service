package security

import "golang.org/x/crypto/bcrypt"

type PasswordService struct {}

func NewPasswordService() PasswordService {
	return PasswordService{}
}

func (pass PasswordService) Hash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

func (pass PasswordService) Check(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}