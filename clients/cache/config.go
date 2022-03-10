package cache

const (
	TypeFieldName = "cache.type"

	TypeDefault = GoCacheType

	RedisType   = "redis"
	GoCacheType = "app"
)

type Config struct {
	Type     string
	OnlyType string
}

func NewConfig() *Config {
	return &Config{}
}
