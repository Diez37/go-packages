package sqlite

import (
	"database/sql"
	"github.com/diez37/go-packages/configurator"
	"github.com/diez37/go-packages/log"
	"github.com/doug-martin/goqu/v9"
	_ "modernc.org/sqlite"
)

func WithConfigurator(config *Config, configurator configurator.Configurator, informer log.Informer) (goqu.SQLDatabase, error) {
	config = Configuration(config, configurator)

	return NewSQLite(config, informer)
}

func Configuration(config *Config, configurator configurator.Configurator) *Config {
	configurator.SetDefault(DsnFieldName, DsnDefault)
	if dsn := configurator.GetString(DsnFieldName); dsn != "" && config.Dsn == "" {
		config.Dsn = dsn
	}

	return config
}

func NewSQLite(config *Config, informer log.Informer) (goqu.SQLDatabase, error) {
	informer.Infof("sqlite: dsn - %s", config.Dsn)

	return sql.Open("sqlite", config.Dsn)
}
