// Package filter provides controls for test environment
package filter

import "fmt"

func OneOfTagExample() string    { return "aTag" }
func RequiredTagExample() string { return "+requiredTag" }
func ExcludeTagExample() string  { return "-excludedTag" }

func StringExample() string {
	return fmt.Sprintf("%s,%s,%s", OneOfTagExample(), RequiredTagExample(), ExcludeTagExample())
}
