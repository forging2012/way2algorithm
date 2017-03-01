package timsort

import (
	"way2algorithm/binarysearch"
	"way2algorithm/sort/mergesort"
	"way2algorithm/util"
)

func computeMinRun(n int) int {
	r := 0
	for n >= 64 {
		r |= n & 1
		n >>= 1
	}
	return n + r
}

func runLength(r run) int {
	return r.end - r.start + 1
}

func maintainStack(array []int, stack *[]run) {
	for {
		n := len(*stack)
		if n <= 1 {
			return
		}

		r1, r2 := (*stack)[n-1], (*stack)[n-2]

		switch {
		case n > 2 && runLength((*stack)[n-3]) <= runLength(r2)+runLength(r1):
			r3 := (*stack)[n-3]
			if runLength(r3) < runLength(r1) {
				mergesort.GallopingMerge(array, r3.start, r3.end, r2.end)
				(*stack)[n-3].end = r2.end
				(*stack)[n-2] = r1
				*stack = (*stack)[:n-1]
			} else {
				mergesort.GallopingMerge(array, r2.start, r2.end, r1.end)
				(*stack)[n-2].end = r1.end
				*stack = (*stack)[:n-1]
			}
		case runLength(r2) <= runLength(r1):
			mergesort.GallopingMerge(array, r2.start, r2.end, r1.end)
			(*stack)[n-2].end = r1.end
			*stack = (*stack)[:n-1]
		default:
			return
		}
	}
}

func TimSort(array []int) []int {
	stack, n := make([]run, 0), len(array)
	minRun := computeMinRun(n)

	for i := 0; i < n; i++ {
		start := i

		if i+1 >= n {
			stack = append(stack, run{start: i, end: i})
			goto Maintain
		}

		if array[i+1] >= array[i] {
			for i+1 < n && array[i+1] >= array[i] {
				i++
			}

		} else {
			for i+1 < n && array[i+1] < array[i] {
				i++
			}
			util.Reverse(array, start, i)
		}

		for ; i+1 < n && i-start+1 < minRun; i++ {
			target := array[i+1]
			j := binarysearch.BinarySearchRange(array, start, i, target)
			copy(array[j+1:i+2], array[j:i+1])
			array[j] = target
		}
		stack = append(stack, run{start: start, end: i})

	Maintain:
		maintainStack(array, &stack)
	}

	for n := len(stack); n > 1; n = len(stack) {
		r1, r2 := stack[n-1], stack[n-2]
		mergesort.GallopingMerge(array, r2.start, r2.end, r1.end)
		stack[n-2].end = r1.end
		stack = stack[:n-1]
	}

	return array
}
