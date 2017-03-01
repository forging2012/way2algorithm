package fenwicktree

// 从最小位到最高位找到第一个为1的bit
func lowbit(x int) int {
	return x & (-x)
}

// 树状数组
type FenwickTree struct {
	array    []int
	chunkSum []int
}

func New(array []int) *FenwickTree {
	tree := &FenwickTree{
		array:    make([]int, len(array)),
		chunkSum: make([]int, len(array)+1),
	}

	// go里边的slice是引用，为防止array被意外修改，这里copy一份
	copy(tree.array, array)

	// 初始化chunkSum，chunkSum[n]代表的含义为array前n项和
	for i := 1; i <= len(array); i++ {
		for j := i - 1; j >= i-lowbit(i); j-- {
			tree.chunkSum[i] += array[j]
		}
	}

	return tree
}

// 求前n项和
func (tree *FenwickTree) Sum(n int) int {
	sum := 0
	for n > 0 {
		sum += tree.chunkSum[n]
		n -= lowbit(n)
	}
	return sum
}

// 更新数组的第i(从0开始)位，即array[i]=num
func (tree *FenwickTree) Update(i, num int) {
	n := i + 1
	for n <= len(tree.array) {
		tree.chunkSum[n] += num - tree.array[i]
		n += lowbit(n)
	}
	tree.array[i] = num
}
