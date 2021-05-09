package handler

import "github.com/thehungry-dev/log/message"

type passthroughType struct{}

func (passthroughType) Handle(msg message.Message) message.Message { return msg }

var Passthrough Handler

func init() {
	Passthrough = passthroughType{}
}
