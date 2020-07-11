package hello

import (
	"time"

	"github.com/go-kit/kit/metrics"
)

// InstrumentingMiddleware Hello
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

// Hello Hello
func (s *InstrumentingMiddleware) Hello(world string) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "hello").Add(1)
		s.requestLatency.With("method", "hello").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Hello(world)
}
