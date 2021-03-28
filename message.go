package log

type Message struct {
	halted bool
}

func (msg Message) Halt() Message {
	msg.halted = true
	return msg
}

func (msg Message) IsHalted() bool {
	return msg.halted
}
