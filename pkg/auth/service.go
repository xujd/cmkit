package auth

import (
	"cmkit/pkg/utils"

	"github.com/jinzhu/gorm"
)

// Service 服务接口
type Service interface {
	AddUser() (string, error)
	Login(name, pwd string) (string, error)
	Renewval(token string) (string, error)
}

// AuthService 权限服务
type AuthService struct {
	DbHandler *gorm.DB
}

// AddUser 添加用户
func (s AuthService) AddUser() (string, error) {
	if !s.DbHandler.HasTable(&User{}) {
		if err := s.DbHandler.CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
	}
	s.DbHandler.Create(&User{Name: "xujd", PassWd: "12435454", Status: 1, Remark: "测试"})
	return "", nil
}

// Login 登录
func (s AuthService) Login(name, pwd string) (string, error) {
	s.AddUser()
	if name == "name" && pwd == "pwd" {
		token, err := Sign(name, pwd)
		return token, err
	}

	return "", utils.ErrUserPwdDismatch
}

// Renewval Token续订
func (s AuthService) Renewval(oldToken string) (string, error) {
	if oldToken != "" {
		token, err := Resign(oldToken)
		return token, err
	}

	return "", utils.ErrBadQueryParams
}
