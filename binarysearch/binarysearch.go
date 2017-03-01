package binarysearch

// 函数描述:
//
// 在array的 [start, end] 闭区间中寻找target的位置index，使得
// array[index] >= target，且index最小。
//
// 用通俗点的话说：在array的[start, end] 区间中找出第一个不小于target的数的位置.
func BinarySearchRange(array []int, start, end int, target int) int {
	low, high := start, end

	for low <= high {
		mid := low + (high-low)>>1
		if array[mid] < target {
			low = mid + 1
		} else {
			if mid == 0 || array[mid-1] < target {
				return mid
			}
			high = mid - 1
		}
	}
	return low
}

// 是上述函数的特例，即start=0, end=len(array)-1
func BinarySearch(array []int, target int) int {
	return BinarySearchRange(array, 0, len(array)-1, target)
}
