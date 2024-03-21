package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const secretKey = "secret"

type CustomClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateJwtToken(userId string, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		ID:    userId,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    userId,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	})
	accessToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func ParseToken(value string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(value, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("unknown claims type, cannot proceed")
	}
	return claims, err
}
