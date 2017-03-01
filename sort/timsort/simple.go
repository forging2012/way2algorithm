package timsort

import (
	"way2algorithm/sort/mergesort"
	"way2algorithm/util"
)

// 基本思想： 遍历一遍数组，找出所有的run，并把run放到队列里边，然后把队列里边
// 的runs进行合并。
//
// 优化：
//   - Merge次数尽量少？Run的个数、Merge策略（什么时候Merge，以怎样的顺序）
//   - Merge怎样才能运行更快？Run的长度、比较次数（算法优化）
func SimpleTimSort(array []int) []int {
	queue, n := make([][]int, 0), len(array)

	for i := 0; i < n; i++ {
		// 只剩最后一个元素
		if i+1 >= n {
			queue = append(queue, array[i:])
			break
		}

		// 如果序列是递增的
		if array[i+1] >= array[i] {
			start := i
			for i+1 < n && array[i+1] >= array[i] {
				i++
			}
			queue = append(queue, array[start:i+1])
			continue
		}

		// 否则是递减的
		start := i
		for i+1 < n && array[i+1] < array[i] {
			i++
		}
		util.Reverse(array, start, i)
		queue = append(queue, array[start:i+1])
	}

	// 对队列中的runs进行Merge
	r := make([]int, 0)
	for _, item := range queue {
		r = mergesort.Merge(r, item)
	}
	return r
}
