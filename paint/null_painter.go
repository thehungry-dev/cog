package paint

import (
	"io"

	"github.com/thehungry-dev/cog/message"
)

type nullPainter struct{}

func (nullPainter) BeginMessage(_ message.Message)   {}
func (nullPainter) EndMessage(_ message.Message)     {}
func (nullPainter) BeginTimestamp(_ message.Message) {}
func (nullPainter) EndTimestamp(_ message.Message)   {}
func (nullPainter) BeginTags(_ message.Message)      {}
func (nullPainter) EndTags(_ message.Message)        {}
func (nullPainter) BeginLevel(_ message.Message)     {}
func (nullPainter) EndLevel(_ message.Message)       {}
func (nullPainter) BeginBody(_ message.Message)      {}
func (nullPainter) EndBody(_ message.Message)        {}
func (nullPainter) BeginFields(_ message.Message)    {}
func (nullPainter) EndFields(_ message.Message)      {}

func NewNullPainter(_ io.StringWriter) Painter { return NullPainter }
