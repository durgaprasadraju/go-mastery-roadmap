package monotonic_stack_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/36-monotonic-stack"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = monotonic_stack.New()
	}
}
