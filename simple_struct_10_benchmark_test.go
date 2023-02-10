package kamino_test

import (
	"testing"

	"github.com/LastPossum/kamino"
	barkimedes "github.com/barkimedes/go-deepcopy"
	jinzhu "github.com/jinzhu/copier"
	mohae "github.com/mohae/deepcopy"
)

type benchSimpleStruct10 struct {
	A  int
	B  float64
	C  string
	D  byte
	E  complex64
	A1 int
	B2 float64
	C3 string
	D4 byte
	E5 complex64
}

var benchSimpleStruct10Instance = benchSimpleStruct10{1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i}

//go:noinline
func plainCopySimpleStruct10(src benchSimpleStruct10) benchSimpleStruct10 {
	return src
}

func BenchmarkPlainCopySimpleStruct10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plainCopySimpleStruct10(benchSimpleStruct10Instance)
	}
}

func BenchmarkCloneBarkimedesSimpleStruct10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		barkimedes.Anything(benchSimpleStruct10Instance)
	}
}

func BenchmarkCloneMohaeSimpleStruct10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mohae.Copy(benchSimpleStruct10Instance)
	}
}

func BenchmarkCloneJinzhuSimpleStruct10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var res benchSimpleStruct10
		jinzhu.CopyWithOption(&res, &benchSimpleStruct10Instance, jinzhu.Option{IgnoreEmpty: true, DeepCopy: true})
	}
}

func BenchmarkCloneJsonSimpleStruct10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CloneJSON(benchSimpleStruct10Instance)
	}
}

func BenchmarkCloneMsgPackSimpleStruct10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cloneMsgPack(benchSimpleStruct10Instance)
	}
}

func BenchmarkCloneSimpleStruct10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kamino.Clone(benchSimpleStruct10Instance)
	}
}
