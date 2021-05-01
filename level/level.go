package level

import "strconv"

type Level int8

const (
	Trace Level = iota
	Debug
	Info
	Warn
	Error
	Fatal
)

func (lvl Level) String() string {
	var output string

	switch lvl {
	case Trace:
		output = "trace"
	case Debug:
		output = "debug"
	case Info:
		output = "info"
	case Warn:
		output = "warn"
	case Error:
		output = "error"
	case Fatal:
		output = "fatal"
	default:
		output = strconv.Itoa(int(lvl))
	}

	return output
}
