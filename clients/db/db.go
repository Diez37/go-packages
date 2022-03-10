package db

import (
	"errors"
	"fmt"
	"github.com/diez37/go-packages/clients/db/mysql"
	"github.com/diez37/go-packages/clients/db/sqlite"
	"github.com/diez37/go-packages/configurator"
	"github.com/diez37/go-packages/log"
	"github.com/doug-martin/goqu/v9"
)

func WithConfigurator(
	config *Config,
	configurator configurator.Configurator,
	informer log.Informer,
	mysqlConfig *mysql.Config,
	sqliteConfig *sqlite.Config,
) (goqu.SQLDatabase, error) {
	if driver := configurator.GetString(DriverFieldName); driver != "" && config.Driver == "" {
		config.Driver = driver
	}

	if config.Driver == "" {
		return nil, errors.New("db: driver cannot be empty")
	}

	mysqlConfig = mysql.Configuration(mysqlConfig, configurator)
	sqliteConfig = sqlite.Configuration(sqliteConfig, configurator)

	return NewDB(config, informer, mysqlConfig, sqliteConfig)
}

func NewDB(config *Config, informer log.Informer, mysqlConfig *mysql.Config, sqliteConfig *sqlite.Config) (goqu.SQLDatabase, error) {
	switch config.Driver {
	case MySQLDriver:
		informer.Info("db: mysql usage")

		return mysql.NewMySQL(mysqlConfig, informer)
	case SQLiteDriver:
		informer.Info("db: sqlite usage")

		return sqlite.NewSQLite(sqliteConfig, informer)
	}

	return nil, errors.New(fmt.Sprintf("db: driver '%s' unknown", config.Driver))
}
