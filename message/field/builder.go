package field

import "math"

type Builder struct{ Private bool }

func (buf Builder) String(name string, value string) Field {
	return Field{Type: String, Name: name, Private: buf.Private, ValueString: value}
}
func (buf Builder) Int64(name string, value int64) Field {
	return Field{Type: Int64, Name: name, Private: buf.Private, ValueNumeric: value}
}
func (buf Builder) Int(name string, value int) Field {
	return buf.Int64(name, int64(value))
}
func (buf Builder) Int32(name string, value int32) Field {
	return Field{Type: Int32, Name: name, Private: buf.Private, ValueNumeric: int64(value)}
}
func (buf Builder) Int16(name string, value int16) Field {
	return Field{Type: Int16, Name: name, Private: buf.Private, ValueNumeric: int64(value)}
}
func (buf Builder) Int8(name string, value int8) Field {
	return Field{Type: Int8, Name: name, Private: buf.Private, ValueNumeric: int64(value)}
}
func (buf Builder) Bool(name string, value bool) Field {
	var internalValue int64
	if value {
		internalValue = 1
	}
	return Field{Type: Bool, Name: name, Private: buf.Private, ValueNumeric: internalValue}
}
func (buf Builder) Complex128(name string, value complex128) Field {
	return Field{Type: Complex128, Name: name, Private: buf.Private, ValueInterface: value}
}
func (buf Builder) Complex64(name string, value complex64) Field {
	return Field{Type: Complex64, Name: name, Private: buf.Private, ValueInterface: value}
}
func (buf Builder) Float64(name string, value float64) Field {
	return Field{Type: Float64, Name: name, Private: buf.Private, ValueNumeric: int64(math.Float64bits(value))}
}
func (buf Builder) Float32(name string, value float32) Field {
	return Field{Type: Float32, Name: name, Private: buf.Private, ValueNumeric: int64(math.Float32bits(value))}
}
func (buf Builder) Nil(name string, value interface{}) Field {
	return Field{Type: Nil, Name: name, Private: buf.Private, ValueInterface: value}
}
func (buf Builder) Object(name string, value interface{}) Field {
	return Field{Type: Object, Name: name, Private: buf.Private, ValueInterface: value}
}
func (buf Builder) Array(name string, value interface{}) Field {
	return Field{Type: Array, Name: name, Private: buf.Private, ValueInterface: value}
}
