# Cog

Cog provides the building blocks for creating the logging mechanism
tailored to the need of your software.

## Usage

```go
package mypkg

import (
	"github.com/thehungry-dev/cog"
	"github.com/thehungry-dev/cog/handler"
	"github.com/thehungry-dev/cog/message/field"
	"github.com/thehungry-dev/cog/message/transform"
	"github.com/thehungry-dev/cog/settings"
)

var Log cog.Writer

func init() {
	set := settings.Getenv()
	device := set.Device()

  // Builds a writer that can filter log output by level, by tag and outputs colored log messages
	Log = cog.With(
		handler.BuildPipe(
			set.LevelFilter(),
			set.TagFilter(),
			handler.BuildOutput(device, handler.OutputText),
		),
	)
}
```

```go
package main

func main() {
  log := mypkg.Log
  fields := field.Builder{}
  log.Info("some message")
  log.Tags("tag1", "tag2").Warn("some message")
  log.
    Tags("tag1",  "data").
    Data(
      fields.String("Name", "some name"),
      fields.Bool("SomeBool", true)
    ).
    Fatalf("Interpolated %s", "message")
}
```
