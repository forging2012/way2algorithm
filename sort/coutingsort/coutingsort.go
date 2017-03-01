package coutingsort

import (
	"way2algorithm/util"
)

// 计数排序
//
// Time: O(n)
// Space: O(n)
//
// Note: n为数组内的数字取值区间
func CoutingSort(array []int) []int {
	// 找出最大值和最小值，从而确定界限，减少下面count数组内存开销
	min, max := util.Min(array...), util.Max(array...)

	// 统计array里边每个数字的个数, count的下标index和array里边的元素num之间的
	// 关系为 index=num-min
	count := make([]int, max-min+1)
	for _, num := range array {
		count[num-min]++
	}

	// 根据统计结果对原数组进行填充
	pointer := 0
	for index, cnt := range count {
		for i := 0; i < cnt; i++ {
			array[pointer] = index + min
			pointer++
		}
	}

	return array
}
