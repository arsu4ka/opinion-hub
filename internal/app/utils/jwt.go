package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtCustomClaims struct {
	UserId uint `json:"userId"`
	jwt.RegisteredClaims
}

func GenerateJWT(secret string, expirationHours int, userId uint) (string, error) {
	claims := &JwtCustomClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expirationHours))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
