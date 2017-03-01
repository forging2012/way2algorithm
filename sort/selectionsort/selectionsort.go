package selectionsort

import (
	"math"
)

// 选择排序
func SelectionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		min, index := math.MaxInt64, i
		for j := i; j < len(array); j++ {
			if array[j] < min {
				min = array[j]
				index = j
			}
		}
		array[i], array[index] = array[index], array[i]
	}
	return array
}
