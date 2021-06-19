package filter_test

import (
	"testing"

	. "github.com/thehungry-dev/asserting"

	ctrls "github.com/thehungry-dev/cog/ctrls/level/filter"
	"github.com/thehungry-dev/cog/level"
	"github.com/thehungry-dev/cog/level/filter"
)

func TestLevelFilterParse(t *testing.T) {
	t.Run("Level", func(t *testing.T) {
		t.Run("Filter", func(t *testing.T) {
			t.Run("Parse", func(t *testing.T) {
				t.Run("Invalid", func(t *testing.T) {
					levelFilterString := ctrls.InvalidLevelStringExample()

					_, ok := filter.Parse(levelFilterString)

					t.Run("Not Parsed", func(t *testing.T) {
						Assert(t, !ok)
					})
				})

				t.Run("Trace", func(t *testing.T) {
					levelFilterString := filter.TraceKeyword

					levelFilter, ok := filter.Parse(levelFilterString)

					t.Run("Parsed", func(t *testing.T) {
						Assert(t, ok)
					})

					selected := levelFilter.Select(level.Trace)

					t.Run("Select Trace", func(t *testing.T) {
						Assert(t, selected)
					})
				})

				t.Run("Debug", func(t *testing.T) {
					levelFilterString := filter.DebugKeyword

					levelFilter, ok := filter.Parse(levelFilterString)

					t.Run("Parsed", func(t *testing.T) {
						Assert(t, ok)
					})

					selected := levelFilter.Select(level.Debug)

					t.Run("Select Debug", func(t *testing.T) {
						Assert(t, selected)
					})

					rejected := levelFilter.Reject(level.Trace)

					t.Run("Reject Trace", func(t *testing.T) {
						Assert(t, rejected)
					})
				})

				t.Run("Info", func(t *testing.T) {
					levelFilterString := filter.InfoKeyword

					levelFilter, ok := filter.Parse(levelFilterString)

					t.Run("Parsed", func(t *testing.T) {
						Assert(t, ok)
					})

					selected := levelFilter.Select(level.Info)

					t.Run("Select Info", func(t *testing.T) {
						Assert(t, selected)
					})

					rejected := levelFilter.Reject(level.Debug)

					t.Run("Reject Debug", func(t *testing.T) {
						Assert(t, rejected)
					})
				})

				t.Run("Warn", func(t *testing.T) {
					levelFilterString := filter.WarnKeyword

					levelFilter, ok := filter.Parse(levelFilterString)

					t.Run("Parsed", func(t *testing.T) {
						Assert(t, ok)
					})

					selected := levelFilter.Select(level.Warn)

					t.Run("Select Warn", func(t *testing.T) {
						Assert(t, selected)
					})

					rejected := levelFilter.Reject(level.Info)

					t.Run("Reject Info", func(t *testing.T) {
						Assert(t, rejected)
					})
				})

				t.Run("Error", func(t *testing.T) {
					levelFilterString := filter.ErrorKeyword

					levelFilter, ok := filter.Parse(levelFilterString)

					t.Run("Parsed", func(t *testing.T) {
						Assert(t, ok)
					})

					selected := levelFilter.Select(level.Error)

					t.Run("Select Error", func(t *testing.T) {
						Assert(t, selected)
					})

					rejected := levelFilter.Reject(level.Info)

					t.Run("Reject Info", func(t *testing.T) {
						Assert(t, rejected)
					})
				})

				t.Run("Fatal", func(t *testing.T) {
					levelFilterString := filter.FatalKeyword

					levelFilter, ok := filter.Parse(levelFilterString)

					t.Run("Parsed", func(t *testing.T) {
						Assert(t, ok)
					})

					selected := levelFilter.Select(level.Fatal)

					t.Run("Select Fatal", func(t *testing.T) {
						Assert(t, selected)
					})

					rejected := levelFilter.Reject(level.Error)

					t.Run("Reject Error", func(t *testing.T) {
						Assert(t, rejected)
					})
				})
			})
		})
	})
}
