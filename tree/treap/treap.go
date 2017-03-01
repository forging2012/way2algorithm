package treap

import (
	"fmt"
)

// 树堆
type Treap struct {
	root *Node
}

// 插入一个key到树堆中
func (treap *Treap) Insert(key int) {
	treap.root = treap.root.insert(key)
}

// 从树堆中删除一个key
func (treap *Treap) Delete(key int) {
	treap.root = treap.root.delete(key)
}

// 查询key
func (treap *Treap) Search(key int) *Node {
	root := treap.root
	for root != nil && root.key != key {
		if key < root.key {
			root = root.left
		} else {
			root = root.right
		}
	}
	return root
}

// NOTE: this is only for debug
func (treap *Treap) traverse() {
	queue := []*Node{treap.root}
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
