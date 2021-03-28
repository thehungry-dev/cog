// Package filter provides a data structure to filter data based on tags
package filter

import (
	"strings"

	"github.com/thehungry-dev/log/tag"
)

const TagSeparator = ","
const TagRequiredToken = "+"
const TagExcludedToken = "-"
const AllTagsKeyword = "_all"
const UntaggedKeyword = "_untagged"

var Nothing *TagFilter

type TagFilter struct {
	requiredTags tag.Set
	oneOfTags    tag.Set
	excludedTags tag.Set
	untagged     bool
}

func Parse(filterText string) *TagFilter {
	text := strings.TrimSpace(filterText)

	if text == "" || text == AllTagsKeyword {
		return Nothing
	}

	tags := strings.Split(text, TagSeparator)
	maxTags := len(tags)

	requiredTags := tag.BuildSetOfSize(maxTags)
	oneOfTags := tag.BuildSetOfSize(maxTags)
	excludedTags := tag.BuildSetOfSize(maxTags)
	filter := &TagFilter{requiredTags, oneOfTags, excludedTags, false}

	for _, tag := range tags {
		tag = strings.TrimSpace(tag)

		if tag == "" {
			continue
		}

		switch {
		case tag == UntaggedKeyword:
			filter.untagged = true
		case strings.HasPrefix(tag, TagRequiredToken):
			requiredTags.Add(tag)
		case strings.HasPrefix(tag, TagExcludedToken):
			excludedTags.Add(tag)
		default:
			oneOfTags.Add(tag)
		}
	}

	return filter
}

func (tagFilter *TagFilter) IsRequiredTag(tag string) bool {
	return tagFilter.requiredTags.Include(tag)
}
func (tagFilter *TagFilter) IsExcludedTag(tag string) bool {
	return tagFilter.excludedTags.Include(tag)
}
func (tagFilter *TagFilter) IsRequiredOneOfTag(tag string) bool {
	return tagFilter.oneOfTags.Include(tag)
}
