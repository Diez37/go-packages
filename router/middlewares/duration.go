package middlewares

import (
	"github.com/diez37/go-packages/log"
	"github.com/diez37/go-packages/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
)

// HttpDurationMiddleware return MiddlewareHandler for add information to metric on request time
func HttpDurationMiddleware(metrics *metrics.Metrics, informer log.Informer) MiddlewareHandler {
	informer.Info("http.middleware: add http duration")

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			timer := prometheus.NewTimer(metrics.HttpDuration.WithLabelValues(r.URL.Path))
			defer timer.ObserveDuration()

			next.ServeHTTP(w, r)
		})
	}
}
