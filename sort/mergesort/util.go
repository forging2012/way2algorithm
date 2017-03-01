package mergesort

import (
	"way2algorithm/binarysearch"
)

const MinGalloping = 7

func grow(array []int, pointer *int, data *[]int, length int) {
	copy(array[*pointer:*pointer+length], (*data)[:length])
	*pointer += length
	*data = (*data)[length:]
}

func galloping(array []int, pointer *int, arrayA, arrayB *[]int) (
	shouldStop bool) {

	if len(*arrayA) == 0 || len(*arrayB) == 0 {
		return true
	}

	var k uint = 0

	for index := 0; index < len(*arrayA) &&
		(*arrayA)[index] < (*arrayB)[0]; index = 2<<k - 1 {
		k++
	}

	j := 0

	if k > 0 {
		start := 2<<(k-1) - 1
		if k == 1 {
			start = 0
		}

		end := 2<<k - 1
		if len(*arrayA) <= end {
			end = len(*arrayA) - 1
		}

		j = binarysearch.BinarySearchRange(
			*arrayA, start, end, (*arrayB)[0],
		)
	}

	grow(array, pointer, arrayA, j)
	grow(array, pointer, arrayB, 1)

	return j < MinGalloping
}

func forwardStep(array []int, slice *[]int, pointer, x, y *int) {
	grow(array, pointer, slice, 1)
	(*x)++
	(*y) = 0
}

func GallopingMerge(array []int, start, mid, end int) {
	s1, e1, s2, e2 := start, mid, mid+1, end
	if s1 > e1 || s2 > e2 {
		return
	}

	s1 = binarysearch.BinarySearchRange(
		array, s1, e1, array[s2],
	)

	index := binarysearch.BinarySearchRange(array, s2, e2, array[e1])
	if index < e2 {
		e2 = index
	}

	arrayA := make([]int, e1-s1+1)
	arrayB := array[s2 : e2+1]
	copy(arrayA, array[s1:e1+1])

	countA, countB, pointer := 0, 0, s1

Loop:
	for len(arrayA) > 0 && len(arrayB) > 0 {
		for countA >= MinGalloping || countB >= MinGalloping {
			if galloping(array, &pointer, &arrayA, &arrayB) &&
				galloping(array, &pointer, &arrayB, &arrayA) {
				countA, countB = 0, 0
				goto Loop
			}
		}

		if arrayA[0] < arrayB[0] {
			forwardStep(array, &arrayA, &pointer, &countA, &countB)
		} else {
			forwardStep(array, &arrayB, &pointer, &countB, &countA)
		}
	}

	lenA, lenB := len(arrayA), len(arrayB)
	copy(array[pointer:pointer+lenA], arrayA)
	copy(array[pointer+lenA:pointer+lenA+lenB], arrayB)
}
