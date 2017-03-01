package binarytree

import (
	"math/rand"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	tree := new(BinaryTree)

	keys := rand.Perm(1000)
	for _, key := range keys {
		tree.Insert(key)
		if node := tree.Search(key); node == nil || node.key != key {
			t.Fail()
		}
	}

	for _, key := range keys {
		tree.Delete(key)
		if node := tree.Search(key); node != nil {
			t.Fatal(node)
		}
	}
}
