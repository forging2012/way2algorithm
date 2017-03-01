package bucketsort

// 计算num所在的bucket编号
func toBucketIndex(num int, totalBuckets, bucketSize int, min int) int {
	index := (num - min) / bucketSize
	if index >= totalBuckets {
		index = totalBuckets - 1
	}
	return index
}
