package segment_tree_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/29-segment-tree"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = segment_tree.New()
	}
}
