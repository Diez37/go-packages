package container

import (
	"github.com/diez37/go-packages/app"
	"github.com/diez37/go-packages/clients/cache"
	"github.com/diez37/go-packages/clients/cache/gocache"
	"github.com/diez37/go-packages/clients/cache/redis"
	"github.com/diez37/go-packages/clients/db"
	"github.com/diez37/go-packages/clients/db/mysql"
	"github.com/diez37/go-packages/clients/db/sqlite"
	httpClient "github.com/diez37/go-packages/clients/http"
	"github.com/diez37/go-packages/closer"
	"github.com/diez37/go-packages/configurator"
	"github.com/diez37/go-packages/log"
	"github.com/diez37/go-packages/metrics"
	"github.com/diez37/go-packages/migrator"
	"github.com/diez37/go-packages/repeater"
	"github.com/diez37/go-packages/router"
	"github.com/diez37/go-packages/server/http"
	"github.com/diez37/go-packages/server/http/helpers"
	"github.com/diez37/go-packages/tracer"
	"go.uber.org/dig"
	"go.uber.org/multierr"
)

// digWrapper wrapper from dig.Container implements Container
type digWrapper struct {
	dig *dig.Container
}

func NewDigWrapper() (Container, error) {
	container := &digWrapper{dig: dig.New()}

	var err error

	err = container.dig.Provide(
		log.WithConfigurator,
		dig.As(
			new(log.Debuger),
			new(log.Informer),
			new(log.Warner),
			new(log.Printer),
			new(log.Logger),
		))
	if err != nil {
		return nil, err
	}

	return container, container.Provides(
		app.NewConfig,
		closer.NewOsSignal,
		configurator.NewViper,
		log.NewConfig,
		metrics.NewMetrics,

		tracer.NewConfig,
		tracer.WithConfigurator,

		// http
		router.WithConfigurator,
		http.NewConfig,
		http.WithConfigurator,
		helpers.NewError,
		httpClient.NewConfig,
		httpClient.WithConfigurator,

		// cache
		gocache.NewConfig,
		gocache.WithConfigurator,
		redis.NewConfig,
		redis.WithConfigurator,
		cache.NewConfig,
		cache.WithConfigurator,

		// data base
		mysql.NewConfig,
		sqlite.NewConfig,
		db.NewConfig,
		db.WithConfigurator,

		migrator.NewConfig,
		migrator.WithConfigurator,

		repeater.New,
	)
}

func (container *digWrapper) Provide(constructor interface{}) error {
	return container.dig.Provide(constructor)
}

func (container *digWrapper) Provides(constructors ...interface{}) error {
	var errs error

	for _, constructor := range constructors {
		if err := container.Provide(constructor); err != nil {
			errs = multierr.Append(errs, err)
		}
	}

	return errs
}

func (container *digWrapper) Invoke(function interface{}) error {
	return container.dig.Invoke(function)
}
