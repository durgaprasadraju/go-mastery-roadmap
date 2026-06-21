package fenwick_tree_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/30-fenwick-tree"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fenwick_tree.New()
	}
}
