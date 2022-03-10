package http

import (
	"github.com/diez37/go-packages/clients/http/middlewares"
	"github.com/diez37/go-packages/configurator"
	"github.com/diez37/go-packages/log"
	"github.com/diez37/go-packages/metrics"
	"net/http"
)

func WithConfigurator(configurator configurator.Configurator, config *Config, informer log.Informer, metrics *metrics.Metrics) Doer {
	configurator.SetDefault(TimeoutFieldName, TimeoutDefault)
	configurator.SetDefault(MaxIdleConnsFieldName, MaxIdleConnsDefault)
	configurator.SetDefault(MaxIdleConnsPerHostFieldName, MaxIdleConnsPerHostDefault)
	configurator.SetDefault(MaxConnsPerHostFieldName, MaxConnsPerHostDefault)

	if timeout := configurator.GetDuration(TimeoutFieldName); timeout > 0 && config.Timeout == 0 {
		config.Timeout = timeout
	}

	if maxIdleConns := configurator.GetInt(MaxIdleConnsFieldName); maxIdleConns > 0 && config.MaxIdleConns == 0 {
		config.MaxIdleConns = maxIdleConns
	}

	if maxIdleConnsPerHost := configurator.GetInt(MaxIdleConnsPerHostFieldName); maxIdleConnsPerHost > 0 && config.MaxIdleConnsPerHost == 0 {
		config.MaxIdleConnsPerHost = maxIdleConnsPerHost
	}

	if maxConnsPerHost := configurator.GetInt(MaxConnsPerHostFieldName); maxConnsPerHost > 0 && config.MaxConnsPerHost == 0 {
		config.MaxConnsPerHost = maxConnsPerHost
	}

	return NewHttp(config, informer, metrics)
}

func NewHttp(config *Config, informer log.Informer, metrics *metrics.Metrics) *http.Client {
	informer.Infof("http.client: timeout - %s", config.Timeout)
	informer.Infof("http.client: maxIdleConns - %d", config.MaxIdleConns)
	informer.Infof("http.client: maxConnsPerHost - %d", config.MaxConnsPerHost)
	informer.Infof("http.client: maxIdleConnsPerHost - %d", config.MaxIdleConnsPerHost)

	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxIdleConns = config.MaxIdleConns
	transport.MaxConnsPerHost = config.MaxIdleConnsPerHost
	transport.MaxIdleConnsPerHost = config.MaxIdleConnsPerHost

	return &http.Client{
		Timeout:   config.Timeout,
		Transport: middlewares.NewMetricRoundTripper(transport, metrics, informer),
	}
}
