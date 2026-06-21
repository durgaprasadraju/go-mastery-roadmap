// Package hashtables implements HashMap and HashSet from scratch using chaining.
package hashtables

import "fmt"

const defaultCapacity = 16

// HashMap is a generic string→int map for learning (production uses Go maps).
type HashMap struct {
	buckets [][]entry
	size    int
}

type entry struct {
	key   string
	value int
}

// NewHashMap creates a HashMap with given capacity.
func NewHashMap(capacity int) *HashMap {
	if capacity < 1 {
		capacity = defaultCapacity
	}
	return &HashMap{buckets: make([][]entry, capacity)}
}

func (m *HashMap) hash(key string) int {
	h := 0
	for _, c := range key {
		h = (h*31 + int(c)) % len(m.buckets)
	}
	if h < 0 {
		h = -h
	}
	return h
}

// Put inserts or updates a key-value pair. O(1) average.
func (m *HashMap) Put(key string, value int) {
	idx := m.hash(key)
	for i, e := range m.buckets[idx] {
		if e.key == key {
			m.buckets[idx][i].value = value
			return
		}
	}
	m.buckets[idx] = append(m.buckets[idx], entry{key: key, value: value})
	m.size++
}

// Get retrieves a value. O(1) average.
func (m *HashMap) Get(key string) (int, bool) {
	idx := m.hash(key)
	for _, e := range m.buckets[idx] {
		if e.key == key {
			return e.value, true
		}
	}
	return 0, false
}

// Delete removes a key. O(1) average.
func (m *HashMap) Delete(key string) bool {
	idx := m.hash(key)
	bucket := m.buckets[idx]
	for i, e := range bucket {
		if e.key == key {
			m.buckets[idx] = append(bucket[:i], bucket[i+1:]...)
			m.size--
			return true
		}
	}
	return false
}

// Size returns the number of entries.
func (m *HashMap) Size() int { return m.size }

// HashSet is a set of strings.
type HashSet struct {
	m *HashMap
}

// NewHashSet creates an empty set.
func NewHashSet(capacity int) *HashSet {
	return &HashSet{m: NewHashMap(capacity)}
}

// Add inserts an element.
func (s *HashSet) Add(key string) {
	s.m.Put(key, 1)
}

// Contains reports membership.
func (s *HashSet) Contains(key string) bool {
	_, ok := s.m.Get(key)
	return ok
}

// Remove deletes an element.
func (s *HashSet) Remove(key string) bool {
	return s.m.Delete(key)
}

// String returns a debug representation.
func (m *HashMap) String() string {
	return fmt.Sprintf("HashMap(size=%d)", m.size)
}
