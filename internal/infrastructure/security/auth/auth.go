package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
	Secret string
}

func NewJwtService(secret string) JwtService {
	return JwtService{
		Secret: secret,
	}
}

func (service JwtService) Generate(userId int) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"exp": time.Now().UTC().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(service.Secret))
}

func (service JwtService) Validate(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (any, error) {
		return []byte(service.Secret), nil
	})
}