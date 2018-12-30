package binarysearchtree

import "github.com/miliya612/ods/util"

type node struct {
	left, right, parent *node
	x interface{}
}

func newNode() *node {
	return &node{}
}

type BinarySearchTree struct {
	r *node
}

func NewBST() *BinarySearchTree {
	return &BinarySearchTree{
		r: newNode(),
	}
}

// FindEQ seeks and returns the given element in the set
// If that is found, return that, otherwise return nil
func (bst BinarySearchTree) FindEQ(x int) interface{} {
	w := bst.r
	for w != nil {
		comp := util.Compare(x, w.x)
		if comp < 0 {
			w = w.left
		} else if comp > 0 {
			w = w.right
		} else {
			return w.x
		}
	}
	return nil
}

// Find implements SSet#Find interface
// Find seeks and returns the given element in the set if it exists.
// If not, find an element y in the set such that y equals x. Return y, or nil if no such element exists.
func (bst BinarySearchTree) Find(x interface{}) *interface{} {
	var w, z *node = bst.r, nil
	for w != nil {
		comp := util.Compare(x, w.x)
		if comp < 0 {
			z = w
			w = w.left
		} else if comp > 0 {
			w = w.right
		} else {
			return &w.x
		}
	}
	if x == nil {
		return nil
	}
	return &z.x
}


