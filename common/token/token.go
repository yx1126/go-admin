package token

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/yx1126/go-admin/DB"
	"github.com/yx1126/go-admin/common/redis"
	"github.com/yx1126/go-admin/config"
)

type TokenClaims struct {
	Uuid     string `json:"uuid"`
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 生成token
func GenToken(id int, username string) (string, error) {
	claims := TokenClaims{
		Uuid:     uuid.New().String(),
		Id:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "yx1126",
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.Token.Secret))
	if err != nil {
		return "", err
	}
	err = DB.Redis.Set(redis.UserTokenKey+claims.Uuid, token, time.Minute*time.Duration(config.Token.ExpireTime)).Err()
	if err != nil {
		return "", err
	}
	return token, nil
}

// 解析token
func ParseToken(value string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(value, &TokenClaims{}, func(_ *jwt.Token) (interface{}, error) {
		return []byte(config.Token.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// 刷新token
func RefreshToken(token string) error {
	claims, err := ParseToken(token)
	if err != nil {
		return err
	}
	return DB.Redis.Set(redis.UserTokenKey+claims.Uuid, token, time.Minute*time.Duration(config.Token.ExpireTime)).Err()
}

// 删除token
func DeleteToken(token string) error {
	claims, err := ParseToken(token)
	if err != nil {
		return err
	}
	return DB.Redis.Del(redis.UserTokenKey + claims.Uuid).Err()
}

// 解析请求头
func ParseHeaderToken(authorization string) (string, error) {
	if authorization == "" {
		return "", errors.New("请先登录")
	}
	tokenSplit := strings.Split(authorization, " ")
	if len(tokenSplit) != 2 || tokenSplit[0] != "Bearer" {
		return "", errors.New("token格式错误")
	}
	return tokenSplit[1], nil
}
