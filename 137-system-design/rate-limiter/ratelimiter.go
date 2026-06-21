// Package ratelimiter implements token bucket and sliding window rate limiters.
package ratelimiter

import (
	"sync"
	"time"
)

// TokenBucket is a token bucket rate limiter.
type TokenBucket struct {
	mu         sync.Mutex
	rate       float64   // tokens per second
	capacity   float64
	tokens     float64
	lastRefill time.Time
}

// NewTokenBucket creates a rate limiter allowing rate tokens/sec with burst capacity.
func NewTokenBucket(rate float64, capacity int) *TokenBucket {
	return &TokenBucket{
		rate:       rate,
		capacity:   float64(capacity),
		tokens:     float64(capacity),
		lastRefill: time.Now(),
	}
}

// Allow reports whether a request is permitted.
func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill).Seconds()
	tb.tokens += elapsed * tb.rate
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}
	tb.lastRefill = now
	if tb.tokens >= 1 {
		tb.tokens--
		return true
	}
	return false
}

// SlidingWindow counts requests in a rolling time window.
type SlidingWindow struct {
	mu       sync.Mutex
	limit    int
	window   time.Duration
	requests []time.Time
}

// NewSlidingWindow creates a sliding window limiter.
func NewSlidingWindow(limit int, window time.Duration) *SlidingWindow {
	return &SlidingWindow{limit: limit, window: window}
}

// Allow reports whether a request is within the limit.
func (sw *SlidingWindow) Allow() bool {
	sw.mu.Lock()
	defer sw.mu.Unlock()
	now := time.Now()
	cutoff := now.Add(-sw.window)
	// Remove expired entries
	valid := sw.requests[:0]
	for _, t := range sw.requests {
		if t.After(cutoff) {
			valid = append(valid, t)
		}
	}
	sw.requests = valid
	if len(sw.requests) >= sw.limit {
		return false
	}
	sw.requests = append(sw.requests, now)
	return true
}
