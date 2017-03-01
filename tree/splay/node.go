package splay

type Node struct {
	key   int
	left  *Node
	right *Node
}

func NewNode(key int) *Node {
	return &Node{
		key: key,
	}
}

func (node *Node) insert(key int) *Node {
	if node == nil {
		return NewNode(key)
	}

	if key == node.key {
		return node
	}

	if key < node.key {
		node.left = node.left.insert(key)
	} else {
		node.right = node.right.insert(key)
	}
	return node
}
