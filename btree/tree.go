package btree

import (
	"fmt"
	"sync"
)

// BinarySearchTree the binary search tree of Items
type BinarySearchTree struct {
	root *Node
	lock sync.RWMutex
}

func (bst *BinarySearchTree) Insert(key int, value interface{}) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	n := &Node{key, value, nil, nil, nil}
	if bst.root == nil {
		bst.root = n
	} else {
		insertNode(bst.root, n)
	}
}

func (bst *BinarySearchTree) InOrderTraverse(f func(interface{})) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	bst.root.inOrderTraverse(f)
}

func (bst *BinarySearchTree) PreOrderTraverse(f func(interface{})) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	bst.root.preOrderTraverse(f)
}

func (bst *BinarySearchTree) PostOrderTraverse(f func(interface{})) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	bst.root.postOrderTraverse(f)
}

func (bst *BinarySearchTree) Min() *interface{} {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	return &bst.root.min().value
}

func (bst *BinarySearchTree) Max() *interface{} {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	return &bst.root.max().value
}

func (bst *BinarySearchTree) Search(key int) bool {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	n := bst.root.search(key)
	if n == nil {
		return false
	}
	return true
}

func (bst *BinarySearchTree) Remove(key int) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	remove(bst.root, key)
}

// String prints a visual representation of the tree
func (bst *BinarySearchTree) String() {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	fmt.Println("------------------------------------------------")
	bst.root.stringify(0)
	fmt.Println("------------------------------------------------")
}
