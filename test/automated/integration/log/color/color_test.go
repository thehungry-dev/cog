package color_test

import (
	"testing"

	. "github.com/thehungry-dev/asserting"

	"github.com/thehungry-dev/cog/color"
	ctrls "github.com/thehungry-dev/cog/ctrls/log"
)

func TestLogColor(t *testing.T) {
	t.Run("Log", func(t *testing.T) {
		t.Run("Color", func(t *testing.T) {
			t.Parallel()

			writer := ctrls.LogColoredTextDeviceExample()

			t.Run("Trace", func(t *testing.T) {
				writer.Trace(ctrls.BodyExample)

				uncolored := color.IsUncolored(writer.Last)

				t.Run("None", func(t *testing.T) {
					Assert(t, uncolored)
				})
			})

			t.Run("Debug", func(t *testing.T) {
				writer.Debug(ctrls.BodyExample)

				colored := color.IsApplied(writer.Last, color.Green)

				t.Run("Green", func(t *testing.T) {
					Assert(t, colored)
				})
			})

			t.Run("Info", func(t *testing.T) {
				writer.Info(ctrls.BodyExample)

				colored := color.IsApplied(writer.Last, color.White)

				t.Run("White", func(t *testing.T) {
					Assert(t, colored)
				})
			})

			t.Run("Warn", func(t *testing.T) {
				writer.Warn(ctrls.BodyExample)

				colored := color.IsApplied(writer.Last, color.Yellow)

				t.Run("Yellow", func(t *testing.T) {
					Assert(t, colored)
				})
			})

			t.Run("Error", func(t *testing.T) {
				writer.Error(ctrls.BodyExample)

				colored := color.IsApplied(writer.Last, color.Red)

				t.Run("Red", func(t *testing.T) {
					Assert(t, colored)
				})
			})

			t.Run("Fatal", func(t *testing.T) {
				writer.Fatal(ctrls.BodyExample)

				colored := color.IsApplied(writer.Last, color.Red+color.Bold)

				t.Run("Red + Bold", func(t *testing.T) {
					Assert(t, colored)
				})
			})
		})
	})
}
