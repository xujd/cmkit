package hello

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

// NewLoggingMiddleware New
func NewLoggingMiddleware(logger log.Logger, s Service) Service {
	return &LoggingMiddleware{logger, s}
}

// Hello Hello
func (mw LoggingMiddleware) Hello(world string) (token string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "Hello",
			"input", fmt.Sprintf("world=%s", world),
			"result", token,
			"took", time.Since(begin),
		)
	}(time.Now())
	token, err = mw.Service.Hello(world)
	return
}
