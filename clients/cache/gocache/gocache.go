package gocache

import (
	"github.com/diez37/go-packages/configurator"
	"github.com/diez37/go-packages/log"
	"github.com/patrickmn/go-cache"
)

func WithConfigurator(configurator configurator.Configurator, config *Config, informer log.Informer) *cache.Cache {
	config = Configuration(config, configurator)

	return NewGoCache(config, informer)
}

func Configuration(config *Config, configurator configurator.Configurator) *Config {
	configurator.SetDefault(ExpirationFieldName, ExpirationDefault)
	if expiration := configurator.GetDuration(ExpirationFieldName); expiration > 0 && (config.Expiration == 0 || config.Expiration == ExpirationDefault) {
		config.Expiration = expiration
	}

	configurator.SetDefault(CleanupIntervalFieldName, CleanupIntervalDefault)
	if cleanupInterval := configurator.GetDuration(CleanupIntervalFieldName); cleanupInterval > 0 && (config.CleanupInterval == 0 || config.CleanupInterval == CleanupIntervalDefault) {
		config.CleanupInterval = cleanupInterval
	}

	return config
}

func NewGoCache(config *Config, informer log.Informer) *cache.Cache {
	informer.Infof("go-cache: expiration - %s", config.Expiration)
	informer.Infof("go-cache: cleanupInterval - %s", config.CleanupInterval)

	return cache.New(config.Expiration, config.CleanupInterval)
}
