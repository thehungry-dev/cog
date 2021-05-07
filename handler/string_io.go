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
	format IOFormat
	config transform.Config
}

func BuildStringIOJSON(writer io.StringWriter) StringIO {
	return StringIO{writer, JSONFormat, transform.DefaultConfig}
}

func BuildStringIOTextConfigured(writer io.StringWriter, config transform.Config) StringIO {
	return StringIO{writer, TextFormat, config}
}

func BuildStringIOText(writer io.StringWriter) StringIO {
	return BuildStringIOTextConfigured(writer, transform.DefaultConfig)
}

func (handler StringIO) Handle(msg message.Message) message.Message {
	if msg.IsHalted() {
		return msg
	}

	var output string

	switch handler.format {
	case TextFormat:
		output = transform.ToTextConfigured(msg, handler.config)
	case JSONFormat:
		output = transform.ToJSON(msg)
	}

	handler.writer.WriteString(output)

	return msg
}
