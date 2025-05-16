package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/yx1126/go-admin/config"
)

type CustomClaims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenToken(id int, username string) (string, error) {
	claims := CustomClaims{
		Id:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "yx1126",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(config.TokenConf.Secret)
}

func ParseToken(value string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(value, &CustomClaims{}, func(_ *jwt.Token) (interface{}, error) {
		return config.TokenConf.Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
