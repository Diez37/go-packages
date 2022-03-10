package tracer

const (
	// UseProfileDefault is added path to http router for pprof on default for Config.UseProfile
	UseProfileDefault = true

	// UseProfileFieldName field name in configuration file or ENV name for value of Config.UseProfile
	UseProfileFieldName = "trace.pprof"

	// JaegerDSNFieldName field name in configuration file or ENV name for value of Config.JaegerDSN
	JaegerDSNFieldName = "trace.jaeger.dsn"
)

// Config setup params for tracer
type Config struct {
	// UseProfile if true then added path to http router for pprof
	UseProfile bool

	// JaegerDSN Jaeger connection string
	JaegerDSN string
}

// NewConfig creating and return new structure instance Config
func NewConfig() *Config {
	return &Config{}
}
