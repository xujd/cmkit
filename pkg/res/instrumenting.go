package res

import (
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
