package transform

import (
	"encoding/json"

	"github.com/thehungry-dev/log/message"
)

type messageJSON struct {
	Level   string                 `json:"level"`
	Message string                 `json:"message"`
	Tags    []string               `json:"tags"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

func ToJSON(msg message.Message) string {
	jsonMsg := messageJSON{}
	jsonMsg.Level = msg.Level.String()
	jsonMsg.Message = msg.Body
	jsonMsg.Tags = msg.Tags

	if msg.HasFields() {
		data := make(map[string]interface{}, len(msg.Fields))

		for _, fld := range msg.Fields {
			if !fld.IsOutput() {
				continue
			}

			value, ok := FieldToJSONType(fld)

			if !ok {
				continue
			}

			data[fld.Name] = value
		}

		jsonMsg.Data = data
	}

	output, err := json.Marshal(jsonMsg)

	if err != nil {
		panic(err)
	}

	return string(output)
}
