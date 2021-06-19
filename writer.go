package cog

import (
	"github.com/thehungry-dev/cog/handler"
	"github.com/thehungry-dev/cog/level"
	"github.com/thehungry-dev/cog/message"
	"github.com/thehungry-dev/cog/message/field"
)

type Writer struct {
	builder message.Builder
	hndl    handler.Handler
}

func With(hndl handler.Handler) Writer { return Writer{hndl: hndl} }

func (writer Writer) Tags(tags ...string) Writer {
	writer.builder = writer.builder.Tags(tags...)
	return writer
}
func (writer Writer) Data(fields ...field.Field) Writer {
	writer.builder = writer.builder.Data(fields...)
	return writer
}
func (writer Writer) Put(lvl level.Level, body string) {
	writer.builder = writer.builder.Put(lvl, body)
	msg := writer.builder.Message()
	writer.hndl.Handle(msg)
}
func (writer Writer) Putf(lvl level.Level, format string, a ...interface{}) {
	writer.builder = writer.builder.Putf(lvl, format, a...)
	msg := writer.builder.Message()
	writer.hndl.Handle(msg)
}
func (writer Writer) PutLevel(lvl level.Level) {
	writer.builder = writer.builder.PutLevel(lvl)
	msg := writer.builder.Message()
	writer.hndl.Handle(msg)
}

func (writer Writer) Traced()                           { writer.PutLevel(level.Trace) }
func (writer Writer) Trace(body string)                 { writer.Put(level.Trace, body) }
func (writer Writer) Tracef(f string, a ...interface{}) { writer.Putf(level.Trace, f, a...) }

func (writer Writer) Debugd()                           { writer.PutLevel(level.Debug) }
func (writer Writer) Debug(body string)                 { writer.Put(level.Debug, body) }
func (writer Writer) Debugf(f string, a ...interface{}) { writer.Putf(level.Debug, f, a...) }

func (writer Writer) Infod()                           { writer.PutLevel(level.Info) }
func (writer Writer) Info(body string)                 { writer.Put(level.Info, body) }
func (writer Writer) Infof(f string, a ...interface{}) { writer.Putf(level.Info, f, a...) }

func (writer Writer) Warnd()                           { writer.PutLevel(level.Warn) }
func (writer Writer) Warn(body string)                 { writer.Put(level.Warn, body) }
func (writer Writer) Warnf(f string, a ...interface{}) { writer.Putf(level.Warn, f, a...) }

func (writer Writer) Errord()                           { writer.PutLevel(level.Error) }
func (writer Writer) Error(body string)                 { writer.Put(level.Error, body) }
func (writer Writer) Errorf(f string, a ...interface{}) { writer.Putf(level.Error, f, a...) }

func (writer Writer) Fatald()                           { writer.PutLevel(level.Fatal) }
func (writer Writer) Fatal(body string)                 { writer.Put(level.Fatal, body) }
func (writer Writer) Fatalf(f string, a ...interface{}) { writer.Putf(level.Fatal, f, a...) }
