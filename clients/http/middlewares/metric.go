package middlewares

import (
	"github.com/diez37/go-packages/log"
	"github.com/diez37/go-packages/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"strconv"
)

type metricRoundTripper struct {
	originalTransport http.RoundTripper
	metrics           *metrics.Metrics
	informer          log.Informer
}

func NewMetricRoundTripper(originalTransport http.RoundTripper, metrics *metrics.Metrics, informer log.Informer) http.RoundTripper {
	return &metricRoundTripper{originalTransport: originalTransport, metrics: metrics, informer: informer}
}

func (transport *metricRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	timer := prometheus.NewTimer(transport.metrics.HttpClientDuration.WithLabelValues(request.URL.Host))
	defer timer.ObserveDuration()

	transport.informer.Infof("http.client: sending request to %s", request.URL.String())

	response, err := transport.originalTransport.RoundTrip(request)

	transport.metrics.HttpClientRequestTotal.With(prometheus.Labels{
		"code":    strconv.Itoa(response.StatusCode),
		"host":    request.URL.Host,
		"handler": request.URL.RequestURI(),
	}).Inc()

	return response, err
}
