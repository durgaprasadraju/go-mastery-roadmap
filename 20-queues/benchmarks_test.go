package queues_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/20-queues"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = queues.New()
	}
}
