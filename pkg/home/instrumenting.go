package home

import (
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

// StatAllRes 统计资源数
func (s InstrumentingMiddleware) StatAllRes() (*[]map[string]interface{}, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "StatAllRes").Add(1)
		s.requestLatency.With("method", "StatAllRes").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.StatAllRes()
}

// StatSlingByTon 按吨位统计吊索具
func (s InstrumentingMiddleware) StatSlingByTon() (*[]map[string]interface{}, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "StatSlingByTon").Add(1)
		s.requestLatency.With("method", "StatSlingByTon").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.StatSlingByTon()
}

// GetSlingUsedTop 取使用次数最多的top10吊索具
func (s InstrumentingMiddleware) GetSlingUsedTop() (*[]map[string]interface{}, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "GetSlingUsedTop").Add(1)
		s.requestLatency.With("method", "GetSlingUsedTop").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.GetSlingUsedTop()
}

// StatSlingByStatus 获取状态统计
func (s InstrumentingMiddleware) StatSlingByStatus() (*[]map[string]interface{}, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "StatSlingByStatus").Add(1)
		s.requestLatency.With("method", "StatSlingByStatus").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.StatSlingByStatus()
}
