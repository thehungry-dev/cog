// Package paint provides an interface to set terminal output colors for parts or the entirety of a log message
package paint

import (
	"github.com/thehungry-dev/cog/color"
	"github.com/thehungry-dev/cog/level"
	"github.com/thehungry-dev/cog/message"
)

type Painter interface {
	BeginMessage(msg message.Message)
	EndMessage(msg message.Message)
	BeginTimestamp(msg message.Message)
	EndTimestamp(msg message.Message)
	BeginTags(msg message.Message)
	EndTags(msg message.Message)
	BeginLevel(msg message.Message)
	EndLevel(msg message.Message)
	BeginBody(msg message.Message)
	EndBody(msg message.Message)
	BeginFields(msg message.Message)
	EndFields(msg message.Message)
}

var NullPainter Painter

func Color(lvl level.Level) string {
	clr, ok := palette[lvl]

	if !ok {
		return color.None
	}

	return clr
}

func init() {
	NullPainter = nullPainter{}
}
