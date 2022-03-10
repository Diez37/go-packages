package redis

import "time"

const (
	HostFieldName           = "cache.redis.host"
	PortFieldName           = "cache.redis.port"
	UserFieldName           = "cache.redis.auth.user"
	PasswordFieldName       = "cache.redis.auth.password"
	ReadTimeoutFieldName    = "cache.redis.timeout.read"
	WriteTimeoutFieldName   = "cache.redis.timeout.write"
	ConnectTimeoutFieldName = "cache.redis.timeout.connect"

	HostDefault           = "localhost"
	PortDefault           = uint(6379)
	ReadTimeoutDefault    = 300 * time.Millisecond
	WriteTimeoutDefault   = 300 * time.Millisecond
	ConnectTimeoutDefault = 500 * time.Millisecond
)

type Config struct {
	Host string
	Port uint

	User     string
	Password string

	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	ConnectTimeout time.Duration
}

func NewConfig() *Config {
	return &Config{}
}
