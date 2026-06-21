// Package caching implements LRU and LFU caches from scratch.
package caching

import (
	"container/list"
	"errors"
	"sync"
)

var ErrKeyNotFound = errors.New("key not found")

// LRUCache is a thread-safe LRU cache with O(1) get/put.
type LRUCache struct {
	mu       sync.Mutex
	capacity int
	ll       *list.List
	items    map[string]*list.Element
}

type lruEntry struct {
	key   string
	value any
}

// NewLRU creates an LRU cache with the given capacity.
func NewLRU(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		ll:       list.New(),
		items:    make(map[string]*list.Element),
	}
}

// Get retrieves a value and marks it recently used.
func (c *LRUCache) Get(key string) (any, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if el, ok := c.items[key]; ok {
		c.ll.MoveToFront(el)
		return el.Value.(*lruEntry).value, nil
	}
	return nil, ErrKeyNotFound
}

// Put inserts or updates a value.
func (c *LRUCache) Put(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if el, ok := c.items[key]; ok {
		el.Value.(*lruEntry).value = value
		c.ll.MoveToFront(el)
		return
	}
	if c.ll.Len() >= c.capacity {
		back := c.ll.Back()
		if back != nil {
			entry := back.Value.(*lruEntry)
			delete(c.items, entry.key)
			c.ll.Remove(back)
		}
	}
	el := c.ll.PushFront(&lruEntry{key: key, value: value})
	c.items[key] = el
}

// Size returns current entry count.
func (c *LRUCache) Size() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ll.Len()
}

// LFUCache is a simplified LFU cache.
type LFUCache struct {
	mu       sync.Mutex
	capacity int
	minFreq  int
	keyMap   map[string]*lfuNode
	freqMap  map[int]*list.List
}

type lfuNode struct {
	key, value any
	freq       int
	elem       *list.Element
}

// NewLFU creates an LFU cache.
func NewLFU(capacity int) *LFUCache {
	return &LFUCache{
		capacity: capacity,
		keyMap:   make(map[string]*lfuNode),
		freqMap:  make(map[int]*list.List),
	}
}

// Get retrieves and increments frequency.
func (c *LFUCache) Get(key string) (any, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	node, ok := c.keyMap[key]
	if !ok {
		return nil, ErrKeyNotFound
	}
	c.updateFreq(node)
	return node.value, nil
}

// Put inserts or updates.
func (c *LFUCache) Put(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if node, ok := c.keyMap[key]; ok {
		node.value = value
		c.updateFreq(node)
		return
	}
	if len(c.keyMap) >= c.capacity {
		list := c.freqMap[c.minFreq]
		back := list.Back()
		if back != nil {
			evict := back.Value.(*lfuNode)
			list.Remove(back)
			delete(c.keyMap, evict.key.(string))
		}
	}
	node := &lfuNode{key: key, value: value, freq: 1}
	c.keyMap[key] = node
	if c.freqMap[1] == nil {
		c.freqMap[1] = list.New()
	}
	node.elem = c.freqMap[1].PushFront(node)
	c.minFreq = 1
}

func (c *LFUCache) updateFreq(node *lfuNode) {
	c.freqMap[node.freq].Remove(node.elem)
	if node.freq == c.minFreq && c.freqMap[node.freq].Len() == 0 {
		c.minFreq++
	}
	node.freq++
	if c.freqMap[node.freq] == nil {
		c.freqMap[node.freq] = list.New()
	}
	node.elem = c.freqMap[node.freq].PushFront(node)
}
