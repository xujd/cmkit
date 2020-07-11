package auth

import (
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

// Renew 续订
func (mw LoggingMiddleware) Renew(oldToken string) (token string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "Renew",
			"input", fmt.Sprintf("oldToken=%s", oldToken),
			"result", token,
			"took", time.Since(begin),
		)
	}(time.Now())
	token, err = mw.Service.Renew(oldToken)
	return
}
