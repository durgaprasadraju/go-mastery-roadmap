package greedy_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/40-greedy"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = greedy.New()
	}
}
