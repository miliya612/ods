package util

import (
	"fmt"
	"hash/fnv"
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type HashCodeFunc func(interface{}) uint64

var hcs =  make(map[string]HashCodeFunc)

func init() {
	intHash := func(x interface{}) uint64 {
		v := x.(int)
		return uint64(v)
	}
	stringHash := func(x interface{}) uint64 {
		h := fnv.New64a()
		_, _ = h.Write([]byte(x.(string)))
		return h.Sum64()
	}

	hcs["int"] = intHash
	hcs["int32"] = intHash
	hcs["int64"] = intHash
	hcs["uint"] = intHash
	hcs["uint32"] = intHash
	hcs["uint64"] = intHash
	hcs["rune"] = intHash
	hcs["string"] = stringHash
}

func HashCode(v interface{}) uint64 {
	t := fmt.Sprintf("%T", v)
	f, ok := hcs[t]
	if !ok {
		panic("unsupported type")
	}
	return f(v)

}
