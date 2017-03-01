package unionfind

import (
	"testing"
)

func TestUnionFind(t *testing.T) {
	cases := [][]int{
		{0, 1},
		{2, 3},
		{4, 5},
		{6, 7},
		{8, 9},
	}

	uf := New(10)
	for i, pair := range cases {
		uf.Union(pair[0], pair[1])
		if uf.Find(pair[0]) != pair[1] || uf.count != 10-i-1 {
			t.Fail()
		}
	}

	uf.Union(1, 9)
	if uf.Find(0) != 9 || uf.Find(1) != 9 || uf.Find(8) != 9 ||
		uf.Find(9) != 9 || uf.count != 4 {
		t.Fail()
	}

	uf.Union(2, 8)
	if uf.Find(2) != 9 || uf.Find(3) != 9 || uf.count != 3 {
		t.Fail()
	}
}
