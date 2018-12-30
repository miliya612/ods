package binarytree

import "github.com/miliya612/ods/util"

// recSize returns number of nodes composing the tree with recursion
// NOTICE: it can cause stack over flow
func (n *node) recSize() int {
	if n == nil {
		return 0
	}
	return 1 + (n.left).recSize() + (n.right).recSize()
}

func (n *node) recHeight() int {
	if n == nil {
		return -1
	}
	return 1 + util.Max((n.left).recHeight(), (n.right).recHeight())
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