package http

import "time"

const (
	// InterfaceFieldName field name in configuration file or ENV name for value of HttpConfig.Interface
	InterfaceFieldName = "server.http.interface"

	// PortFieldName field name in configuration file or ENV name for value of HttpConfig.Port
	PortFieldName = "server.http.port"

	ShutdownTimeoutFieldName = "server.http.timeout.shutdown"

	// InterfaceDefault address for listen on default
	InterfaceDefault string = "0.0.0.0"

	// PortDefault port for listen on default
	PortDefault uint = 8080

	ShutdownTimeoutDefault = 30 * time.Second
)

// Config setup params for http web server
type Config struct {
	// Interface address for listen
	Interface string

	// Port port for listen
	Port uint

	ShutdownTimeout time.Duration
}

// NewConfig creating and return new structure instance HttpConfig
func NewConfig() *Config {
	return &Config{}
}
