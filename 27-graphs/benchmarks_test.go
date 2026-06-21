package graphs_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/27-graphs"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = graphs.New()
	}
}
