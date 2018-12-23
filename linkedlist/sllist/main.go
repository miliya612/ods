package sllist

type node struct {
	x    interface{}
	next *node
}

func NewNode(x interface{}) *node {
	return &node{
		x:    x,
		next: nil,
	}
}

// SLList represents singly-linked list
type SLList struct {
	head *node
	tail *node
	len  int
}

func NewSLList() SLList {
	return SLList{}
}

/*
	List interface
 */

 func (sl *SLList) Size() int {
 	return sl.len
 }

func (sl *SLList) Get(i int) interface{} {
	if i > sl.len || i < 0 {panic("index out bound of array!")}
	u := sl.head
	for j := 0; j < sl.len; j++ {
		if i == j {
			return u.x
		}
		u = u.next
	}
	return nil
}

func (sl *SLList) Set(i int, x interface{}) interface{} {
	if i > sl.len || i < 0 {panic("index out bound of array!")}
	var old interface{}

	u := sl.head
	for j := 0; j < sl.len; j++ {
		if i == j {
			old = u.x
			u.x = x
			break
		}
		u = u.next
	}
	return old
}

 func (sl *SLList) Add(i int, x interface{}) {
 	p := NewNode(x)
	 u := sl.head
	 for j := 0; j < sl.len; j++ {
		 if j == (i-1) {
			 p.next = u.next
			 u.next = p
			 break
		 }
		 u = u.next
	 }
 }

 func (sl *SLList) Remove(i int) interface{} {
	 x := sl.tail.x
	 u := sl.head
	 for j := 0; j < i; j++ {
		 if j == i-1 {
			 u.next = u.next.next
			 break
		 }
		 u = u.next
	 }
	 return x
 }

/*
	Stack interface(LIFO)
 */

// Push implements the push method in Stack interface
// add element to the head of list
// O(1)
func (sl *SLList) Push(x interface{}) interface{} {
	u := NewNode(x)
	u.next = sl.head
	sl.head = u
	if sl.len == 0 {
		sl.tail = u
	}
	sl.len++
	return x
}

// Pop implements the pop method in Stack interface
// remove element at the head of list
// O(1)
func (sl *SLList) Pop() interface{} {
	if sl.len == 0 {
		return nil
	}
	x := sl.head.x
	sl.head = sl.head.next
	if (sl.len - 1) == 0 {
		sl.tail = nil
	}
	return x
}

/*
	Queue interface(FIFO)
 */

// Enqueue implements the add method in Queue interface
// add element to the tail of list
// O(1)
func (sl *SLList) Enqueue(x interface{}) bool {
	u := NewNode(x)
	if sl.len == 0 {
		sl.head = u
	} else {
		sl.tail.next = u
	}
	sl.tail = u
	sl.len++
	return true
}

// Dequeue implements the remove method in Queue interface
// remove element at the head of list
// its implementation is same as the one of Pop()
// O(1)
func (sl *SLList) Dequeue() interface{} {
	if sl.len == 0 {
		return nil
	}
	x := sl.head.x
	sl.head = sl.head.next
	if (sl.len - 1) == 0 {
		sl.tail = nil
	}
	return x
}

/*
	Deque interface
 */

 func (sl *SLList) AddFirst(x interface{}) {
	 u := NewNode(x)
	 u.next = sl.head
	 sl.head = u
	 if sl.len == 0 {
		 sl.tail = u
	 }
	 sl.len++
 }

func (sl *SLList) RemoveFirst() interface{} {
	if sl.len == 0 {
		return nil
	}
	x := sl.head.x
	sl.head = sl.head.next
	if (sl.len - 1) == 0 {
		sl.tail = nil
	}
	return x
}
func (sl *SLList) AddLast(x interface{}){
	u := NewNode(x)
	if sl.len == 0 {
		sl.head = u
	} else {
		sl.tail.next = u
	}
	sl.tail = u
	sl.len++
}

func (sl *SLList) RemoveLast() interface{}{
	x := sl.tail.x
	u := sl.head
	for j := 0; j < sl.len; j++ {
		if j == sl.len-2 {
			sl.tail =  u
			break
		}
		u = u.next
	}
	return x
}