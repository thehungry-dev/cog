// Package handler provides controls for test environment
package handler

import (
	filterctrls "github.com/thehungry-dev/log/ctrls/tag/filter"
	"github.com/thehungry-dev/log/handler"
)

func TagFilterExample() handler.Handler {
	tagFilterText := filterctrls.StringExample()
	return handler.BuildTagFilter(tagFilterText)
}
