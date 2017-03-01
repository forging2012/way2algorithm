package binarytree

func PreOrderRecursive(root *Node) []int {
	if root == nil {
		return []int{}
	}

	left := PreOrderRecursive(root.left)
	right := PreOrderRecursive(root.right)

	r := make([]int, 0, len(left)+len(right)+1)
	r = append(append(append(r, root.key), left...), right...)
	return r
}

func PreOrderIterative(root *Node) []int {
	stack, r := make([]*Node, 0), make([]int, 0)
	for root != nil || len(stack) > 0 {
		if root != nil {
			r = append(r, root.key)
			stack = append(stack, root)
			root = root.left
		} else {
			root = stack[len(stack)-1].right
			stack = stack[:len(stack)-1]
		}
	}
	return r
}
