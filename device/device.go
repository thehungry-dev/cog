// Package device provides different kind of devices valid for log output
package device

import "io"

var Null io.StringWriter

type nullStringWriter struct{}

func (nullStringWriter) WriteString(_ string) (int, error) { return 0, nil }

func init() {
	Null = nullStringWriter{}
}
