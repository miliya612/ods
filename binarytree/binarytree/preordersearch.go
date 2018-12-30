package binarytree

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
		} else {
			next = n.parent
		}
	} else {
		next = n.parent
	}
	return
}

// poSize returns number of nodes composing the tree
// it searches a tree like following: parent -> left -> right -> parent
func (n *node) poSize() int {
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

	// TODO: なんとかしたい
	if n != nil && s == 0 {
		return 1
	}
	return s
}

func (n *node) poHeight() int {
	if n == nil {
		return -1
	}

	var prev, next *node
	var i, h int

	u := n
	for u != nil {
		if prev == u.parent {
			i++
			if i > h {
				h = i
			}
			if n.left != nil {
				next = u.left
			} else if u.right != nil {
				next = u.right
			} else {
				next = u.parent
			}
		} else if prev == u.left {
			if u.right != nil {
				next = u.right
			} else {
				next = u.parent
			}
		} else {
			i--
			next = u.parent
		}
		prev = u
		u = next
	}
	return h
}

// poTraverse traverses a tree with pre-order without recursion
// it searches a tree like following: parent -> left -> right -> parent
func (n *node) poTraverse() {
	var prev, next *node
	u := n
	for u != nil {
		next = u.nextWithPO(prev)
		prev = u
		u = next
	}
}
