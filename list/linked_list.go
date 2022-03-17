package list

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func (l *LinkedList) Size() (int, error) {
	if l == nil {
		return 0, ErrNilLinkedList
	}
	return l.size, nil
}

func (l *LinkedList) PushHead(data interface{}) error {
	if l == nil {
		return ErrNilLinkedList
	}

	node := newNode(data)

	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.next, l.head, l.head.prev = l.head, node, node
	}

	l.size++

	if l.size == 2 {
		l.tail.prev = node
	}

	return nil
}

func (l *LinkedList) PushTail(data interface{}) error {
	if l == nil {
		return ErrNilLinkedList
	}

	node := newNode(data)

	if l.tail == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail, l.tail.next, node.prev = node, node, l.tail
	}

	l.size++

	if l.size == 2 {
		l.head.next = node
	}
	return nil
}

func (l *LinkedList) PopHead() error {
	if l == nil {
		return ErrNilLinkedList
	}

	if l.head == nil {
		return nil
	}

	if l.head == l.tail {
		l.head, l.tail = nil, nil
	} else {
		l.head, l.head.next.prev = l.head.next, nil
	}

	l.size--
	return nil
}

func (l *LinkedList) PopTail() error {
	if l == nil {
		return ErrNilLinkedList
	}

	if l.tail == nil {
		return nil
	}

	if l.head == l.tail {
		l.head, l.tail = nil, nil
	} else {
		l.tail, l.tail.prev.next = l.tail.prev, nil
	}

	l.size--
	return nil
}

func (l *LinkedList) ForEach(f func(node *Node)) error {
	if l == nil {
		return ErrNilLinkedList
	}
	if f == nil {
		return ErrNilFunction
	}

	if l.size == 0 {
		return nil
	}

	curr := l.head
	for {
		f(curr)
		curr = curr.next
		if curr == nil {
			break
		}
	}

	return nil
}
