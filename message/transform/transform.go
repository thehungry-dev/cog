// Package transform converts a build message into different formats
package transform

type Config struct {
	Timestamp bool
	Tags      bool
	Level     bool
	Fields    bool
}

var DefaultConfig Config
var EverythingConfig Config

func init() {
	DefaultConfig = Config{
		Timestamp: true,
		Tags:      false,
		Level:     true,
		Fields:    true,
	}

	EverythingConfig = Config{
		Timestamp: true,
		Tags:      true,
		Level:     true,
		Fields:    true,
	}
}
