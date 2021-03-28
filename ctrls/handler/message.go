package handler

import (
	filterctrls "github.com/thehungry-dev/log/ctrls/tag/filter"
	"github.com/thehungry-dev/log/handler"
)

func MessageTagsExample() []string {
	return filterctrls.TagsMatchingExample()
}
func MessageTagsExcludedExample() []string {
	return filterctrls.TagsExcludedMatchingExample()
}

func MessageExample() handler.Message {
	return handler.BuildMessage(MessageTagsExample())
}
func MessageExcludedExample() handler.Message {
	return handler.BuildMessage(MessageTagsExcludedExample())
}
