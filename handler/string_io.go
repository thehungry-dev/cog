package handler

import (
	"io"

	"github.com/thehungry-dev/log/message"
	"github.com/thehungry-dev/log/message/transform"
)

type IOFormat int8

const (
	TextFormat IOFormat = iota
	JSONFormat
)

type StringIO struct {
	writer io.StringWriter
	Format IOFormat
}

func BuildStringIO(writer io.StringWriter) StringIO {
	return BuildStringIOFormatted(writer, TextFormat)
}

func BuildStringIOFormatted(writer io.StringWriter, format IOFormat) StringIO {
	return StringIO{writer, format}
}

func (handler StringIO) Handle(msg message.Message) message.Message {
	var output string

	switch handler.Format {
	case TextFormat:
		output = transform.ToText(msg)
	case JSONFormat:
		output = transform.ToJSON(msg)
	}

	handler.writer.WriteString(output)

	return msg
}
