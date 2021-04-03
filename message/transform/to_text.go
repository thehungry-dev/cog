package transform

import (
	"strings"

	"github.com/thehungry-dev/log/message"
)

func ToText(msg message.Message) string {
	text, ok := TextMessage(msg)

	if !ok {
		panic("Message is not a text message")
	}

	var b strings.Builder
	var hasData bool

	b.WriteString(text)

	for index, fld := range msg.Fields() {
		if index == 0 {
			b.WriteString(" (")
			hasData = true
			continue
		}

		if index != 1 {
			b.WriteString(", ")
		}

		b.WriteString(fld.Name)
		b.WriteString(": ")

		output, _ := FieldToString(fld)
		b.WriteString(output)
	}

	if hasData {
		b.WriteString(")")
	}

	return b.String()
}
