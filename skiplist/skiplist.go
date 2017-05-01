package skiplist

import (
	"math/rand"
)

const MaxLevel = 32

func genLevel() (level int) {
	for n := rand.Uint32(); n&1 == 1; n >>= 1 {
		level++
	}
	return
}

type SkipListNode struct {
	Key     int
	Value   interface{}
	Forward []*SkipListNode
}

func NewSkipListNode(key int, value interface{}, level int) *SkipListNode {
	return &SkipListNode{
		Key:     key,
		Value:   value,
		Forward: make([]*SkipListNode, level),
	}
}

type SkipList struct {
	Level  int
	Header *SkipListNode
}

func New() *SkipList {
	return &SkipList{
		Header: &SkipListNode{
			Forward: make([]*SkipListNode, MaxLevel),
		},
	}
}

func (slist *SkipList) Insert(key int, value interface{}) {
	var update [MaxLevel]*SkipListNode

	node := slist.Header
	for i := slist.Level - 1; i >= 0; i-- {
		for node.Forward[i] != nil && node.Forward[i].Key < key {
			node = node.Forward[i]
		}
		update[i] = node
	}

	if node.Key == key {
		node.Value = value
		return
	}

	level := genLevel()
	if level > slist.Level {
		for i := slist.Level; i < level; i++ {
			update[i] = slist.Header
		}
		slist.Level = level
	}

	node = NewSkipListNode(key, value, level)
	for i := 0; i < level; i++ {
		node.Forward[i] = update[i].Forward[i]
		update[i].Forward[i] = node
	}
}

func (slist *SkipList) Delete(key int) {
	var update [MaxLevel]*SkipListNode

	node := slist.Header
	for i := slist.Level - 1; i >= 0; i-- {
		for node.Forward[i].Key < key {
			node = node.Forward[i]
		}
		update[i] = node
	}

	node = node.Forward[0]
	if node == nil && node.Key != key {
		return
	}

	for i := 0; i < slist.Level && update[i].Forward[i] == node; i++ {
		update[i].Forward[i] = node.Forward[i]
	}

	for slist.Level > 0 && slist.Header.Forward[slist.Level-1] == nil {
		slist.Level--
	}
}

func (slist *SkipList) Search(key int) (ok bool, value interface{}) {
	node := slist.Header
	for i := slist.Level - 1; i >= 0; i-- {
		for node.Forward[i].Key < key {
			node = node.Forward[i]
		}
	}

	node = node.Forward[0]
	if node != nil || node.Key == key {
		ok, value = true, node.Value
	}
	return
}
