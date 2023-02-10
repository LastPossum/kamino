package kamino_test

import (
	"testing"

	"github.com/LastPossum/kamino"
	libdc1 "github.com/barkimedes/go-deepcopy"
	libdc3 "github.com/jinzhu/copier"
	libdc2 "github.com/mohae/deepcopy"
)

type nestedStructs struct {
	L1 struct {
		L2 struct {
			L3 struct {
				base *int
			}
		}
	}
}

var nestedStructsInstance = nestedStructs{
	L1: struct {
		L2 struct{ L3 struct{ base *int } }
	}{
		L2: struct{ L3 struct{ base *int } }{
			L3: struct{ base *int }{
				ptrTo(20),
			},
		},
	},
}

func BenchmarkCloneLibDC1NestedStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		libdc1.Anything(nestedStructsInstance)
	}
}

func BenchmarkCloneLibDC2NestedStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		libdc2.Copy(nestedStructsInstance)
	}
}

func BenchmarkCloneLibDC3NestedStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var res nestedStructs
		libdc3.CopyWithOption(&res, &nestedStructsInstance, libdc3.Option{IgnoreEmpty: true, DeepCopy: true})
	}
}

func BenchmarkCloneJsonNestedStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CloneJSON(nestedStructsInstance)
	}
}

func BenchmarkCloneNestedStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kamino.Clone(nestedStructsInstance)
	}
}
