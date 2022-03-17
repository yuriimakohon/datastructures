package btree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchSuccessor(t *testing.T) {
	tree := BinarySearchTree{}
	tree.Insert(4, "4")
	tree.Insert(2, "2")
	tree.Insert(3, "3")
	tree.Insert(1, "1")
	tree.Insert(6, "6")
	tree.Insert(7, "7")
	tree.Insert(5, "5")

	assert.Equal(t, "4", tree.root.search(3).searchSuccessor().value)
	assert.Equal(t, "2", tree.root.search(1).searchSuccessor().value)
	assert.Equal(t, "5", tree.root.search(4).searchSuccessor().value)
	assert.Nil(t, tree.root.search(7).searchSuccessor())
}

func TestSearchPredecessor(t *testing.T) {
	tree := BinarySearchTree{}
	tree.Insert(4, "4")
	tree.Insert(2, "2")
	tree.Insert(3, "3")
	tree.Insert(1, "1")
	tree.Insert(6, "6")
	tree.Insert(7, "7")
	tree.Insert(5, "5")

	assert.Equal(t, "2", tree.root.search(3).searchPredecessor().value)
	assert.Equal(t, "6", tree.root.search(7).searchPredecessor().value)
	assert.Equal(t, "3", tree.root.search(4).searchPredecessor().value)
	assert.Nil(t, tree.root.search(1).searchPredecessor())
}
