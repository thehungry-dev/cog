// Package handler provides controls for test environment
package handler

import (
	tagfilterctrls "github.com/thehungry-dev/cog/ctrls/tag/filter"
	"github.com/thehungry-dev/cog/handler"
	levelfilter "github.com/thehungry-dev/cog/level/filter"
)

func TagFilterExample() handler.Handler {
	tagFilterText := tagfilterctrls.StringExample()
	return handler.BuildTagFilter(tagFilterText)
}

func LevelFilterExample() handler.Handler {
	return handler.BuildLevelFilter(levelfilter.InfoKeyword)
}
