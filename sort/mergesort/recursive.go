package mergesort

func Merge(a, b []int) []int {
	r, i, j := make([]int, len(a)+len(b)), 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}
	}

	copy(r[i+j:], a[i:])
	copy(r[len(a)+j:], b[j:])

	return r
}

// 归并排序，递归版本
func MergeSortRecursive(array []int) []int {
	if len(array) < 2 {
		return array
	}

	mid := len(array) >> 1

	return Merge(
		MergeSortRecursive(array[:mid]),
		MergeSortRecursive(array[mid:]),
	)
}
