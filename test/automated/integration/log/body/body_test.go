package body_test

import (
	"reflect"
	"strings"
	"testing"

	. "github.com/thehungry-dev/asserting"

	ctrls "github.com/thehungry-dev/cog/ctrls/log"
	"github.com/thehungry-dev/cog/level"
	"github.com/thehungry-dev/cog/message/transform/text"
)

func TestLogBody(t *testing.T) {
	t.Run("Log", func(t *testing.T) {
		t.Run("Body", func(t *testing.T) {
			t.Parallel()

			writer := ctrls.LogTaggedDataTextDeviceExample()

			t.Run("Attributes", func(t *testing.T) {
				t.Parallel()

				writer.Trace(ctrls.BodyExample)

				data := text.ParseUncolored(writer.Last)

				t.Logf("Log message:\n%+v", writer.Last)
				t.Logf("Parsed message:\n%+v", data)

				t.Run("Timestamp", func(t *testing.T) {
					isSet := !data.Timestamp.IsZero()

					t.Run("Set", func(t *testing.T) {
						Assert(t, isSet)
					})
				})

				t.Run("Tags", func(t *testing.T) {
					isEqual := reflect.DeepEqual(data.Tags, ctrls.TagsExample)

					t.Run("Set", func(t *testing.T) {
						Assert(t, isEqual)
					})
				})

				t.Run("Level", func(t *testing.T) {
					t.Parallel()

					t.Run("Trace", func(t *testing.T) {
						writer.Trace(ctrls.BodyExample)
						lvlData := text.ParseUncolored(writer.Last)

						Assert(t, lvlData.Level == level.Trace)
					})
					t.Run("Debug", func(t *testing.T) {
						writer.Debug(ctrls.BodyExample)
						lvlData := text.ParseUncolored(writer.Last)

						Assert(t, lvlData.Level == level.Trace)
					})
					t.Run("Info", func(t *testing.T) {
						writer.Info(ctrls.BodyExample)
						lvlData := text.ParseUncolored(writer.Last)

						Assert(t, lvlData.Level == level.Trace)
					})
					t.Run("Warn", func(t *testing.T) {
						writer.Warn(ctrls.BodyExample)
						lvlData := text.ParseUncolored(writer.Last)

						Assert(t, lvlData.Level == level.Trace)
					})
					t.Run("Error", func(t *testing.T) {
						writer.Error(ctrls.BodyExample)
						lvlData := text.ParseUncolored(writer.Last)

						Assert(t, lvlData.Level == level.Trace)
					})
					t.Run("Fatal", func(t *testing.T) {
						writer.Fatal(ctrls.BodyExample)
						lvlData := text.ParseUncolored(writer.Last)

						Assert(t, lvlData.Level == level.Trace)
					})
				})

				t.Run("Body", func(t *testing.T) {
					isMsgBody := strings.HasPrefix(data.Body, ctrls.BodyExample)

					t.Run("Begin with message body", func(t *testing.T) {
						Assert(t, isMsgBody)
					})
				})

				t.Run("Fields", func(t *testing.T) {
					fieldsText := data.FieldsText(ctrls.BodyExample)
					areIncluded := true

					for _, field := range ctrls.DataExample {
						areIncluded = areIncluded && strings.Contains(fieldsText, field.Name)
					}

					t.Run("Included in body", func(t *testing.T) {
						Assert(t, areIncluded)
					})
				})
			})
		})
	})
}
