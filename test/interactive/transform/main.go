package main

import (
	"fmt"

	ctrls "github.com/thehungry-dev/cog/ctrls/message"
	"github.com/thehungry-dev/cog/message"
	"github.com/thehungry-dev/cog/message/transform"
)

func main() {
	msg := ctrls.MessageWithFieldsExample()
	var text string
	text = transform.ToJSON(msg)
	fmt.Print(text)
	text = transform.ToText(msg)
	fmt.Print(text)
	text = transform.ToTextConfigured(msg, transform.EverythingConfig)
	fmt.Print(text)
	msg.Content = message.DataContent
	text = transform.ToText(msg)
	fmt.Print(text)
	text = transform.ToTextConfigured(msg, transform.EverythingConfig)
	fmt.Print(text)
	msg.Body = ""
	text = transform.ToText(msg)
	fmt.Print(text)
	text = transform.ToJSON(msg)
	fmt.Print(text)
}
