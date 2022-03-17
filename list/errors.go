package list

import "errors"

var (
	ErrNilLinkedList = errors.New("linked list is nil")
	ErrNilFunction   = errors.New("given function is nil")
)
