// Package filter provides controls for test environment
package filter

import (
	"fmt"

	"github.com/thehungry-dev/log/tag/filter"
)

func TagNameOneOfExample() string    { return "aTag" }
func TagNameRequiredExample() string { return "requiredTag" }
func TagNameExcludedExample() string { return "excludedTag" }

func OneOfTagExample() string    { return TagNameOneOfExample() }
func RequiredTagExample() string { return fmt.Sprintf("+%s", TagNameRequiredExample()) }
func ExcludeTagExample() string  { return fmt.Sprintf("-%s", TagNameExcludedExample()) }

func TagsMatchingExample() []string { return []string{TagNameOneOfExample(), TagNameRequiredExample()} }
func TagsExcludedMatchingExample() []string {
	return []string{TagNameOneOfExample(), TagNameRequiredExample(), TagNameExcludedExample()}
}

func TagsNonMatchingMissingRequiredExample() []string { return []string{TagNameOneOfExample()} }
func TagsNonMatchingMissingAllRequiredOneOfExample() []string {
	return []string{TagNameRequiredExample()}
}

func StringExample() string {
	return fmt.Sprintf("%s,%s,%s", OneOfTagExample(), RequiredTagExample(), ExcludeTagExample())
}

func TagFilterExample() *filter.TagFilter {
	return filter.Parse(StringExample())
}
func TagFilterNothingExample() *filter.TagFilter {
	return filter.Nothing
}
