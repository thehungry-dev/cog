package level

type Level int

const (
	Trace Level = iota
	Debug
	Info
	Warn
	Error
	Fatal
)
