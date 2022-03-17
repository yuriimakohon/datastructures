package btree

import "fmt"

// Node a single node that composes the tree
type Node struct {
	key    int
	value  interface{}
	left   *Node // left
	right  *Node // right
	parent *Node // parent
}

func (n *Node) inOrderTraverse(f func(interface{})) {
	if n != nil {
		n.left.inOrderTraverse(f)
		f(n.value)
		n.right.inOrderTraverse(f)
	}
}

func (n *Node) preOrderTraverse(f func(interface{})) {
	if n != nil {
		f(n.value)
		n.left.preOrderTraverse(f)
		n.right.preOrderTraverse(f)
	}
}

func (n *Node) postOrderTraverse(f func(interface{})) {
	if n != nil {
		n.left.postOrderTraverse(f)
		n.right.postOrderTraverse(f)
		f(n.value)
	}
}

func insertNode(node, newNode *Node) {
	if newNode.key < node.key {
		if node.left == nil {
			node.left, newNode.parent = newNode, node
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right, newNode.parent = newNode, node
		} else {
			insertNode(node.right, newNode)
		}
	}
}

func (n *Node) search(key int) *Node {
	if n == nil || key == n.key {
		return n
	}
	if key < n.key {
		return n.left.search(key)
	} else {
		return n.right.search(key)
	}
}

func (n *Node) max() *Node {
	if n == nil {
		return nil
	}

	curr := n
	for curr.right != nil {
		curr = curr.right
	}
	return curr
}

func (n *Node) min() *Node {
	if n == nil {
		return nil
	}

	curr := n
	for curr.left != nil {
		curr = curr.left
	}
	return curr
}

func (n *Node) searchSuccessor() *Node {
	if n.right != nil {
		return n.right.min()
	}
	curr := n

	for curr.parent != nil && curr != curr.parent.left {
		curr = curr.parent
	}
	return curr.parent
}

func (n *Node) searchPredecessor() *Node {
	if n.left != nil {
		return n.left.max()
	}
	curr := n

	for curr.parent != nil && curr != curr.parent.right {
		curr = curr.parent
	}
	return curr.parent
}

func remove(n *Node, key int) *Node {
	if n == nil {
		return nil
	}
	if key < n.key {
		n.left = remove(n.left, key)
		return n
	}
	if key > n.key {
		n.right = remove(n.right, key)
		return n
	}
	if n.left == nil && n.right == nil {
		n = nil
		return nil
	}
	if n.left == nil {
		n, n.right.parent = n.right, n.parent
		return n
	}
	if n.right == nil {
		n, n.left.parent = n.left, n.parent
		return n
	}
	leftmostrightside := n.right
	for {
		if leftmostrightside != nil && leftmostrightside.left != nil {
			leftmostrightside = leftmostrightside.left
		} else {
			break
		}
	}
	n.key, n.value = leftmostrightside.key, leftmostrightside.value
	n.right = remove(n.right, n.key)
	return n
}

func (n *Node) stringify(level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		n.left.stringify(level)
		fmt.Printf(format+"%d\n", n.key)
		n.right.stringify(level)
	}
}
