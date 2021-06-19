// Package terminal provides functions to test terminal for supported features
package terminal

import (
	"io"

	"github.com/mattn/go-isatty"
)

type filedescriptor interface {
	Fd() uintptr
}

func IsColorSupported(device io.StringWriter) bool {
	descriptor, ok := device.(filedescriptor)

	if !ok {
		return false
	}

	supported := isatty.IsTerminal(descriptor.Fd())
	supported = supported || isatty.IsCygwinTerminal(descriptor.Fd())

	return supported
}
