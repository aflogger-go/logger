package logger

//go:generate stringer -type=Level
type Level int

// levels
const (
	Debug Level = iota
	Info
	Warn
	Err
	Crit
)

// basic config
type Config struct {

	// ConstFields are displayed at every log entry
	ConstFields map[string]string

	// Debug shows debug messages
	Debug bool
}

// Logger just what I want to see a logger have
type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})

	Info(v ...interface{})
	Infof(format string, v ...interface{})

	Warn(v ...interface{})
	Warnf(format string, v ...interface{})

	Err(v ...interface{})
	Errf(format string, v ...interface{})

	Crit(v ...interface{})
	Critf(format string, v ...interface{})

	Level(lvl Level, v ...interface{})
	Levelf(lvl Level, format string, v ...interface{})
}