package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeap_Heapify(t *testing.T) {
	h := Heap{
		arr:  []int{1, 2, 3, 4, 5, 6},
		size: 6,
		cmp:  greater,
	}

	assert.ErrorIs(t, h.heapify(6), ErrIdxOutOfRange)
	assert.ErrorIs(t, h.heapify(-1), ErrIdxOutOfRange)

	assert.Nil(t, h.heapify(2)) // swap 3 and 6
	assert.Nil(t, h.heapify(1)) // swap 2 and 5
	assert.Nil(t, h.heapify(0)) // swap 1 and 6
	assert.Nil(t, h.heapify(2)) // swap 3 and 1

	for i := h.size/2 - 1; i >= 0; i-- {
		if i*2+1 < h.size {
			assert.Greater(t, h.arr[i], h.arr[i*2+1], "parent must be greater than left child")
		}
		if i*2+2 < h.size {
			assert.Greater(t, h.arr[i], h.arr[i*2+2], "parent must ge greater than right child")
		}
	}
}

func TestMinHeap_Heapify(t *testing.T) {
	h := Heap{
		arr:  []int{6, 5, 4, 3, 2, 1},
		size: 6,
		cmp:  less,
	}

	assert.Error(t, ErrIdxOutOfRange, h.heapify(4))
	assert.Error(t, ErrIdxOutOfRange, h.heapify(-1))

	assert.Nil(t, h.heapify(2)) // swap 4 and 1
	assert.Nil(t, h.heapify(1)) // swap 2 and 5
	assert.Nil(t, h.heapify(0)) // swap 1 and 6
	assert.Nil(t, h.heapify(2)) // swap 4 and 1

	for i := h.size/2 - 1; i >= 0; i-- {
		if i*2+1 < h.size {
			assert.Less(t, h.arr[i], h.arr[i*2+1], "parent must be less than left child")
		}
		if i*2+2 < h.size {
			assert.Less(t, h.arr[i], h.arr[i*2+2], "parent must ge less than right child")
		}
	}
}

func TestMaxHeap_MakeHeap(t *testing.T) {
	h := Heap{
		arr:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15, 11, 13, 12, 14, 10},
		size: 15,
		cmp:  greater,
	}

	h.MakeHeap()

	for i := h.size/2 - 1; i >= 0; i-- {
		if i*2+1 < h.size {
			assert.Greater(t, h.arr[i], h.arr[i*2+1], "parent must be greater than left child")
		}
		if i*2+2 < h.size {
			assert.Greater(t, h.arr[i], h.arr[i*2+2], "parent must ge greater than right child")
		}
	}
}

func TestMinHeap_MakeHeap(t *testing.T) {
	h := Heap{
		arr:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15, 11, 13, 12, 14, 10},
		size: 15,
		cmp:  less,
	}

	h.MakeHeap()

	for i := h.size/2 - 1; i >= 0; i-- {
		if i*2+1 < h.size {
			assert.Less(t, h.arr[i], h.arr[i*2+1], "parent must be less than left child")
		}
		if i*2+2 < h.size {
			assert.Less(t, h.arr[i], h.arr[i*2+2], "parent must ge less than right child")
		}
	}
}

func TestHeap_Insert(t *testing.T) {
	h := NewMaxHeap()

	h.Insert(1)
	h.Insert(2)
	h.Insert(3)

	assert.Equal(t, 3, h.size)
	assert.Equal(t, 3, h.arr[0])
	assert.Equal(t, 1, h.arr[1])
	assert.Equal(t, 2, h.arr[2])

	h.Insert(9)
	h.Insert(10)

	assert.Equal(t, 5, h.size)
	assert.Equal(t, 10, h.arr[0])
	assert.Equal(t, 9, h.arr[1])
}

func TestHeap_Delete(t *testing.T) {
	toDelete := []int{100, 7, 10, 12, 5, 9, 72, 100, 88, 53, 53}
	h := NewMaxHeap()

	for i := 1; i <= 100; i++ {
		h.Insert(i)
	}

	for _, del := range toDelete {
		h.Delete(del)
		for _, i := range h.arr {
			assert.NotEqualf(t, del, i, "item %v must be deleted", del)
		}
	}
}

func TestHeap_Peek(t *testing.T) {
	h := NewMaxHeap()

	res, err := h.Peek()
	assert.Equal(t, 0, res)
	assert.ErrorIs(t, err, ErrHeapIsEmpty)

	h.Insert(5)
	h.Insert(9)
	h.Insert(14)
	h.Insert(2)

	res, err = h.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 14, res)

	h.cmp = less

	h.Insert(10)
	h.Insert(3)

	res, err = h.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 2, res)
}

func TestHeap_Extract(t *testing.T) {
	h := NewMaxHeap()

	res, err := h.Extract()
	assert.Equal(t, 0, res)
	assert.ErrorIs(t, err, ErrHeapIsEmpty)

	h.Insert(8)
	h.Insert(10)
	h.Insert(4)
	h.Insert(7)

	res, err = h.Extract()
	assert.NoError(t, err)
	assert.Equal(t, 10, res)

	for _, i := range h.arr {
		assert.NotEqual(t, 10, i)
	}

	h.cmp = less

	h.Insert(31)
	h.Insert(92)

	res, err = h.Extract()
	assert.NoError(t, err)
	assert.Equal(t, 4, res)

	for _, i := range h.arr {
		assert.NotEqual(t, 4, i)
	}
}
