package kamino_benchmark

import (
	"testing"

	"github.com/LastPossum/kamino"
	barkimedes "github.com/barkimedes/go-deepcopy"
	jinzhu "github.com/jinzhu/copier"
	mohae "github.com/mohae/deepcopy"
)

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

func BenchmarkCloneBarkimedesSimpleStruct5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		barkimedes.Anything(benchSimpleStruct5Instance)
	}
}

func BenchmarkCloneMohaeSimpleStruct5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mohae.Copy(benchSimpleStruct5Instance)
	}
}

func BenchmarkCloneJinzhuSimpleStruct5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var res benchSimpleStruct5
		jinzhu.CopyWithOption(&res, &benchSimpleStruct5Instance, jinzhu.Option{IgnoreEmpty: true, DeepCopy: true})
	}
}

func BenchmarkCloneJsonSimpleStruct5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cloneJSON(benchSimpleStruct5Instance)
	}
}

func BenchmarkCloneMsgPackSimpleStruct5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cloneMsgPack(benchSimpleStruct5Instance)
	}
}

func BenchmarkCloneSimpleStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kamino.Clone(benchSimpleStruct5Instance)
	}
}

type singleStruct struct {
	i int64
	// i2 int64
}

func BenchmarkClone1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mohae.Copy(singleStruct{1})
	}
}

func BenchmarkClone1_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kamino.Clone(singleStruct{1})
	}
}
