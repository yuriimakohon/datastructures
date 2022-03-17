package heap

type Heap struct {
	arr  []int
	cmp  func(int, int) bool
	size int
}

func greater(a, b int) bool {
	return a > b
}
func less(a, b int) bool {
	return a < b
}

func NewMaxHeap() Heap {
	return Heap{cmp: greater}
}

func NewMinHeap() Heap {
	return Heap{cmp: less}
}

func (h *Heap) Insert(value int) {
	h.arr = append(h.arr, value)
	h.size++
	h.MakeHeap()
}

// Delete value from heap
func (h *Heap) Delete(value int) {
	last := h.size - 1

	for i := 0; i < h.size; i++ {
		if h.arr[i] == value {
			h.arr[i], h.arr[last] = h.arr[last], h.arr[i]
			h.arr = h.arr[:last]
			h.size--
			h.MakeHeap()
		}
	}
}

// Peek returns max or min(according to type) value of heap.
// If heap is empty - returns error
func (h *Heap) Peek() (int, error) {
	if h.size == 0 {
		return 0, ErrHeapIsEmpty
	}

	return h.arr[0], nil
}

// Extract return max or min(according to type) value of heap and remove it.
// If heap is empty - return error
func (h *Heap) Extract() (int, error) {
	if h.size == 0 {
		return 0, ErrHeapIsEmpty
	}

	value := h.arr[0]
	h.Delete(value)

	return value, nil
}

// Heapify heap starting from i index
func (h *Heap) heapify(i int) error {
	if i >= h.size || i < 0 {
		return ErrIdxOutOfRange
	}

	swap := i
	left := i*2 + 1
	right := i*2 + 2

	// find swap
	if left < h.size &&
		h.cmp(h.arr[left], h.arr[swap]) {
		swap = left
	}
	if right < h.size &&
		h.cmp(h.arr[right], h.arr[swap]) {
		swap = right
	}

	// if current node wasn't swap - swap and heapify the smaller or greatest child
	if i != swap {
		h.arr[swap], h.arr[i] = h.arr[i], h.arr[swap]
		return h.heapify(swap) // now swap contain the smallest child
	}
	return nil
}

// MakeHeap heapify all heap starting from first non-leaf node
func (h *Heap) MakeHeap() {
	for i := h.size/2 - 1; i >= 0; i-- {
		_ = h.heapify(i)
	}
}
