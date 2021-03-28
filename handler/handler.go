// Package handler provides functions to process log messages for output
package handler

type Handler interface {
	Handle(msg Message) Message
}

type HandlerFunc func(msg Message) Message

func (handle HandlerFunc) Handle(msg Message) Message {
	return handle(msg)
}
