package handler

import "github.com/thehungry-dev/log/level"

type Message struct {
	halted bool
	tags   []string
	level  level.Level
}

func BuildMessage(tags []string) Message {
	return Message{tags: tags}
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
