package handler

import (
	"io"

	"github.com/thehungry-dev/log/message"
	"github.com/thehungry-dev/log/message/transform"
)

type StringIO struct {
	writer io.StringWriter
}

func BuildStringIO(writer io.StringWriter) StringIO {
	return StringIO{writer}
}

func (handler StringIO) Handle(msg message.Message) message.Message {
	output := transform.ToString(msg)
	handler.writer.WriteString(output)
	return msg
}
