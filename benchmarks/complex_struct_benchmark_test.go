package kamino_benchmark

import (
	"testing"

	barkimedes "github.com/barkimedes/go-deepcopy"
	mohae "github.com/mohae/deepcopy"

	"github.com/LastPossum/kamino"
)

type barer interface {
	Bar() int
}

type simpleStructBenchmark struct {
	Int     int
	Float64 float64
	String  string
}

func (s *simpleStructBenchmark) Bar() int {
	return s.Int
}

type complexStructForBench struct {
	Int     int
	Float64 float64
	String  string

	Struct    simpleStructBenchmark
	Interface barer

	PointerToInt     *int
	PointerToFloat64 *float64
	PointerToString  *string
	PointerToStruct  *simpleStructBenchmark

	ArrayOfInt          [10]int
	ArrayOfSimpleStruct [5]simpleStructBenchmark

	SliceOfInt          []int
	SliceOfSimpleStruct []simpleStructBenchmark
	SliceOfPtrsToInt    []*int
	SliceOfStructPtrs   []*simpleStructBenchmark
}

func simpleStructForBenchCp() simpleStructBenchmark {
	return simpleStructForBenchInstance
}

var (
	simpleStructForBenchInstance = simpleStructBenchmark{
		Int:     1,
		Float64: 2,
		String:  "3",
	}
	complexStructForBenchInstance = complexStructForBench{
		Int:                 1,
		Float64:             2,
		String:              "3",
		Struct:              simpleStructForBenchInstance,
		Interface:           ptrTo(simpleStructForBenchCp()),
		PointerToInt:        ptrTo(4),
		PointerToFloat64:    ptrTo(5.0),
		PointerToString:     ptrTo("6"),
		PointerToStruct:     ptrTo(simpleStructForBenchCp()),
		ArrayOfInt:          [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		ArrayOfSimpleStruct: [...]simpleStructBenchmark{simpleStructForBenchInstance, simpleStructForBenchInstance, simpleStructForBenchInstance, simpleStructForBenchInstance, simpleStructForBenchInstance},
		SliceOfInt:          []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		SliceOfSimpleStruct: []simpleStructBenchmark{simpleStructForBenchInstance, simpleStructForBenchInstance, simpleStructForBenchInstance, simpleStructForBenchInstance, simpleStructForBenchInstance},
		SliceOfPtrsToInt:    []*int{ptrTo(1), ptrTo(2), ptrTo(3), ptrTo(4), ptrTo(5)},
		SliceOfStructPtrs:   []*simpleStructBenchmark{ptrTo(simpleStructForBenchCp()), ptrTo(simpleStructForBenchCp()), ptrTo(simpleStructForBenchCp())},
	}
)

func BenchmarkCloneBarkimedesComplexStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		barkimedes.Anything(complexStructForBenchInstance)
	}
}

func BenchmarkCloneMohaeComplexStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mohae.Copy(complexStructForBenchInstance)
	}
}

func BenchmarkCloneJsonComplexStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cloneJSON(complexStructForBenchInstance)
	}
}

func BenchmarkCloneKaminoComplexStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kamino.Clone(complexStructForBenchInstance)
	}
}
