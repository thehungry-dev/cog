package message

import (
	"fmt"
	"strings"
	"time"

	"github.com/thehungry-dev/cog/level"
	"github.com/thehungry-dev/cog/message/field"
)

type Builder struct {
	msg Message
}

func (builder Builder) Message() Message { return builder.msg }

func (builder Builder) Tags(tags ...string) Builder {
	if len(builder.msg.Tags) == 0 {
		builder.msg.Tags = tags
	} else {
		builder.msg.Tags = append(builder.msg.Tags, tags...)
	}

	return builder
}

func (builder Builder) Data(fields ...field.Field) Builder {
	if len(builder.msg.Fields) == 0 {
		builder.msg.Fields = fields
	} else {
		builder.msg.Fields = append(builder.msg.Fields, fields...)
	}

	return builder
}

func (builder Builder) PutLevel(lvl level.Level) Builder {
	builder.msg.Level = lvl
	builder.msg.Timestamp = time.Now().UTC()

	return builder
}

func (builder Builder) Put(lvl level.Level, body string) Builder {
	builder.msg.Level = lvl
	builder.msg.Timestamp = time.Now().UTC()
	builder.msg.Body = body
	builder.msg.Content = TextContent

	return builder
}

func (builder Builder) Putf(lvl level.Level, format string, a ...interface{}) Builder {
	var buf strings.Builder

	base := fmt.Sprintf(format, a...)

	buf.WriteString(base)
	if !strings.HasSuffix(base, "\n") {
		buf.WriteString("\n")
	}

	body := buf.String()

	return builder.Put(lvl, body)
}
