package binarytree

import "testing"

func dummyBT() *BinaryTree {
	n1 := newNode()
	n2 := newNode()
	n3 := newNode()
	n4 := newNode()

	n1.left = n2
	n2.parent = n1
	n1.right = n3
	n3.parent = n1
	n2.right = n4
	n4.parent = n2

	return &BinaryTree{
		r: n1,
	}
}

func TestDepth(t *testing.T) {
	bt := dummyBT()

	r := bt.r
	tests := []struct {
		arg      *node
		expected int
	}{
		{r, 0},
		{r.left, 1},
		{r.right, 1},
		{r.left.right, 2},
	}
	for _, test := range tests {
		if ret := (r).depth(test.arg); ret != test.expected {
			t.Errorf("bt.Depth(%v) = %v", test.arg, ret)
		}
	}
}

func TestSize(t *testing.T) {
	bt := dummyBT()
	if ret := bt.Size(); ret != 4 {
		t.Errorf("bt.Size = %v", ret)
	}
}

func TestHeight(t *testing.T) {
	bt := dummyBT()
	if ret := bt.Height(); ret != 2 {
		t.Errorf("bt.Height() = %v", ret)
	}
}

/**
 * Recursive implementation
 */

func Test_recSize(t *testing.T) {
	bt := dummyBT()

	r := bt.r
	tests := []struct {
		arg      *node
		expected int
	}{
		{r, 4},
		{r.left, 2},
		{r.right, 1},
		{r.left.right, 1},
		{r.left.left, 0},
	}
	for _, test := range tests {
		if ret := (test.arg).recSize(); ret != test.expected {
			t.Errorf("size_subtree(%v) = %v", test.arg, ret)
		}
	}
}

func Test_recHeight(t *testing.T) {
	bt := dummyBT()

	r := bt.r
	tests := []struct {
		arg      *node
		expected int
	}{
		{r, 2},
		{r.left, 1},
		{r.right, 0},
		{r.left.right, 0},
	}
	for _, test := range tests {
		if ret := (test.arg).recHeight(); ret != test.expected {
			t.Errorf("recHeight(%v) = %v", test.arg, ret)
		}
	}
}

/**
 * Un-recursive implementation with pre-order search
 */

func Test_poSize(t *testing.T) {
	bt := dummyBT()

	r := bt.r
	tests := []struct {
		arg      *node
		expected int
	}{
		{r, 4},
		{r.left, 2},
		{r.right, 1},
		{r.left.right, 1},
		{r.left.left, 0},
	}
	for _, test := range tests {
		if ret := (test.arg).poSize(); ret != test.expected {
			t.Errorf("poSize(%v) = %v", test.arg, ret)
		}
	}
}

func Test_poHeight(t *testing.T) {
	bt := dummyBT()

	r := bt.r
	tests := []struct {
		arg      *node
		expected int
	}{
		{r, 2},
		{r.left, 1},
		{r.right, 0},
		{r.left.right, 0},
	}
	for _, test := range tests {
		if ret := (test.arg).poHeight(); ret != test.expected {
			t.Errorf("poHeight(%v) = %v", test.arg, ret)
		}
	}
}
