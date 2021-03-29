package handler

import (
	tagfilterctrls "github.com/thehungry-dev/log/ctrls/tag/filter"
	"github.com/thehungry-dev/log/handler"
	"github.com/thehungry-dev/log/level"
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

func MessageExample() handler.Message {
	msg := handler.BuildMessage(MessageTagsExample())
	msg = msg.SetLevel(MessageLevelExample())

	return msg
}
func MessageExcludedExample() handler.Message {
	return handler.BuildMessage(MessageTagsExcludedExample())
}
func MessageLevelExcludedExample() handler.Message {
	msg := handler.BuildMessage(MessageTagsExample())
	msg = msg.SetLevel(MessageLevelAlternateExample())

	return msg
}
