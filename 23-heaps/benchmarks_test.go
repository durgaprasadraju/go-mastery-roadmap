package heaps_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/23-heaps"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = heaps.New()
	}
}
