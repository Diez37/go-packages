package gocache

import "time"

const (
	ExpirationFieldName      = "cache.app.expiration"
	CleanupIntervalFieldName = "cache.app.cleanup"

	ExpirationDefault      = 5 * time.Minute
	CleanupIntervalDefault = 10 * time.Minute
)

type Config struct {
	Expiration      time.Duration
	CleanupInterval time.Duration
}

func NewConfig() *Config {
	return &Config{}
}
