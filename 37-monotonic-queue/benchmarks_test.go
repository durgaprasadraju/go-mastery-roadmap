package monotonic_queue_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/37-monotonic-queue"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = monotonic_queue.New()
	}
}
