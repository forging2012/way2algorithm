package mergesort

func MergeSortIterativeImproved(array []int) []int {
	n := len(array)

	for radius := 1; radius < n; radius <<= 1 {
		for start, diameter := 0, radius<<1; start < n; start += diameter {
			end := start + diameter
			if n < end {
				end = n
			}
			GallopingMerge(array, start, start+radius-1, end-1)
		}
	}

	return array
}
