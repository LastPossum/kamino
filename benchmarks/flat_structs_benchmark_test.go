package kamino_benchmark

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/LastPossum/kamino"

	barkimedes "github.com/barkimedes/go-deepcopy"
	mohae "github.com/mohae/deepcopy"
)

type benchSimpleStruct5 struct {
	I1 int
	F1 float64
	S1 string
	B1 byte
	L1 bool
}

type benchSimpleStruct10 struct {
	I1 int
	F1 float64
	S1 string
	B1 byte
	L1 bool
	I2 int
	F2 float64
	S2 string
	B2 byte
	L2 bool
}

type benchSimpleStruct15 struct {
	I1 int
	F1 float64
	S1 string
	B1 byte
	L1 bool
	I2 int
	F2 float64
	S2 string
	B2 byte
	L2 bool
	I3 int
	F3 float64
	S3 string
	B3 byte
	L3 bool
}

type benchSimpleStruct20 struct {
	I1 int
	F1 float64
	S1 string
	B1 byte
	L1 bool
	I2 int
	F2 float64
	S2 string
	B2 byte
	L2 bool
	I3 int
	F3 float64
	S3 string
	B3 byte
	L3 bool
	I4 int
	F4 float64
	S4 string
	B4 byte
	L4 bool
}
type benchSimpleStruct25 struct {
	I1 int
	F1 float64
	S1 string
	B1 byte
	L1 bool
	I2 int
	F2 float64
	S2 string
	B2 byte
	L2 bool
	I3 int
	F3 float64
	S3 string
	B3 byte
	L3 bool
	I4 int
	F4 float64
	S4 string
	B4 byte
	L4 bool
	I5 int
	F5 float64
	S5 string
	B5 byte
	L5 bool
}
type benchSimpleStruct30 struct {
	I1 int
	F1 float64
	S1 string
	B1 byte
	L1 bool
	I2 int
	F2 float64
	S2 string
	B2 byte
	L2 bool
	I3 int
	F3 float64
	S3 string
	B3 byte
	L3 bool
	I4 int
	F4 float64
	S4 string
	B4 byte
	L4 bool
	I5 int
	F5 float64
	S5 string
	B5 byte
	L5 bool
	I6 int
	F6 float64
	S6 string
	B6 byte
	L6 bool
}
type benchSimpleStruct35 struct {
	I1 int
	F1 float64
	S1 string
	B1 byte
	L1 bool
	I2 int
	F2 float64
	S2 string
	B2 byte
	L2 bool
	I3 int
	F3 float64
	S3 string
	B3 byte
	L3 bool
	I4 int
	F4 float64
	S4 string
	B4 byte
	L4 bool
	I5 int
	F5 float64
	S5 string
	B5 byte
	L5 bool
	I6 int
	F6 float64
	S6 string
	B6 byte
	L6 bool
	I7 int
	F7 float64
	S7 string
	B7 byte
	L7 bool
}

type benchSimpleStruct40 struct {
	I1 int
	F1 float64
	S1 string
	B1 byte
	L1 bool
	I2 int
	F2 float64
	S2 string
	B2 byte
	L2 bool
	I3 int
	F3 float64
	S3 string
	B3 byte
	L3 bool
	I4 int
	F4 float64
	S4 string
	B4 byte
	L4 bool
	I5 int
	F5 float64
	S5 string
	B5 byte
	L5 bool
	I6 int
	F6 float64
	S6 string
	B6 byte
	L6 bool
	I7 int
	F7 float64
	S7 string
	B7 byte
	L7 bool
	I8 int
	F8 float64
	S8 string
	B8 byte
	L8 bool
}

type benchSimpleStruct45 struct {
	I1 int
	F1 float64
	S1 string
	B1 byte
	L1 bool
	I2 int
	F2 float64
	S2 string
	B2 byte
	L2 bool
	I3 int
	F3 float64
	S3 string
	B3 byte
	L3 bool
	I4 int
	F4 float64
	S4 string
	B4 byte
	L4 bool
	I5 int
	F5 float64
	S5 string
	B5 byte
	L5 bool
	I6 int
	F6 float64
	S6 string
	B6 byte
	L6 bool
	I7 int
	F7 float64
	S7 string
	B7 byte
	L7 bool
	I8 int
	F8 float64
	S8 string
	B8 byte
	L8 bool
	I9 int
	F9 float64
	S9 string
	B9 byte
	L9 bool
}

