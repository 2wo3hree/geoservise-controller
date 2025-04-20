package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	HTTPRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	HTTPRequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "path"},
	)

	CacheDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "cache_duration_seconds",
			Help:    "Duration of cache operations.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"operation"},
	)

	DBDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "db_duration_seconds",
			Help:    "Duration of database operations.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"operation"},
	)

	ExternalAPIDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "external_api_duration_seconds",
			Help:    "Duration of external API calls.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"operation"},
	)
)

func Init() {
	prometheus.MustRegister(
		HTTPRequestDuration,
		HTTPRequestCount,
		CacheDuration,
		DBDuration,
		ExternalAPIDuration,
	)
}
