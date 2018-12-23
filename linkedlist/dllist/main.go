package dllist

type node struct {
	x          interface{}
	prev, next *node
}

func NewNode(x interface{}) *node {
	return &node{
		x:    x,
		prev: nil,
		next: nil,
	}
}

// DLList represents doubly-linked list
type DLList struct {
	// dummy is an empty node which is laid on after the tail of the list and before the head of the list
	dummy *node
	len   int
}

func NewDLList() *DLList {
	dummy := new(node)
	dummy.prev = dummy
	dummy.next = dummy
	return &DLList{
		dummy: dummy,
		len:   0,
	}
}

/*
	List interface
 */

func (dl *DLList) Size() int {
	return dl.len
}

func (dl *DLList) Get(i int) interface{} {
	return dl.getNode(i)
}

func (dl *DLList) Set(i int, x interface{}) interface{} {
	u := dl.getNode(i)
	old := u.x
	u.x = x
	return old
}

func (dl *DLList) Add(i int, x interface{}) {
	dl.addBefore(dl.getNode(i), x)
}

func (dl *DLList) Remove(i int) interface{} {
	w := dl.getNode(i)
	x := w.x
	dl.removeNode(w)
	return x
}

/*
	Queue interface(FIFO)
 */

func (dl *DLList) Enqueue(x interface{}) {
	dl.Add(dl.len, x)
}

func (dl *DLList) Dequeue() interface{} {
	return dl.Remove(0)
}

/*
	Stack interface(LIFO)
 */

func (dl *DLList) Push(x interface{}) {
	dl.Add(dl.len, x)
}

func (dl *DLList) Pop() interface{} {
	return dl.Remove(dl.len - 1)
}

func (dl *DLList) getNode(i int) *node {
	var p *node
	if i < (dl.len / 2) {
		p = dl.dummy.next
		for j := 0; j < i; j++ {
			p = p.next
		}
	} else {
		p = dl.dummy
		for j := dl.len; j > i; j-- {
			p = p.prev
		}
	}
	return p
}

func (dl *DLList) addBefore(w *node, x interface{}) *node {
	u := NewNode(x)
	u.prev = w.prev
	u.next = w
	u.next.prev = u
	u.prev.next = u
	dl.len++
	return u
}

func (dl *DLList) removeNode(w *node) {
	w.prev.next = w.next
	w.next.prev = w.prev
	w = nil
	dl.len--
}
