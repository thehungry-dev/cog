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
	painter := config.BuildPainter(&buf)

	painter.BeginMessage(msg)

	if config.Timestamp {
		painter.BeginTimestamp(msg)
		putTimestamp(&buf, msg)
		painter.EndTimestamp(msg)
	}

	if config.Tags {
		painter.BeginTags(msg)
		putTags(&buf, msg)
		painter.EndTags(msg)
	}

	if config.Level {
		painter.BeginLevel(msg)
		putLevel(&buf, msg)
		painter.EndLevel(msg)
	}

	painter.BeginBody(msg)
	buf.WriteString(text)
	painter.EndBody(msg)

	if config.Fields {
		painter.BeginFields(msg)
		putFields(&buf, msg)
		painter.EndFields(msg)
	}

	painter.EndMessage(msg)

	buf.WriteString("\n")

	return buf.String()
}

func putTimestamp(buf io.StringWriter, msg message.Message) {
	buf.WriteString("[")
	data := msg.Timestamp.Format(time.RFC3339)
	buf.WriteString(data)
	buf.WriteString("] ")
}

func putTags(buf io.StringWriter, msg message.Message) {
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

func putLevel(buf io.StringWriter, msg message.Message) {
	lvlText := msg.Level.String()
	lvlText = strings.ToUpper(lvlText)
	lvlText = fmt.Sprintf("%-5s ", lvlText)
	buf.WriteString(lvlText)
}

func putFields(buf io.StringWriter, msg message.Message) {
	switch msg.Content {
	case message.DataContent:
		data := FieldsToTextJSON(msg)
		buf.WriteString(data)
	case message.TextContent:
		data := FieldsToText(msg)
		buf.WriteString(data)
	}
}
