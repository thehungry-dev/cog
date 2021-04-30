package main

import (
	"fmt"

	ctrls "github.com/thehungry-dev/log/ctrls/message"
	"github.com/thehungry-dev/log/message/transform"
)

func main() {
	msg := ctrls.MessageWithFieldsExample()
	text := transform.ToString(msg)
	text = transform.ToJSON(msg)
	fmt.Println(text)
}
