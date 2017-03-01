package mergesort

func mergeSortRecursiveImproved(array []int, start, end int) {
	if end-start <= 0 {
		return
	}

	mid := start + (end-start)>>1

	mergeSortRecursiveImproved(array, start, mid)
	mergeSortRecursiveImproved(array, mid+1, end)

	GallopingMerge(array, start, mid, end)
}

func MergeSortRecursiveImproved(array []int) []int {
	mergeSortRecursiveImproved(array, 0, len(array)-1)
	return array
}
