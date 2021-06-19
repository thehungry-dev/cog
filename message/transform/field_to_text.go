package transform

import (
	"encoding/json"
	"math"
	"strconv"
	"strings"

	"github.com/thehungry-dev/cog/message"
	"github.com/thehungry-dev/cog/message/field"
)

func FieldToText(fld field.Field) (string, bool) {
	success := true
	var output string

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
		output = strconv.FormatInt(fld.ValueNumeric, 10)
	case field.Bool:
		output = strconv.FormatBool(fld.ValueNumeric == 1)
	case field.Complex128:
		fallthrough
	case field.Complex64:
		cvalue := fld.ValueInterface.(complex128)
		output = strconv.FormatComplex(cvalue, 'f', 2, 128)
	case field.Float64:
		uvalue := uint64(fld.ValueNumeric)
		output = strconv.FormatFloat(math.Float64frombits(uvalue), 'f', 2, 64)
	case field.Float32:
		uvalue := uint32(fld.ValueNumeric)
		output = strconv.FormatFloat(float64(math.Float32frombits(uvalue)), 'f', 2, 64)
	case field.Nil:
		output = "<nil>"
	case field.Object:
		output = "<object>"
	case field.Array:
		output = "<array>"
	default:
		output = "<unrecognized>"
		success = false
	}

	return output, success
}

func FieldsToText(msg message.Message) string {
	var b strings.Builder

	var publicFieldsCount = 0

	for _, fld := range msg.Fields {
		if !fld.IsOutput() {
			continue
		}

		publicFieldsCount += 1

		if publicFieldsCount == 1 {
			b.WriteString(" (")
		}

		if publicFieldsCount > 1 {
			b.WriteString(", ")
		}

		b.WriteString(fld.Name)
		b.WriteString(": ")

		output, _ := FieldToText(fld)
		b.WriteString(output)
	}

	if publicFieldsCount > 0 {
		b.WriteString(")")
	}

	return b.String()
}

func FieldsToTextJSON(msg message.Message) string {
	if !msg.HasFields() {
		return ""
	}

	var buf strings.Builder

	jsonData := FieldsToJSONType(msg)

	output, err := json.MarshalIndent(jsonData, "", "  ")

	if err != nil {
		panic(err)
	}

	if msg.Body != "" {
		buf.WriteString("\n")
	}
	buf.Write(output)

	return buf.String()
}
