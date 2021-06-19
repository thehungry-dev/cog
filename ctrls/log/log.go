package log

import (
	"github.com/thehungry-dev/cog"
	"github.com/thehungry-dev/cog/handler"
	"github.com/thehungry-dev/cog/message/field"
	"github.com/thehungry-dev/cog/message/transform"
	"github.com/thehungry-dev/cog/paint"
)

const BodyExample = "A log message"

var TagsExample []string
var DataExample []field.Field

type TelemetricWriter struct {
	Last string
	cog.Writer
}

func (writer *TelemetricWriter) WriteString(text string) (int, error) {
	writer.Last = text
	return len(text), nil
}

func LogColoredTextExample() cog.Writer {
	return cog.TagsWriter
}

func LogColoredTextDeviceExample() *TelemetricWriter {
	telemetricWriter := &TelemetricWriter{}

	telemetricWriter.Writer = cog.With(
		handler.BuildPipe(
			handler.Output{
				Device: telemetricWriter,
				Format: handler.OutputText,
				Config: transform.EverythingConfig,
			},
		),
	)

	return telemetricWriter
}

func LogTextDeviceExample() *TelemetricWriter {
	telemetricWriter := &TelemetricWriter{}

	config := transform.EverythingConfig
	config.BuildPainter = paint.NewNullPainter

	telemetricWriter.Writer = cog.With(
		handler.BuildPipe(
			handler.Output{
				Device: telemetricWriter,
				Format: handler.OutputText,
				Config: config,
			},
		),
	)

	return telemetricWriter
}

func LogTaggedDataTextDeviceExample() *TelemetricWriter {
	writer := LogTextDeviceExample()
	writer.Writer = writer.
		Tags(TagsExample...).
		Data(DataExample...)
	return writer
}

func init() {
	TagsExample = []string{"foo", "bar", "baz"}

	builder := field.Builder{}
	DataExample = []field.Field{
		builder.String("Name", "a name"),
		builder.Int("Number", 1),
		builder.Bool("Boolean", true),
	}
}
