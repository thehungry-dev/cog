package handler

import (
	"github.com/thehungry-dev/log/level/filter"
	"github.com/thehungry-dev/log/message"
)

type LevelFilter struct {
	levelFilter *filter.LevelFilter
}

func ParseLevelFilter(filterText string) (LevelFilter, bool) {
	levelFilter, ok := filter.Parse(filterText)
	return LevelFilter{levelFilter}, ok
}

func BuildLevelFilter(filterText string) LevelFilter {
	levelFilter, ok := ParseLevelFilter(filterText)

	if !ok {
		panic("Level filter built with invalid string")
	}

	return levelFilter
}

func (handler LevelFilter) Handle(msg message.Message) message.Message {
	if handler.levelFilter.Select(msg.Level) {
		return msg
	}

	return msg.Halt()
}
