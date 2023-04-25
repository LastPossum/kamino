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
	C1 complex64
}

type benchSimpleStruct10 struct {
	I1 int
	F1 float64
	S1 string
	B1 byte
	C1 complex64
	I2 int
	F2 float64
	S2 string
	B2 byte
	C2 complex64
}

type benchSimpleStruct15 struct {
	I1 int
	F1 float64
	S1 string
	B1 byte
	C1 complex64
	I2 int
	F2 float64
	S2 string
	B2 byte
	C2 complex64
	I3 int
	F3 float64
	S3 string
	B3 byte
	C3 complex64
}

type benchSimpleStruct20 struct {
	I1 int
	F1 float64
	S1 string
	B1 byte
	C1 complex64
	I2 int
	F2 float64
	S2 string
	B2 byte
	C2 complex64
	I3 int
	F3 float64
	S3 string
	B3 byte
	C3 complex64
	I4 int
	F4 float64
	S4 string
	B4 byte
	C4 complex64
}
type benchSimpleStruct25 struct {
	I1 int
	F1 float64
	S1 string
	B1 byte
	C1 complex64
	I2 int
	F2 float64
	S2 string
	B2 byte
	C2 complex64
	I3 int
	F3 float64
	S3 string
	B3 byte
	C3 complex64
	I4 int
	F4 float64
	S4 string
	B4 byte
	C4 complex64
	I5 int
	F5 float64
	S5 string
	B5 byte
	C5 complex64
}
type benchSimpleStruct30 struct {
	I1 int
	F1 float64
	S1 string
	B1 byte
	C1 complex64
	I2 int
	F2 float64
	S2 string
	B2 byte
	C2 complex64
	I3 int
	F3 float64
	S3 string
	B3 byte
	C3 complex64
	I4 int
	F4 float64
	S4 string
	B4 byte
	C4 complex64
	I5 int
	F5 float64
	S5 string
	B5 byte
	C5 complex64
	I6 int
	F6 float64
	S6 string
	B6 byte
	C6 complex64
}
type benchSimpleStruct35 struct {
	I1 int
	F1 float64
	S1 string
	B1 byte
	C1 complex64
	I2 int
	F2 float64
	S2 string
	B2 byte
	C2 complex64
	I3 int
	F3 float64
	S3 string
	B3 byte
	C3 complex64
	I4 int
	F4 float64
	S4 string
	B4 byte
	C4 complex64
	I5 int
	F5 float64
	S5 string
	B5 byte
	C5 complex64
	I6 int
	F6 float64
	S6 string
	B6 byte
	C6 complex64
	I7 int
	F7 float64
	S7 string
	B7 byte
	C7 complex64
}

type benchSimpleStruct40 struct {
	I1 int
	F1 float64
	S1 string
	B1 byte
	C1 complex64
	I2 int
	F2 float64
	S2 string
	B2 byte
	C2 complex64
	I3 int
	F3 float64
	S3 string
	B3 byte
	C3 complex64
	I4 int
	F4 float64
	S4 string
	B4 byte
	C4 complex64
	I5 int
	F5 float64
	S5 string
	B5 byte
	C5 complex64
	I6 int
	F6 float64
	S6 string
	B6 byte
	C6 complex64
	I7 int
	F7 float64
	S7 string
	B7 byte
	C7 complex64
	I8 int
	F8 float64
	S8 string
	B8 byte
	C8 complex64
}

type benchSimpleStruct45 struct {
	I1 int
	F1 float64
	S1 string
	B1 byte
	C1 complex64
	I2 int
	F2 float64
	S2 string
	B2 byte
	C2 complex64
	I3 int
	F3 float64
	S3 string
	B3 byte
	C3 complex64
	I4 int
	F4 float64
	S4 string
	B4 byte
	C4 complex64
	I5 int
	F5 float64
	S5 string
	B5 byte
	C5 complex64
	I6 int
	F6 float64
	S6 string
	B6 byte
	C6 complex64
	I7 int
	F7 float64
	S7 string
	B7 byte
	C7 complex64
	I8 int
	F8 float64
	S8 string
	B8 byte
	C8 complex64
	I9 int
	F9 float64
	S9 string
	B9 byte
	C9 complex64
}

