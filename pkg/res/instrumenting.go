package res

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

// ListSlings 查询吊索具
func (s InstrumentingMiddleware) ListSlings(name string, slingType uint, maxTon uint, useStatus uint, inspectStatus uint, pageIndex int, pageSize int) (*models.SearchResult, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "ListSlings").Add(1)
		s.requestLatency.With("method", "ListSlings").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.ListSlings(name, slingType, maxTon, useStatus, inspectStatus, pageIndex, pageSize)
}

// AddSling 添加吊索具
func (s InstrumentingMiddleware) AddSling(sling Sling) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "AddSling").Add(1)
		s.requestLatency.With("method", "AddSling").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.AddSling(sling)
}

// UpdateSling 修改吊索具
func (s InstrumentingMiddleware) UpdateSling(sling Sling) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "UpdateSling").Add(1)
		s.requestLatency.With("method", "UpdateSling").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.UpdateSling(sling)
}

// DeleteSling 删除吊索具
func (s InstrumentingMiddleware) DeleteSling(id uint) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "DeleteSling").Add(1)
		s.requestLatency.With("method", "DeleteSling").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.DeleteSling(id)
}

// ListCabinets 查询智能柜
func (s InstrumentingMiddleware) ListCabinets(name string, pageIndex int, pageSize int) (*models.SearchResult, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "ListCabinets").Add(1)
		s.requestLatency.With("method", "ListCabinets").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.ListCabinets(name, pageIndex, pageSize)
}

// AddCabinet 添加智能柜
func (s InstrumentingMiddleware) AddCabinet(cabinet Cabinet) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "AddCabinet").Add(1)
		s.requestLatency.With("method", "AddCabinet").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.AddCabinet(cabinet)
}

// UpdateCabinet 修改智能柜
func (s InstrumentingMiddleware) UpdateCabinet(cabinet Cabinet) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "UpdateCabinet").Add(1)
		s.requestLatency.With("method", "UpdateCabinet").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.UpdateCabinet(cabinet)
}

// DeleteCabinet 删除智能柜
func (s InstrumentingMiddleware) DeleteCabinet(id uint) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "DeleteCabinet").Add(1)
		s.requestLatency.With("method", "DeleteCabinet").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.DeleteCabinet(id)
}

// ListCabinetGrids 查询箱格状态
func (s InstrumentingMiddleware) ListCabinetGrids(cabinetID uint) (*models.SearchResult, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "ListCabinetGrids").Add(1)
		s.requestLatency.With("method", "ListCabinetGrids").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.ListCabinetGrids(cabinetID)
}

// Store 存
func (s InstrumentingMiddleware) Store(cabinetID uint, gridNo uint, resID uint) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Store").Add(1)
		s.requestLatency.With("method", "Store").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Store(cabinetID, gridNo, resID)
}

// TakeReturn 取还
func (s InstrumentingMiddleware) TakeReturn(cabinetID uint, gridNo uint, flag int) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "TakeReturn").Add(1)
		s.requestLatency.With("method", "TakeReturn").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.TakeReturn(cabinetID, gridNo, flag)
}

// TakeReturnByResID 按资源ID取还
func (s InstrumentingMiddleware) TakeReturnByResID(useLog UseLog) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "TakeReturnByResID").Add(1)
		s.requestLatency.With("method", "TakeReturnByResID").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.TakeReturnByResID(useLog)
}

// SaveTakeReturnLog 保存取还日志
func (s InstrumentingMiddleware) SaveTakeReturnLog(useLog UseLog) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "SaveTakeReturnLog").Add(1)
		s.requestLatency.With("method", "SaveTakeReturnLog").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.SaveTakeReturnLog(useLog)
}

// GetTakeReturnLog 查询取还日志
func (s InstrumentingMiddleware) GetTakeReturnLog(resName string, takeStaff uint, returnStaff uint,
	takeStartTime string, takeEndTime string, returnFlag int,
	pageIndex int, pageSize int) (*models.SearchResult, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "GetTakeReturnLog").Add(1)
		s.requestLatency.With("method", "GetTakeReturnLog").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.GetTakeReturnLog(resName, takeStaff, returnStaff, takeStartTime, takeEndTime, returnFlag, pageIndex, pageSize)
}
