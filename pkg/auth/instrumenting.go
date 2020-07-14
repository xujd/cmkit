package auth

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

// Login 登录
func (s *InstrumentingMiddleware) Login(name, pwd string) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "login").Add(1)
		s.requestLatency.With("method", "login").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Login(name, pwd)
}

// Renewval 续订
func (s *InstrumentingMiddleware) Renewval(oldToken string) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Renewval").Add(1)
		s.requestLatency.With("method", "Renewval").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Renewval(oldToken)
}
