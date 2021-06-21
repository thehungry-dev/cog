// Package settings configures build for logging
package settings

import (
	"io"
	"os"

	"github.com/thehungry-dev/cog/device"
	"github.com/thehungry-dev/cog/handler"
	"github.com/thehungry-dev/cog/message/transform"
	"github.com/thehungry-dev/cog/paint"
	"github.com/thehungry-dev/cog/terminal"
)

type Builder struct {
	TagFilterText    string
	LevelFilterText  string
	DeviceText       string
	TagsOutputText   string
	OutputFormatText string
	ColorText        string
}

func Getenv() Builder {
	return Builder{
		TagFilterText:    os.Getenv("LOG_TAGS"),
		LevelFilterText:  os.Getenv("LOG_LEVEL"),
		DeviceText:       os.Getenv("LOG_DEVICE"),
		TagsOutputText:   os.Getenv("LOG_TAGS_OUTPUT"),
		OutputFormatText: os.Getenv("LOG_OUTPUT_FORMAT"),
		ColorText:        os.Getenv("LOG_COLOR"),
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
	config.BuildPainter = builder.PainterBuilder()

	return handler.Output{
		Device: builder.Device(),
		Format: builder.OutputFormat(),
		Config: config,
	}
}
func (builder Builder) OutputTags() bool { return builder.TagsOutputText == "on" }
func (builder Builder) Color() bool {
	var color bool

	switch builder.ColorText {
	case "off":
		color = false
	case "on":
		color = true
	case "":
		fallthrough
	default:
		color = terminal.IsColorSupported(builder.Device())
	}

	return color
}
func (builder Builder) PainterBuilder() func(io.StringWriter) paint.Painter {
	if !builder.Color() {
		return paint.NewNullPainter
	}

	return paint.NewFullPainter
}
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
		output = device.Null
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
