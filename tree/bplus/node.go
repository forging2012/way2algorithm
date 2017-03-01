package bplus

const T = 256

type Node struct {
	n        int
	keys     [T<<1 - 1]int
	children [T << 1]interface{}
	isLeaf   bool
}

func NewNode(isLeaf bool) *Node {
	return &Node{
		isLeaf: isLeaf,
	}
}

func (node *Node) isFull() bool {
	return node.n == T<<1-1
}

func (node *Node) isRich() bool {
	return node.n > T-1
}

// 当前叶子节点的右兄弟（或者叫后继）
func (node *Node) rightLeaf() *Node {
	if !node.isLeaf {
		return nil
	}
	return node.children[T<<1-1].(*Node)
}

// 找到keys中第一个比key大的数的位置
func (node *Node) indexKey(key int) int {
	low, high := 0, node.n-1
	for low <= high {
		mid := low + (high-low)>>1
		if node.keys[mid] <= key {
			low = mid + 1
		} else {
			if mid == 0 || node.keys[mid-1] <= key {
				return mid
			}
			high = mid - 1
		}
	}
	return low
}

// 在i处插入keys
func (node *Node) insertKeys(i int, keys ...int) {
	copy(node.keys[i+len(keys):], node.keys[i:node.n])
	copy(node.keys[i:], keys)
	node.n += len(keys)
}

// 在i处插入children
func (node *Node) insertChildren(i int, children ...interface{}) {
	copy(node.children[i+len(children):], node.children[i:node.n+1])
	copy(node.children[i:], children)
}

// 从i处移除n个key
func (node *Node) removeKeys(i, n int) {
	copy(node.keys[i:], node.keys[i+n:node.n])
	node.n -= n
}

// 从i处移除n个child
func (node *Node) removeChildren(i, n int) {
	copy(node.children[i:], node.children[i+n:node.n+1])
	for j := 0; j < n; j++ {
		node.children[node.n-j] = nil
	}
}

// 当node keys满的时候进行split, 此时node.n == T<<1 - 1
func (node *Node) split() *Node {
	father := NewNode(false)

	rightBro := NewNode(node.isLeaf)
	if node.isLeaf {
		rightBro.insertChildren(0, node.children[T-1:]...)
		rightBro.insertKeys(0, node.keys[T-1:]...)

		father.insertChildren(0, node, rightBro)
		father.insertKeys(0, rightBro.keys[0])

		node.removeChildren(T-1, T)
		node.removeKeys(T-1, T)

		// 把左右子节点连接起来
		node.children[T<<1-1] = rightBro
	} else {
		rightBro.insertChildren(0, node.children[T:]...)
		rightBro.insertKeys(0, node.keys[T:]...)

		father.insertChildren(0, node, rightBro)
		father.insertKeys(0, node.keys[T-1])

		node.removeChildren(T, T)
		node.removeKeys(T-1, T)
	}
	return father
}

// 合并node.children[i]和node.children[i+1]
func (node *Node) mergeChildren(i int) {
	left, right := node.children[i].(*Node), node.children[i+1].(*Node)

	leftN, rightN := left.n, right.n
	if !left.isLeaf {
		leftN++
		rightN++
	} else {
		// 因为要把右孩子merge到左孩子，然后删除右孩子，因此需要把左孩子和右
		// 孩子的右兄弟连接起来。
		left.children[T<<1-1] = right.children[T<<1-1]
	}

	left.insertChildren(leftN, right.children[:rightN]...)
	left.insertKeys(left.n, right.keys[:right.n]...)

	node.removeChildren(i+1, 1)
	node.removeKeys(i, 1)
}

// 往一个节点中插入，核心思想在于保证被插入的节点一定不是满的。
// 具体做法是：如果将要被插入的节点是满的，那么就split。
func (node *Node) insert(key int, value interface{}) {
	i := node.indexKey(key)

	if node.isLeaf {
		if i-1 >= 0 && node.keys[i-1] == key {
			node.children[i-1] = value
		} else {
			node.insertChildren(i, value)
			node.insertKeys(i, key)
		}
		return
	}

	if child := node.children[i].(*Node); child.isFull() {
		father := child.split()
		node.insertChildren(i+1, father.children[1])
		node.insertKeys(i, father.keys[0])

		if key >= father.keys[0] {
			i++
		}
	}

	node.children[i].(*Node).insert(key, value)
}

// 从一个节点中删除一个key, 核心思想在于保证要被删除的节点的key的个数必须大于
// keys的最少数目，即T-1。这样即使被删除后也不必向上维护。
func (node *Node) delete(key int) {
	i := node.indexKey(key)

	// 如果是子节点
	if node.isLeaf {
		if i != 0 && node.keys[i-1] == key {
			node.removeChildren(i-1, 1)
			node.removeKeys(i-1, 1)
		}
		return
	}

	child := node.children[i].(*Node)
	if !child.isRich() {
		if i > 0 && node.children[i-1].(*Node).isRich() {
			// 如果左兄弟是rich的，从左兄弟借一个

			leftBro := node.children[i-1].(*Node)
			if child.isLeaf {
				child.insertChildren(0, leftBro.children[leftBro.n-1])
				child.insertKeys(0, leftBro.keys[leftBro.n-1])

				leftBro.removeChildren(leftBro.n-1, 1)
				leftBro.removeKeys(leftBro.n-1, 1)

				node.keys[i-1] = child.keys[0]
			} else {
				child.insertChildren(0, leftBro.children[leftBro.n])
				child.insertKeys(0, node.keys[i-1])

				node.keys[i-1] = leftBro.keys[leftBro.n-1]

				leftBro.removeChildren(leftBro.n, 1)
				leftBro.removeKeys(leftBro.n-1, 1)
			}
		} else if i < node.n && node.children[i+1].(*Node).isRich() {
			// 如果右兄弟是rich的，那么从右兄弟借一个

			rightBro := node.children[i+1].(*Node)
			if child.isLeaf {
				child.insertChildren(child.n, rightBro.children[0])
				child.insertKeys(child.n, rightBro.keys[0])

				rightBro.removeChildren(0, 1)
				rightBro.removeKeys(0, 1)

				node.keys[i] = rightBro.keys[0]
			} else {
				child.insertChildren(child.n+1, rightBro.children[0])
				child.insertKeys(child.n, node.keys[i])

				node.keys[i] = rightBro.keys[0]

				rightBro.removeChildren(0, 1)
				rightBro.removeKeys(0, 1)
			}
		} else {
			// 否则，合并右兄弟
			if i == node.n {
				i--
			}
			node.mergeChildren(i)
		}
	}

	node.children[i].(*Node).delete(key)

	// 维护非子节点
	if i != 0 && node.keys[i-1] == key {
		child := node.children[i].(*Node)
		for !child.isLeaf {
			child = child.children[0].(*Node)
		}
		node.keys[i-1] = child.keys[0]
	}
}
