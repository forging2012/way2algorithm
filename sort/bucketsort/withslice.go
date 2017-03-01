package bucketsort

import (
	"way2algorithm/sort/coutingsort"
	"way2algorithm/util"
)

// 桶排序，桶数据结构是数组
func BucketSortWithSlice(array []int) []int {
	n := len(array)
	if n < 2 {
		return array
	}

	min, max := util.Min(array...), util.Max(array...)
	bucketSize, buckets := (max-min+1)/n, make([][]int, n)

	// 把array里边的元素放到相应的桶里边
	for _, num := range array {
		index := toBucketIndex(num, n, bucketSize, min)
		// 数组元素不固定，需要伴随扩容
		buckets[index] = append(buckets[index], num)
	}

	// 对每个桶进行排序
	//
	// 注意：桶排序算法并没有规定桶内排序算法，这里使用了计数排序，你也可以
	// 使用其他的排序算法
	for i := 0; i < n; i++ {
		buckets[i] = coutingsort.CoutingSort(buckets[i])
	}

	// 把所有的桶合并起来
	pointer := 0
	for _, bucket := range buckets {
		copy(array[pointer:pointer+len(bucket)], bucket)
		pointer += len(bucket)
	}

	return array
}
