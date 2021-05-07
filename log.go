// Package log provides structured and unstructured logging functions
package log

import (
	"os"

	"github.com/thehungry-dev/log/handler"
	"github.com/thehungry-dev/log/message/field"
	"github.com/thehungry-dev/log/message/transform"
)

var Private field.Builder
var fields field.Builder
var DefaultWriter Writer
var TagsWriter Writer
var JSONWriter Writer

func String(name string, value string) field.Field         { return fields.String(name, value) }
func Int64(name string, value int64) field.Field           { return fields.Int64(name, value) }
func Int(name string, value int) field.Field               { return fields.Int(name, value) }
func Int32(name string, value int32) field.Field           { return fields.Int32(name, value) }
func Int16(name string, value int16) field.Field           { return fields.Int16(name, value) }
func Int8(name string, value int8) field.Field             { return fields.Int8(name, value) }
func Bool(name string, value bool) field.Field             { return fields.Bool(name, value) }
func Complex128(name string, value complex128) field.Field { return fields.Complex128(name, value) }
func Complex64(name string, value complex64) field.Field   { return fields.Complex64(name, value) }
func Float64(name string, value float64) field.Field       { return fields.Float64(name, value) }
func Float32(name string, value float32) field.Field       { return fields.Float32(name, value) }
func Nil(name string, value interface{}) field.Field       { return fields.Nil(name, value) }
func Object(name string, value interface{}) field.Field    { return fields.Object(name, value) }
func Array(name string, value interface{}) field.Field     { return fields.Array(name, value) }

func Tags(tags ...string) Writer        { return DefaultWriter.Tags(tags...) }
func Data(fields ...field.Field) Writer { return DefaultWriter.Data(fields...) }

func Trace(body string)                 { DefaultWriter.Trace(body) }
func Tracef(f string, a ...interface{}) { DefaultWriter.Tracef(f, a...) }
func Debug(body string)                 { DefaultWriter.Debug(body) }
func Debugf(f string, a ...interface{}) { DefaultWriter.Debugf(f, a...) }
func Info(body string)                  { DefaultWriter.Info(body) }
func Infof(f string, a ...interface{})  { DefaultWriter.Infof(f, a...) }
func Warn(body string)                  { DefaultWriter.Warn(body) }
func Warnf(f string, a ...interface{})  { DefaultWriter.Warnf(f, a...) }
func Error(body string)                 { DefaultWriter.Error(body) }
func Errorf(f string, a ...interface{}) { DefaultWriter.Errorf(f, a...) }
func Fatal(body string)                 { DefaultWriter.Fatal(body) }
func Fatalf(f string, a ...interface{}) { DefaultWriter.Fatalf(f, a...) }

// // ParseEnvTagFilter
// // ParseEnvLevelFilter

func init() {
	Private = field.Builder{Private: true}

	DefaultWriter = With(
		handler.BuildPipe(
			handler.BuildLevelFilter("_min"),
			handler.BuildTagFilter(""),
			handler.BuildStringIOText(os.Stderr),
		),
	)

	JSONWriter = With(
		handler.BuildPipe(
			handler.BuildLevelFilter("_min"),
			handler.BuildTagFilter(""),
			handler.BuildStringIOJSON(os.Stderr),
		),
	)

	TagsWriter = With(
		handler.BuildPipe(
			handler.BuildLevelFilter("_min"),
			handler.BuildTagFilter(""),
			handler.BuildStringIOTextConfigured(os.Stderr, transform.EverythingConfig),
		),
	)
}
