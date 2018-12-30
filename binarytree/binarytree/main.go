package binarytree

import "github.com/miliya612/ods/util"

type node struct {
	left, right, parent *node
}

func newNode() *node {
	return &node{}
}

func (n *node) nextWithPO(prev *node) (next *node) {
	if prev == n.parent {
		if n.left != nil {
			next = n.left
		} else if n.right != nil {
			next = n.right
		} else {
			next = n.parent
		}
	} else if prev == n.left {
		if n.right != nil {
			next = n.right
		}
	} else {
		next = n.parent
	}
	return
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

// recSize returns number of nodes composing the tree with recursion
// NOTICE: it can cause stack over flow
func (n *node) recSize() int {
	if n == nil {
		return 0
	}
	return 1 + (n.left).recSize() + (n.right).recSize()
}

// sizePO returns number of nodes composing the tree
// it searches a tree like following: parent -> left -> right -> parent
func (n *node) sizePO() int {
	var prev, next *node
	var s int
	u := n
	for u != nil {
		if prev == u.parent {
			s++
		}
		next = u.nextWithPO(prev)
		prev = u
		u = next
	}
	return s
}

// Height returns the height of the tree
// it traverses a tree recursively
func (bt BinaryTree) Height() int {
	return (bt.r).recHeight()
}

func (n *node) recHeight() int {
	if n == nil {
		return -1
	}
	return 1 + util.Max((n.left).recHeight(), (n.right).recHeight())
}

func (bt BinaryTree) Traverse() {
	(bt.r).recTraverse()
}

// recTraverse traverses a tree with recursion
// NOTICE: it can cause stack over flow
func (n *node) recTraverse() {
	if n == nil {
		return
	}
	(n.left).recTraverse()
	(n.right).recTraverse()
}

// poTraverse traverses a tree with pre-order without recursion
// it searches a tree like following: parent -> left -> right -> parent
func (n *node) poTraverse() {
	var prev, next *node
	for n != nil {
		next = n.nextWithPO(prev)
		prev = n
		n = next
	}
}

