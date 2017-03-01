package btree

const T = 256

type entry struct {
	key int
	val interface{}
}

type BTreeNode struct {
	n        int // n代表key即entries的个数
	entries  [T<<1 - 1]*entry
	children []*BTreeNode
	isLeaf   bool
}

func (node *BTreeNode) isRich() bool {
	return node.n >= T
}

func (node *BTreeNode) isFull() bool {
	return node.n == len(node.entries)
}

// 在node中找到第一个不小于key的元素的下标
func (node *BTreeNode) indexKey(key int) int {
	low, high := 0, node.n-1
	for low <= high {
		mid := low + (high-low)>>1
		if node.entries[mid].key < key {
			low = mid + 1
		} else {
			if mid == 0 || node.entries[mid-1].key < key {
				return mid
			}
			high = mid - 1
		}
	}
	return low
}

// 在i处插入entries
func (node *BTreeNode) insertEntries(i int, entries ...*entry) {
	copy(node.entries[i+len(entries):], node.entries[i:node.n])
	copy(node.entries[i:], entries)
	node.n += len(entries)
}

// 在i处插入children
func (node *BTreeNode) insertChildren(i int, children ...*BTreeNode) {
	if node.children == nil {
		node.children = make([]*BTreeNode, T<<1)
	}

	copy(node.children[i+len(children):], node.children[i:node.n+1])
	copy(node.children[i:], children)
}

// 从i处删除n个entry
func (node *BTreeNode) removeEntries(i, n int) {
	copy(node.entries[i:], node.entries[i+n:node.n])

	for i := 0; i < n; i++ {
		node.n--
		node.entries[node.n] = nil
	}
}

// 从i处删除n个child
func (node *BTreeNode) removeChildren(i, n int) {
	copy(node.children[i:], node.children[i+n:node.n+1])

	for i := 0; i < n; i++ {
		node.children[node.n-i] = nil
	}
}

// 当节点是满的时候，即node.n=T<<1 - 1的时候
func (node *BTreeNode) split() (root *BTreeNode) {
	root = &BTreeNode{isLeaf: false}
	root.insertEntries(0, node.entries[T-1])

	rightBro := &BTreeNode{isLeaf: node.isLeaf}
	if !node.isLeaf {
		rightBro.insertChildren(0, node.children[T:node.n+1]...)
		for i := T; i < node.n+1; i++ {
			node.children[i] = nil
		}
	}
	rightBro.insertEntries(0, node.entries[T:node.n]...)
	for i := T - 1; i < node.n; i++ {
		node.entries[i] = nil
	}

	node.n = T - 1
	root.insertChildren(0, node, rightBro)
	return
}

// 插入算法的核心：必须保证当前访问的节点不是满的。这样的话就保证了要插入
// 叶节点时，直接插入即可，避免了递归向上分裂的情况。
func (node *BTreeNode) insert(e *entry) {
	i := node.indexKey(e.key)
	if i < node.n && node.entries[i].key == e.key {
		node.entries[i] = e
		return
	}

	if node.isLeaf {
		node.insertEntries(i, e)
		return
	}

	if child := node.children[i]; child.isFull() {
		father := child.split()
		node.insertChildren(i+1, father.children[1])
		node.insertEntries(i, father.entries[0])

		if e.key == node.entries[i].key {
			node.entries[i] = e
			return
		}

		if e.key > node.entries[i].key {
			i++
		}
	}

	node.children[i].insert(e)
}

func (node *BTreeNode) mergeLeft(root *BTreeNode, i int) {
	leftBro := root.children[i-1]

	if !node.isLeaf {
		node.insertChildren(0, leftBro.children[:leftBro.n+1]...)
	}
	node.insertEntries(0, root.entries[i-1])
	node.insertEntries(0, leftBro.entries[:leftBro.n]...)

	root.removeChildren(i-1, 1)
	root.removeEntries(i-1, 1)
}

func (node *BTreeNode) mergeRight(root *BTreeNode, i int) {
	rightBro := root.children[i+1]

	if !node.isLeaf {
		node.insertChildren(
			node.n+1, rightBro.children[:rightBro.n+1]...,
		)
	}
	node.insertEntries(node.n, root.entries[i])
	node.insertEntries(node.n, rightBro.entries[:rightBro.n]...)

	root.removeChildren(i+1, 1)
	root.removeEntries(i, 1)
}

// 删除算法的核心在于：保证接下来要访问的节点中的key的个数至少为T。这样的话，
// 就避免了删除子节点后需要向上递归合并的情形.
func (node *BTreeNode) delete(key int) {
	i := node.indexKey(key)

	// 如果找到key
	if i < node.n && node.entries[i].key == key {
		// 在叶子节点，直接删除
		if node.isLeaf {
			node.removeEntries(i, 1)
			return
		}

		leftChild, rightChild := node.children[i], node.children[i+1]

		if leftChild.isRich() {
			// 如果左孩子富有，找到其前驱替换自己，然后删除前驱
			precursor := leftChild
			for !precursor.isLeaf {
				precursor = precursor.children[precursor.n]
			}
			node.entries[i] = precursor.entries[precursor.n-1]

			leftChild.delete(node.entries[i].key)
		} else if rightChild.isRich() {
			// 如果右孩子富有，找到其后继替换自己，然后删除后继
			successor := rightChild
			for !successor.isLeaf {
				successor = successor.children[0]
			}
			node.entries[i] = successor.entries[0]

			rightChild.delete(node.entries[i].key)
		} else {
			// 否则的话，将左右孩子merge到一起，然后在merge后的节点中删除key
			leftChild.mergeRight(node, i)
			leftChild.delete(key)
		}
		return
	}

	// 如果是叶节点的话，孩子不存在,删除结束
	if node.isLeaf {
		return
	}

	child := node.children[i]

	if !child.isRich() {
		// 如果要删除的节点的左兄弟是富有的，那么就向左兄弟借一个entry
		if i > 0 && node.children[i-1].isRich() {
			leftBro := node.children[i-1]

			if !child.isLeaf {
				child.insertChildren(0, leftBro.children[leftBro.n])
			}
			child.insertEntries(0, node.entries[i-1])

			node.entries[i-1] = leftBro.entries[leftBro.n-1]

			if !child.isLeaf {
				leftBro.removeChildren(leftBro.n, 1)
			}
			leftBro.removeEntries(leftBro.n-1, 1)
		} else if i < node.n && node.children[i+1].isRich() {
			// 如果要删除的节点的右兄弟是富有的，那么向右兄弟借一个entry
			rightBro := node.children[i+1]

			if !child.isLeaf {
				child.insertChildren(child.n+1, rightBro.children[0])
			}
			child.insertEntries(child.n, node.entries[i])

			node.entries[i] = rightBro.entries[0]

			if !child.isLeaf {
				rightBro.removeChildren(0, 1)
			}
			rightBro.removeEntries(0, 1)
		} else {
			// 否则的话，和左右兄弟中的一个进行合并
			if i < node.n {
				child.mergeRight(node, i)
			} else {
				child.mergeLeft(node, i)
			}
		}
	}

	child.delete(key)
}
