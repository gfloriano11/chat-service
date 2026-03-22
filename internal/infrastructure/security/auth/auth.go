package auth

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
	Secret string
}

func NewJwtService() JwtService {
	secret := os.Getenv("SECRET")
	if secret == "" {
		panic("SECRET not set")
	}
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

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(service.Secret), nil
	})
}

func (service JwtService) AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			splittedAuthHeader := strings.Split(authHeader, " ")

			if len(splittedAuthHeader) != 2 || splittedAuthHeader[0] != "Bearer" {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			tokenToString := splittedAuthHeader[1]

			token, err := service.Validate(tokenToString)
			if err != nil || !token.Valid {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			claims, isTokenOk := token.Claims.(jwt.MapClaims)

			if !isTokenOk {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			userId := int(claims["userId"].(float64))

			userContext := context.WithValue(r.Context(), "userId", userId)

			next.ServeHTTP(w, r.WithContext(userContext))
		})
	}
}

func GetUserIdFromContext(securityContext context.Context) (int, bool) {
	userId, ok := securityContext.Value("userId").(int)
	return userId, ok
}