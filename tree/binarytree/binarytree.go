package binarytree

import (
	"fmt"
)

// 二叉搜索树
type BinaryTree struct {
	root *Node
}

// 插入
func (tree *BinaryTree) Insert(key int) {
	tree.root = tree.root.insert(key)
}

// 删除
func (tree *BinaryTree) Delete(key int) {
	tree.root = tree.root.delete(key)
}

// 查找
func (tree *BinaryTree) Search(key int) *Node {
	root := tree.root
	for root != nil {
		if key == root.key {
			return root
		}

		if key < root.key {
			root = root.left
		} else {
			root = root.right
		}
	}
	return nil
}

// NOTE: only for debugging
func (tree *BinaryTree) traverse() {
	queue := []*Node{tree.root}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			fmt.Print(node)
			if node != nil {
				queue = append(queue, node.left, node.right)
			}
		}
		fmt.Println()
	}
}
