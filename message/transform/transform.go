// Package transform converts a log message into different formats
package transform

import (
	"github.com/thehungry-dev/log/message"
)

func ToString(msg message.Message) string {
	var output string

	switch msg.Content {
	case message.TextContent:
		output = ToText(msg)
	case message.DataContent:
		output = ToJSON(msg)
	default:
		output = "<content not implemented>"
	}

	return output
}
