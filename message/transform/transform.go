// Package transform converts a build message into different formats
package transform

import (
	"io"

	"github.com/thehungry-dev/cog/paint"
)

type Config struct {
	Timestamp    bool
	Tags         bool
	Level        bool
	Fields       bool
	BuildPainter func(io.StringWriter) paint.Painter
}

var DefaultConfig Config
var EverythingConfig Config

func init() {
	DefaultConfig = Config{
		Timestamp:    true,
		Tags:         false,
		Level:        true,
		Fields:       true,
		BuildPainter: paint.NewFullPainter,
	}

	EverythingConfig = Config{
		Timestamp:    true,
		Tags:         true,
		Level:        true,
		Fields:       true,
		BuildPainter: paint.NewFullPainter,
	}
}
