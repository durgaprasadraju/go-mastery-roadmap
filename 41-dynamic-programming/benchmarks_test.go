package dynamic_programming_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/41-dynamic-programming"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = dynamic_programming.New()
	}
}
