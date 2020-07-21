package auth

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

// AddUser 添加用户
func (s *InstrumentingMiddleware) AddUser(user models.User) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "AddUser").Add(1)
		s.requestLatency.With("method", "AddUser").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.AddUser(user)
}

// UpdateUser 修改用户
func (s *InstrumentingMiddleware) UpdateUser(user models.User) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "UpdateUser").Add(1)
		s.requestLatency.With("method", "UpdateUser").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.UpdateUser(user)
}

// DeleteUser 删除用户
func (s *InstrumentingMiddleware) DeleteUser(id uint) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "DeleteUser").Add(1)
		s.requestLatency.With("method", "DeleteUser").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.DeleteUser(id)
}
