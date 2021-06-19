// Package cog provides the building blocks to build the logging mechanism tailored to the need of your software
package cog

import (
	"github.com/thehungry-dev/cog/handler"
	"github.com/thehungry-dev/cog/message/transform"
	"github.com/thehungry-dev/cog/settings"
)

var DefaultWriter Writer
var TagsWriter Writer
var JSONWriter Writer

func init() {
	set := settings.Getenv()
	device := set.Device()

	DefaultWriter = With(
		handler.BuildPipe(
			set.LevelFilter(),
			set.TagFilter(),
			handler.BuildOutput(device, handler.OutputText),
		),
	)

	JSONWriter = With(
		handler.BuildPipe(
			set.LevelFilter(),
			set.TagFilter(),
			handler.BuildOutput(device, handler.OutputJSON),
		),
	)

	TagsWriter = With(
		handler.BuildPipe(
			set.LevelFilter(),
			set.TagFilter(),
			handler.Output{
				Device: device,
				Format: handler.OutputText,
				Config: transform.EverythingConfig,
			},
		),
	)
}
