package handler_test

import (
	"testing"

	. "github.com/thehungry-dev/asserting"

	ctrls "github.com/thehungry-dev/cog/ctrls/handler"
	msgctrls "github.com/thehungry-dev/cog/ctrls/message"
)

func TestTagFilterHandler(t *testing.T) {
	t.Run("Tag Filter Handler", func(t *testing.T) {
		t.Run("Matching", func(t *testing.T) {
			handler := ctrls.TagFilterExample()
			msg := msgctrls.MessageExample()

			Assert(t, !msg.IsHalted())

			msg = handler.Handle(msg)

			t.Run("Not halted", func(t *testing.T) {
				Assert(t, !msg.IsHalted())
			})
		})

		t.Run("Not Matching", func(t *testing.T) {
			handler := ctrls.TagFilterExample()
			msg := msgctrls.MessageExcludedExample()

			Assert(t, !msg.IsHalted())

			msg = handler.Handle(msg)

			t.Run("Halted", func(t *testing.T) {
				Assert(t, msg.IsHalted())
			})
		})
	})
}
