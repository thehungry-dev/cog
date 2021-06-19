package transform

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/thehungry-dev/cog/message"
)

type messageJSON struct {
	Timestamp string                 `json:"timestamp"`
	Level     string                 `json:"level"`
	Message   string                 `json:"message"`
	Tags      []string               `json:"tags"`
	Fields    map[string]interface{} `json:"data,omitempty"`
}

func ToJSON(msg message.Message) string {
	jsonMsg := messageJSON{}
	jsonMsg.Timestamp = msg.Timestamp.Format(time.RFC3339)
	jsonMsg.Level = msg.Level.String()
	jsonMsg.Message = msg.Body
	jsonMsg.Tags = msg.Tags
	jsonMsg.Fields = FieldsToJSONType(msg)

	output, err := json.Marshal(jsonMsg)

	if err != nil {
		panic(err)
	}

	var buf strings.Builder

	buf.Write(output)
	buf.WriteString("\n")

	return buf.String()
}
