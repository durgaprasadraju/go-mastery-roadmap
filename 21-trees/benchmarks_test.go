package trees_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/21-trees"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = trees.New()
	}
}
