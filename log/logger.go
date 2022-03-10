package log

type Debuger interface {
	Debugf(string, ...interface{})
	Debug(...interface{})
}

type Informer interface {
	Infof(string, ...interface{})
	Info(...interface{})
}

type Warner interface {
	Warnf(string, ...interface{})
	Warn(...interface{})
}

type Printer interface {
	Print(...interface{})
}

// Logger general of logger interface for any implementation
type Logger interface {
	Debuger
	Informer
	Warner
	Printer

	Errorf(string, ...interface{})
	Error(...interface{})
}
