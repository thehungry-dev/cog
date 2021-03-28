package log

type Pipe []Handler

func (handlers Pipe) Handle(msg Message) Message {
	for _, handler := range handlers {
		msg = handler.Handle(msg)
	}

	return msg
}
