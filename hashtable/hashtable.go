//
// Revision History:
//     Initial: 2019-02-09 14:37    Jon Snow

// This package is a complete hashtable structure, but benchmark is not
// well, just for view, use "sync.Map" instead this package.

package hashtable

import (
	"container/list"
	"fmt"
	"hash/crc32"
)

type item struct {
	key, value interface{}
}

type Hashtable struct {
	capacity int

	table []*list.List
}

func NewHashTable(capacity int) *Hashtable {
	if capacity < 1 {
		return nil
	}
	return &Hashtable{
		capacity: capacity,
		table:    make([]*list.List, capacity),
	}
}

// TODO: find a better method to instead fmt.Sprint(k)
func (t *Hashtable) getIndex(k interface{}) int {
	return int(crc32.ChecksumIEEE([]byte(fmt.Sprint(k)))) % t.capacity
}

// Set insert a new element or modify the existed one.
// New element always inserted at the front of list.
func (t *Hashtable) Set(k, v interface{}) {
	data := &item{k, v}
	idx := t.getIndex(k)
	if t.table[idx] == nil {
		t.table[idx] = list.New()
		t.table[idx].PushFront(data)
		return
	}

	for e := t.table[idx].Front(); e != nil; e = e.Next() {
		if e.Value.(*item).key == k {
			e.Value = data
			return
		}
	}

	t.table[idx].PushFront(data)
}

// Get a element. If can not match the key k, return nil.
func (t *Hashtable) Get(k interface{}) (v interface{}) {
	idx := t.getIndex(k)
	if t.table[idx] == nil {
		return
	}

	for e := t.table[idx].Front(); e != nil; e = e.Next() {
		item := e.Value.(*item)
		if item.key == k {
			return item.value
		}
	}

	return
}

// Del a element from the matched list, but never delete the list if it existed.
func (t *Hashtable) Del(k interface{}) {
	idx := t.getIndex(k)
	if t.table[idx] == nil {
		return
	}

	if t.table[idx].Len() == 0 {
		return
	}

	for e := t.table[idx].Front(); e != nil; e = e.Next() {
		item := e.Value.(*item)
		if item.key == k {
			t.table[idx].Remove(e)
			break
		}
	}
}
