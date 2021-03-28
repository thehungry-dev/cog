package filter_test

import (
	"testing"

	. "github.com/thehungry-dev/asserting"

	ctrls "github.com/thehungry-dev/log/ctrls/tag/filter"
	"github.com/thehungry-dev/log/tag/filter"
)

func TestTagFilterParse(t *testing.T) {
	t.Run("Tag", func(t *testing.T) {
		t.Run("Filter", func(t *testing.T) {
			t.Run("Parse", func(t *testing.T) {
				tagFilterString := ctrls.StringExample()

				t.Run("One of tags", func(t *testing.T) {
					tagFilter := filter.Parse(tagFilterString)
					oneOfTag := ctrls.OneOfTagExample()

					t.Run("Included", func(t *testing.T) {
						Assert(t, tagFilter.IsRequiredOneOfTag(oneOfTag))
					})
				})

				t.Run("Required tags", func(t *testing.T) {
					tagFilter := filter.Parse(tagFilterString)
					requiredTag := ctrls.RequiredTagExample()

					t.Run("Included", func(t *testing.T) {
						Assert(t, tagFilter.IsRequiredTag(requiredTag))
					})
				})

				t.Run("Excluded tags", func(t *testing.T) {
					tagFilter := filter.Parse(tagFilterString)
					excludeTag := ctrls.ExcludeTagExample()

					t.Run("Included", func(t *testing.T) {
						Assert(t, tagFilter.IsExcludedTag(excludeTag))
					})
				})
			})
		})
	})
}
