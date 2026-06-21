package sorting_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/31-sorting"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sorting.New()
	}
}
