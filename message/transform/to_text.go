package transform

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/thehungry-dev/cog/message"
)

func ToText(msg message.Message) string {
	return ToTextConfigured(msg, DefaultConfig)
}

func ToTextConfigured(msg message.Message, config Config) string {
	text := msg.Body

	var buf strings.Builder

	if config.Timestamp {
		timestamp(&buf, msg)
	}

	if config.Tags {
		tags(&buf, msg)
	}

	if config.Level {
		level(&buf, msg)
	}

	buf.WriteString(text)

	if config.Fields {
		fields(&buf, msg)
	}

	buf.WriteString("\n")

	return buf.String()
}

func timestamp(buf io.StringWriter, msg message.Message) {
	buf.WriteString("[")
	data := msg.Timestamp.Format(time.RFC3339)
	buf.WriteString(data)
	buf.WriteString("] ")
}

func tags(buf io.StringWriter, msg message.Message) {
	if len(msg.Tags) == 0 {
		return
	}

	buf.WriteString("|")

	for index, tag := range msg.Tags {
		if index != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(tag)
	}

	buf.WriteString("| ")
}

func level(buf io.StringWriter, msg message.Message) {
	lvlText := msg.Level.String()
	lvlText = strings.ToUpper(lvlText)
	lvlText = fmt.Sprintf("%-5s ", lvlText)
	buf.WriteString(lvlText)
}

func fields(buf io.StringWriter, msg message.Message) {
	switch msg.Content {
	case message.DataContent:
		data := FieldsToTextJSON(msg)
		buf.WriteString(data)
	case message.TextContent:
		data := FieldsToText(msg)
		buf.WriteString(data)
	}
}
