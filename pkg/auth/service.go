package auth

import "errors"

type Service interface {
	Login(name, pwd string) (string, error)
}
type AuthService struct {
}

func (s AuthService) Login(name, pwd string) (string, error) {
	if name == "name" && pwd == "pwd" {
		token, err := Sign(name, pwd)
		return token, err
	}

	return "", errors.New("Your name or password dismatch")
}
