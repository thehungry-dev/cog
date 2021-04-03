package level

type Level int8

const (
	Trace Level = iota
	Debug
	Info
	Warn
	Error
	Fatal
)
