package redblack

import (
	"math/rand"
	"testing"
	"time"
)

type entry struct {
	key int
	val int
}

func TestRedBlackTree(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	entries := make([]entry, 10000)
	for i := 0; i < 10000; i++ {
		entries[i] = entry{
			key: rand.Intn(100000),
			val: rand.Intn(100000),
		}
	}

	tree := New()
	for _, entry := range entries {
		tree.Insert(entry.key, entry.val)
		tree.root.validate()

		if val, ok := tree.Search(entry.key); !ok || val.(int) != entry.val {
			t.FailNow()
		}
	}

	for _, entry := range entries {
		tree.Delete(entry.key)
		tree.root.validate()

		if _, ok := tree.Search(entry.key); ok {
			t.FailNow()
		}
	}
}
