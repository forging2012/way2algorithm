package quicksort

func quickSortThreeWay(array []int, start, end int) {
	if start >= end {
		return
	}

	low, high, i, pivot := start, end, start, array[start]
	for i <= high {
		if array[i] < pivot {
			array[low], array[i] = array[i], array[low]
			low++
			i++
		} else if array[i] > pivot {
			array[high], array[i] = array[i], array[high]
			high--
		} else {
			i++
		}
	}

	quickSortThreeWay(array, start, low-1)
	quickSortThreeWay(array, high+1, end)
}

func QuickSortThreeWay(array []int) []int {
	quickSortThreeWay(array, 0, len(array)-1)
	return array
}
