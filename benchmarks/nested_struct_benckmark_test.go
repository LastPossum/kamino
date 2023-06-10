package kamino_benchmark

import (
	"fmt"
	"testing"

	"github.com/LastPossum/kamino"
	barkimedes "github.com/barkimedes/go-deepcopy"
	mohae "github.com/mohae/deepcopy"
)

type NestedLevel1 struct {
	Payload int
}

type NestedLevel2 struct {
	Payload NestedLevel1
}
type NestedLevel3 struct {
	Payload NestedLevel2
}
type NestedLevel4 struct {
	Payload NestedLevel3
}

type NestedLevel5 struct {
	Payload NestedLevel4
}

type NestedLevel6 struct {
	Payload NestedLevel5
}

type NestedLevel7 struct {
	Payload NestedLevel6
}

type NestedLevel8 struct {
	Payload NestedLevel7
}

type NestedLevel9 struct {
	Payload NestedLevel8
}

type NestedLevel10 struct {
	Payload NestedLevel9
}

var (
	i         = 1
	nestedMap = map[int]any{
		1:  NestedLevel1{i},
		2:  NestedLevel2{NestedLevel1{i}},
		3:  NestedLevel3{NestedLevel2{NestedLevel1{i}}},
		4:  NestedLevel4{NestedLevel3{NestedLevel2{NestedLevel1{i}}}},
		5:  NestedLevel5{NestedLevel4{NestedLevel3{NestedLevel2{NestedLevel1{i}}}}},
		6:  NestedLevel6{NestedLevel5{NestedLevel4{NestedLevel3{NestedLevel2{NestedLevel1{i}}}}}},
		7:  NestedLevel7{NestedLevel6{NestedLevel5{NestedLevel4{NestedLevel3{NestedLevel2{NestedLevel1{i}}}}}}},
		8:  NestedLevel8{NestedLevel7{NestedLevel6{NestedLevel5{NestedLevel4{NestedLevel3{NestedLevel2{NestedLevel1{i}}}}}}}},
		9:  NestedLevel9{NestedLevel8{NestedLevel7{NestedLevel6{NestedLevel5{NestedLevel4{NestedLevel3{NestedLevel2{NestedLevel1{i}}}}}}}}},
		10: NestedLevel10{NestedLevel9{NestedLevel8{NestedLevel7{NestedLevel6{NestedLevel5{NestedLevel4{NestedLevel3{NestedLevel2{NestedLevel1{i}}}}}}}}}},
	}

	nestedSingle = NestedLevel8{NestedLevel7{NestedLevel6{NestedLevel5{NestedLevel4{NestedLevel3{NestedLevel2{NestedLevel1{i}}}}}}}}
)

func BenchmarkNestedKamino(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kamino.Clone(nestedSingle)
	}
}

func BenchmarkNestedMohae(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mohae.Copy(nestedSingle)
	}
}

func BenchmarkNestedStruct(b *testing.B) {
	for k := 1; k <= 10; k++ {
		st := nestedMap[k]
		b.Run(fmt.Sprintf("barkimedes for %d fiels nested struct", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				barkimedes.Anything(st)
			}
		})
	}
	for k := 1; k <= 10; k++ {
		st := nestedMap[k]
		b.Run(fmt.Sprintf("mohae for %d fiels nested struct", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mohae.Copy(st)
			}
		})
	}
	for k := 1; k <= 10; k++ {
		st := nestedMap[k]
		b.Run(fmt.Sprintf("json for %d fiels nested struct", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				cloneJSON(st)
			}
		})
	}

	for k := 1; k <= 10; k++ {
		st := nestedMap[k]
		b.Run(fmt.Sprintf("msgpack for %d fiels nested struct", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				cloneMsgPack(st)
			}
		})
	}

	for k := 1; k <= 10; k++ {
		st := nestedMap[k]
		b.Run(fmt.Sprintf("kamino for %d fiels nested struct", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				kamino.Clone(st)
			}
		})
	}
}

func BenchmarkNestedStructKamino(b *testing.B) {
	for k := 1; k <= 10; k++ {
		st := nestedMap[k]
		b.Run(fmt.Sprintf("kamino for %d fiels nested struct", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				kamino.Clone(st)
			}
		})
	}
}
