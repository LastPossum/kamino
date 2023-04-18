package kamino_benchmark

import (
	"fmt"
	"testing"

	barkimedes "github.com/barkimedes/go-deepcopy"
	mohae "github.com/mohae/deepcopy"

	"github.com/LastPossum/kamino"
)

func mkByteSlice(l int) []byte {
	res := make([]byte, l)
	res[0] = 'b'
	return res
}

func BenchmarkBytesSlice(b *testing.B) {
	for i := 5; i < 10; i++ {
		k := 1 << i
		bytes := mkByteSlice(k)

		b.Run(fmt.Sprintf("barkimedes for %d bytes slice", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				barkimedes.Anything(bytes)
			}
		})
	}

	for i := 5; i < 10; i++ {
		k := 1 << i
		bytes := mkByteSlice(k)

		b.Run(fmt.Sprintf("mohae for %d bytes slice", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mohae.Copy(bytes)
			}
		})
	}

	for i := 5; i < 10; i++ {
		k := 1 << i
		bytes := mkByteSlice(k)

		b.Run(fmt.Sprintf("json for %d bytes slice", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				cloneJSON(bytes)
			}
		})
	}

	for i := 5; i < 10; i++ {
		k := 1 << i
		bytes := mkByteSlice(k)

		b.Run(fmt.Sprintf("msgpack for %d bytes slice", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				cloneMsgPack(bytes)
			}
		})
	}

	for i := 5; i < 10; i++ {
		k := 1 << i
		bytes := mkByteSlice(k)

		b.Run(fmt.Sprintf("kamino for %d bytes slice", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				kamino.Clone(bytes)
			}
		})
	}
}

func BenchmarkBytesSliceKamino(b *testing.B) {
	for i := 5; i < 10; i++ {
		k := 1 << i
		bytes := mkByteSlice(k)

		b.Run(fmt.Sprintf("kamino for %d bytes slice", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				kamino.Clone(bytes)
			}
		})
	}
}
