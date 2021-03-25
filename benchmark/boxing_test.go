package benchmark

import (
	"testing"
)

var ensure interface{}

type boxed interface {
	Do()
}

type unboxed struct {
	someField int
}

type genmap map[string]interface{}
type fixedmap map[string]int

func (u unboxed) Do() {
	ensure = u.someField
}

func BenchmarkUnboxing(b *testing.B) {
	b.Run("Unboxing", func(b *testing.B) {
		b.Run("Cast", func(b *testing.B) {
			var data boxed = unboxed{0}
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				udata := data.(unboxed)
				udata.Do()
			}
		})

		b.Run("Direct", func(b *testing.B) {
			var data unboxed = unboxed{0}
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				data.Do()
			}
		})
	})
}

func BenchmarkUnboxingMap(b *testing.B) {
	b.Run("Map", func(b *testing.B) {
		b.Run("Unboxing", func(b *testing.B) {
			b.Run("Cast", func(b *testing.B) {
				data := genmap{"foo": 123, "bar": "text"}
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					ensure = data["foo"].(int)
					ensure = data["bar"].(string)
				}
			})

			b.Run("Direct", func(b *testing.B) {
				data := fixedmap{"foo": 123, "bar": 547}
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					ensure = data["foo"]
					ensure = data["bar"]
				}
			})
		})
	})
}
