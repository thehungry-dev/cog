package benchmark

import (
	"testing"
)

type anInterface interface {
	Copy() anInterface
}

type aStruct struct {
	OtherField        string
	SomeOtherField    int
	FinallyOtherField int
	Field             int
}

func (a aStruct) Copy() anInterface {
	ensure = a
	a.Field += 1
	return a
}

type aPtrStruct struct {
	OtherField        string
	SomeOtherField    int
	FinallyOtherField int
	Field             int
}

func (a *aPtrStruct) Copy() anInterface {
	ensure = a
	a.Field += 1
	return a
}

func BenchmarkInterfaceCopy(b *testing.B) {
	b.Run("Interface Copy", func(b *testing.B) {
		b.Run("Copying", func(b *testing.B) {
			var data anInterface = aStruct{}
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				udata := data.Copy()
				data = udata.Copy()
			}
		})

		b.Run("Modifying", func(b *testing.B) {
			var data anInterface = &aPtrStruct{}
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				udata := data.Copy()
				data = udata.Copy()
			}
		})
	})
}
