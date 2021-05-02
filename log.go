// Package log provides structured and unstructured logging functions
package log

import "github.com/thehungry-dev/log/message/field"

var Private field.Builder
var fields field.Builder

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

// var DefaultLogger Handler

// // ParseEnvTagFilter
// // ParseEnvLevelFilter

// func init() {
// 	DefaultLogger = Pipe{}
// }
func init() {
	Private = field.Builder{Private: true}
}
