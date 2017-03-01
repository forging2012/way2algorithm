package bplus

import (
	"math/rand"
	"testing"
)

type pair struct {
	key int
	val int
}

const N = 10000

func TestBPlusTree(t *testing.T) {
	tree := New()

	pairs := make([]pair, N)
	for i := 0; i < N; i++ {
		pairs[i] = pair{
			key: rand.Intn(100000),
			val: rand.Intn(100000),
		}

		tree.Insert(pairs[i].key, pairs[i].val)
		if val, ok := tree.Search(pairs[i].key); !ok ||
			val.(int) != pairs[i].val {
			t.FailNow()
		}
	}

	for _, pair := range pairs {
		tree.Delete(pair.key)
		if _, ok := tree.Search(pair.key); ok {
			t.FailNow()
		}
	}
}
