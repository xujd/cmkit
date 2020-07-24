package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	kitjwt "github.com/go-kit/kit/auth/jwt"
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
	expAt := time.Now().Add(time.Duration(30) * time.Minute).Unix()

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
	claims, err := ParseToken(oldToken)
	if err != nil {
		return "", err
	}

	return Sign(claims.Name, claims.UserId)
}

// ParseToken 解析token
func ParseToken(token string) (*AuthClaims, error) {
	data, err := jwt.ParseWithClaims(strings.Replace(token, "Bearer ", "", -1), &AuthClaims{}, JwtKeyFunc)
	if err != nil {
		return nil, err
	}
	if claims, ok := data.Claims.(*AuthClaims); ok && data.Valid {
		return claims, nil
	}

	return nil, kitjwt.ErrTokenInvalid
}
