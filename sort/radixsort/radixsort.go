package radixsort

import (
	"strconv"
	"strings"
)

// 合并桶并把结果填充到arrayString。
func mergeBuckets(arrayString []string, buckets [][]string) {
	pointer := 0
	for _, bucket := range buckets {
		copy(arrayString[pointer:pointer+len(bucket)], bucket)
		pointer += len(bucket)
	}
}

// 重置桶
func resetBuckets(buckets [][]string) {
	for i := 0; i < len(buckets); i++ {
		buckets[i] = buckets[i][:0]
	}
}

// 基数排序
func RadixSort(array []int) []int {
	arrayString, maxLen := make([]string, len(array)), -1

	// 把数组里边的数字转化为string类型, 并统计生成的string最大的长度maxLen
	for i, num := range array {
		arrayString[i] = strconv.Itoa(num)
		if l := len(arrayString[i]); l > maxLen {
			maxLen = l
		}
	}

	// 把所有的string右对齐，使所有的string都具有相同的长度maxLen。长度不够的
	// 在左边填充0。
	for i, s := range arrayString {
		arrayString[i] = strings.Repeat("0", maxLen-len(s)) + s
	}

	// 对于i = maxLen - 1 到 0, 把arrayString里边的数据item根据item[i] - '0'
	// 放到对应的桶里边。然后合并桶，并把合并后的结果填充到arrayString。
	// 依次这样，直到 i < 0。最后把arrayString的数据转换为in，便得到结果。
	buckets := make([][]string, 10)
	for i := maxLen - 1; i >= 0; i-- {
		for _, item := range arrayString {
			index := item[i] - '0'
			buckets[index] = append(buckets[index], item)
		}
		mergeBuckets(arrayString, buckets)
		resetBuckets(buckets)
	}

	for i, s := range arrayString {
		num, _ := strconv.ParseInt(s, 10, 64)
		array[i] = int(num)
	}

	return array
}
