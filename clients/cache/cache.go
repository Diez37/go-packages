package cache

import (
	"fmt"
	"github.com/diez37/go-packages/app"
	"github.com/diez37/go-packages/clients/cache/gocache"
	"github.com/diez37/go-packages/clients/cache/redis"
	"github.com/diez37/go-packages/configurator"
	"github.com/diez37/go-packages/log"
	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/metrics"
	"github.com/eko/gocache/v2/store"
	"gopkg.in/errgo.v2/errors"
	"strings"
)

var (
	UnknownCacheTypeError = errors.New("is unknown cache type")
)

func WithConfigurator(
	configurator configurator.Configurator,
	config *Config,
	appConfig *app.Config,
	redisConfig *redis.Config,
	goCacheConfig *gocache.Config,
	informer log.Informer,
) (cache.CacheInterface, error) {
	if _type := configurator.GetString(TypeFieldName); _type != "" && config.Type == TypeDefault {
		config.Type = _type
	}

	if config.OnlyType != "" && config.Type != config.OnlyType {
		return nil, errors.New(fmt.Sprintf("cache: type '%s' not supported", config.Type))
	}

	redisConfig = redis.Configuration(redisConfig, configurator)
	goCacheConfig = gocache.Configuration(goCacheConfig, configurator)

	return NewCache(config, appConfig, redisConfig, goCacheConfig, informer)
}

func NewCache(config *Config, appConfig *app.Config, redisConfig *redis.Config, goCacheConfig *gocache.Config, informer log.Informer) (cache.CacheInterface, error) {
	var cacheStore store.StoreInterface

	switch strings.ToLower(config.Type) {
	case RedisType:
		cacheStore = store.NewRedis(redis.NewRedis(redisConfig, informer), nil)
	case GoCacheType:
		cacheStore = store.NewGoCache(gocache.NewGoCache(goCacheConfig, informer), nil)
	default:
		return nil, UnknownCacheTypeError
	}

	informer.Infof("cache: usage store - %s", config.Type)

	return cache.NewMetric(metrics.NewPrometheus(appConfig.Name), cache.New(cacheStore)), nil
}
