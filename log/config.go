package log

const (
	// VerboseDefault is using debug log level on default for Config.InfoLevel
	VerboseDefault = false

	// InfoLevelFieldName field name in configuration file or ENV name for value of Config.InfoLevel
	InfoLevelFieldName  = "logger.level.info"
	DebugLevelFieldName = "logger.level.debug"

	// SentryDSNFieldName field name in configuration file or ENV name for value of Config.SentryDSN
	SentryDSNFieldName = "logger.sentry.dsn"
)

// Config setup params for logger
type Config struct {
	// InfoLevel if true then logger using debug log level
	InfoLevel  bool
	DebugLevel bool

	// SentryDSN Sentry connection string
	SentryDSN string
}

// NewConfig creating and return new structure instance Config
func NewConfig() *Config {
	return &Config{}
}
