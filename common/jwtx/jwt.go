package jwtx

/**
 * @Author elastic·H
 * @Date 2024-09-12
 * @File: jwt.go
 * @Description:
 */

import (
	"github.com/golang-jwt/jwt"
)

func GetToken(secretKey string, iat, seconds, uid int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["uid"] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
