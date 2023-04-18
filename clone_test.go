package kamino_test

import (
	"fmt"
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/assert"

	"github.com/LastPossum/kamino"
)

func ExampleClone() {
	tests := []any{
		200000,
		`units are ready, with a`,
		1000000,
		[]string{
			"more well",
			"on the way.",
		},
		map[string]any{
			"I don't like sand": "It's coarse and rough and irritating and it gets everywhere.",
		},
	}

	for _, expected := range tests {
		actual, err := kamino.Clone(expected, kamino.WithErrOnUnsupported())
		if err != nil {
			fmt.Println("got error:", err)
		}
		fmt.Println(actual)
	}
	// Output:
	// 200000
	// units are ready, with a
	// 1000000
	// [more well on the way.]
	// map[I don't like sand:It's coarse and rough and irritating and it gets everywhere.]
}

type simpleStruct struct {
	Int     int
	Float64 float64
	String  string
}

type alltogether struct {
	Bool                bool
	Int                 int
	Float64             float64
	String              string
	ArrayOfInt          [10]int
	ArrayOfSimpleStruct [5]simpleStruct
	SliceOfInt          []int
	SliceOfSimpleStruct []simpleStruct
	Nested              simpleStruct

	IntPtrs    []*int
	StructPtrs []*simpleStruct
}

