// Package field represents a single message field for a log message
package field

type Type uint8

type Field struct {
	Type           Type
	Name           string
	ValueString    string
	ValueNumeric   int64
	ValueInterface interface{}
}

const (
	None Type = iota
	String
	Int64
	Int32
	Int16
	Int8
	Bool
	Complex128
	Complex64
	Float64
	Float32
	Nil
)
