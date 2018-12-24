package dsif

type List interface {
	Size() int
	Get(i int) interface{}
	Set(i int, x interface{}) interface{}
	Add(i int, x interface{})
	Remove(i int) interface{}
}

// Queue represents FIFO queue
type Queue interface {
	// Enqueue add an element at the tail of array
	Enqueue(x interface{})
	// Dequeue remove an element at the head of array
	Dequeue() interface{}
}

// Stack represents LIFO queue
type Stack interface {
	// Push add an element at the tail of array
	Push(x interface{})
	// Pop remove an element at the tail of array
	Pop() interface{}
}

// Dequeue represents two-way queue
type Deque interface {
	AddFirst(x interface{})
	RemoveFirst() interface{}
	AddLast(x interface{})
	RemoveLast() interface{}
}