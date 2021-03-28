package benchmark

import (
	"testing"
)

type anotherStruct struct {
	OtherField        string
	SomeOtherField    int
	FinallyOtherField int
	Field             int
}

func BenchmarkStructCopy(b *testing.B) {
	b.Run("Struct Copy", func(b *testing.B) {
		b.Run("Copying", func(b *testing.B) {
			data := anotherStruct{}
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				newData := data
				newData.Field += 1
			}
		})
	})
}
