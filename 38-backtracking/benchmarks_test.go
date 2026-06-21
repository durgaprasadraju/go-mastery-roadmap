package backtracking_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/38-backtracking"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = backtracking.New()
	}
}
