// Package log provides structured and unstructured logging functions
package log

var DefaultLogger Handler

// ParseEnvTagFilter
// ParseEnvLevelFilter

func init() {
	DefaultLogger = Pipe{}
}
