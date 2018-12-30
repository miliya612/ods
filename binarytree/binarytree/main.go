package binarytree

import "github.com/miliya612/ods/arraylist/arraydeque"

type node struct {
	left, right, parent *node
}

func newNode() *node {
	return &node{}
}

type BinaryTree struct {
	r *node
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{
		r: newNode(),
	}
}

func (bt BinaryTree) Depth(l *node) int {
	return (bt.r).depth(l)
}

func (n *node) depth(l *node) int {
	var d int
	for l != n {
		l = l.parent
		d++
	}
	return d
}

// Size returns number of nodes composing the tree
func (bt BinaryTree) Size() int {
	return (bt.r).recSize()
}

// Height returns the height of the tree
// it traverses a tree recursively
func (bt BinaryTree) Height() int {
	return (bt.r).recHeight()
}

func (bt BinaryTree) Traverse() {
	(bt.r).recTraverse()
}

func (bt BinaryTree) BFTraverse() {
	var q arraydeque.ArrayDeque
	if bt.r != nil {
		q.Add(q.Size(), bt.r)
	}
	for q.Size() > 0 {
		n := q.Remove(q.Size()-1).(*node)
		if n.left != nil {
			q.Add(q.Size(), n.left)
		}
		if n.right != nil {
			q.Add(q.Size(), n.right)
		}
	}
}
