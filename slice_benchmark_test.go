package kamino_test

import (
	"testing"

	libdc1 "github.com/barkimedes/go-deepcopy"
	libdc3 "github.com/jinzhu/copier"
	libdc2 "github.com/mohae/deepcopy"

	"github.com/LastPossum/kamino"
)

const bytes = 1024

var bigByteSlice = make([]byte, bytes)

func plainCopyByteSlice(src []byte) []byte {
	r := make([]byte, len(src))
	copy(r, src)
	return r
}

func BenchmarkPlainCopySlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plainCopyByteSlice(bigByteSlice)
	}
}

func BenchmarkCloneLibDC1Slice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		libdc1.Anything(bigByteSlice)
	}
}

func BenchmarkCloneLibDC2Slice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		libdc2.Copy(bigByteSlice)
	}
}

func BenchmarkCloneLibDC3Slice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var res []byte
		libdc3.CopyWithOption(&res, &bigByteSlice, libdc3.Option{IgnoreEmpty: true, DeepCopy: true})
	}
}

func BenchmarkCloneJsonSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CloneJSON(bigByteSlice)
	}
}

func BenchmarkCloneSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kamino.Clone(bigByteSlice)
	}
}
