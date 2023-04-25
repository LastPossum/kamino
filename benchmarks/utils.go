package kamino_benchmark

import (
	"encoding/json"

	"github.com/vmihailenco/msgpack"
)

func ptrTo[T any](v T) *T {
	return &v
}

func cloneJSON[T any](obj T) T {
	x, _ := json.Marshal(obj)
	var res T
	json.Unmarshal(x, &res)
	return res
}

func cloneMsgPack[T any](x T) T {
	b, _ := msgpack.Marshal(x)
	var res T
	msgpack.Unmarshal(b, &res)
	return res
}
