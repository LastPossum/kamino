package kamino_test

import (
	"testing"

	"github.com/LastPossum/kamino"
	libdc1 "github.com/barkimedes/go-deepcopy"
	libdc3 "github.com/jinzhu/copier"
	libdc2 "github.com/mohae/deepcopy"
)

type benchSimpleStruct5 struct {
	A int
	B float64
	C string
	D byte
	E complex64
}

var benchSimpleStruct5Instance = benchSimpleStruct5{1, 2, "3", 4, 5i}

//go:noinline
func plainCopySimpleStruct5(src benchSimpleStruct5) benchSimpleStruct5 {
	return src
}

func BenchmarkPlainCopySimpleStruct5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plainCopySimpleStruct5(benchSimpleStruct5Instance)
	}
}

func BenchmarkCloneLibDC1SimpleStruct5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		libdc1.Anything(benchSimpleStruct5Instance)
	}
}

func BenchmarkCloneLibDC2SimpleStruct5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		libdc2.Copy(benchSimpleStruct5Instance)
	}
}

func BenchmarkCloneLibDC3SimpleStruct5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var res benchSimpleStruct5
		libdc3.CopyWithOption(&res, &benchSimpleStruct5Instance, libdc3.Option{IgnoreEmpty: true, DeepCopy: true})
	}
}

func BenchmarkCloneJsonSimpleStruct5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CloneJSON(benchSimpleStruct5Instance)
	}
}

func BenchmarkCloneSimpleStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kamino.Clone(benchSimpleStruct5Instance)
	}
}