type benchSimpleStruct50 struct {
	I1  int
	F1  float64
	S1  string
	B1  byte
	L1  bool
	I2  int
	F2  float64
	S2  string
	B2  byte
	L2  bool
	I3  int
	F3  float64
	S3  string
	B3  byte
	L3  bool
	I4  int
	F4  float64
	S4  string
	B4  byte
	L4  bool
	I5  int
	F5  float64
	S5  string
	B5  byte
	L5  bool
	I6  int
	F6  float64
	S6  string
	B6  byte
	L6  bool
	I7  int
	F7  float64
	S7  string
	B7  byte
	L7  bool
	I8  int
	F8  float64
	S8  string
	B8  byte
	L8  bool
	I9  int
	F9  float64
	S9  string
	B9  byte
	L9  bool
	I10 int
	F10 float64
	S10 string
	B10 byte
	L10 bool
}

var flats = map[int]any{
	5:  benchSimpleStruct5{1, 2, "3", 4, true},
	10: benchSimpleStruct10{1, 2, "3", 4, true, 1, 2, "3", 4, true},
	15: benchSimpleStruct15{1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true},
	20: benchSimpleStruct20{1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true},
	25: benchSimpleStruct25{1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true},
	30: benchSimpleStruct30{1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true},
	35: benchSimpleStruct35{1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true},
	40: benchSimpleStruct40{1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true},
	45: benchSimpleStruct45{1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true},
	50: benchSimpleStruct50{1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true, 1, 2, "3", 4, true},
}

var structFieldsKeys = []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50}

func BenchmarkFlatStruct(b *testing.B) {
	for _, k := range structFieldsKeys {
		st := flats[k]
		b.Run(fmt.Sprintf("barkimedes for %d fiels flat struct", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				barkimedes.Anything(st)
			}
		})
	}

	for _, k := range structFieldsKeys {
		st := flats[k]
		b.Run(fmt.Sprintf("mohae for %d fiels flat struct", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mohae.Copy(st)
			}
		})
	}

	for _, k := range structFieldsKeys {
		st := flats[k]
		b.Run(fmt.Sprintf("json for %d fiels flat struct", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				cloneJSON(st)
			}
		})
	}

	for _, k := range structFieldsKeys {
		st := flats[k]
		b.Run(fmt.Sprintf("msgpack for %d fiels flat struct", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				cloneMsgPack(st)
			}
		})
	}

	for _, k := range structFieldsKeys {
		st := flats[k]
		b.Run(fmt.Sprintf("kamino for %d fiels flat struct", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				kamino.Clone(st)
			}
		})
	}
}

func BenchmarkFlatStructKamino(b *testing.B) {
	for _, k := range structFieldsKeys {
		st := flats[k]
		b.Run(fmt.Sprintf("kamino for %d fiels flat struct", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				kamino.Clone(st)
			}
		})
	}
}

func cpInit(x any) any {
	original := reflect.ValueOf(x)

	cpy := reflect.New(original.Type()).Elem()
	cp(original, cpy)

	return cpy.Interface()
}

func isBoxedNil(src any) bool {
	return src == nil || src == any(nil)
}

func cpInitT[T any](src T) (T, error) {
	if isBoxedNil(src) {
		return src, nil
	}

	valPtr := reflect.ValueOf(&src)
	err := cloneNested(valPtr)
	return src, err
}

func cp(original, cpy reflect.Value) {
	if original.Kind() != reflect.Struct {
		cpy.Set(original)
		return
	}
	for i := 0; i < original.NumField(); i++ {
		if !cpy.CanSet() {
			continue
		}

		cp(original.Field(i), cpy.Field(i))
	}
}

func cloneNested(valPtr reflect.Value) error {
	v := valPtr.Elem()

	switch v.Kind() {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {

			wField := v.Field(i)

			if wField.CanSet() {
				if err := cloneNested(wField.Addr()); err != nil {
					return err
				}
				continue
			}
		}
	}

	return nil
}

func BenchmarkX(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cpInit(nestedSingle)
	}
}

func BenchmarkY(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cpInitT(nestedSingle)
	}
}
