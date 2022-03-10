package mysql

const (
	HostFieldName     = "db.mysql.host"
	PortFieldName     = "db.mysql.port"
	UserFieldName     = "db.mysql.auth.user"
	PasswordFieldName = "db.mysql.auth.password"
	DataBaseFieldName = "db.mysql.name"

	HostDefault string = "127.0.0.1"
	PortDefault uint32 = 3306
)

type Config struct {
	Host string
	Port uint32

	User     string
	Password string

	DataBase string
}

func NewConfig() *Config {
	return &Config{}
}
