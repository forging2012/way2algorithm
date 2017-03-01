package binarytree

import (
	"way2algorithm/util"
)

func PostOrderRecursive(root *Node) []int {
	if root == nil {
		return []int{}
	}

	left := InOrderRecursive(root.left)
	right := InOrderRecursive(root.right)

	r := make([]int, 0, len(left)+len(right)+1)
	r = append(append(append(r, left...), right...), root.key)

	return r
}

func PostOrderIterative(root *Node) []int {
	stack, r := make([]*Node, 0), make([]int, 0)
	for root != nil || len(stack) > 0 {
		if root != nil {
			r = append(r, root.key)
			stack = append(stack, root)
			root = root.right
		} else {
			root = stack[len(stack)-1].left
			stack = stack[:len(stack)-1]
		}
	}

	util.Reverse(r, 0, len(r)-1)
	return r
}
