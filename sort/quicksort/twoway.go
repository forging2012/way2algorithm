package quicksort

func quickSortTwoWay(array []int, start, end int) {
	if start >= end {
		return
	}

	pivot, i, j := array[start], start, end
	for {
		for i <= end && array[i] <= pivot {
			i++
		}

		for j >= start && array[j] > pivot {
			j--
		}

		if i >= j {
			break
		}
		array[i], array[j] = array[j], array[i]
	}
	array[start], array[j] = array[j], array[start]

	quickSortTwoWay(array, start, j-1)
	quickSortTwoWay(array, j+1, end)
}

// 快排，two way，即有两个指针
func QuickSortTwoWay(array []int) []int {
	quickSortTwoWay(array, 0, len(array)-1)
	return array
}
