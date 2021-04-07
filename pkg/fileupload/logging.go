package fileupload

import (
	"fmt"
	"net/http"
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

// Upload 上传
func (mw LoggingMiddleware) Upload(r *http.Request) (result *FileData, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "Upload",
			"input", fmt.Sprintf("fileName=%s", r.Form.Get("fileName")),
			"result", fmt.Sprintf("%+v", result),
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.Upload(r)
	return
}
