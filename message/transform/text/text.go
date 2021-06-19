// Package text provides utilities to parse back a message converted to text into machine-parsable data
package text

import (
	"regexp"
	"strings"
	"time"

	"github.com/thehungry-dev/cog/level"
)

type MessageData struct {
	Timestamp time.Time
	Tags      []string
	Level     level.Level
	Body      string
}

func (data MessageData) FieldsText(body string) string {
	return strings.TrimPrefix(data.Body, body)
}

var uncoloredRegex *regexp.Regexp

func ParseUncolored(text string) MessageData {
	var data MessageData
	matches := uncoloredRegex.FindStringSubmatch(text)

	data.Timestamp = ParseTimestamp(matches[1])
	data.Tags = ParseTags(matches[2])
	data.Level = ParseLevel(matches[3])
	data.Body = matches[4]

	return data
}

func ParseTimestamp(text string) time.Time {
	timestamp, err := time.Parse(time.RFC3339, text)

	if err != nil {
		return time.Time{}
	}

	return timestamp
}

func ParseTags(text string) []string {
	return strings.Split(text, ", ")
}

func ParseLevel(text string) level.Level {
	lvl, err := level.Parse(text)

	if err != nil {
		return level.Trace
	}

	return lvl
}

func init() {
	uncoloredRegex = regexp.MustCompile(`(?s)(?:\[(?P<timestamp>.+)\]\s)?(?:\|(?P<tags>.+)\|\s)?(?:(?P<level>TRACE|DEBUG|INFO|WARN|ERROR|FATAL)\s+)?(?P<body>.+)`)
}
