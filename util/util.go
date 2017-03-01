package util

import (
	"math"
)

// 反转s[start:end+1]
func Reverse(s []int, start, end int) {
	for ; start < end; start, end = start+1, end-1 {
		s[start], s[end] = s[end], s[start]
	}
}

// Max返回数组里边最大的数
func Max(array ...int) int {
	max := math.MinInt64
	for _, num := range array {
		if num > max {
			max = num
		}
	}
	return max
}

// Min返回数组里边最小的数
func Min(array ...int) int {
	min := math.MaxInt64
	for _, num := range array {
		if num < min {
			min = num
		}
	}
	return min
}

// Abs返回n的绝对值
func Abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}
