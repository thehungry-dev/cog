package filter_test

import (
	"testing"

	. "github.com/thehungry-dev/asserting"

	ctrls "github.com/thehungry-dev/cog/ctrls/tag/filter"
)

func TestTagFilterNothingSelect(t *testing.T) {
	t.Run("Tag", func(t *testing.T) {
		t.Run("Filter", func(t *testing.T) {
			t.Run("Nothing", func(t *testing.T) {
				tagFilter := ctrls.TagFilterNothingExample()
				tags := ctrls.TagsMatchingExample()

				selected := tagFilter.Select(tags)

				t.Run("Selected", func(t *testing.T) {
					Assert(t, selected)
				})
			})
		})
	})
}
