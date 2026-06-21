package goroutines_test

import (
	"context"
	"testing"
	"time"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/49-goroutines"
)

func TestBasicDemo(t *testing.T) {
	ctx := context.Background()
	results := goroutines.BasicDemo(ctx, 5)
	if results[2] != 4 {
		t.Fatalf("got %v", results)
	}
}

func TestWorkerPool(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	jobs := make(chan int, 10)
	go func() {
		for i := 1; i <= 5; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	out := goroutines.WorkerPool(ctx, jobs, 3)
	count := 0
	for range out {
		count++
	}
	if count != 5 {
		t.Fatalf("processed %d jobs", count)
	}
}

func TestAtomicCounter(t *testing.T) {
	var c goroutines.AtomicCounter
	for i := 0; i < 100; i++ {
		c.Inc()
	}
	if c.Value() != 100 {
		t.Fatalf("count = %d", c.Value())
	}
}
