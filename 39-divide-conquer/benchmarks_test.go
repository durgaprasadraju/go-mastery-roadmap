package divide_conquer_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/39-divide-conquer"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = divide_conquer.New()
	}
}
