package btree

import (
	"fmt"
)

type BTree struct {
	root *BTreeNode
}

func New() *BTree {
	return &BTree{
		root: &BTreeNode{
			isLeaf: true,
		},
	}
}

func (tree *BTree) Insert(key int, val interface{}) {
	if tree.root.isFull() {
		tree.root = tree.root.split()
	}
	tree.root.insert(&entry{key: key, val: val})
}

// 查找方式类似BST，只不过节点的查询用二分来查找
func (btree *BTree) Search(key int) (val interface{}, ok bool) {
	root := btree.root
	for {
		i := root.indexKey(key)
		if i < root.n && root.entries[i].key == key {
			return root.entries[i].val, true
		}

		if root.isLeaf {
			return
		}

		root = root.children[i]
	}
}

// 从B树中删除一个key
func (tree *BTree) Delete(key int) {
	tree.root.delete(key)
	if tree.root.n == 0 && len(tree.root.children) > 0 {
		tree.root = tree.root.children[0]
	}
}

// BFS遍历树的结构，仅做debug之用
// TODO: delete
func (btree *BTree) traverse() {
	var translate = func(entries []*entry) []int {
		r := make([]int, len(entries))
		for i := 0; i < len(entries); i++ {
			r[i] = entries[i].key
		}
		return r
	}

	queue := []*BTreeNode{btree.root}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			head := queue[0]
			queue = queue[1:]

			fmt.Print(translate(head.entries[:head.n]))

			if !head.isLeaf {
				queue = append(queue, head.children[:head.n+1]...)
			}
		}
		fmt.Println()
	}
}