type benchSimpleStruct50 struct {
	I1  int
	F1  float64
	S1  string
	B1  byte
	C1  complex64
	I2  int
	F2  float64
	S2  string
	B2  byte
	C2  complex64
	I3  int
	F3  float64
	S3  string
	B3  byte
	C3  complex64
	I4  int
	F4  float64
	S4  string
	B4  byte
	C4  complex64
	I5  int
	F5  float64
	S5  string
	B5  byte
	C5  complex64
	I6  int
	F6  float64
	S6  string
	B6  byte
	C6  complex64
	I7  int
	F7  float64
	S7  string
	B7  byte
	C7  complex64
	I8  int
	F8  float64
	S8  string
	B8  byte
	C8  complex64
	I9  int
	F9  float64
	S9  string
	B9  byte
	C9  complex64
	I10 int
	F10 float64
	S10 string
	B10 byte
	C10 complex64
}

type benchSimpleStruct100 struct {
	I1  int
	F1  float64
	S1  string
	B1  byte
	C1  complex64
	I2  int
	F2  float64
	S2  string
	B2  byte
	C2  complex64
	I3  int
	F3  float64
	S3  string
	B3  byte
	C3  complex64
	I4  int
	F4  float64
	S4  string
	B4  byte
	C4  complex64
	I5  int
	F5  float64
	S5  string
	B5  byte
	C5  complex64
	I6  int
	F6  float64
	S6  string
	B6  byte
	C6  complex64
	I7  int
	F7  float64
	S7  string
	B7  byte
	C7  complex64
	I8  int
	F8  float64
	S8  string
	B8  byte
	C8  complex64
	I9  int
	F9  float64
	S9  string
	B9  byte
	C9  complex64
	I10 int
	F10 float64
	S10 string
	B10 byte
	C10 complex64
	I11 int
	F11 float64
	S11 string
	B11 byte
	C11 complex64
	I12 int
	F12 float64
	S12 string
	B12 byte
	C12 complex64
	I13 int
	F13 float64
	S13 string
	B13 byte
	C13 complex64
	I14 int
	F14 float64
	S14 string
	B14 byte
	C14 complex64
	I15 int
	F15 float64
	S15 string
	B15 byte
	C15 complex64
	I16 int
	F16 float64
	S16 string
	B16 byte
	C16 complex64
	I17 int
	F17 float64
	S17 string
	B17 byte
	C17 complex64
	I18 int
	F18 float64
	S18 string
	B18 byte
	C18 complex64
	I19 int
	F19 float64
	S19 string
	B19 byte
	C19 complex64
	I20 int
	F20 float64
	S20 string
	B20 byte
	C20 complex64
}

var flats = map[int]any{
	5:  benchSimpleStruct5{1, 2, "3", 4, 5i},
	10: benchSimpleStruct10{1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i},
	15: benchSimpleStruct15{1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i},
	20: benchSimpleStruct20{1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i},
	25: benchSimpleStruct25{1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i},
	30: benchSimpleStruct30{1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i},
	35: benchSimpleStruct35{1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i},
	40: benchSimpleStruct40{1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i},
	45: benchSimpleStruct45{1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i},
	50: benchSimpleStruct50{1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i, 1, 2, "3", 4, 5i},
}

func BenchmarkFlatStruct(b *testing.B) {
	for k, st := range flats {

		b.Run(fmt.Sprintf("barkimedes for %d fiels flat struct", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				barkimedes.Anything(st)
			}
		})

		b.Run(fmt.Sprintf("mohae for %d fiels flat struct", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mohae.Copy(st)
			}
		})

		b.Run(fmt.Sprintf("json for %d fiels flat struct", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				cloneJSON(st)
			}
		})

		b.Run(fmt.Sprintf("msgpack for %d fiels flat struct", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				cloneMsgPack(st)
			}
		})

		b.Run(fmt.Sprintf("kamino for %d fiels flat struct", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				kamino.Clone(st)
			}
		})
	}
}

func BenchmarkFlatStructKamino(b *testing.B) {
	for k, st := range flats {
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
