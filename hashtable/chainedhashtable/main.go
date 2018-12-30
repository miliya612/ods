package chainedhashtable

import (
	"github.com/miliya612/ods/arraylist/arraystack"
	"github.com/miliya612/ods/util"
	"math/rand"
)

type ChainedHashTable struct {
	buf []arraystack.ArrayStack
	len int
	z, d uint64
}

func New() *ChainedHashTable {
	return &ChainedHashTable{
		buf: make([]arraystack.ArrayStack, 0),
		z: rand.Uint64() | uint64(1),
		d: 1,
	}
}

func (cht *ChainedHashTable) Size() int {
	return cht.len
}

func (cht *ChainedHashTable) Add(x interface{}) bool {
	if cht.Find(x) != nil {
		return false
	}
	if cht.hash(x)+1 > uint64(cap(cht.buf)) {
		cht.resize()
	}
	cht.buf[cht.hash(x)].Push(x)
	cht.len++
	return true
}

func (cht *ChainedHashTable) Remove(x interface{}) interface{} {
	if cht.len == 0 {
		return nil
	}
	j := cht.hash(x)
	for i := 0; int(j) < cap(cht.buf) && i < cht.buf[j].Size(); i++ {
		y := cht.buf[j].Get(i)
		if x == *y {
			_ = cht.buf[j].Remove(i)
			cht.len--
			return y
		}
	}
	return nil
}

func (cht ChainedHashTable) Find(x interface{}) *interface{} {
	if cht.len == 0 {
		return nil
	}
	j := cht.hash(x)
	for i := 0; int(j) < cap(cht.buf) && i < cht.buf[j].Size(); i ++ {
		if x == *(cht.buf[j].Get(i)) {
			return cht.buf[j].Get(i)
		}
	}
	return nil
}

func (cht ChainedHashTable) hash(x interface{}) uint64 {
	c := util.HashCode(x)
	h := (cht.z * c) >> (64 - cht.d)
	return h
}

func (cht *ChainedHashTable) resize() {
	newBuf := make([]arraystack.ArrayStack, util.Max(2*cap(cht.buf), 1))
	copy(newBuf, cht.buf)
	cht.buf = newBuf
}
