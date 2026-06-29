package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(720 * 2 * time.Hour).Unix(),
	})
	return token.SignedString([]byte(secret))
}
