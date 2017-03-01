package quicksort

func quickSortDualPivot(array []int, start, end int) {
	if start >= end {
		return
	}

	if array[start] > array[end] {
		array[start], array[end] = array[end], array[start]
	}

	low, high, i := start+1, end-1, start+1
	for i <= high {
		if array[i] < array[start] {
			array[i], array[low] = array[low], array[i]
			low++
			i++
		} else if array[i] > array[end] {
			array[i], array[high] = array[high], array[i]
			high--
		} else {
			i++
		}
	}

	array[start], array[low-1] = array[low-1], array[start]
	array[end], array[high+1] = array[high+1], array[end]

	quickSortDualPivot(array, start, low-2)
	quickSortDualPivot(array, low, high)
	quickSortDualPivot(array, high+2, end)
}

// 双枢纽快排
func QuickSortDualPivot(array []int) []int {
	quickSortDualPivot(array, 0, len(array)-1)
	return array
}
