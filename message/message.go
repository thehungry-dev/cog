// Package message helps building a log message
package message

import (
	"github.com/thehungry-dev/log/level"
	"github.com/thehungry-dev/log/message/field"
)

type Message struct {
	halted bool
	tags   []string
	level  level.Level
	fields []field.Field
}

func New() Message {
	return Message{}
}

func (msg Message) Halt() Message {
	msg.halted = true
	return msg
}
func (msg Message) IsHalted() bool {
	return msg.halted
}

func (msg Message) Tags() []string {
	return msg.tags
}
func (msg Message) SetTags(tags []string) Message {
	msg.tags = tags
	return msg
}

func (msg Message) Level() level.Level {
	return msg.level
}
func (msg Message) SetLevel(level level.Level) Message {
	msg.level = level
	return msg
}

func (msg Message) Fields() []field.Field {
	return msg.fields
}
func (msg Message) Field(index int) field.Field {
	return msg.fields[index]
}
func (msg Message) SetFields(fields []field.Field) Message {
	msg.fields = fields
	return msg
}

func (msg Message) HasFields() bool {
	return len(msg.fields) > 0
}
