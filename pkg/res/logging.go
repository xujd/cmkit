package res

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

// ListSlings 查询吊索具
func (mw LoggingMiddleware) ListSlings(name string, slingType uint, maxTon uint, useStatus uint, inspectStatus uint, pageIndex int, pageSize int) (result *models.SearchResult, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "ListSlings",
			"input", fmt.Sprintf("name=%s, slingType=%d, maxTon=%d, useStatus=%d, inspectStatus=%d, pageIndex=%d, pageSize=%d", name, slingType, maxTon, useStatus, inspectStatus, pageIndex, pageSize),
			"result", fmt.Sprintf("%+v", result),
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.ListSlings(name, slingType, maxTon, useStatus, inspectStatus, pageIndex, pageSize)
	return
}

// AddSling 添加吊索具
func (mw LoggingMiddleware) AddSling(sling Sling) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "AddSling",
			"input", fmt.Sprintf("sling=%+v", sling),
			"result", fmt.Sprintf("%+v", result),
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.AddSling(sling)
	return
}

// UpdateSling 修改吊索具
func (mw LoggingMiddleware) UpdateSling(sling Sling) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "UpdateSling",
			"input", fmt.Sprintf("sling=%+v", sling),
			"result", fmt.Sprintf("%+v", result),
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.UpdateSling(sling)
	return
}

// DeleteSling 删除吊索具
func (mw LoggingMiddleware) DeleteSling(id uint) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "DeleteSling",
			"input", fmt.Sprintf("id=%d", id),
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.DeleteSling(id)
	return
}

// ListCabinets 查询智能柜
func (mw LoggingMiddleware) ListCabinets(name string, pageIndex int, pageSize int) (result *models.SearchResult, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "ListCabinets",
			"input", fmt.Sprintf("name=%s, pageIndex=%d, pageSize=%d", name, pageIndex, pageSize),
			"result", fmt.Sprintf("%+v", result),
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.ListCabinets(name, pageIndex, pageSize)
	return
}

// AddCabinet 添加智能柜
func (mw LoggingMiddleware) AddCabinet(cabinet Cabinet) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "AddCabinet",
			"input", fmt.Sprintf("cabinet=%+v", cabinet),
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.AddCabinet(cabinet)
	return
}

// UpdateCabinet 修改智能柜
func (mw LoggingMiddleware) UpdateCabinet(cabinet Cabinet) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "UpdateCabinet",
			"input", fmt.Sprintf("cabinet=%+v", cabinet),
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.UpdateCabinet(cabinet)
	return
}

// DeleteCabinet 删除智能柜
func (mw LoggingMiddleware) DeleteCabinet(id uint) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "DeleteCabinet",
			"input", fmt.Sprintf("id=%d", id),
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.DeleteCabinet(id)
	return
}

// ListCabinetGrids 查询箱格状态
func (mw LoggingMiddleware) ListCabinetGrids(cabinetID uint) (result *models.SearchResult, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "ListCabinetGrids",
			"input", fmt.Sprintf("cabinetID=%d", cabinetID),
			"result", fmt.Sprintf("%+v", result),
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.ListCabinetGrids(cabinetID)
	return
}

// Store 存
func (mw LoggingMiddleware) Store(cabinetID uint, gridNo uint, resID uint) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "Store",
			"input", fmt.Sprintf("cabinetID=%d, gridNo=%d, resID=%d", cabinetID, gridNo, resID),
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.Store(cabinetID, gridNo, resID)
	return
}

// TakeReturn 取还
func (mw LoggingMiddleware) TakeReturn(cabinetID uint, gridNo uint, flag int) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "TakeReturn",
			"input", fmt.Sprintf("cabinetID=%d, gridNo=%d, flag=%d", cabinetID, gridNo, flag),
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.TakeReturn(cabinetID, gridNo, flag)
	return
}

// TakeReturnByResID 按资源ID取还
func (mw LoggingMiddleware) TakeReturnByResID(useLog UseLog) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "TakeReturnByResID",
			"input", fmt.Sprintf("useLog=%+v", useLog),
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.TakeReturnByResID(useLog)
	return
}

// SaveTakeReturnLog 保存取还日志
func (mw LoggingMiddleware) SaveTakeReturnLog(useLog UseLog) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "SaveTakeReturnLog",
			"input", fmt.Sprintf("useLog=%+v", useLog),
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.SaveTakeReturnLog(useLog)
	return
}

// GetTakeReturnLog 查询取还日志
func (mw LoggingMiddleware) GetTakeReturnLog(resName string, takeStaff uint, returnStaff uint,
	takeStartTime string, takeEndTime string, returnFlag int,
	pageIndex int, pageSize int) (result *models.SearchResult, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "GetTakeReturnLog",
			"input", fmt.Sprintf("resName=%s, takeStaff=%d, returnStaff=%d, takeStartTime=%s, takeEndTime=%s, returnFlag=%d, pageIndex=%d, pageSize=%d",
				resName, takeStaff, returnStaff, takeStartTime, takeEndTime, returnFlag, pageIndex, pageSize),
			"result", fmt.Sprintf("%+v", result),
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.GetTakeReturnLog(resName, takeStaff, returnStaff, takeStartTime, takeEndTime, returnFlag, pageIndex, pageSize)
	return
}
