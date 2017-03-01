package treap

import (
	"math/rand"
	"testing"
)

func TestTreap(t *testing.T) {
	treap := new(Treap)

	keys := rand.Perm(1000)
	for _, key := range keys {
		treap.Insert(key)
		if node := treap.Search(key); node == nil || node.key != key {
			t.FailNow()
		}
	}

	for _, key := range keys {
		treap.Delete(key)
		if node := treap.Search(key); node != nil {
			t.FailNow()
		}
	}
}
