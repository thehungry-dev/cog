// Package handler provides functions to process log messages for output
package handler

import "github.com/thehungry-dev/log/message"

type Handler interface {
	Handle(msg message.Message) message.Message
}

type HandlerFunc func(msg message.Message) message.Message

func (handle HandlerFunc) Handle(msg message.Message) message.Message {
	return handle(msg)
}

var Passthrough Pipe
