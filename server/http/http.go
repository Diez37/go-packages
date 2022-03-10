package http

import (
	"fmt"
	"github.com/diez37/go-packages/configurator"
	"github.com/diez37/go-packages/log"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func WithConfigurator(configurator configurator.Configurator, config *Config, router chi.Router, informer log.Informer) *http.Server {
	configurator.SetDefault(InterfaceFieldName, InterfaceDefault)
	configurator.SetDefault(PortFieldName, PortDefault)
	configurator.SetDefault(ShutdownTimeoutFieldName, ShutdownTimeoutDefault)

	if _interface := configurator.GetString(InterfaceFieldName); _interface != InterfaceDefault && config.Interface == "" {
		config.Interface = _interface
	}

	if port := configurator.GetUint(PortFieldName); port != PortDefault && config.Port == 0 {
		config.Port = port
	}

	if shutdownTimeout := configurator.GetDuration(ShutdownTimeoutFieldName); shutdownTimeout > 0 && config.ShutdownTimeout == 0 {
		config.ShutdownTimeout = shutdownTimeout
	}

	return NewHttp(config, router, informer)
}

func NewHttp(config *Config, router chi.Router, informer log.Informer) *http.Server {
	informer.Infof("http.server: interface - %s, port - %d", config.Interface, config.Port)

	return &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Interface, config.Port),
		Handler: router,
	}
}
