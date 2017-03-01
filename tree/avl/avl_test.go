package avl

import (
	"math/rand"
	"testing"
	"time"
)

func TestAVLTree(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	tree := new(AVLTree)

	keys := make([]int, 1000)
	for i := 0; i < len(keys); i++ {
		keys[i] = rand.Intn(10000)
		tree.Insert(keys[i])
		if node := tree.Search(keys[i]); node == nil || node.key != keys[i] {
			t.FailNow()
		}
	}

	for _, key := range keys {
		tree.Delete(key)
		if node := tree.Search(key); node != nil {
			t.FailNow()
		}
	}
}
