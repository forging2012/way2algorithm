package segmenttree

import (
	"testing"
)

func TestSegmentTree(t *testing.T) {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	tree := New(array)

	cases := [][]int{
		{0, 9, 45},
		{1, 1, 1},
		{1, 0, 0},
		{-100, 100, 45},
		{0, 2, 3},
	}

	for _, item := range cases {
		if sum := tree.Sum(item[0], item[1]); sum != item[2] {
			t.FailNow()
		}
	}

	tree.Update(0, 2)
	if sum := tree.Sum(0, 9); sum != 47 {
		t.FailNow()
	}

	if sum := tree.Sum(1, 9); sum != 45 {
		t.FailNow()
	}
}
