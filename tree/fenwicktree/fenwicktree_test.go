package fenwicktree

import (
	"testing"
)

func TestSegmentTree(t *testing.T) {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	tree := New(array)

	cases := [][]int{
		{10, 45},
		{0, 0},
		{1, 0},
		{3, 3},
	}

	for _, item := range cases {
		if sum := tree.Sum(item[0]); sum != item[1] {
			t.FailNow()
		}
	}

	tree.Update(3, 100)
	if sum := tree.Sum(3); sum != 3 {
		t.FailNow()
	}

	if sum := tree.Sum(4); sum != 103 {
		t.FailNow()
	}
}
