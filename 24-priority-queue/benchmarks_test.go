package priority_queue_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/24-priority-queue"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = priority_queue.New()
	}
}
