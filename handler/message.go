package handler

type Message struct {
	halted bool
	tags   []string
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
