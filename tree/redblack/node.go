package redblack

import (
	"fmt"
)

// 哨兵节点, 也叫dummy node, 是链表/tree等数据结构常见的一种技巧. 它能够使得
// 代码更加简洁.
var sentinel *Node

func init() {
	sentinel = new(Node)
	sentinel.left = sentinel
	sentinel.right = sentinel
}

// 节点
type Node struct {
	isRed bool
	key   int
	val   interface{}
	left  *Node
	right *Node
}

func NewNode(key int, val interface{}) *Node {
	return &Node{
		key:   key,
		val:   val,
		isRed: true,
		left:  sentinel,
		right: sentinel,
	}
}

// 如果左右节点都是哨兵节点, 那么该节点是叶节点
func (node *Node) isLeaf() bool {
	return node.left == sentinel && node.right == sentinel
}

// 前驱
func (node *Node) precursor() *Node {
	precursor := node.left
	for precursor.right != sentinel {
		precursor = precursor.right
	}
	return precursor
}

// 后继
func (node *Node) successor() *Node {
	successor := node.right
	for successor.left != sentinel {
		successor = successor.left
	}
	return successor
}

// 左旋
func (node *Node) rotateLeft() *Node {
	right := node.right
	node.right = right.left
	right.left = node

	right.isRed, node.isRed = node.isRed, right.isRed
	return right
}

// 右旋
func (node *Node) rotateRight() *Node {
	left := node.left
	node.left = left.right
	left.right = node

	left.isRed, node.isRed = node.isRed, left.isRed
	return left
}

// 反转颜色
func (node *Node) flipColor() {
	node.isRed = !node.isRed
	node.left.isRed = !node.left.isRed
	node.right.isRed = !node.right.isRed
}

// fixup主要的作用是把4节点给去掉, 即红色节点不能连续
func (node *Node) fixup() *Node {
	if node.left.isRed {
		if node.left.right.isRed {
			node.left = node.left.rotateLeft()
		}

		if node.left.left.isRed {
			node = node.rotateRight()
		}
	} else if node.right.isRed {
		if node.right.left.isRed {
			node.right = node.right.rotateRight()
		}

		if node.right.right.isRed {
			node = node.rotateLeft()
		}
	}

	if node.left.isRed && node.right.isRed {
		node.flipColor()
	}

	return node
}

// 插入算法同BST, 唯一不同的是需要维护红黑树的性质
func (node *Node) insert(key int, val interface{}) *Node {
	if node == sentinel {
		return NewNode(key, val)
	}

	if key == node.key {
		node.val = val
	} else if key < node.key {
		node.left = node.left.insert(key, val)
	} else {
		node.right = node.right.insert(key, val)
	}

	return node.fixup()
}

// 使左孩子至少是3节点, 其过程类似B-Tree
func (node *Node) makeLeftRich() *Node {
	if node.left.isRed || node.left.left.isRed || node.left.right.isRed {
		return node
	}

	if node.right.isRed {
		node = node.rotateLeft()
		return node
	}

	node.flipColor()

	if node.right.left.isRed {
		node.right = node.right.rotateRight()
	}

	if node.right.right.isRed {
		node = node.rotateLeft()
		node.flipColor()
	}

	return node
}

// 使右孩子至少为3节点, 其过程同样类似B-Tree
func (node *Node) makeRightRich() *Node {
	if node.right.isRed || node.right.left.isRed || node.right.right.isRed {
		return node
	}

	if node.left.isRed {
		node = node.rotateRight()
		return node
	}

	node.flipColor()

	if node.left.right.isRed {
		node.left = node.left.rotateLeft()
	}

	if node.left.left.isRed {
		node = node.rotateRight()
		node.flipColor()
	}

	return node
}

// 删除算法, 同B-Tree的删除过程类似
func (node *Node) delete(key int) *Node {
	if key == node.key {
		if node.isLeaf() {
			return sentinel
		}

		if node.left != sentinel {
			precursor := node.precursor()
			node.key = precursor.key
			node.val = precursor.val

			node = node.makeLeftRich()
			node.left = node.left.delete(precursor.key)
		} else {
			successor := node.successor()
			node.key = successor.key
			node.val = successor.val

			node = node.makeRightRich()
			node.right = node.right.delete(successor.key)
		}
	} else if key < node.key && node.left != sentinel {
		node = node.makeLeftRich()
		node.left = node.left.delete(key)
	} else if key > node.key && node.right != sentinel {
		node = node.makeRightRich()
		node.right = node.right.delete(key)
	}

	return node.fixup()
}

// 验证红黑树的性质, 即
// - 不存在红色节点相连的情况, 即不存在4节点
// - 每条路径上的黑色节点个数相同
func (root *Node) validate() int {
	if root == sentinel {
		return 0
	}

	if root.isRed && (root.left.isRed || root.right.isRed) {
		panic("root and root.left are both red")
	}

	leftHeight, rightHeight := root.left.validate(), root.right.validate()
	if leftHeight != rightHeight {
		panic("not the same height")
	}

	if !root.isRed {
		leftHeight++
	}

	return leftHeight
}

// 定制String
func (root *Node) String() string {
	color := "black"
	if root.isRed {
		color = "red"
	}

	s := ""
	if root == sentinel {
		s = "sentinel"
	}
	return fmt.Sprintf("<Node %d %v %v>", root.key, color, s)
}
