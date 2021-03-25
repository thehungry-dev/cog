package log

type Logger interface {
	Put(msg Message) interface{}
}
