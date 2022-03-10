package bind_flags

const (
	Cache            Component = iota
	CacheOnlyRedis   Component = iota
	CacheOnlyGoCache Component = iota
	DataBase         Component = iota
	MySQL            Component = iota
	SQLite           Component = iota
	Logger           Component = iota
	HttpServer       Component = iota
	HttpClient       Component = iota
	Tracer           Component = iota
	Migrator         Component = iota
)

type Component uint32
