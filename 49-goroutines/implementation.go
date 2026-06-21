// Package goroutines demonstrates goroutine fundamentals and the GMP scheduler model.
package goroutines

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// SchedulerInfo returns current goroutine and CPU counts.
func SchedulerInfo() string {
	return fmt.Sprintf("GOMAXPROCS=%d goroutines=%d",
		runtime.GOMAXPROCS(0), runtime.NumGoroutine())
}

// BasicDemo spawns goroutines and waits with WaitGroup.
func BasicDemo(ctx context.Context, n int) []int {
	results := make([]int, n)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				results[idx] = idx * idx
			}
		}(i)
	}
	wg.Wait()
	return results
}

// WorkerPool processes jobs with a fixed worker count.
func WorkerPool(ctx context.Context, jobs <-chan int, workers int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	worker := func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case j, ok := <-jobs:
				if !ok {
					return
				}
				out <- j * 2
			}
		}
	}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go worker()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// FanOut distributes work to multiple goroutines.
func FanOut(ctx context.Context, input <-chan int, workers int) []<-chan int {
	outputs := make([]<-chan int, workers)
	for i := 0; i < workers; i++ {
		ch := make(chan int)
		outputs[i] = ch
		go func(out chan<- int) {
			defer close(out)
			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-input:
					if !ok {
						return
					}
					out <- v
				}
			}
		}(ch)
	}
	return outputs
}

// FanIn merges multiple channels into one.
func FanIn(ctx context.Context, channels ...<-chan int) <-chan int {
	merged := make(chan int)
	var wg sync.WaitGroup
	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-c:
					if !ok {
						return
					}
					select {
					case merged <- v:
					case <-ctx.Done():
						return
					}
				}
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(merged)
	}()
	return merged
}

// AtomicCounter demonstrates lock-free increment.
type AtomicCounter struct {
	val atomic.Int64
}

func (c *AtomicCounter) Inc() int64 {
	return c.val.Add(1)
}

func (c *AtomicCounter) Value() int64 {
	return c.val.Load()
}
