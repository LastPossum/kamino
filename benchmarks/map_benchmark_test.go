package kamino_benchmark

import (
	"testing"

	"github.com/LastPossum/kamino"
	barkimedes "github.com/barkimedes/go-deepcopy"
	mohae "github.com/mohae/deepcopy"
)

var mp = map[string]any{
	"int":    1,
	"float":  2.,
	"string": "3",
	"bytes":  make([]byte, 128),
	"structs": []*simpleStruct{
		{1, 2., "3"},
		{1, 2., "3"},
		{1, 2., "3"},
	},
	"nested": map[string]any{
		"int":   1,
		"float": 2.,
	},
}

func BenchmarkCloneBarkimedesMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		barkimedes.Anything(mp)
	}
}

func BenchmarkCloneMohaeMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mohae.Copy(mp)
	}
}

func BenchmarkCloneJsonMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cloneJSON(mp)
	}
}

func BenchmarkCloneMsgPacknMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cloneMsgPack(mp)
	}
}

func BenchmarkCloneMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kamino.Clone(mp)
	}
}
