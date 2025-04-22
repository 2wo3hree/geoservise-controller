package middleware

import (
	"geoservise-jwt/internal/infrastructure/metrics"
	"net/http"
	"time"
)

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start).Seconds()

		method := r.Method
		path := r.URL.Path

		metrics.HTTPRequestDuration.WithLabelValues(method, path).Observe(duration)
		metrics.HTTPRequestCount.WithLabelValues(method, path).Inc()
	})
}
