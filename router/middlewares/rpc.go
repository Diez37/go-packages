package middlewares

import (
	"github.com/diez37/go-packages/log"
	"github.com/diez37/go-packages/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"strconv"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (writer *loggingResponseWriter) WriteHeader(code int) {
	writer.statusCode = code
	writer.ResponseWriter.WriteHeader(code)
}

func HttpRequestTotalMiddleware(metrics *metrics.Metrics, informer log.Informer) MiddlewareHandler {
	informer.Info("http.middleware: add http request total")

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			loggingResponseWriter := &loggingResponseWriter{
				ResponseWriter: w,
				statusCode:     0,
			}

			next.ServeHTTP(loggingResponseWriter, r)

			metrics.HttpRequestTotal.With(prometheus.Labels{
				"code":    strconv.Itoa(loggingResponseWriter.statusCode),
				"handler": r.URL.Path,
			}).Inc()
		})
	}
}
