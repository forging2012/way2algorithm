package binarytree

func InOrderRecursive(root *Node) []int {
	if root == nil {
		return []int{}
	}

	left := InOrderRecursive(root.left)
	right := InOrderRecursive(root.right)

	r := make([]int, 0, len(left)+len(right)+1)
	r = append(append(append(r, left...), root.key), right...)

	return r
}

func InOrderIterative(root *Node) []int {
	stack, r := make([]*Node, 0), make([]int, 0)
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.left
		} else {
			r = append(r, root.key)
			root = stack[len(stack)-1].right
			stack = stack[:len(stack)-1]
		}
	}
	return r
}
