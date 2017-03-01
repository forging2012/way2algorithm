package redblack

import "fmt"

type RedBlackTree struct {
	root *Node
}

func New() *RedBlackTree {
	return &RedBlackTree{
		root: sentinel,
	}
}

func (tree *RedBlackTree) Search(key int) (val interface{}, ok bool) {
	root := tree.root
	for root != sentinel {
		if key == root.key {
			return root.val, true
		}

		if key < root.key {
			root = root.left
		} else {
			root = root.right
		}
	}
	return
}

func (tree *RedBlackTree) Insert(key int, val interface{}) {
	tree.root = tree.root.insert(key, val)
	tree.root.isRed = false
}

func (tree *RedBlackTree) Delete(key int) {
	if tree.root == sentinel {
		return
	}

	tree.root = tree.root.delete(key)
	tree.root.isRed = false
}

// NOTE: only for debugging
func (tree *RedBlackTree) traverse() {
	queue := []*Node{tree.root}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			fmt.Print(node)
			if node != sentinel {
				queue = append(queue, node.left, node.right)
			}
		}
		fmt.Println()
	}
}
