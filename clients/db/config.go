package db

const (
	DriverFieldName = "db.driver"

	MySQLDriver  = "mysql"
	SQLiteDriver = "sqlite"
)

type Config struct {
	Driver string
}

func NewConfig() *Config {
	return &Config{}
}
