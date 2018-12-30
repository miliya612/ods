package binarytree

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


