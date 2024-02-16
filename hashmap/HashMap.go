package hashmap

import "sort"

type kv struct {
	key   int64
	value interface{}
}

type Bucket struct {
	data []kv
}

func (b *Bucket) Insert(key int64, value interface{}) {
	for i := range b.data {
		if b.data[i].key == key {
			b.data[i].value = value
			return
		}
	}

	b.data = append(b.data, kv{key: key, value: value})
	sort.Slice(b.data, func(i, j int) bool {
		return b.data[i].key < b.data[j].key
	})
}

func (b *Bucket) Find(key int64) interface{} {
	// Binary search for efficiency
	start, end := 0, len(b.data)-1
	for start <= end {
		mid := start + (end-start)/2
		if b.data[mid].key == key {
			return b.data[mid].value
		}
		if b.data[mid].key < key {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return nil
}

func (b *Bucket) Remove(key int64) {
	start, end := 0, len(b.data)-1
	for start <= end {
		mid := start + (end-start)/2
		if b.data[mid].key == key {
			b.data = append(b.data[:mid], b.data[mid+1:]...)
			return
		}
		if b.data[mid].key < key {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
}

type HashMap struct {
	size int64
	data map[int64]*Bucket
}

func NewHashMap(size int64) *HashMap {
	return &HashMap{
		size: size,
		data: make(map[int64]*Bucket, size),
	}
}

func (hm *HashMap) HashFunction(key int64) int64 {
	return key % hm.size
}

func (hm *HashMap) Insert(key int64, value interface{}) {
	hash := hm.HashFunction(key)
	if _, ok := hm.data[hash]; !ok {
		hm.data[hash] = &Bucket{}
	}
	hm.data[hash].Insert(key, value)
}

func (hm *HashMap) Find(key int64) interface{} {
	hash := hm.HashFunction(key)
	if b, ok := hm.data[hash]; ok {
		return b.Find(key)
	}
	return nil
}

func (hm *HashMap) Remove(key int64) {
	hash := hm.HashFunction(key)
	if b, ok := hm.data[hash]; ok {
		b.Remove(key)
	}
}
