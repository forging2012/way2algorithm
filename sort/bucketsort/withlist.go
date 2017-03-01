package bucketsort

import (
	"way2algorithm/util"
)

// 桶里边的每个元素
type bucketElement struct {
	num  int
	next *bucketElement
}

// 把元素e插入到桶buckets[i]内
func insertToBucket(buckets []*bucketElement, i int, e *bucketElement) {
	root, pre := buckets[i], (*bucketElement)(nil)
	for root != nil && e.num >= root.num {
		pre, root = root, root.next
	}

	if pre == nil {
		buckets[i] = e
	} else {
		pre.next = e
	}
	e.next = root
}

// 桶排序，桶的数据结构为链表
func BucketSortWithList(array []int) []int {
	n := len(array)
	if n < 2 {
		return array
	}

	min, max := util.Min(array...), util.Max(array...)
	bucketSize, buckets := (max-min+1)/n, make([]*bucketElement, n)

	// 把array里边的元素放到相应的桶里边，并用插入排序放到合适的位置
	for _, num := range array {
		index := toBucketIndex(num, n, bucketSize, min)
		insertToBucket(buckets, index, &bucketElement{num: num})
	}

	// 把所有的桶合并起来
	pointer := 0
	for _, bucket := range buckets {
		for root := bucket; root != nil; root = root.next {
			array[pointer] = root.num
			pointer++
		}
	}

	return array
}
