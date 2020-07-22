package auth

import (
	"cmkit/pkg/models"
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
)

// LoggingMiddleware Make a new type
// that contains Service interface and logger instance
type LoggingMiddleware struct {
	logger log.Logger
	Service
}

// NewLoggingMiddleware 日志
func NewLoggingMiddleware(logger log.Logger, s Service) Service {
	return &LoggingMiddleware{logger, s}
}

// Login 登录
func (mw LoggingMiddleware) Login(name, pwd string) (token string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "Login",
			"input", fmt.Sprintf("name=%s, password=%s", name, pwd),
			"result", token,
			"took", time.Since(begin),
		)
	}(time.Now())
	token, err = mw.Service.Login(name, pwd)
	return
}

// Renewval 续订
func (mw LoggingMiddleware) Renewval(oldToken string) (token string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "Renewval",
			"input", fmt.Sprintf("oldToken=%s", oldToken),
			"result", token,
			"took", time.Since(begin),
		)
	}(time.Now())
	token, err = mw.Service.Renewval(oldToken)
	return
}

// AddUser 添加用户
func (mw LoggingMiddleware) AddUser(user models.User) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "AddUser",
			"input", fmt.Sprintf("user=%s", user),
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.AddUser(user)
	return
}

// UpdateUser 修改用户
func (mw LoggingMiddleware) UpdateUser(user models.User) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "UpdateUser",
			"input", fmt.Sprintf("user=%s", user),
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.UpdateUser(user)
	return
}

// DeleteUser 删除用户
func (mw LoggingMiddleware) DeleteUser(id uint) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "DeleteUser",
			"input", fmt.Sprintf("id=%d", id),
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.DeleteUser(id)
	return
}

// QueryUserByID 查询用户
func (mw LoggingMiddleware) QueryUserByID(id uint) (result *models.User, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "QueryUserByID",
			"input", fmt.Sprintf("id=%d", id),
			"result", fmt.Sprintf("%+v", result),
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.QueryUserByID(id)
	return
}

// ListUsers 查询用户列表
func (mw LoggingMiddleware) ListUsers(name string, pageIndex int, pageSize int) (result *models.SearchResult, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "ListUsers",
			"input", fmt.Sprintf("name=%s, pageIndex=%d, pageSize=%d", name, pageIndex, pageSize),
			"result", fmt.Sprintf("%+v", result),
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.ListUsers(name, pageIndex, pageSize)
	return
}
