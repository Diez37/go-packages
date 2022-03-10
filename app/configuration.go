package app

import (
	"github.com/diez37/go-packages/configurator"
	"os"
	"runtime"
)

type Option func(config *Config) *Config

// Configuration set maximum number process in goroutine manager pool
func Configuration(config *Config, configurator configurator.Configurator, options ...Option) {
	config.Name = NameDefault

	config.PID = os.Getpid()

	configurator.SetDefault(MaxProcFieldName, MaxProcDefault)
	maxProc := configurator.GetInt(MaxProcFieldName)
	runtime.GOMAXPROCS(maxProc)

	for _, option := range options {
		config = option(config)
	}
}

func WithAppName(name string) Option {
	return func(config *Config) *Config {
		config.Name = name

		return config
	}
}
