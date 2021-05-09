package handler

import (
	"io"

	"github.com/thehungry-dev/log/message"
	"github.com/thehungry-dev/log/message/transform"
)

type OutputFormat int8

const (
	OutputText OutputFormat = iota
	OutputJSON
)

type Output struct {
	Device io.StringWriter
	Format OutputFormat
	transform.Config
}

func BuildOutput(device io.StringWriter, format OutputFormat) Output {
	return Output{Device: device, Format: format, Config: transform.DefaultConfig}
}

func (handler Output) Transform(msg message.Message) string {
	var text string

	switch handler.Format {
	case OutputText:
		text = transform.ToTextConfigured(msg, handler.Config)
	case OutputJSON:
		text = transform.ToJSON(msg)
	}

	return text
}

func (handler Output) Handle(msg message.Message) message.Message {
	if msg.IsHalted() {
		return msg
	}

	output := handler.Transform(msg)
	handler.Device.WriteString(output)

	return msg
}
