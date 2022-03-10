package app

const (
	// NameDefault a name for application on default
	NameDefault string = "app"

	// MaxProcFieldName ENV name for value of maximum number process in goroutine manager pool
	MaxProcFieldName = "GOMAXPROC"

	// MaxProcDefault maximum number process in goroutine manager pool on default
	MaxProcDefault uint32 = 2
)

type Config struct {
	// Name this is a name for application
	Name string

	PID int
}

func NewConfig() *Config {
	return &Config{}
}
