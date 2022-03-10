package http

import "time"

const (
	TimeoutFieldName             = "client.http.timeout"
	MaxIdleConnsFieldName        = "client.http.maximum.idle.connections"
	MaxIdleConnsPerHostFieldName = "client.http.maximum.idle.host"
	MaxConnsPerHostFieldName     = "client.http.maximum.connections"

	TimeoutDefault             = 30 * time.Second
	MaxIdleConnsDefault        = 100
	MaxIdleConnsPerHostDefault = 100
	MaxConnsPerHostDefault     = 100
)

type Config struct {
	Timeout             time.Duration
	MaxIdleConns        int
	MaxIdleConnsPerHost int
	MaxConnsPerHost     int
}

func NewConfig() *Config {
	return &Config{}
}
