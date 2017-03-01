package hashtable

import "way2algorithm/encryption/murmurhash"

// bucket里边的每一个元素
type bucket struct {
	key  string
	val  interface{}
	pre  *bucket
	next *bucket
}

// HashTable使用链地址法实现了最简单的哈希表
type HashTable struct {
	buckets []*bucket
}

func New() *HashTable {
	return &HashTable{
		buckets: make([]*bucket, 64),
	}
}

// 这里哈希使用了MurmurHash3
func (table *HashTable) index(key string) uint32 {
	hash := murmurhash.MurmurHash3_x86_32([]byte(key), 0)
	return hash % uint32(len(table.buckets))
}

func (table *HashTable) search(key string) (uint32, *bucket) {
	index := table.index(key)
	for bkt := table.buckets[index]; bkt != nil; bkt = bkt.next {
		if bkt.key == key {
			return index, bkt
		}
	}
	return index, nil
}

func (table *HashTable) Set(key string, val interface{}) {
	index, bkt := table.search(key)
	if bkt != nil {
		bkt.val = val
		return
	}

	bkt = &bucket{
		key:  key,
		val:  val,
		next: table.buckets[index],
	}

	if table.buckets[index] != nil {
		table.buckets[index].pre = bkt
	}

	table.buckets[index] = bkt
}

func (table *HashTable) Get(key string) (interface{}, bool) {
	_, bkt := table.search(key)
	if bkt != nil {
		return bkt.val, true
	}
	return nil, false
}

func (table *HashTable) Delete(key string) {
	index, bkt := table.search(key)
	if bkt == nil {
		return
	}

	if bkt.pre == nil {
		table.buckets[index] = bkt.next
		return
	}

	bkt.pre.next = bkt.next
	if bkt.next != nil {
		bkt.next.pre = bkt.pre
	}
}
