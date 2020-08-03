package sys

import (
	"cmkit/pkg/models"
	"time"

	"github.com/go-kit/kit/metrics"
)

// InstrumentingMiddleware Instrumenting
type InstrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

// NewInstrumentingMiddleware New
func NewInstrumentingMiddleware(counter metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &InstrumentingMiddleware{
		requestCount:   counter,
		requestLatency: latency,
		Service:        s,
	}
}

// ListCompanys 查询公司
func (s InstrumentingMiddleware) ListCompanys(name string, pageIndex int, pageSize int) (*models.SearchResult, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "ListCompanys").Add(1)
		s.requestLatency.With("method", "ListCompanys").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.ListCompanys(name, pageIndex, pageSize)
}

// ListDepartments 查询部门
func (s InstrumentingMiddleware) ListDepartments(name string, companyID uint, pageIndex int, pageSize int) (*models.SearchResult, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "ListDepartments").Add(1)
		s.requestLatency.With("method", "ListDepartments").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.ListDepartments(name, companyID, pageIndex, pageSize)
}

// AddStaff 添加员工
func (s InstrumentingMiddleware) AddStaff(staff models.Staff) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "AddStaff").Add(1)
		s.requestLatency.With("method", "AddStaff").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.AddStaff(staff)
}

// UpdateStaff 修改员工
func (s InstrumentingMiddleware) UpdateStaff(staff models.Staff) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "UpdateStaff").Add(1)
		s.requestLatency.With("method", "UpdateStaff").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.UpdateStaff(staff)
}

// DeleteStaff 删除员工
func (s InstrumentingMiddleware) DeleteStaff(id uint) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "DeleteStaff").Add(1)
		s.requestLatency.With("method", "DeleteStaff").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.DeleteStaff(id)
}

// ListStaffs 查询员工列表
func (s InstrumentingMiddleware) ListStaffs(name string, companyID uint, departmentID uint, pageIndex int, pageSize int) (*models.SearchResult, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "ListStaffs").Add(1)
		s.requestLatency.With("method", "ListStaffs").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.ListStaffs(name, companyID, departmentID, pageIndex, pageSize)
}

// ListDict 获取字典列表
func (s InstrumentingMiddleware) ListDict(scene string, dictType string) (*[]models.DictData, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "ListDict").Add(1)
		s.requestLatency.With("method", "ListDict").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.ListDict(scene, dictType)
}
