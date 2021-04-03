package handler

import "github.com/thehungry-dev/log/message"

type Pipe []Handler

func BuildPipe(handlers ...Handler) Pipe {
	return handlers
}

func (handlers Pipe) Handle(msg message.Message) message.Message {
	for _, handler := range handlers {
		msg = handler.Handle(msg)
	}

	return msg
}
