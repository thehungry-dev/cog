package main

import (
	"github.com/thehungry-dev/log"
)

func main() {
	msg := log.Message{}
	logger := log.DefaultLogger
	logger.Handle(msg)
}
