package services

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func Auth(password string) (string, error) {
	const op = "services.Auth"
	signingKey := []byte(password)
	token := jwt.New(jwt.SigningMethodHS256)
	signedToken, err := token.SignedString(signingKey)
	if err != nil {
		return "", fmt.Errorf("Проблема с генерацией токена. %s:%v", op, err)
	}
	err = os.Setenv("JWT_TOKEN", signedToken)
	if err != nil {
		return "", fmt.Errorf("Проблема с записью токена в переменные окружения %s:%v", op, err)
	}
	return signedToken, nil
}
