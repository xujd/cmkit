package sys

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

// ListCompanys 查询公司
func (mw LoggingMiddleware) ListCompanys(name string, pageIndex int, pageSize int) (result *models.SearchResult, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "ListCompanys",
			"input", fmt.Sprintf("name=%s, pageIndex=%d, pageSize=%d", name, pageIndex, pageSize),
			"result", fmt.Sprintf("%+v", result),
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.ListCompanys(name, pageIndex, pageSize)
	return
}

// ListDepartments 查询部门
func (mw LoggingMiddleware) ListDepartments(name string, companyID uint, pageIndex int, pageSize int) (result *models.SearchResult, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "ListDepartments",
			"input", fmt.Sprintf("name=%s, companyID=%d, pageIndex=%d, pageSize=%d", name, companyID, pageIndex, pageSize),
			"result", fmt.Sprintf("%+v", result),
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.ListDepartments(name, companyID, pageIndex, pageSize)
	return
}

// AddStaff 添加员工
func (mw LoggingMiddleware) AddStaff(staff models.Staff) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "AddStaff",
			"input", fmt.Sprintf("staff=%+v", staff),
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.AddStaff(staff)
	return
}

// UpdateStaff 修改员工
func (mw LoggingMiddleware) UpdateStaff(staff models.Staff) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "UpdateStaff",
			"input", fmt.Sprintf("staff=%+v", staff),
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.UpdateStaff(staff)
	return
}

// DeleteStaff 删除员工
func (mw LoggingMiddleware) DeleteStaff(id uint) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "DeleteStaff",
			"input", fmt.Sprintf("id=%d", id),
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.DeleteStaff(id)
	return
}

// ListStaffs 查询员工列表
func (mw LoggingMiddleware) ListStaffs(name string, companyID uint, departmentID uint, pageIndex int, pageSize int) (result *models.SearchResult, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "ListStaffs",
			"input", fmt.Sprintf("name=%s, companyID=%d, departmentID=%d, pageIndex=%d, pageSize=%d", name, companyID, departmentID, pageIndex, pageSize),
			"result", fmt.Sprintf("%+v", result),
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.ListStaffs(name, companyID, departmentID, pageIndex, pageSize)
	return
}

// ListDict 获取字典列表
func (mw LoggingMiddleware) ListDict(scene string, dictType string) (result *[]models.DictData, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "ListDict",
			"input", fmt.Sprintf("scene=%s, dictType=%s", scene, dictType),
			"result", fmt.Sprintf("%+v", result),
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.ListDict(scene, dictType)
	return
}
