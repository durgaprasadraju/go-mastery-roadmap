package fast_slow_pointers_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/35-fast-slow-pointers"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fast_slow_pointers.New()
	}
}
