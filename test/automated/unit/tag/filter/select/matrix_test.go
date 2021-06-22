package filter_test

import (
	_ "embed"
	"testing"

	"github.com/thehungry-dev/cog/cogtest"
)

//go:embed matrix.csv
var matrixFile string

func TestTagFilterSelectMatrix(t *testing.T) {
	t.Run("Tag", func(t *testing.T) {
		t.Run("Filter", func(t *testing.T) {
			t.Run("Select", func(t *testing.T) {
				t.Run("Matrix", func(t *testing.T) {
					cogtest.AssertMatrix(t, matrixFile)
				})
			})
		})
	})
}
