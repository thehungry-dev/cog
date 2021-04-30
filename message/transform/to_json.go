package transform

import (
	"io"
	"strings"

	"github.com/francoispqt/gojay"
	"github.com/thehungry-dev/log/level"
	"github.com/thehungry-dev/log/message"
)

type messageMarshaler struct{ msg message.Message }

func (m messageMarshaler) MarshalJSONObject(encoder *gojay.Encoder) {
	msg := m.msg

	// encoder.ArrayKey("tags", ) HOW?
	encoder.StringKey("message", msg.Body)
	// TODO: Other parts of message
}

func (m messageMarshaler) IsNil() bool {
	msg := m.msg

	isNil := msg.Body == ""
	isNil = isNil && len(msg.Fields) == 0
	isNil = isNil && len(msg.Tags) == 0
	isNil = isNil && msg.Level == level.Trace
	isNil = isNil && msg.Content == message.DataContent

	return isNil
}

func ToJSONWriter(msg message.Message, writer io.Writer) error {
	encoder := gojay.BorrowEncoder(writer)
	defer encoder.Release()

	marshaler := messageMarshaler{msg}

	encoder.Encode(marshaler)

	return nil
}

func ToJSON(msg message.Message) string {
	var textBuilder strings.Builder

	if err := ToJSONWriter(msg, &textBuilder); err != nil {
		return ""
	}

	return textBuilder.String()
}
