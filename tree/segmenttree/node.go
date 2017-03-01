package segmenttree

type Node struct {
	start int
	end   int
	value int
	left  *Node
	right *Node
}

func NewNode(start, end int) *Node {
	return &Node{
		start: start,
		end:   end,
	}
}

func (node *Node) sum(start, end int) int {
	if node == nil || start > end || start > node.end || end < node.start {
		return 0
	}

	if start <= node.start && node.end <= end {
		return node.value
	}

	return node.left.sum(start, end) + node.right.sum(start, end)
}

func (node *Node) update(i, num int) {
	if node == nil || i < node.start || node.end < i {
		return
	}

	if node.start == node.end && node.start == i {
		node.value = num
		return
	}

	node.left.update(i, num)
	node.right.update(i, num)

	node.value = node.left.value + node.right.value
}
