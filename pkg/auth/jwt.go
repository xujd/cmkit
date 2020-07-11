package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//secret key
var secretKey = []byte("abcd1234!@#$")

// AuthClaims 自定义声明
type AuthClaims struct {
	UserId string `json:"userId"`
	Name   string `json:"name"`

	jwt.StandardClaims
}

// jwtKeyFunc 返回密钥
func JwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return secretKey, nil
}

// Sign 生成token
func Sign(name, uid string) (string, error) {
	// 过期时间
	expAt := time.Now().Add(time.Duration(2) * time.Minute).Unix()

	claims := AuthClaims{
		UserId: uid,
		Name:   name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expAt,
			Issuer:    "system",
		},
	}

	//创建token，指定加密算法为HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//生成token
	tokenString, err := token.SignedString(secretKey)
	return fmt.Sprintf("Bearer %s", tokenString), err
}

// Resign 续订token
func Resign(oldToken string) (string, error) {
	token, err := jwt.ParseWithClaims(strings.Replace(oldToken, "Bearer ", "", -1), &AuthClaims{}, JwtKeyFunc)
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return Sign(claims.Name, claims.UserId)
	}
	return "", err
}
