package kamino_test

import (
	"strconv"
	"testing"

	libdc1 "github.com/barkimedes/go-deepcopy"
	libdc3 "github.com/jinzhu/copier"
	libdc2 "github.com/mohae/deepcopy"

	"github.com/LastPossum/kamino"
)

const lenStructSlice = 1024

var bigStructSlice = genSliceOfStruct()

func genSliceOfStruct() []simpleStruct {
	bigStructSlice := make([]simpleStruct, lenStructSlice)
	for i := 0; i < lenStructSlice; i++ {
		bigStructSlice[i].Int = i
		bigStructSlice[i].Float64 = float64(i)
		bigStructSlice[i].String = strconv.Itoa(i)
	}
	return bigStructSlice
}

func BenchmarkCloneLibDC1SliceOfStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		libdc1.Anything(bigStructSlice)
	}
}

func BenchmarkCloneLibDC2SliceOfStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		libdc2.Copy(bigStructSlice)
	}
}

func BenchmarkCloneLibDC3SliceOfStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var res []benchSimpleStruct5
		libdc3.CopyWithOption(&res, &bigStructSlice, libdc3.Option{IgnoreEmpty: true, DeepCopy: true})
	}
}

func BenchmarkCloneJsonSliceOfStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CloneJSON(bigStructSlice)
	}
}

func BenchmarkCloneSliceOfStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kamino.Clone(bigStructSlice)
	}
}
