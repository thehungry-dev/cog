package log

type Handler interface {
	Handle(msg Message) Message
}
