package unionfind

// 并查集
type UnionFind struct {
	id    []int
	size  []int
	count int
}

func New(n int) *UnionFind {
	uf := &UnionFind{
		id:    make([]int, n),
		size:  make([]int, n),
		count: n,
	}

	// 初始化:
	// 每个元素的父节点为它自己，且每个连通集的个数为1
	for i := 0; i < n; i++ {
		uf.id[i] = i
		uf.size[i] = 1
	}

	return uf
}

// 查找的时候同时把元素改成父节点的父节点，也即我们所说的路径压缩
func (uf *UnionFind) Find(key int) int {
	for key != uf.id[key] {
		uf.id[key] = uf.id[uf.id[key]]
		key = uf.id[key]
	}
	return key
}

// 把两个连通集合并成一个，需要把较小的集合合并到较大的集合中，这样做的原因
// 是路径压缩的数量就会变少
func (uf *UnionFind) Union(key1, key2 int) {
	root1, root2 := uf.Find(key1), uf.Find(key2)
	if root1 == root2 {
		return
	}

	small, large := root1, root2
	if uf.size[root1] > uf.size[root2] {
		small, large = root2, root1
	}

	uf.id[small] = uf.id[large]
	uf.size[large] += uf.size[small]
	uf.count--
}
