package avl

import (
	"fmt"
)

// AVL树
type AVLTree struct {
	root *Node
}

func (tree *AVLTree) Insert(key int) {
	tree.root = tree.root.insert(key)
}

func (tree *AVLTree) Delete(key int) {
	tree.root = tree.root.delete(key)
}

// 查找同BST
func (tree *AVLTree) Search(key int) *Node {
	root := tree.root
	for root != nil {
		if root.key == key {
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

// TODO: delete
func (tree *AVLTree) traverse() {
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
