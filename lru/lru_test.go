/*
 * Revision History:
 *     Initial: 2019-02-15 10:32    Jon Snow
 */
package lru

import (
	"strconv"
	"testing"
)

type robot struct {
	Id   int
	Name string
}

func TestLRUCache_Set(t *testing.T) {
	cache := NewLRUCache(10)
	for i := 0; i < 20; i++ {
		cache.Set(i, &robot{i, "robot_" + strconv.Itoa(i)})
	}

	if cache.Len() != 10 {
		t.Errorf("error LRUCache.Len(): data loss")
	}

	for i := 0; i < 20; i++ {
		element := cache.Get(i)
		if element == nil {
			continue
		}
		if element.(*robot).Id != i {
			t.Errorf("error LRUCache.Get(): not match")
		}
	}

	cache.Get(10)                         // use the least recently used element
	cache.Set(20, &robot{20, "robot_20"}) // add a new one
	if element := cache.Get(11); element != nil {
		t.Errorf("error LRU cache failed")
	}

	cache.Clear()
	for i := 0; i < 20; i++ {
		element := cache.Get(i)
		if element != nil {
			t.Errorf("error Clear() failed")
		}
	}
}

func BenchmarkLRUCache_Set(b *testing.B) {
	cache := NewLRUCache(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Set(i, &robot{i, "robot"})
	}
}

func BenchmarkLRUCache_Get(b *testing.B) {
	capacity := 1000000
	cache := NewLRUCache(capacity)
	for i := 0; i < capacity; i++ {
		cache.Set(i, &robot{i, "robot"})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Get(i)
	}
}
