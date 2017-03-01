package mergesort

// 归并排序，迭代版本
func MergeSortIterative(array []int) []int {
	n := len(array)

	for radius := 1; radius < n; radius <<= 1 {
		for start, diameter := 0, radius<<1; start < n; start += diameter {
			end := start + diameter
			if n < end {
				end = n
			}

			for i, j := start, start+radius; i < j && j < end; i++ {
				if array[i] > array[j] {
					tmp := array[j]
					copy(array[i+1:j+1], array[i:j])
					array[i] = tmp
					j++
				}
			}
		}
	}

	return array
}
