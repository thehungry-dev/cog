// Package handler provides controls for test environment
package handler

import (
	tagfilterctrls "github.com/thehungry-dev/log/ctrls/tag/filter"
	"github.com/thehungry-dev/log/handler"
	levelfilter "github.com/thehungry-dev/log/level/filter"
)

func TagFilterExample() handler.Handler {
	tagFilterText := tagfilterctrls.StringExample()
	return handler.BuildTagFilter(tagFilterText)
}

func LevelFilterExample() handler.Handler {
	return handler.BuildLevelFilter(levelfilter.InfoKeyword)
}
