package two_pointers_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/34-two-pointers"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = two_pointers.New()
	}
}
