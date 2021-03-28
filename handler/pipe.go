package handler

type Pipe []Handler

func BuildPipe(handlers ...Handler) Pipe {
	return handlers
}

func (handlers Pipe) Handle(msg Message) Message {
	for _, handler := range handlers {
		msg = handler.Handle(msg)
	}

	return msg
}
