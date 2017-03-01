package heapsort

// 从root到叶节点，如果子节点大于父节点，那么交换父节点和子节点，从而保持
// 堆得性质。
func heapify(array []int, start, end int) {
	father := start

	for {
		son := father<<1 + 1
		if son > end {
			break
		}

		if son+1 <= end && array[son+1] > array[son] {
			son++
		}

		if array[son] < array[father] {
			break
		}

		array[father], array[son] = array[son], array[father]
		father = son
	}
}

// 建堆
func build(array []int) {
	n := len(array)
	for i := n >> 1; i >= 0; i-- {
		heapify(array, i, n-1)
	}
}

// 堆排序
func HeapSort(array []int) []int {
	build(array)

	for i := len(array) - 1; i > 0; i-- {
		array[0], array[i] = array[i], array[0]
		heapify(array, 0, i-1)
	}

	return array
}
