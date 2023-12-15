package middleware

import (
	"fmt"
	"net/http"
	logging "server/internal/logging"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PrometheusMiddleware interface {
	// WrapHandler wraps the given HTTP handler for instrumentation.
	WrapHandler(name string, next http.Handler) http.HandlerFunc
}

type middleware struct {
	logger   logging.ILogger
	buckets  []float64
	registry prometheus.Registerer
}

// WrapHandler wraps the given HTTP handler for instrumentation:
// It registers four metric collectors (if not already done) and reports HTTP
// metrics to the (newly or already) registered collectors.
// Each has a constant label named "handler" with the provided handlerName as
// value.
func (m *middleware) WrapHandler(name string, next http.Handler) http.HandlerFunc {
	reg := prometheus.WrapRegistererWith(prometheus.Labels{"handler": name}, m.registry)
	m.logger.DebugFmt("Set prometheus handler label "+name, "", "WrapHandler", "middleware")

	requestsTotal := promauto.With(reg).NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Tracks the number of HTTP requests.",
		}, []string{"method", "code"},
	)
	requestDuration := promauto.With(reg).NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Tracks the latencies for HTTP requests.",
			Buckets: m.buckets,
		},
		[]string{"method", "code"},
	)
	requestSize := promauto.With(reg).NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "http_request_size_bytes",
			Help: "Tracks the size of HTTP requests.",
		},
		[]string{"method", "code"},
	)
	responseSize := promauto.With(reg).NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "http_response_size_bytes",
			Help: "Tracks the size of HTTP responses.",
		},
		[]string{"method", "code"},
	)

	// Wraps the provided http.Handler to observe the request result with the provided metrics.
	base := promhttp.InstrumentHandlerCounter(
		requestsTotal,
		promhttp.InstrumentHandlerDuration(
			requestDuration,
			promhttp.InstrumentHandlerRequestSize(
				requestSize,
				promhttp.InstrumentHandlerResponseSize(
					responseSize,
					http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						next.ServeHTTP(w, r)
					}),
				),
			),
		),
	)

	return base
}

// New returns a Middleware interface.
func NewPromMiddleware(logger logging.ILogger, registry prometheus.Registerer, buckets []float64) PrometheusMiddleware {
	if buckets == nil {
		buckets = prometheus.ExponentialBuckets(0.004, 1.75, 10)
	}

	logger.DebugFmt(fmt.Sprintf("Buckets: %v", buckets), "", "NewPromMiddleware", "middleware")

	return &middleware{
		logger:   logger,
		buckets:  buckets,
		registry: registry,
	}
}
