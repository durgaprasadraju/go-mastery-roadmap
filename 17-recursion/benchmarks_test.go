package recursion_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/17-recursion"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = recursion.New()
	}
}
