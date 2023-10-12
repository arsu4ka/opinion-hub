package utils

import (
	"github.com/aru4ka/opinion-hub/internal/app/configs"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtCustomClaims struct {
	UserId uint `json:"userId"`
	jwt.RegisteredClaims
}

func GenerateJWT(userId uint, conf *configs.JwtConfig) (string, error) {
	claims := &JwtCustomClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(conf.TimeToExpire)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(conf.Secret))
}
