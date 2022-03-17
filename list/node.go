package list

type Node struct {
	next *Node
	prev *Node
	Data interface{}
}

func newNode(data interface{}) *Node {
	return &Node{nil, nil, data}
}
