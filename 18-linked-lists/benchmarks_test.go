package linked_lists_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/18-linked-lists"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = linked_lists.New()
	}
}
