package bubblesort

// 冒泡排序
func BubbleSort(array []int) []int {
	n := len(array)
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
