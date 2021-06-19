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

func Parse(text string) (Level, error) {
	var output Level

	switch text {
	case Trace.String():
		output = Trace
	case Debug.String():
		output = Debug
	case Info.String():
		output = Info
	case Warn.String():
		output = Warn
	case Error.String():
		output = Error
	case Fatal.String():
		output = Fatal
	default:
		num, err := strconv.Atoi(text)

		if err != nil {
			return output, err
		}

		output = Level(num)
	}

	return output, nil
}
