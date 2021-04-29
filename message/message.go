// Package message helps building a log message
package message

import (
	"github.com/thehungry-dev/log/level"
	"github.com/thehungry-dev/log/message/field"
)

type Message struct {
	halted bool

	Tags    []string
	Level   level.Level
	Content Content
	Body    string

	Fields []field.Field
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

func (msg Message) Field(index int) field.Field {
	return msg.Fields[index]
}

func (msg Message) HasFields() bool {
	return len(msg.Fields) > 0
}
