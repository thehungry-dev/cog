package settings

import (
	"io"
	"os"

	"github.com/thehungry-dev/log/handler"
)

const DefaultTagFilterText = "_all"

type nullStringWriter struct{}

func (nullStringWriter) WriteString(_ string) (int, error) { return 0, nil }

type Builder struct {
	TagFilterText   string
	LevelFilterText string
	DeviceText      string
	FormattersText  string
}

func Getenv() Builder {
	return Builder{
		TagFilterText:   os.Getenv("LOG_TAGS"),
		LevelFilterText: os.Getenv("LOG_LEVEL"),
		DeviceText:      os.Getenv("LOG_DEVICE"),
		FormattersText:  os.Getenv("LOG_FORMATTERS"),
	}
}

func (builder Builder) TagFilter() handler.Handler {
	return handler.BuildTagFilter(builder.TagFilterText)
}
func (builder Builder) LevelFilter() handler.Handler {
	if levelFilter, ok := handler.ParseLevelFilter(builder.LevelFilterText); ok {
		return levelFilter
	} else {
		return handler.Passthrough
	}
}
func (builder Builder) Formatters() bool { return builder.FormattersText == "on" }
func (builder Builder) Device() io.StringWriter {
	var output io.StringWriter

	switch builder.DeviceText {
	case "stdout":
		output = os.Stdout
	case "stderr":
		fallthrough
	case "":
		output = os.Stderr
	default:
		output = nullStringWriter{}
	}

	return output
}
