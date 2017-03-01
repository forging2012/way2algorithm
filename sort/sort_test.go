package sort

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"

	"way2algorithm/sort/bubblesort"
	"way2algorithm/sort/bucketsort"
	"way2algorithm/sort/coutingsort"
	"way2algorithm/sort/heapsort"
	"way2algorithm/sort/insertionsort"
	"way2algorithm/sort/mergesort"
	"way2algorithm/sort/quicksort"
	"way2algorithm/sort/radixsort"
	"way2algorithm/sort/selectionsort"
	"way2algorithm/sort/timsort"
)

func randSlice() []int {
	rand.Seed(time.Now().UnixNano())

	r := make([]int, rand.Intn(500))
	for i := 0; i < len(r); i++ {
		r[i] = rand.Intn(1000)
	}
	return r
}

func copySlice(s []int) []int {
	clone := make([]int, len(s))
	copy(clone, s)
	return clone
}

type SortFunc func([]int) []int

var sortFuncs = []SortFunc{
	coutingsort.CoutingSort,
	bucketsort.BucketSortWithSlice,
	bucketsort.BucketSortWithList,
	radixsort.RadixSort,
	bubblesort.BubbleSort,
	insertionsort.InsertionSortSwap,
	insertionsort.InsertionSortFindAndSwap,
	insertionsort.InsertionSortBinary,
	insertionsort.InsertionSortBinaryFinal,
	selectionsort.SelectionSort,
	heapsort.HeapSort,
	mergesort.MergeSortRecursiveImproved,
	mergesort.MergeSortIterativeImproved,
	mergesort.MergeSortIterative,
	quicksort.QuickSortTwoWay,
	quicksort.QuickSortThreeWay,
	quicksort.QuickSortDualPivot,
	timsort.SimpleTimSort,
	timsort.ImprovedTimSort,
	timsort.TimSort,
}

func TestSort(t *testing.T) {
	for i := 0; i < 1000; i++ {
		array := randSlice()

		standard := copySlice(array)
		sort.Ints(standard)

		for i, f := range sortFuncs {
			r := f(copySlice(array))
			if !reflect.DeepEqual(standard, r) {
				t.Fatal("Not equeal ", array, standard, r, i)
			}
		}
	}
}
