package handler

import (
	tagfilterctrls "github.com/thehungry-dev/log/ctrls/tag/filter"
	"github.com/thehungry-dev/log/level"
	"github.com/thehungry-dev/log/message"
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
	msg := message.Build(MessageTagsExample())
	msg = msg.SetLevel(MessageLevelExample())

	return msg
}
func MessageExcludedExample() message.Message {
	return message.Build(MessageTagsExcludedExample())
}
func MessageLevelExcludedExample() message.Message {
	msg := message.Build(MessageTagsExample())
	msg = msg.SetLevel(MessageLevelAlternateExample())

	return msg
}
