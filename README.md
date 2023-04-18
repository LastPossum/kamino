#  Kamino - deep copy library

Kamino is a programming library that provides faster way to [deep copy](https://developer.mozilla.org/en-US/docs/Glossary/Deep_copy) data structures in Golang.
It designed to be fast and avoid unnecessary allocations. See [benchmarks](#benchmarks). Also It uses generics to avoid  type assertions in client's code.
Although kamino designed to be fast, it handle circular pointers. 
By default kamino makes a shallow copy for unsupported kinds of values (chans and funcs) and unexported struct fields. Buit in can be changed via functional options. 

## Installation

`go get github.com/LastPossum/kamino`

## Usage

To use the library, you need to import it into your project:

`import "github.com/LastPossum/kamino"`

Once you have imported the library, you can use the `func Clone[T any](src T, opts ...funcOptions) (T, error)` function to create a deep copy of your data structure.

### Sample Program

```go
package main

import (
	"fmt"
	"reflect"

	"github.com/LastPossum/kamino"
)

type Foo struct {
	A string
	B int
	C *float64
	D any
}

func main() {
	var f float64 = 1
	original := Foo{
		A: "a string",
		B: 1114,
		C: &f,
	}

    original.D = &original

	copy, err := kamino.Clone(original)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(original, original.D)
	fmt.Println(copy, copy.D)
	fmt.Println(reflect.DeepEqual(original, copy))
}
```

### Sample Output

_Note that the values are all the same, but the addresses are all different. Note also that circular references are handled._

```go
{a string 1114 0xc00001c030 0xc0000161b0} &{a string 1114 0xc00001c030 0xc0000161b0}
{a string 1114 0xc00001c038 0xc0000162a0} &{a string 1114 0xc00001c038 0xc0000162a0}
true
```

## Options

Kamino supports several following options:
- `WithExpectedPtrsCount` - Use it to set the initial capacity for the pointers map used during the cloning process. If the amount of pointers in the source object is known and rather big, this setting could reduce allocations and slightly improve performance.
- `WithForceUnexported` - when this setting is enabled, unexported fields will be cloned forcefully.
- `WithZeroUnexported` -  when this setting is enabled, unexported fields will be forcefully set to zero value of their kind.
- `WithErrOnUnsupported` - when this setting is enabled, attempting to clone channels and functions, even if they are nested, will result in an error.

## Benchmarks

Results as of April 25, 2023 with Go go1.20.3 on OSx Intel Core i5 (4 core 2,4 GHz)

```
BenchmarkCloneKaminoComplexStruct-8   	                                  277168	        3943 ns/op	    2013 B/op	      27 allocs/op
BenchmarkNestedStructKamino/kamino_for_7_fiels_nested_struct-8         	 4228910	       281.3 ns/op	      56 B/op	       4 allocs/op
BenchmarkCloneMap-8   	                                                  331911	        3482 ns/op	    1424 B/op	      38 allocs/op

```
## Limitation

For now, Kamino does not support circular slices and maps. Additionally, if two slices in the source object point to the same underlying array, this feature will not be kept in the copy object. Therefore, these cases will cause a stack overflow:

```
func TestCircularSlice(t *testing.T) {
	a := []any{nil}
	a[0] = a
	cp, err := kamino.Clone(a)

	assert.NoError(t, err)
	assert.Equal(t, cp, a)
}

func TestCircularMap(t *testing.T) {
	a := map[string]any{"a": nil}
	a["a"] = a
	cp, err := kamino.Clone(a)

	assert.NoError(t, err)
	assert.Equal(t, cp, a)
}

```

## License

This library is released under the MIT License. You are free to use, modify, and distribute it as you see fit. See the LICENSE file for more information.