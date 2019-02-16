/*
 * Revision History:
 *     Initial: 2019-02-14 22:21    Jon Snow
 */
package lru

import (
	"container/list"
	"sync"
)

// LRU(Least Recently Used)
// This algorithm auto discard last item when the new one added,
// it is used for avoid memory overload.

// item used for store data
type item struct {
	key, value interface{}
}

type LRUCache struct {
	capacity int

	table map[interface{}]*list.Element // hash table
	list  *list.List                    // double link list
	mu    *sync.Mutex
}

func NewLRUCache(capacity int) *LRUCache {
	if capacity < 1 {
		return nil
	}

	return &LRUCache{
		capacity: capacity,
		table:    make(map[interface{}]*list.Element),
		list:     list.New(),
		mu:       &sync.Mutex{},
	}
}

func (c *LRUCache) Len() int { return c.list.Len() }
func (c *LRUCache) Cap() int { return c.capacity }

// removeElement remove the element from list
func (c *LRUCache) removeElement(e *list.Element) (v interface{}) {
	c.mu.Lock()
	value := e.Value
	if value == nil {
		c.mu.Unlock()
		return
	}
	node := value.(*item)
	delete(c.table, node.key)
	c.list.Remove(e)
	c.mu.Unlock()
	return node.value
}

// Set add the new value if the key not existed or rewrite the value
// if the key is existed.
// If LRUCache is full before add a new one, it will delete the
// least-recently-used one and return it.
func (c *LRUCache) Set(k, v interface{}) (rmItem interface{}) {
	c.mu.Lock()
	if element, ok := c.table[k]; ok {
		c.list.MoveToFront(element)
		element.Value.(*item).value = v
		c.mu.Unlock()
		return nil
	}

	newElement := c.list.PushFront(&item{k, v})
	c.table[k] = newElement
	c.mu.Unlock()

	if c.list.Len() > c.capacity {
		lastElement := c.list.Back()
		if lastElement != nil {
			return c.removeElement(lastElement)
		}
	}

	return nil
}

// Get use a key get a element
func (c *LRUCache) Get(k interface{}) (v interface{}) {
	c.mu.Lock()

	if element, ok := c.table[k]; ok {
		c.list.MoveToFront(element)
		c.mu.Unlock()
		return element.Value.(*item).value
	}

	c.mu.Unlock()
	return nil
}

// Del use a key del a element
func (c *LRUCache) Del(k interface{}) {
	if element, ok := c.table[k]; ok {
		c.removeElement(element)
	}
}

// Clear remove all elements
func (c *LRUCache) Clear() {
	for _, v := range c.table {
		c.removeElement(v)
	}
}
