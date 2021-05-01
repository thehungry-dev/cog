package handler

import (
	tagfilterctrls "github.com/thehungry-dev/log/ctrls/tag/filter"
	"github.com/thehungry-dev/log/level"
	"github.com/thehungry-dev/log/message"
	"github.com/thehungry-dev/log/message/field"
)

func MessageTagsExample() []string {
	return tagfilterctrls.TagsMatchingExample()
}
func MessageLevelExample() level.Level {
	return level.Info
}
func MessageLevelAlternateExample() level.Level {
	return level.Trace
}
func MessageTagsExcludedExample() []string {
	return tagfilterctrls.TagsExcludedMatchingExample()
}

func MessageExample() message.Message {
	msg := message.New()
	msg.Tags = MessageTagsExample()
	msg.Level = MessageLevelExample()

	return msg
}
func MessageExcludedExample() message.Message {
	msg := message.New()
	msg.Tags = MessageTagsExcludedExample()

	return msg
}
func MessageLevelExcludedExample() message.Message {
	msg := message.New()
	msg.Tags = MessageTagsExample()
	msg.Level = MessageLevelAlternateExample()

	return msg
}

func MessageWithFieldsExample() message.Message {
	msg := MessageExample()
	msg.Content = message.TextContent
	msg.Body = "Example Message"
	msg.Fields = []field.Field{
		{Type: field.Int64, Name: "id", ValueNumeric: 12},
		{Type: field.Bool, Name: "success", ValueNumeric: 1},
	}

	return msg
}
