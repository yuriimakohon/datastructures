package heap

import "errors"

var (
	ErrIdxOutOfRange = errors.New("index is out of range")
	ErrHeapIsEmpty   = errors.New("heap is empty")
)
