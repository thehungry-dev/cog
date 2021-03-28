package filter_test

import (
	"testing"

	. "github.com/thehungry-dev/asserting"

	ctrls "github.com/thehungry-dev/log/ctrls/tag/filter"
)

func TestTagFilterSelect(t *testing.T) {
	t.Run("Tag", func(t *testing.T) {
		t.Run("Filter", func(t *testing.T) {
			t.Run("Select", func(t *testing.T) {
				tagFilter := ctrls.TagFilterExample()

				t.Run("Matches filter", func(t *testing.T) {
					t.Run("No excluded tags", func(t *testing.T) {
						tags := ctrls.TagsMatchingExample()

						selected := tagFilter.Select(tags)

						t.Run("Selected", func(t *testing.T) {
							Assert(t, selected)
						})
					})

					t.Run("Excluded tags", func(t *testing.T) {
						tags := ctrls.TagsExcludedMatchingExample()

						rejected := tagFilter.Reject(tags)

						t.Run("Rejected", func(t *testing.T) {
							Assert(t, rejected)
						})
					})
				})

				t.Run("Not matching filter", func(t *testing.T) {
					t.Run("Missing required tags", func(t *testing.T) {
						tags := ctrls.TagsNonMatchingMissingRequiredExample()

						rejected := tagFilter.Reject(tags)

						t.Run("Rejected", func(t *testing.T) {
							Assert(t, rejected)
						})
					})

					t.Run("Missing required one of tags", func(t *testing.T) {
						tags := ctrls.TagsNonMatchingMissingAllRequiredOneOfExample()

						rejected := tagFilter.Reject(tags)

						t.Run("Rejected", func(t *testing.T) {
							Assert(t, rejected)
						})
					})
				})
			})
		})
	})
}
