package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/37-monotonic-queue/exercises/solutions"
)

func TestQueue(t *testing.T) {
	q := solutions.NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	v, _ := q.Dequeue()
	if v != 1 {
		t.Fatalf("got %d", v)
	}
}
