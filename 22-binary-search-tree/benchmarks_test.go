package binary_search_tree_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/22-binary-search-tree"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = binary_search_tree.New()
	}
}
