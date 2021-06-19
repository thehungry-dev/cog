// Package color provides a list of escape sequences to output colored text to terminal
package color

import "io"

const (
	None   = ""
	Reset  = "\033[0m"
	Bold   = "\033[1m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

func Apply(writer io.StringWriter, color string, do func()) {
	if color == None {
		do()
		return
	}

	writer.WriteString(color)

	do()

	writer.WriteString(Reset)
}
