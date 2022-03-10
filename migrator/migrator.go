package migrator

import (
	"database/sql"
	"errors"
	"github.com/diez37/go-packages/clients/db"
	"github.com/diez37/go-packages/configurator"
	"github.com/doug-martin/goqu/v9"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func WithConfigurator(configurator configurator.Configurator, config *Config, dbConfig *db.Config, db goqu.SQLDatabase) (*migrate.Migrate, error) {
	configurator.SetDefault(SourceFieldName, SourceDefault)
	if directory := configurator.GetString(SourceFieldName); directory != "" && config.Source == "" {
		config.Source = directory
	}

	return NewMigrator(config, dbConfig, db)
}

func NewMigrator(config *Config, dbConfig *db.Config, sqlDatabase goqu.SQLDatabase) (*migrate.Migrate, error) {
	dbInstance, ok := sqlDatabase.(*sql.DB)
	if !ok {
		return nil, errors.New("migrator: db instance unknown")
	}

	var driver database.Driver
	var err error

	switch dbConfig.Driver {
	case db.MySQLDriver:
		driver, err = mysql.WithInstance(dbInstance, &mysql.Config{})
	case db.SQLiteDriver:
		driver, err = sqlite.WithInstance(dbInstance, &sqlite.Config{})
	}

	if err != nil {
		return nil, err
	}

	return migrate.NewWithDatabaseInstance(config.Source, dbConfig.Driver, driver)
}
