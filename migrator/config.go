package migrator

const (
	SourceFieldName = "migrator.source"

	SourceDefault = "file://./migrations"
)

type Config struct {
	Source string
}

func NewConfig() *Config {
	return &Config{}
}
