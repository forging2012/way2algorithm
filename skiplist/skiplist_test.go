package skiplist

import (
	"math/rand"
	"testing"
)

var slist *SkipList

func init() {
	slist = New()
}

const N = 100

type Pair struct {
	Key   int
	Value int
}

func SkipListTest(t *testing.T) {
	pairs := make([]Pair, N)

	for i := 0; i < N; i++ {
		pairs[i] = Pair{
			Key:   rand.Int(),
			Value: rand.Int(),
		}
		slist.Insert(pairs[i].Key, pairs[i].Value)
		if ok, value := slist.Search(pairs[i].Key); !ok ||
			pairs[i].Value != value.(int) {
			t.FailNow()
		}
	}

	for i := 0; i < N; i++ {
		slist.Delete(pairs[i].Key)
		if ok, _ := slist.Search(pairs[i].Key); ok {
			t.FailNow()
		}
	}
}
