package timsort

import (
	"way2algorithm/binarysearch"
	"way2algorithm/sort/mergesort"
	"way2algorithm/util"
)

const MinRun = 10
const MinGalloping = 7

func mergeAll(array []int, queue []run) {
	if len(queue) < 2 {
		return
	}

	newQueue := make([]run, 0)

	for i := 0; i < len(queue); i += 2 {
		if i+1 >= len(queue) {
			newQueue = append(newQueue, queue[i])
			break
		}

		mergesort.GallopingMerge(
			array, queue[i].start, queue[i].end, queue[i+1].end,
		)

		newQueue = append(
			newQueue,
			run{start: queue[i].start, end: queue[i+1].end},
		)
	}

	mergeAll(array, newQueue)
}

func ImprovedTimSort(array []int) []int {
	queue, n, minRun := make([]run, 0), len(array), MinRun

	for i := 0; i < n; i++ {
		if i+1 >= n {
			queue = append(queue, run{start: i, end: i})
			break
		}

		start := i
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
		queue = append(queue, run{start: start, end: i})
	}

	mergeAll(array, queue)
	return array
}
