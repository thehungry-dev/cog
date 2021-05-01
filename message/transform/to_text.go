package transform

import (
	"strings"

	"github.com/thehungry-dev/log/message"
)

func ToText(msg message.Message) string {
	text := msg.Body

	var b strings.Builder

	b.WriteString(text)

	var publicFieldsCount = 0

	for _, fld := range msg.Fields {
		if !fld.IsOutput() {
			continue
		}

		publicFieldsCount += 1

		if publicFieldsCount == 1 {
			b.WriteString(" (")
		}

		if publicFieldsCount > 1 {
			b.WriteString(", ")
		}

		b.WriteString(fld.Name)
		b.WriteString(": ")

		output, _ := FieldToString(fld)
		b.WriteString(output)
	}

	if publicFieldsCount > 0 {
		b.WriteString(")")
	}

	return b.String()
}
