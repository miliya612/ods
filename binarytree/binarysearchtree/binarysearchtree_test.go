package binarysearchtree

import (
	"testing"
)

func TestNew(t *testing.T) {
	if ret := NewBST().len; ret != 0 {
		t.Errorf("BinarySearchTree.NewBST().len = %v", ret)
	}
}

func TestAdd(t *testing.T) {
	n := 10
	bst := NewBST()

	for i := 0; i < n; i++ {
		ret := bst.Add(i)
		if bst.len != i+1 {
			t.Errorf("bst.len = %v at %v th Add", bst.len, i+1)
		}
		if ret == false {
			t.Errorf("Add returned false unexpectedly")
		}
	}

	for i := 0; i < n; i++ {
		ret := bst.Add(i)
		if bst.len != n {
			t.Errorf("bst.len = %v", bst.len)
		}
		if ret == true {
			t.Errorf("Add returned true unexpectedly")
		}
	}
}

func TestFindEQ(t *testing.T) {
	n := 10
	bst := NewBST()

	for i := 0; i < n; i++ {
		ret := bst.findEQ(i)
		if ret == true {
			t.Errorf("Add returned non-nil unexpectedly, ret=%v", ret)
		}
	}

	for i := 0; i < n; i++ {
		bst.Add(i)
		ret := bst.findEQ(i)
		if ret == false {
			t.Errorf("Add returned nil unexpectedly")
		}
	}

	ret := bst.findEQ(n + 123)
	if ret == true {
		t.Errorf("Add returned nil unexpectedly")
	}
}

func TestFind(t *testing.T) {
	n := 10
	bst := NewBST()

	for i := 0; i < n; i++ {
		ret := bst.Find(i)
		if ret != nil {
			t.Errorf("non-nil unexpectedly, ret=%d", ret)
		}
	}

	for i := 0; i < n; i++ {
		bst.Add(i)
		ret := bst.Find(i)
		if ret == nil {
			t.Errorf("Find returned nil unexpectedly")
		}
	}

	ret := bst.Find(n - 123)
	if *ret != 0 {
		t.Errorf("unexpectedly returned %v", *ret)
	}
}

func TestRemove(t *testing.T) {
	x := -12345
	bst := NewBST()

	if bst.Remove(x) == true {
		t.Errorf("Add returned true unexpectedly")
	}

	bst.Add(x)
	if bst.Remove(x) == false {
		t.Errorf("Add returned false unexpectedly")
	}

	if bst.Find(x) != nil {
		t.Errorf("Add returned non-nil unexpectedly")
	}
}
