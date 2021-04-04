package handler

import (
	"github.com/thehungry-dev/log/message"
	"github.com/thehungry-dev/log/tag/filter"
)

type TagFilter struct {
	tagFilter *filter.TagFilter
}

func BuildTagFilter(filterText string) TagFilter {
	tagFilter := filter.Parse(filterText)
	return TagFilter{tagFilter}
}

func (handler TagFilter) Handle(msg message.Message) message.Message {
	if handler.tagFilter.Select(msg.Tags) {
		return msg
	}

	return msg.Halt()
}
