package main

import (
	"fmt"

	ctrls "github.com/thehungry-dev/log/ctrls/message"
	"github.com/thehungry-dev/log/message"
	"github.com/thehungry-dev/log/message/transform"
)

func main() {
	msg := ctrls.MessageWithFieldsExample()
	var text string
	text = transform.ToJSON(msg)
	fmt.Println(text)
	text = transform.ToText(msg)
	fmt.Println(text)
	text = transform.ToTextConfigured(msg, transform.EverythingConfig)
	fmt.Println(text)
	msg.Content = message.DataContent
	text = transform.ToText(msg)
	fmt.Println(text)
	text = transform.ToTextConfigured(msg, transform.EverythingConfig)
	fmt.Println(text)
	msg.Body = ""
	text = transform.ToText(msg)
	fmt.Println(text)
	text = transform.ToJSON(msg)
	fmt.Println(text)
}
