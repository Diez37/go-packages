package sqlite

const (
	DsnFieldName = "db.sqlite.dsn"

	DsnDefault = "/tmp/db"
)

type Config struct {
	Dsn string
}

func NewConfig() *Config {
	return &Config{}
}
