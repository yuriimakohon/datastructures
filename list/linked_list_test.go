package list

import "testing"
import "github.com/stretchr/testify/assert"

func TestLinkedList_PushHead(t *testing.T) {
	var l *LinkedList = nil

	err := l.PushHead(nil)
	assert.Equal(t, err, ErrNilLinkedList)

	l = new(LinkedList)

	assert.NoError(t, l.PushHead(43))

	assert.Equal(t, 1, l.size)
	assert.Equal(t, 43, l.head.Data)
	assert.Equal(t, 43, l.tail.Data)
	assert.Nil(t, l.head.next)
	assert.Nil(t, l.head.prev)
	assert.Nil(t, l.tail.next)
	assert.Nil(t, l.tail.prev)

	assert.NoError(t, l.PushHead(21))

	assert.Equal(t, 2, l.size)
	assert.Equal(t, 21, l.head.Data)
	assert.NotNil(t, l.head.next)
	assert.NotNil(t, l.tail.prev)
	assert.Equal(t, l.head.next, l.tail)
	assert.Equal(t, l.tail.prev, l.head)
	assert.Equal(t, 43, l.tail.Data)

	err = l.PushHead(80)
	assert.NoError(t, err)

	assert.Equal(t, 3, l.size)
	assert.Equal(t, 80, l.head.Data)
	assert.Equal(t, 21, l.head.next.Data)
	assert.Equal(t, 43, l.head.next.next.Data)
	assert.NotNil(t, l.head.next.prev)
	assert.NotNil(t, l.head.next.next.prev)
}

func TestLinkedList_PushTail(t *testing.T) {
	var l *LinkedList = nil

	err := l.PushTail(nil)
	assert.Equal(t, err, ErrNilLinkedList)

	l = new(LinkedList)

	assert.NoError(t, l.PushTail(43))

	assert.Equal(t, 1, l.size)
	assert.Equal(t, 43, l.head.Data)
	assert.Equal(t, 43, l.tail.Data)
	assert.Nil(t, l.head.next)
	assert.Nil(t, l.head.prev)
	assert.Nil(t, l.tail.next)
	assert.Nil(t, l.tail.prev)

	assert.NoError(t, l.PushTail(21))

	assert.Equal(t, 2, l.size)
	assert.Equal(t, 21, l.tail.Data)
	assert.NotNil(t, l.head.next)
	assert.NotNil(t, l.tail.prev)
	assert.Equal(t, l.head.next, l.tail)
	assert.Equal(t, l.tail.prev, l.head)
	assert.Equal(t, 43, l.head.Data)

	assert.NoError(t, l.PushTail(80))

	assert.Equal(t, 3, l.size)
	assert.Equal(t, 80, l.tail.Data)
	assert.Equal(t, 21, l.tail.prev.Data)
	assert.Equal(t, 43, l.tail.prev.prev.Data)
	assert.NotNil(t, l.tail.prev.next)
	assert.NotNil(t, l.tail.prev.prev.next)
}

func TestLinkedList_ForEach(t *testing.T) {
	var l *LinkedList = nil
	test := []int{1, 2, 3, 4, 5}

	err := l.PushTail(nil)
	assert.Equal(t, err, ErrNilLinkedList)

	l = new(LinkedList)

	assert.Error(t, l.ForEach(nil), ErrNilFunction)

	for _, v := range test {
		assert.NoErrorf(t, l.PushTail(v), "no error expected (%v)", v)
	}

	var result []int
	err = l.ForEach(func(node *Node) {
		result = append(result, node.Data.(int))
	})
	assert.NoError(t, err)
	assert.Equal(t, test, result)

	l2 := new(LinkedList)

	assert.NoError(t, l2.ForEach(func(node *Node) {}))
	assert.NoError(t, l2.PushHead(6))

	sum := 2
	err = l2.ForEach(func(node *Node) {
		sum += node.Data.(int)
	})
	assert.NoError(t, err)
	assert.Equal(t, 8, sum)
}

func TestLinkedList_PopHead(t *testing.T) {
	var l *LinkedList = nil
	tests := []int{1, 2, 3, 4, 5}

	err := l.PushTail(nil)
	assert.Equal(t, err, ErrNilLinkedList)

	l = new(LinkedList)
	assert.NoError(t, l.PopHead())

	for _, v := range tests {
		assert.NoError(t, l.PushTail(v))
	}
	assert.NoError(t, l.PopHead())
	assert.NoError(t, l.PopHead())
	assert.NoError(t, l.PopHead())
	var result []int

	assert.NoError(t, l.ForEach(func(node *Node) {
		result = append(result, node.Data.(int))
	}))
	assert.Equal(t, []int{4, 5}, result)

	assert.NoError(t, l.PopHead())
	assert.Nil(t, l.head.prev)
	assert.Nil(t, l.head.next)
	assert.Nil(t, l.tail.next)
	assert.Nil(t, l.tail.prev)
	assert.Equal(t, 5, l.head.Data)

	assert.NoError(t, l.PopHead())
	assert.Equal(t, 0, l.size)
	assert.Nil(t, l.head)
	assert.Nil(t, l.tail)
}

func TestLinkedList_PopTail(t *testing.T) {
	var l *LinkedList = nil
	tests := []int{1, 2, 3, 4, 5}

	err := l.PushTail(nil)
	assert.Equal(t, err, ErrNilLinkedList)

	l = new(LinkedList)
	assert.NoError(t, l.PopTail())

	for _, v := range tests {
		assert.NoError(t, l.PushTail(v))
	}
	assert.NoError(t, l.PopTail())
	assert.NoError(t, l.PopTail())
	assert.NoError(t, l.PopTail())
	var result []int

	assert.NoError(t, l.ForEach(func(node *Node) {
		result = append(result, node.Data.(int))
	}))
	assert.Equal(t, []int{1, 2}, result)

	assert.NoError(t, l.PopTail())
	assert.Nil(t, l.tail.next)
	assert.Nil(t, l.head.next)
	assert.Nil(t, l.tail.next)
	assert.Nil(t, l.tail.prev)
	assert.Equal(t, 1, l.head.Data)

	assert.NoError(t, l.PopTail())
	assert.Equal(t, 0, l.size)
	assert.Nil(t, l.head)
	assert.Nil(t, l.tail)
}

func TestLinkedLisSize(t *testing.T) {
	var l *LinkedList = nil

	_, err := l.Size()
	assert.Equal(t, err, ErrNilLinkedList)

	l = new(LinkedList)

	size, err := l.Size()
	assert.NoError(t, err)
	assert.Equal(t, 0, size)

	assert.NoError(t, l.PushHead(1))
	assert.NoError(t, l.PushHead(2))

	size, err = l.Size()
	assert.NoError(t, err)
	assert.Equal(t, 2, size)

	assert.NoError(t, l.PopTail())
	size, err = l.Size()
	assert.NoError(t, err)
	assert.Equal(t, 1, size)
}
