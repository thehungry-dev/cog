// Package filter provides a data structure to filter messages based on level
package filter

import (
	"strings"

	"github.com/thehungry-dev/log/level"
)

const (
	// NoneKeyword indicates no message will be logged
	NoneKeyword = "_none"
	// AllKeyword indicates that all messages will be logged
	AllKeyword = "_all"
	// MinKeyword indicates minimum verbosity, Fatal level
	MinKeyword = "_min"
	// MaxKeyword indicates maximum verbosity, Trace level
	MaxKeyword   = "_max"
	TraceKeyword = "trace"
	DebugKeyword = "debug"
	InfoKeyword  = "info"
	WarnKeyword  = "warn"
	ErrorKeyword = "error"
	FatalKeyword = "fatal"
)

const (
	None level.Level = level.Level(^uint(0) >> 1)
	All  level.Level = -None - 1
)

var Nothing *LevelFilter

type LevelFilter struct {
	level level.Level
}

func Parse(filterText string) (*LevelFilter, bool) {
	var filter *LevelFilter
	ok := true

	levelText := strings.TrimSpace(filterText)
	levelText = strings.ToLower(levelText)

	if levelText == "" || levelText == AllKeyword {
		return filter, ok
	}

	filter = &LevelFilter{}

	switch {
	case levelText == NoneKeyword:
		filter.level = None
	case levelText == AllKeyword:
		filter.level = All
	case levelText == MinKeyword:
		filter.level = level.Fatal
	case levelText == MaxKeyword:
		filter.level = level.Trace
	case levelText == TraceKeyword:
		filter.level = level.Trace
	case levelText == DebugKeyword:
		filter.level = level.Debug
	case levelText == InfoKeyword:
		filter.level = level.Info
	case levelText == WarnKeyword:
		filter.level = level.Warn
	case levelText == ErrorKeyword:
		filter.level = level.Error
	case levelText == FatalKeyword:
		filter.level = level.Fatal
	default:
		ok = false
	}

	return filter, ok
}

func (levelFilter *LevelFilter) Select(level level.Level) bool {
	return levelFilter.Level() <= level
}
func (levelFilter *LevelFilter) Reject(level level.Level) bool {
	return !levelFilter.Select(level)
}

func (levelFilter *LevelFilter) Level() level.Level {
	if levelFilter == nil {
		return All
	}

	return levelFilter.level
}
