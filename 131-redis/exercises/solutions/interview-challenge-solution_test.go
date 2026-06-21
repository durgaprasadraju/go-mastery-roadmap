package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/131-redis/exercises/solutions"
)

func TestIdempotentConsumer(t *testing.T) {
	p := solutions.NewProcessedSet()
	calls := 0
	fn := func() error { calls++; return nil }
	_ = p.Handle("m1", fn)
	_ = p.Handle("m1", fn)
	if calls != 1 {
		t.Fatal("expected single processing")
	}
}
