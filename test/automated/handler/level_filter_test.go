package handler_test

import (
	"testing"

	. "github.com/thehungry-dev/asserting"

	ctrls "github.com/thehungry-dev/log/ctrls/handler"
	msgctrls "github.com/thehungry-dev/log/ctrls/message"
)

func TestLevelFilterHandler(t *testing.T) {
	t.Run("Level Filter Handler", func(t *testing.T) {
		t.Run("Matching", func(t *testing.T) {
			handler := ctrls.LevelFilterExample()
			msg := msgctrls.MessageExample()

			Assert(t, !msg.IsHalted())

			msg = handler.Handle(msg)

			t.Run("Not halted", func(t *testing.T) {
				Assert(t, !msg.IsHalted())
			})
		})

		t.Run("Not Matching", func(t *testing.T) {
			handler := ctrls.LevelFilterExample()
			msg := msgctrls.MessageLevelExcludedExample()

			Assert(t, !msg.IsHalted())

			msg = handler.Handle(msg)

			t.Run("Halted", func(t *testing.T) {
				Assert(t, msg.IsHalted())
			})
		})
	})
}
