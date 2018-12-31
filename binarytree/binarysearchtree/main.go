package binarysearchtree

import "github.com/miliya612/ods/util"

type node struct {
	left, right, parent *node
	x                   interface{}
}

func newNode() *node {
	return &node{}
}

// BinarySearchTree represents an unbalanced binary tree which satisfies following:
//   - for a node, n, every data value stored in the subtree rooted at n.left is less than n.x
//   - every data value stored in the subtree rooted at n.right is greater than n.x
type BinarySearchTree struct {
	r   *node
	len int
}

func NewBST() *BinarySearchTree {
	return &BinarySearchTree{}
}

// Size implements SSet#Size interface
// it returns the size of tree nodes
func (bst BinarySearchTree) Size() int {
	return bst.len
}

// Find implements SSet#Find interface
// it seeks and returns the given element in the set if it exists.
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
	if z == nil {
		return nil
	}
	return &z.x
}

// Add implements SSet#Add interface
// it inserts a node including a given element to the tree as the leaf
// return true, or return false if the element has already been existed
func (bst *BinarySearchTree) Add(x interface{}) bool {
	p := bst.findLast(x)
	n := newNode()
	n.x = x
	return bst.addChild(p, n)
}

// Remove implements SSet#Remove interface
// it deletes a node including a given element from the tree
// to maintain the property of the tree, replace the node with the appropriate one
func (bst *BinarySearchTree) Remove(x interface{}) interface{} {
	p := bst.findLast(x)
	if p != nil && util.Compare(x, p.x) == 0 {
		bst.removeNode(p)
		return x
	}
	return nil
}

// findEQ seeks and returns the given element in the set
// If that is found, return that, otherwise return nil
func (bst BinarySearchTree) findEQ(x interface{}) interface{} {
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

func (bst BinarySearchTree) findLast(x interface{}) *node {
	var w, prev *node = bst.r, nil
	for w != nil {
		prev = w
		comp := util.Compare(x, w.x)
		if comp < 0 {
			w = w.left
		} else if comp > 0 {
			w = w.right
		} else {
			return w
		}
	}
	return prev
}

func (bst *BinarySearchTree) addChild(p, n *node) bool {
	// add a node to empty tree
	if p == nil {
		bst.r = n
	} else {
		comp := util.Compare(n.x, p.x)
		if comp < 0 {
			p.left = n
		} else if comp > 0 {
			p.right = n
		} else {
			// element had been added
			return false
		}
		n.parent = p
	}
	bst.len++
	return true
}

// splice deletes a node. if it has one child, detach its child to its parent
func (bst *BinarySearchTree) splice(n *node) {
	var c, p *node
	if n.left != nil {
		c = n.left
	} else {
		c = n.right
	}
	// if n is root, remove n and set its child to root
	if n == bst.r {
		bst.r = c
		p = nil
	} else {
		p = n.parent
		if p.left == n {
			p.left = c
		} else {
			p.right = c
		}
	}
	if c != nil {
		c.parent = p
	}
	bst.len--
}

func (bst *BinarySearchTree) removeNode(n *node) {
	// if the node has only one child, remove it and splice its child
	if n.left == nil || n.right == nil {
		bst.splice(n)
		n = nil
	} else {
		// find the smallest value greater than n.left
		// set it at the n's position
		w := n.right
		for w.left != nil {
			w = w.left
		}
		n.x = w.x
		bst.splice(w)
		w = nil
	}
}
