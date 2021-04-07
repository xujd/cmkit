package fileupload

import (
	"net/http"
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

// Upload 文件上传
func (s InstrumentingMiddleware) Upload(r *http.Request) (*FileData, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Upload").Add(1)
		s.requestLatency.With("method", "Upload").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Upload(r)
}
