package dllist

import (
	"testing"
)

func TestNew(t *testing.T) {
	if ret := NewDLList().len; ret != 0 {
		t.Errorf("DLList.New().n = %v", ret)
	}
}

func TestAdd(t *testing.T) {
	n := 10
	dll := NewDLList()

	for i := 0; i < n; i++ {
		dll.Enqueue(i)
		if dll.len != i+1 {
			t.Errorf("dll.n = %v at %v th Enqueue", dll.len, i+1)
		}
	}
}

func TestPush(t *testing.T) {
	n := 10
	dll := NewDLList()

	for i := 0; i < n; i++ {
		dll.Push(i)
		if dll.len != i+1 {
			t.Errorf("dll.n = %v at %v th Push", dll.len, i+1)
		}
	}
}

func TestPushAndPop(t *testing.T) {
	n := 10
	dll := NewDLList()
	for i := 0; i < n; i++ {
		dll.Push(i)
	}

	for i := 0; i < n; i++ {
		if ret := dll.Pop(); ret != n-1-i {
			t.Errorf("%v th dll.Pop() = %v", i, ret)
		}
	}
}

func TestAddAndRemove(t *testing.T) {
	n := 10
	dll := NewDLList()
	for i := 0; i < n; i++ {
		dll.Enqueue(i)
	}

	for i := 0; i < n; i++ {
		if ret := dll.Dequeue(); ret != i {
			t.Errorf("%v th dll.Dequeue() = %v", i, ret)
		}
	}
}
