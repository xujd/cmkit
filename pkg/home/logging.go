package home

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

// StatAllRes 统计资源数
func (mw LoggingMiddleware) StatAllRes() (result *[]map[string]interface{}, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "StatAllRes",
			"input", "无",
			"result", fmt.Sprintf("%+v", result),
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.StatAllRes()
	return
}

// StatSlingByTon 按吨位统计吊索具
func (mw LoggingMiddleware) StatSlingByTon() (result *[]map[string]interface{}, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "StatSlingByTon",
			"input", "无",
			"result", fmt.Sprintf("%+v", result),
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.StatSlingByTon()
	return
}

// GetSlingUsedTop 取使用次数最多的top10吊索具
func (mw LoggingMiddleware) GetSlingUsedTop() (result *[]map[string]interface{}, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "GetSlingUsedTop",
			"input", "无",
			"result", fmt.Sprintf("%+v", result),
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.GetSlingUsedTop()
	return
}

// StatSlingByStatus 获取状态统计
func (mw LoggingMiddleware) StatSlingByStatus() (result *[]map[string]interface{}, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "StatSlingByStatus",
			"input", "无",
			"result", fmt.Sprintf("%+v", result),
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.StatSlingByStatus()
	return
}
