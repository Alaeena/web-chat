package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"server/db/database"
	"time"
)

const secretKey = "secret"

type CustomClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CreateJwtToken(user database.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		ID:       string(user.ID),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    string(user.ID),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	})
	accessToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
