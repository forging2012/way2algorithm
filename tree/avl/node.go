package avl

import (
	"fmt"
	"way2algorithm/util"
)

type Node struct {
	key    int
	height int
	left   *Node
	right  *Node
}

func NewNode(key int) *Node {
	return &Node{
		key:    key,
		height: 1,
	}
}

func (node *Node) Height() int {
	if node == nil {
		return 0
	}
	return node.height
}

func (node *Node) resetHeight() {
	node.height = util.Max(node.left.Height(), node.right.Height()) + 1
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

// 左旋
func (node *Node) rotateLeft() *Node {
	right := node.right
	node.right = right.left
	right.left = node

	node.resetHeight()
	right.resetHeight()
	return right
}

// 右旋
func (node *Node) rotateRight() *Node {
	left := node.left
	node.left = left.right
	left.right = node

	node.resetHeight()
	left.resetHeight()
	return left
}

// 树结构改变后，维护树的性质，处理4种情形
// https://zh.wikipedia.org/wiki/AVL%E6%A0%91
func (node *Node) fixup() *Node {
	if node.left.Height()-node.right.Height() > 1 {
		if node.left.right.Height() > node.left.left.Height() {
			node.left = node.left.rotateLeft()
		}
		node = node.rotateRight()
	} else if node.right.Height()-node.left.Height() > 1 {
		if node.right.left.Height() > node.right.right.Height() {
			node.right = node.right.rotateRight()
		}
		node = node.rotateLeft()
	}

	node.resetHeight()
	return node
}

// 插入方式同BST，额外需要做的工作是维护AVL树的性质
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

	return node.fixup()
}

// 删除方式同BST, 同样需要维护树的性质
func (node *Node) delete(key int) *Node {
	if node == nil {
		return nil
	}

	if key == node.key {
		if node.left == nil && node.right == nil {
			return nil
		}

		if node.left.Height() > node.right.Height() {
			precursor := node.precursor()
			node.key = precursor.key
			node.left = node.left.delete(precursor.key)
		} else {
			successor := node.successor()
			node.key = successor.key
			node.right = node.right.delete(successor.key)
		}
	} else if key < node.key {
		node.left = node.left.delete(key)
	} else {
		node.right = node.right.delete(key)
	}

	return node.fixup()
}

func (node *Node) String() string {
	if node == nil {
		return "<Node: nil>"
	}
	return fmt.Sprintf("<Node: %d>", node.key)
}
