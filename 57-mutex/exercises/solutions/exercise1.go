package solutions

import "sync"

// SafeCounter is mutex-protected.
type SafeCounter struct{ mu sync.Mutex; n int }

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}