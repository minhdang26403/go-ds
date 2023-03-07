package hashtable

import (
	"fmt"
	"strings"

	list "github.com/minhdang26403/algo-ds/linkedlist"
)

const (
	defaultLoadFactor = 0.75
	defaultCapacity   = 8
)

type Entry[K comparable] struct {
	key   K
	value interface{}
}

func (entry *Entry[K]) Key() K {
	return entry.key
}

func (entry *Entry[K]) Value() interface{} {
	return entry.value
}

type HashTable[K comparable] struct {
	capacity, size int
	maxLoadFactor  float32
	table          []list.LinkedList[Entry[K]]
	hashfn         func(key K) int
}

func NewHashTable[K comparable](hash func(key K) int) *HashTable[K] {
	return &HashTable[K]{
		capacity:      defaultCapacity,
		maxLoadFactor: defaultLoadFactor,
		table:         make([]list.LinkedList[Entry[K]], defaultCapacity),
		hashfn:        hash,
	}
}

func (ht *HashTable[K]) Size() int {
	return ht.size
}

func (ht *HashTable[K]) IsEmpty() bool {
	return ht.Size() == 0
}

func (ht *HashTable[K]) Insert(key K, value interface{}) {
	index := ht.keyToIndex(key)
	for node := ht.table[index].Begin(); node != nil; node = node.Next() {
		entry := node.Value()
		if entry.key == key {
			entry.value = value
			return
		}
	}

	if ht.size > int(ht.maxLoadFactor) * ht.capacity {
		ht.growTable()
		index = ht.keyToIndex(key)
	}
	ht.table[index].PushFront(Entry[K]{key, value})
	ht.size++
}

func (ht *HashTable[K]) Delete(key K) {
	index := ht.keyToIndex(key)
	for node := ht.table[index].Begin(); node != nil; node = node.Next() {
		entry := node.Value()
		if entry.key == key {
			ht.table[index].EraseValue(*entry)
			ht.size--
			return
		}
	}
}

func (ht *HashTable[K]) Get(key K) (value interface{}, err error) {
	index := ht.keyToIndex(key)
	for node := ht.table[index].Begin(); node != nil; node = node.Next() {
		entry := node.Value()
		if entry.key == key {
			return entry.value, nil
		}
	}
	return value, fmt.Errorf("The key doesn't exist")
}

func (ht *HashTable[K]) GetKeys() []interface{} {
	keys := make([]interface{}, ht.size)
	i := 0
	for idx := 0; idx < ht.capacity; idx++ {
		for node := ht.table[idx].Begin(); node != nil; node = node.Next() {
			keys[i] = node.Value().key
			i++
		}
	}
	return keys
}

func (ht *HashTable[K]) Contains(key K) bool {
	index := ht.keyToIndex(key)
	for node := ht.table[index].Begin(); node != nil; node = node.Next() {
		entry := node.Value()
		if entry.key == key {
			return true
		}
	}
	return false
}

func (ht *HashTable[K]) String() string {
	if ht.Size() == 0 {
		return "[]"
	}
	var s strings.Builder
	s.WriteByte('[')
	i := 0
	for idx := 0; idx < ht.capacity; idx++ {
		for node := ht.table[idx].Begin(); node != nil; node = node.Next() {
			entry := node.Value()
			i++
			if i < ht.Size() {
				s.WriteString(fmt.Sprintf("{%v,%v}, ", entry.key, entry.value))
			} else {
				s.WriteString(fmt.Sprintf("{%v,%v}]", entry.key, entry.value))
			}
		}
	}
	return s.String()
}

func (ht *HashTable[K]) keyToIndex(key K) int {
	return ht.hashfn(key) % ht.capacity
}

func (ht *HashTable[K]) growTable() {
	oldCapacity := ht.capacity
	ht.capacity *= 2
	newTable := make([]list.LinkedList[Entry[K]], ht.capacity)
	for idx := 0; idx < oldCapacity; idx++ {
		for node := ht.table[idx].Begin(); node != nil; node = node.Next() {
			entry := node.Value()
			newIdx := ht.keyToIndex(entry.key)
			newTable[newIdx].PushFront(*entry)
		}
	}
	ht.table = newTable
}
