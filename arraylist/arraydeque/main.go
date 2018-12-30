package arraydeque

import "github.com/miliya612/ods/util"

// ArrayDeque implements List interface with backing array
// especially, it implements Deque interface effectively
// ArrayDeque is realized by circular array.
type ArrayDeque struct {
	buf []interface{}
	// len represents number of existing elements for its buf
	// cap represents capacity of its buf
	// i represents index to be removed in next removal action, the head of circular array
	len, cap, i int
}

func New(l int) ArrayDeque {
	return ArrayDeque{
		buf: make([]interface{}, l),
		len: 0,
		cap: l,
	}
}

func (ad ArrayDeque) Size() int {
	return ad.len
}

func (ad ArrayDeque) Get(i int) *interface{} {
	if i > ad.len || i < 0 {panic("index out bound of array!")}

	return &ad.buf[(i+ad.i)&ad.cap]
}

func (ad *ArrayDeque) Set(i int, x interface{}) (old interface{}) {
	if i > ad.len || i < 0 {panic("index out bound of array!")}

	old = ad.buf[(i+ad.i)%ad.cap]
	ad.buf[(i+ad.i)%ad.cap] = x
	return old
}

func (ad *ArrayDeque) Add(i int, x interface{}) {
	if i > ad.len || i < 0 {panic("index out bound of array!")}

	if ad.len+1 > ad.cap {
		ad.resize()
	}
	if i < ad.cap/2 {
		ad.shiftHeadToL(i, x)
	} else {
		ad.shiftTailToR(i, x)
	}
	ad.buf[(i+ad.i)%ad.cap] = x
	ad.len++
}

func (ad *ArrayDeque) Remove(i int) interface{} {
	if i > ad.len || i < 0 {panic("index out bound of array!")}

	x := ad.buf[(i +ad.i) % ad.cap]
	// optimize a list operation depending on the position of the given index in the circular array
	if i < ad.cap/2 {
		// shift head to right
		for k := i; k > 0 ; k-- {
			ad.buf[(k+ad.i)%ad.cap] = ad.buf[(k-1+ad.i)%ad.cap]
		}
		ad.i = (ad.i+1) % ad.cap
	} else {
		// shift tail to left
		for k := i; k < ad.len-1; k++ {
			ad.buf[(k+ad.i)%ad.cap] = ad.buf[(k+1+ad.i)%ad.cap]
		}
	}
	ad.len--
	if ad.cap > 3 * ad.len {
		ad.resize()
	}
	return x
}

func (ad *ArrayDeque) shiftHeadToL(i int, x interface{}) {
	// if the head of circular array is on the head of behind array
	if ad.i == 0 {
		// set index at the end of behind array
		ad.i = ad.cap - 1
	} else {
		ad.i -= 1
	}
	for j := 0; j <= i-1; j++ {
		ad.buf[(j+ad.i)%ad.cap] = ad.buf[((j+1)+ad.i)%ad.cap]
	}
}

func (ad *ArrayDeque) shiftTailToR(i int, x interface{}) {
	for j := ad.len; j >= i; j -- {
		ad.buf[((j+1)+ad.i)%ad.cap] = ad.buf[(j+ad.i)%ad.cap]
	}
}

func (ad *ArrayDeque) resize() {
	ad.cap = util.Max(2*ad.len, 1)
	newBuf := make([]interface{}, ad.cap)
	for i := 0; i < ad.len; i++ {
		newBuf[i] = ad.buf[(i+ad.i)%ad.cap]
	}
	ad.buf = newBuf
	ad.i = 0
}