var (
	intInstance1 = 1114
	intInstance2 = 1387

	simpleStructIntance = simpleStruct{-1, -1, "-1"}

	alltogetherInstance = alltogether{
		Bool:       true,
		Int:        10,
		Float64:    20.,
		String:     "30",
		ArrayOfInt: [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		ArrayOfSimpleStruct: [5]simpleStruct{
			{1, 1, "1"},
			{2, 2, "2"},
			{3, 3, "3"},
			{4, 4, "4"},
			{5, 5, "5"},
		},
		SliceOfInt: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		SliceOfSimpleStruct: []simpleStruct{
			{1, 1, "1"},
			{2, 2, "2"},
			{3, 3, "3"},
			{4, 4, "4"},
			{5, 5, "5"},
		},
		Nested: simpleStruct{
			1, 2, "3",
		},
		IntPtrs:    []*int{&intInstance1, &intInstance2, &intInstance1},
		StructPtrs: []*simpleStruct{&simpleStructIntance, &simpleStructIntance},
	}
)

func TestClone2(t *testing.T) {
	t.Run("primitive types", func(t *testing.T) {
		var (
			i64  int64      = 1
			f64  float64    = 1.0
			b               = true
			c128 complex128 = 4i + 1
			s               = "blah-blah"
		)

		goti64, _ := kamino.Clone(i64)
		assert.Equal(t, i64, goti64)

		gotf64, _ := kamino.Clone(f64)
		assert.Equal(t, f64, gotf64)

		gotb, _ := kamino.Clone(b)
		assert.Equal(t, b, gotb)

		gotc128, _ := kamino.Clone(c128)
		assert.Equal(t, c128, gotc128)

		gots, _ := kamino.Clone(s)
		assert.Equal(t, s, gots)
	})

	t.Run("arrays", func(t *testing.T) {
		arrInts := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

		arrIntsClone, _ := kamino.Clone(arrInts)
		assert.Equal(t, arrInts, arrIntsClone)

		arrAny := [...]any{0, 1, 2, 3, nil, 5, 6, &simpleStructIntance, 8, 9, 10, 11, 12, 13, 14, 15}

		gotArrAny, _ := kamino.Clone(arrAny)
		assert.Equal(t, arrAny, gotArrAny)
	})

	t.Run("slices", func(t *testing.T) {
		sliceInts := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

		sliceIntsClone, _ := kamino.Clone(sliceInts)
		assert.Equal(t, sliceInts, sliceIntsClone)

		sliceAny := []any{0, 1, 2, 3, nil, 5, 6, &simpleStructIntance, 8, 9, 10, 11, 12, 13, 14, 15}

		gotSliceAny, _ := kamino.Clone(sliceAny)
		assert.Equal(t, sliceAny, gotSliceAny)

		sliceOfSlice := [][]int{{1, 2, 4}, {4, 5}, {6}}
		gotSliceOfSlice, _ := kamino.Clone(sliceOfSlice)
		assert.Equal(t, sliceOfSlice, gotSliceOfSlice)
		sliceOfSlice[0][1] = 7
		assert.NotEqual(t, sliceOfSlice, gotSliceOfSlice)
	})

	t.Run("map", func(t *testing.T) {
		m := make(map[string]int)
		m["1"] = 1
		m["2"] = 2
		m["3"] = 3

		got, _ := kamino.Clone(m)
		assert.Equal(t, m, got)

		m2 := make(map[[3]int][]int)
		m2[[3]int{1, 2, 3}] = []int{1, 2, 3}
		m2[[3]int{2, 3}] = []int{2, 3}
		m2[[3]int{3}] = []int{3}

		got2, _ := kamino.Clone(m2)
		assert.Equal(t, m2, got2)

		i := 1
		m3 := map[string]*int{
			"1": &i,
		}
		got3, _ := kamino.Clone(m3)
		assert.Equal(t, m3, got3)

		*got3["1"] = 10
		assert.NotEqual(t, m3, got3)
	})

	t.Run("simple struct", func(t *testing.T) {
		type simpleStruct struct {
			Int     int
			Float64 float64
			String  string
		}
		param := simpleStruct{
			10, 20., "30",
		}

		got, _ := kamino.Clone(param)
		assert.Equal(t, param, got)
	})

	t.Run("pointer", func(t *testing.T) {
		var (
			i64     int64 = 1
			sstruct       = simpleStruct{2, 3, "4"}

			i64Ptr     = &i64
			sstructPtr = &sstruct
		)

		i64PtrClone, _ := kamino.Clone(i64Ptr)
		assert.Equal(t, *i64Ptr, *i64PtrClone)

		*i64PtrClone = -1
		assert.Equal(t, *i64Ptr, int64(1))

		sstructPtrClone, _ := kamino.Clone(sstructPtr)
		assert.Equal(t, sstructPtr, sstructPtrClone)

		sstructPtrClone.Int = -2
		sstructPtrClone.Float64 = -3
		sstructPtrClone.String = "-4"

		assert.Equal(t, sstructPtr.Int, 2)
		assert.Equal(t, sstructPtr.Float64, float64(3))
		assert.Equal(t, sstructPtr.String, "4")
	})

	t.Run("ptrCircles", func(t *testing.T) {
		type circled struct {
			PtrA *int
			PtrB *int

			This *circled
		}

		i := 10
		c := circled{PtrA: &i, PtrB: &i}
		c.This = &c

		cPtrClone, _ := kamino.Clone(&c)

		assert.Equal(t, cPtrClone, &c)
		assert.Equal(t, unsafe.Pointer(cPtrClone), unsafe.Pointer(cPtrClone.This))
		assert.Equal(t, unsafe.Pointer(cPtrClone.PtrA), unsafe.Pointer(cPtrClone.PtrB))
	})

	t.Run("alltogether", func(t *testing.T) {
		got, _ := kamino.Clone(alltogetherInstance)
		assert.Equal(t, alltogetherInstance, got)
	})

	t.Run("original is not affected by clone mutation", func(t *testing.T) {
		original := alltogether{
			Int:        1,
			Float64:    2,
			String:     "3",
			ArrayOfInt: [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			SliceOfInt: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			Nested:     simpleStruct{1, 2, "3"},
		}

		originalHandCopy := alltogether{
			Int:        1,
			Float64:    2,
			String:     "3",
			ArrayOfInt: [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			SliceOfInt: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			Nested:     simpleStruct{1, 2, "3"},
		}
		assert.Equal(t, original, originalHandCopy)

		clone, err := kamino.Clone(original)
		assert.NoError(t, err)

		assert.Equal(t, original, clone)
		assert.Equal(t, original, originalHandCopy)

		clone.Int = 2
		clone.Float64 = 3
		clone.String = "another string"
		clone.ArrayOfInt[5] = -1
		clone.SliceOfInt[5] = -5

		assert.NotEqual(t, original, clone)
		assert.Equal(t, original, originalHandCopy)
	})
}

func TestInterface(t *testing.T) {
	x := []any{nil}
	y, _ := kamino.Clone(x)
	assert.Equal(t, x, y)

	var a any
	b, _ := kamino.Clone(a)
	assert.True(t, a == b)
}

func TestAny(t *testing.T) {
	x := any(simpleStructIntance)
	y, _ := kamino.Clone(x)
	assert.Equal(t, x, y)

	ys, ok := y.(simpleStruct)
	assert.True(t, ok)
	assert.Equal(t, ys, simpleStructIntance)
}

func TestTwoNils(t *testing.T) {
	type Foo struct {
		A int
	}
	type Bar struct {
		B int
	}
	type FooBar struct {
		Foo  *Foo
		Bar  *Bar
		Foo2 *Foo
		Bar2 *Bar
	}

	src := &FooBar{
		Foo2: &Foo{1},
		Bar2: &Bar{2},
	}

	dst, _ := kamino.Clone(src)

	assert.Equal(t, src, dst)
}

func ptrTo[T any](v T) *T {
	return &v
}

func TestCloneUnexported(t *testing.T) {
	type Foo struct {
		A    int
		a    int
		Ptr1 *int
		ptr2 *int
	}

	foo := Foo{
		1, 2, ptrTo(3), ptrTo(4),
	}

	bar, _ := kamino.Clone(foo)
	assert.Equal(t, foo, bar)

	assert.NotEqual(t, unsafe.Pointer(foo.Ptr1), unsafe.Pointer(bar.Ptr1))
	assert.Equal(t, unsafe.Pointer(foo.ptr2), unsafe.Pointer(bar.ptr2))
}

func TestForceCloneUnexported(t *testing.T) {
	type Foo struct {
		A    int
		a    int
		Ptr1 *int
		ptr2 *int
	}

	foo := Foo{
		1, 2, ptrTo(3), ptrTo(4),
	}

	bar, _ := kamino.Clone(foo, kamino.WithForceUnexported())
	assert.Equal(t, foo, bar)

	assert.NotEqual(t, unsafe.Pointer(foo.Ptr1), unsafe.Pointer(bar.Ptr1))
	assert.NotEqual(t, unsafe.Pointer(foo.ptr2), unsafe.Pointer(bar.ptr2))
}

func TestForceZeroUnexported(t *testing.T) {
	type Foo struct {
		A    int
		a    int
		Ptr1 *int
		ptr2 *int
	}

	foo := Foo{
		1, 2, ptrTo(3), ptrTo(4),
	}

	bar, _ := kamino.Clone(foo, kamino.WithZeroUnexported())
	assert.Equal(t, foo.A, bar.A)
	assert.Equal(t, foo.Ptr1, bar.Ptr1)

	assert.Equal(t, 2, foo.a)
	assert.Equal(t, ptrTo(4), foo.ptr2)

	var nilintPtr *int
	assert.Equal(t, 0, bar.a)
	assert.Equal(t, nilintPtr, bar.ptr2)
}

type Fooer interface {
	foo() int
}

type fooer struct {
	i int
}

func (f *fooer) foo() int {
	return f.i
}

func TestCopyInterface(t *testing.T) {
	type fooerWrapper struct {
		F Fooer
	}

	fi := &fooer{i: 10}

	fw := fooerWrapper{
		F: fi,
	}

	got, _ := kamino.Clone(fw)

	assert.Equal(t, got, fw)
	fi.i = 20
	assert.NotEqual(t, got.F.foo(), fw.F.foo())
}

func TestCopyNestedTime(t *testing.T) {
	type nestedTime struct {
		T time.Time
	}

	nt := nestedTime{time.Now()}
	got, _ := kamino.Clone(nt)

	assert.Equal(t, got.T, nt.T)
}

func TestCopyNestedNil(t *testing.T) {
	type nestedNil struct {
		X any
	}

	nn := nestedNil{}
	got, _ := kamino.Clone(nn)

	assert.Equal(t, got, nn)
}

func TestWithErrOnUnsupported(t *testing.T) {
	t.Run("errOnUnsupported chans", func(t *testing.T) {
		f := func() {}

		_, err := kamino.Clone(f)
		assert.NoError(t, err)

		_, err = kamino.Clone(f, kamino.WithErrOnUnsupported())
		assert.Error(t, err)
	})

	t.Run("errOnUnsupported funcs", func(t *testing.T) {
		ch := make(chan int)

		_, err := kamino.Clone(ch)
		assert.NoError(t, err)

		_, err = kamino.Clone(ch, kamino.WithErrOnUnsupported())
		assert.Error(t, err)
	})

	t.Run("errOnUnsupported ptrs", func(t *testing.T) {
		ch := make(chan int)
		f := func() {}

		_, err := kamino.Clone(ch, kamino.WithErrOnUnsupported())
		assert.Error(t, err)

		_, err = kamino.Clone(f, kamino.WithErrOnUnsupported())
		assert.Error(t, err)
	})

	t.Run("errOnUnsupported with suported only", func(t *testing.T) {
		_, err := kamino.Clone(alltogetherInstance, kamino.WithErrOnUnsupported())
		assert.NoError(t, err)
	})

	t.Run("errOnUnsupported in structs", func(t *testing.T) {
		type foo struct {
			F func()
		}

		f := foo{F: func() {}}

		_, err := kamino.Clone(f)
		assert.NoError(t, err)

		_, err = kamino.Clone(f, kamino.WithErrOnUnsupported())
		assert.Error(t, err)

		_, err = kamino.Clone(&f, kamino.WithErrOnUnsupported())
		assert.Error(t, err)
	})

	t.Run("errOnUnsupported unexported in structs", func(t *testing.T) {
		type foo struct {
			f func()
		}

		f := foo{f: func() {}}

		_, err := kamino.Clone(f)
		assert.NoError(t, err)

		_, err = kamino.Clone(f, kamino.WithErrOnUnsupported())
		assert.NoError(t, err)

		_, err = kamino.Clone(f, kamino.WithErrOnUnsupported(), kamino.WithForceUnexported())
		assert.Error(t, err)
	})

	t.Run("errOnUnsupported in array", func(t *testing.T) {
		a := [...]any{1, "2", func() {}}

		_, err := kamino.Clone(a)
		assert.NoError(t, err)

		_, err = kamino.Clone(a, kamino.WithErrOnUnsupported())
		assert.Error(t, err)
	})

	t.Run("errOnUnsupported in slice", func(t *testing.T) {
		s := []any{1, "2", func() {}}

		_, err := kamino.Clone(s)
		assert.NoError(t, err)

		_, err = kamino.Clone(s, kamino.WithErrOnUnsupported())
		assert.Error(t, err)
	})

	t.Run("errOnUnsupported in maps", func(t *testing.T) {
		s := map[string]any{"supported": 1, "unsuppotred": func() {}}

		_, err := kamino.Clone(s)
		assert.NoError(t, err)

		_, err = kamino.Clone(s, kamino.WithErrOnUnsupported())
		assert.Error(t, err)
	})
}
