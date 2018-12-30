package arraystack

import "github.com/miliya612/ods/util"

// ArrayStack implements List interface with backing array
// especially, it implements LIFO Stack interface effectively
type ArrayStack struct {
	buf []interface{}
	// len represents number of existing elements for its buf
	// cap represents capacity of its buf
	len, cap int
}

func New(l int) ArrayStack {
	return ArrayStack{
		buf: make([]interface{}, l),
		len: 0,
		cap: l,
	}
}

func (as ArrayStack) Size() int {
	return as.len
}

func (as ArrayStack) Get(i int) *interface{} {
	if i > as.len || i < 0 {panic("index out bound of array!")}

	return &as.buf[i]
}

func (as *ArrayStack) Set(i int, x interface{}) (old interface{}) {
	if i > as.len || i < 0 {panic("index out bound of array!")}
	old, as.buf[i] = as.buf[i], x
	return
}

func (as *ArrayStack) Add(i int, x interface{}) {
	if i > as.len || i < 0 {panic("index out bound of array!")}

	if as.len+1 > as.cap {
		as.resize()
	}
	for j := as.len; j > i; j-- {
		as.buf[j] = as.buf[j-1]
	}
	as.buf[i] = x
	as.len++
}

func (as *ArrayStack) Remove(i int) interface{} {
	x := as.buf[i]
	for j := i; j < as.len-1; j++ {
		as.buf[j] = as.buf[j+1]
	}
	as.len--
	if as.cap >= 3*as.len {
		as.resize()
	}
	return x
}

// Push add an element at the tail of array
func (as *ArrayStack) Push(x interface{}) {
	as.Add(as.len, x)
}

// Pop remove an element at the tail of array
func (as *ArrayStack) Pop() interface{} {
	return as.Remove(as.len - 1)
}

func (as *ArrayStack) resize() {
	as.cap = util.Max(2*as.len, 1)
	newBuf := make([]interface{}, as.cap)

	for i := 0; i < as.len; i++ {
		newBuf[i] = as.buf[i]
	}

	as.buf = newBuf
}
