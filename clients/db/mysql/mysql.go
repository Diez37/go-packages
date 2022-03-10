package mysql

import (
	"database/sql"
	"fmt"
	"github.com/diez37/go-packages/configurator"
	"github.com/diez37/go-packages/log"
	"github.com/doug-martin/goqu/v9"
)

func WithConfigurator(config *Config, configurator configurator.Configurator, informer log.Logger) (goqu.SQLDatabase, error) {
	config = Configuration(config, configurator)

	return NewMySQL(config, informer)
}

func Configuration(config *Config, configurator configurator.Configurator) *Config {
	if host := configurator.GetString(HostFieldName); host != "" && config.Host == HostDefault {
		config.Host = host
	}

	if port := configurator.GetUint32(PortFieldName); port > 0 && config.Port == PortDefault {
		config.Port = port
	}

	if user := configurator.GetString(UserFieldName); user != "" {
		config.User = user
	}

	if password := configurator.GetString(PasswordFieldName); password != "" {
		config.Password = password
	}

	if dataBase := configurator.GetString(DataBaseFieldName); dataBase != "" {
		config.DataBase = dataBase
	}

	return config
}

func NewMySQL(config *Config, informer log.Informer) (goqu.SQLDatabase, error) {
	informer.Infof("mysql: host - %s, port - %d", config.Host, config.Port)
	informer.Infof("mysql: used database - %s", config.DataBase)

	return sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DataBase,
	))
}
