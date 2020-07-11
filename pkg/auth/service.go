package auth

import (
	"cmkit/pkg/utils"
)

type Service interface {
	Login(name, pwd string) (string, error)
	Renew(token string) (string, error)
}
type AuthService struct {
}

func (s AuthService) Login(name, pwd string) (string, error) {
	if name == "name" && pwd == "pwd" {
		token, err := Sign(name, pwd)
		return token, err
	}

	return "", utils.ErrUserPwdDismatch
}

func (s AuthService) Renew(oldToken string) (string, error) {
	if oldToken != "" {
		token, err := Resign(oldToken)
		return token, err
	}

	return "", utils.ErrBadQueryParams
}
