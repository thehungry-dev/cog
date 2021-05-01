package transform

import (
	"math"

	"github.com/thehungry-dev/log/message"
	"github.com/thehungry-dev/log/message/field"
)

func FieldToJSONType(fld field.Field) (interface{}, bool) {
	success := true
	var output interface{}

	switch fld.Type {
	case field.None:
	case field.String:
		output = fld.ValueString
	case field.Int64:
		fallthrough
	case field.Int32:
		fallthrough
	case field.Int16:
		fallthrough
	case field.Int8:
		output = fld.ValueNumeric
	case field.Bool:
		output = fld.ValueNumeric == 1
	case field.Complex128:
		fallthrough
	case field.Complex64:
		output = fld.ValueInterface.(complex128)
	case field.Float64:
		uvalue := uint64(fld.ValueNumeric)
		output = math.Float64frombits(uvalue)
	case field.Float32:
		uvalue := uint32(fld.ValueNumeric)
		output = float64(math.Float32frombits(uvalue))
	case field.Nil:
		output = nil
	case field.Object:
		output = fld.ValueInterface
	case field.Array:
		output = fld.ValueInterface
	default:
		output = "<unrecognized>"
		success = false
	}

	return output, success
}

func FieldsToJSONType(msg message.Message) map[string]interface{} {
	var output map[string]interface{}

	if !msg.HasFields() {
		return output
	}

	if msg.HasFields() {
		data := make(map[string]interface{}, len(msg.Fields))
		hasData := false

		for _, fld := range msg.Fields {
			if !fld.IsOutput() {
				continue
			}

			value, ok := FieldToJSONType(fld)

			if !ok {
				continue
			}

			data[fld.Name] = value
			hasData = true
		}

		if !hasData {
			return output
		}

		output = data
	}

	return output
}
