package util

import (
	"fmt"
	"hash/fnv"
	"strings"
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type HashCodeFunc func(interface{}) uint64
type CompareFunc func(interface{}, interface{}) int

var hcfs = make(map[string]HashCodeFunc)
var cfs = make(map[string]CompareFunc)

func init() {
	initHashCode()
	initCompare()
}

func initHashCode() {
	intHash := func(x interface{}) uint64 {
		v := x.(int)
		return uint64(v)
	}
	stringHash := func(x interface{}) uint64 {
		h := fnv.New64a()
		_, _ = h.Write([]byte(x.(string)))
		return h.Sum64()
	}

	hcfs["int"] = intHash
	hcfs["int32"] = intHash
	hcfs["int64"] = intHash
	hcfs["uint"] = intHash
	hcfs["uint32"] = intHash
	hcfs["uint64"] = intHash
	hcfs["rune"] = intHash
	hcfs["string"] = stringHash
}

func HashCode(v interface{}) uint64 {
	t := fmt.Sprintf("%T", v)
	f, ok := hcfs[t]
	if !ok {
		panic("unsupported type")
	}
	return f(v)
}

func initCompare() {
	intCompare := func(a, b interface{}) int {
		if a == b {
			return 0
		}
		if a.(int) > b.(int) {
			return 1
		}
		return -1
	}
	stringCompare := func(a, b interface{}) int {
		return strings.Compare(a.(string), b.(string))
	}

	cfs["int"] = intCompare
	cfs["int32"] = intCompare
	cfs["int64"] = intCompare
	cfs["uint"] = intCompare
	cfs["uint32"] = intCompare
	cfs["uint64"] = intCompare
	cfs["rune"] = intCompare
	cfs["string"] = stringCompare
}

func Compare(a, b interface{}) int {
	t := fmt.Sprintf("%T", a)
	u := fmt.Sprintf("%T", b)
	if t != u {
		panic(fmt.Sprintf("args types are not matched: %T, %T", a, b))
	}
	f, ok := cfs[t]
	if !ok {
		panic("unsupported type")
	}
	return f(a, b)
}