package metrics

import (
	"github.com/diez37/go-packages/app"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Metrics container for common prometheus metrics
type Metrics struct {
	// ErrorsCounter global instance of prometheus.Counter for errors
	ErrorsCounter prometheus.Counter

	// WarningsCounter global instance of prometheus.Counter for warnings
	WarningsCounter prometheus.Counter

	// HttpDuration global instance of prometheus.HistogramVec for duration of HTTP requests
	HttpDuration *prometheus.HistogramVec

	HttpRequestTotal *prometheus.CounterVec

	HttpClientDuration *prometheus.HistogramVec

	HttpClientRequestTotal *prometheus.CounterVec
}

func NewMetrics(appConfig *app.Config) *Metrics {
	return &Metrics{
		ErrorsCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "errors_count",
			Help: "number of errors",
			ConstLabels: map[string]string{
				"app": appConfig.Name,
			},
		}),
		WarningsCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "warnings_count",
			Help: "number of warnings",
			ConstLabels: map[string]string{
				"app": appConfig.Name,
			},
		}),
		HttpDuration: promauto.NewHistogramVec(prometheus.HistogramOpts{
			Name: "http_duration_seconds",
			Help: "Duration of HTTP requests",
			ConstLabels: map[string]string{
				"app": appConfig.Name,
			},
		}, []string{"path"}),
		HttpRequestTotal: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of HTTP requests",
			ConstLabels: map[string]string{
				"app": appConfig.Name,
			},
		}, []string{"code", "handler"}),
		HttpClientDuration: promauto.NewHistogramVec(prometheus.HistogramOpts{
			Name: "http_client_duration_seconds",
			Help: "Duration of HTTP client requests",
			ConstLabels: map[string]string{
				"app": appConfig.Name,
			},
		}, []string{"host"}),
		HttpClientRequestTotal: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "http_client_requests_total",
			Help: "Number of HTTP client requests",
			ConstLabels: map[string]string{
				"app": appConfig.Name,
			},
		}, []string{"code", "handler", "host"}),
	}
}
