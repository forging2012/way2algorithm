package insertionsort

import (
	"way2algorithm/binarysearch"
	"way2algorithm/util"
)

// 每次从后往前遍历已经排序的部分，遍历的后移一位
func InsertionSortSwap(array []int) []int {
	for i := 0; i < len(array); i++ {
		num, j := array[i], i-1
		for ; j >= 0 && array[j] > num; j-- {
			array[j+1] = array[j]
		}
		array[j+1] = num
	}
	return array
}

// 首先在已排序的部分中查找要插入的数的位置，然后把从该位置开始的部分统统右移
// 一位，最后把要插入的数字插到该位置。
//
// 查找算法为：从后往前遍历已经排序的部分，一直找到相关位置。算法复杂度为O(n).
func InsertionSortFindAndSwap(array []int) []int {
	for i := 0; i < len(array); i++ {
		j, num := i-1, array[i]

		for j >= 0 && num < array[j] {
			j--
		}

		copy(array[j+2:i+1], array[j+1:i])
		array[j+1] = num
	}

	return array
}

// 与上述算法类似，只不过查找算法改为了二分查找。使用二分查找相对上述两种算法
// 的优势在于：减少了比较次数.
func InsertionSortBinary(array []int) []int {
	for i, num := range array {
		j := binarysearch.BinarySearch(array[:i], num)
		copy(array[j+1:i+1], array[j:i])
		array[j] = num
	}

	return array
}

// 查找算法改用二分查找算法复杂度很稳定，都是为O(n^2)。但是我们考虑一个特殊的
// 例子：序列全部是已经排好序的。这时候，普通的插入排序InsertionSortSwap和
// InsertionSortFindAndSwap 均为 O(n)，但是使用了二分查找的依旧为O(n^2)。仔细
// 考虑一下，我们没有必要对开头已经有序的部分使用二分查找进行插入排序。
//
// 此算法对上面算法做了优化，开头有序的部分不再进行插入排序，因此最好的情况下
// 依旧是O(n).
func InsertionSortBinaryFinal(array []int) []int {
	// 边界检查
	n := len(array)
	if n < 2 {
		return array
	}

	// 如果是非递减数列
	i := 1
	for i < n && array[i] >= array[i-1] {
		i++
	}

	// 如果是递减数列
	if i == 1 {
		for i < n && array[i] < array[i-1] {
			i++
		}
		util.Reverse(array, 0, i-1)
	}

	for ; i < n; i++ {
		num := array[i]
		j := binarysearch.BinarySearch(array[:i], num)
		copy(array[j+1:i+1], array[j:i])
		array[j] = num
	}

	return array
}
