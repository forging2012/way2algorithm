package bplus

import (
	"fmt"
)

// B+Tree
type BPlusTree struct {
	root *Node
}

func New() *BPlusTree {
	return &BPlusTree{
		root: NewNode(true),
	}
}

// 插入key-value
func (tree *BPlusTree) Insert(key int, value interface{}) {
	if tree.root.isFull() {
		tree.root = tree.root.split()
	}
	tree.root.insert(key, value)
}

// 删除一个key
func (tree *BPlusTree) Delete(key int) {
	tree.root.delete(key)
	if tree.root.n == 0 && tree.root.children[0] != nil {
		tree.root = tree.root.children[0].(*Node)
	}
}

// 查找key有可能在的节点
func (tree *BPlusTree) SearchNode(key int) *Node {
	root := tree.root
	for {
		if root.isLeaf {
			return root
		}
		root = root.children[root.indexKey(key)].(*Node)
	}
	return nil
}

// 查找一个key
func (tree *BPlusTree) Search(key int) (value interface{}, ok bool) {
	node := tree.SearchNode(key)
	if i := node.indexKey(key); i != 0 && node.keys[i-1] == key {
		value, ok = node.children[i-1], true
	}
	return
}

// 查找闭区间[lowKey, highKey]之间的数据
func (tree *BPlusTree) Range(lowKey, highKey int) (values []interface{}) {
	if lowKey > highKey {
		return
	}

	lowNode, highNode := tree.SearchNode(lowKey), tree.SearchNode(highKey)

	i := lowNode.indexKey(lowKey)
	if i != 0 && lowNode.keys[i-1] == lowKey {
		values = append(values, lowNode.children[i-1])
	}

	for lowNode != highNode {
		values = append(values, lowNode.children[i:T<<1-1]...)
		lowNode, i = lowNode.rightLeaf(), 0
	}

	values = append(values, lowNode.children[i:highNode.indexKey(highKey)]...)
	return values
}

// TODO: DELETE
func (tree *BPlusTree) traverse() {
	queue := []*Node{tree.root}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			head := queue[0]
			queue = queue[1:]

			fmt.Print(head.keys[:head.n])

			if !head.isLeaf {
				for i := 0; i < head.n+1; i++ {
					queue = append(queue, head.children[i].(*Node))
				}
			}
		}
		fmt.Println()
	}
}
