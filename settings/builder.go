package settings

import (
	"io"
	"os"

	"github.com/thehungry-dev/log/handler"
	"github.com/thehungry-dev/log/message/transform"
)

const DefaultTagFilterText = "_all"

type nullStringWriter struct{}

func (nullStringWriter) WriteString(_ string) (int, error) { return 0, nil }

type Builder struct {
	TagFilterText    string
	LevelFilterText  string
	DeviceText       string
	FormattersText   string
	TagsOutputText   string
	OutputFormatText string
}

func Getenv() Builder {
	return Builder{
		TagFilterText:    os.Getenv("LOG_TAGS"),
		LevelFilterText:  os.Getenv("LOG_LEVEL"),
		DeviceText:       os.Getenv("LOG_DEVICE"),
		FormattersText:   os.Getenv("LOG_FORMATTERS"),
		TagsOutputText:   os.Getenv("LOG_TAGS_OUTPUT"),
		OutputFormatText: os.Getenv("LOG_OUTPUT_FORMAT"),
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
func (builder Builder) OutputHandler() handler.Handler {
	config := transform.DefaultConfig
	config.Tags = builder.OutputTags()

	return handler.Output{
		Device: builder.Device(),
		Format: builder.OutputFormat(),
		Config: config,
	}
}
func (builder Builder) OutputTags() bool { return builder.TagsOutputText == "on" }
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
func (builder Builder) OutputFormat() handler.OutputFormat {
	var format handler.OutputFormat

	switch builder.OutputFormatText {
	case "json":
		format = handler.OutputJSON
	case "text":
		fallthrough
	case "":
		fallthrough
	default:
		format = handler.OutputText
	}

	return format
}
