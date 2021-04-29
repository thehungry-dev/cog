// Package transform converts a log message into different formats
package transform

import (
	"github.com/thehungry-dev/log/message"
	"github.com/thehungry-dev/log/message/field/name"
)

func TextMessage(msg message.Message) (string, bool) {
	if !IsTextMessage(msg) {
		return "", false
	}
	return msg.Field(0).ValueString, true
}

func IsTextMessage(msg message.Message) bool {
	return msg.HasFields() && msg.Field(0).Name == name.Body
}

func ToString(msg message.Message) string {
	if IsTextMessage(msg) {
		return ToText(msg)
	}

	return ToJSON(msg)
}
