package segmenttree

type SegmentTree struct {
	root *Node
}

func New(array []int) *SegmentTree {
	tree := new(SegmentTree)
	tree.root = tree.build(array, 0, len(array)-1)
	return tree
}

func (tree *SegmentTree) build(array []int, start, end int) *Node {
	if start > end {
		return nil
	}

	root := NewNode(start, end)
	if start == end {
		root.value = array[start]
		return root
	}

	mid := start + (end-start)>>1

	root.left = tree.build(array, start, mid)
	root.right = tree.build(array, mid+1, end)
	root.value = root.left.value + root.right.value

	return root
}

func (tree *SegmentTree) Sum(start, end int) int {
	return tree.root.sum(start, end)
}

func (tree *SegmentTree) Update(i, num int) {
	tree.root.update(i, num)
}
