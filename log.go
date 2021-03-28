// Package log provides structured and unstructured logging functions
package log

var DefaultLogger Handler

func init() {
	DefaultLogger = Pipe{}
}
