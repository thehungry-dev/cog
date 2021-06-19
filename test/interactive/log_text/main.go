package main

import (
	ctrls "github.com/thehungry-dev/cog/ctrls/log"
)

func main() {
	log := ctrls.LogTextExample()
	log.Tags("foo", "bar", "baz").Trace("A trace message")
	log.Tags("foo", "bar", "baz").Debug("A debug message")
	log.Tags("foo", "bar", "baz").Info("An info message")
	log.Tags("foo", "bar", "baz").Warn("A warn message")
	log.Tags("foo", "bar", "baz").Error("An error message")
	log.Tags("foo", "bar", "baz").Fatal("A fatal message")
}
