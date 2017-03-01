package binarytree

import "fmt"

// 二叉树节点
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

	if key <= node.key {
		node.left = node.left.insert(key)
	} else {
		node.right = node.right.insert(key)
	}

	return node
}

// 前驱
func (node *Node) precursor() *Node {
	precursor := node.left
	for precursor != nil && precursor.right != nil {
		precursor = precursor.right
	}
	return precursor
}

// 后继
func (node *Node) successor() *Node {
	successor := node.right
	for successor != nil && successor.left != nil {
		successor = successor.left
	}
	return successor
}

func (node *Node) delete(key int) *Node {
	if node == nil {
		return nil
	}

	if key == node.key {
		if precursor := node.precursor(); precursor != nil {
			node.key = precursor.key
			node.left = node.left.delete(precursor.key)
		} else if successor := node.successor(); successor != nil {
			node.key = successor.key
			node.right = node.right.delete(successor.key)
		} else {
			return nil
		}
		return node
	}

	if key < node.key {
		node.left = node.left.delete(key)
	} else {
		node.right = node.right.delete(key)
	}

	return node
}

func (node *Node) String() string {
	if node == nil {
		return "<Node: nil>"
	}
	return fmt.Sprintf("<Node: %d>", node.key)
}
