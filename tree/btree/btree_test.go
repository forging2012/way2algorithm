package btree

import (
	"math/rand"
	"testing"
	"time"
)

func TestBTree(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	entries := make([]entry, 10000)
	for i := 0; i < 10000; i++ {
		entries[i] = entry{
			key: rand.Intn(100000),
			val: rand.Intn(100000),
		}
	}

	btree := New()
	for _, entry := range entries {
		btree.Insert(entry.key, entry.val)
		if val, ok := btree.Search(entry.key); !ok || val.(int) != entry.val {
			t.Fail()
		}
	}

	for _, entry := range entries {
		btree.Delete(entry.key)
		if _, ok := btree.Search(entry.key); ok {
			t.Fail()
		}
	}
}
