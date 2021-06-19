package paint

import (
	"io"

	"github.com/thehungry-dev/cog/color"
	"github.com/thehungry-dev/cog/message"
)

type fullPainter struct {
	buf    io.StringWriter
	styled bool
}

func NewFullPainter(buf io.StringWriter) Painter {
	return &fullPainter{buf: buf, styled: false}
}

func (painter *fullPainter) BeginMessage(msg message.Message) {
	clr := Color(msg.Level)

	painter.styled = clr != color.None

	painter.buf.WriteString(clr)
}

func (painter *fullPainter) EndMessage(msg message.Message) {
	if !painter.styled {
		return
	}

	painter.styled = false

	painter.buf.WriteString(color.Reset)
}

func (*fullPainter) BeginTimestamp(_ message.Message) {}
func (*fullPainter) EndTimestamp(_ message.Message)   {}
func (*fullPainter) BeginTags(_ message.Message)      {}
func (*fullPainter) EndTags(_ message.Message)        {}
func (*fullPainter) BeginLevel(_ message.Message)     {}
func (*fullPainter) EndLevel(_ message.Message)       {}
func (*fullPainter) BeginBody(_ message.Message)      {}
func (*fullPainter) EndBody(_ message.Message)        {}
func (*fullPainter) BeginFields(_ message.Message)    {}
func (*fullPainter) EndFields(_ message.Message)      {}
