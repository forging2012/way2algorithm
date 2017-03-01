package treap

import (
	"fmt"
	"math/rand"
)

type Node struct {
	key      int
	left     *Node
	right    *Node
	priority float64
}

func NewNode(key int) *Node {
	return &Node{
		key:      key,
		priority: rand.Float64(),
	}
}

// 左旋
func (node *Node) rotateLeft() *Node {
	right := node.right
	node.right = right.left
	right.left = node
	return right
}

// 右旋
func (node *Node) rotateRight() *Node {
	left := node.left
	node.left = left.right
	left.right = node
	return left
}

// 插入一个key
func (node *Node) insert(key int) *Node {
	if node == nil {
		return NewNode(key)
	}

	if key <= node.key {
		node.left = node.left.insert(key)
		if node.left.priority > node.priority {
			node = node.rotateRight()
		}
	} else {
		node.right = node.right.insert(key)
		if node.right.priority > node.priority {
			node = node.rotateLeft()
		}
	}
	return node
}

// 删除一个key有两种方式: 一种是利用了二叉树的性质，跟删除普通的二叉树方式一样。
// 另外一种是利用了堆的性质，把要被删除的节点向下移动到叶节点，中间伴随左右旋
// 以维持堆的性质，最后删除叶节点即可。这里使用的是第二种方式，第一种方式见二
// 叉树.
func (node *Node) delete(key int) *Node {
	if node == nil {
		return nil
	}

	if key == node.key {
		if node.left == nil && node.right == nil {
			return nil
		}

		if (node.left != nil && node.right != nil &&
			node.left.priority > node.right.priority) || node.right == nil {
			node = node.rotateRight()
			node.right = node.right.delete(key)
		} else {
			node = node.rotateLeft()
			node.left = node.left.delete(key)
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

// 定制化String方法
func (node *Node) String() string {
	if node == nil {
		return "<Node: nil>"
	}
	return fmt.Sprintf("<Node: %d>", node.key)
}
