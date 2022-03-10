package redis

import (
	"fmt"
	"github.com/diez37/go-packages/configurator"
	"github.com/diez37/go-packages/log"
	"github.com/go-redis/redis/v8"
)

func WithConfigurator(configurator configurator.Configurator, config *Config, informer log.Informer) redis.Cmdable {
	config = Configuration(config, configurator)

	return NewRedis(config, informer)
}

func Configuration(config *Config, configurator configurator.Configurator) *Config {
	configurator.SetDefault(HostFieldName, HostDefault)
	if host := configurator.GetString(HostFieldName); host != "" && config.Host == HostDefault {
		config.Host = host
	}

	configurator.SetDefault(PortFieldName, PortDefault)
	if port := configurator.GetUint(PortFieldName); port > 0 && config.Port == PortDefault {
		config.Port = port
	}

	configurator.SetDefault(ReadTimeoutFieldName, ReadTimeoutDefault)
	if timeout := configurator.GetDuration(ReadTimeoutFieldName); timeout > 0 && config.ReadTimeout == ReadTimeoutDefault {
		config.ReadTimeout = timeout
	}

	configurator.SetDefault(WriteTimeoutFieldName, WriteTimeoutDefault)
	if timeout := configurator.GetDuration(WriteTimeoutFieldName); timeout > 0 && config.WriteTimeout == WriteTimeoutDefault {
		config.WriteTimeout = timeout
	}

	configurator.SetDefault(ConnectTimeoutFieldName, ConnectTimeoutDefault)
	if timeout := configurator.GetDuration(ConnectTimeoutFieldName); timeout > 0 && config.ConnectTimeout == ConnectTimeoutDefault {
		config.ConnectTimeout = timeout
	}

	if user := configurator.GetString(UserFieldName); user != "" {
		config.User = user
	}

	if password := configurator.GetString(PasswordFieldName); password != "" {
		config.Password = password
	}

	return config
}

func NewRedis(config *Config, informer log.Informer) redis.Cmdable {
	informer.Infof("redis: host - %s, port - %d", config.Host, config.Port)
	informer.Infof("redis: readTimeout - %s", config.ReadTimeout)
	informer.Infof("redis: writeTimeout - %s", config.WriteTimeout)
	informer.Infof("redis: connectTimeout - %s", config.ConnectTimeout)

	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Username: config.User,
		Password: config.Password,

		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		PoolTimeout:  config.ConnectTimeout,
	})
}
