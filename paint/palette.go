package paint

import (
	"github.com/thehungry-dev/cog/color"
	"github.com/thehungry-dev/cog/level"
)

var palette map[level.Level]string

func init() {
	palette = make(map[level.Level]string, 6)
	palette[level.Trace] = color.None
	palette[level.Debug] = color.Green
	palette[level.Info] = color.White
	palette[level.Warn] = color.Yellow
	palette[level.Error] = color.Red
	palette[level.Fatal] = color.Red + color.Bold
}
