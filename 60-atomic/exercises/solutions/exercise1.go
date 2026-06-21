package solutions

import "sync/atomic"

// AtomicCounter uses atomic operations.
type AtomicCounter struct{ n int64 }

func (c *AtomicCounter) Add(v int64) { atomic.AddInt64(&c.n, v) }
func (c *AtomicCounter) Load() int64 { return atomic.LoadInt64(&c.n) }