package auth

import (
	"cmkit/pkg/utils"
)

// Service 服务接口
type Service interface {
	Login(name, pwd string) (string, error)
	Renew(token string) (string, error)
}

// AuthService 权限服务
type AuthService struct {
}

// Login 登录
func (s AuthService) Login(name, pwd string) (string, error) {
	if name == "name" && pwd == "pwd" {
		token, err := Sign(name, pwd)
		return token, err
	}

	return "", utils.ErrUserPwdDismatch
}

// Renew Token续订
func (s AuthService) Renew(oldToken string) (string, error) {
	if oldToken != "" {
		token, err := Resign(oldToken)
		return token, err
	}

	return "", utils.ErrBadQueryParams
}
