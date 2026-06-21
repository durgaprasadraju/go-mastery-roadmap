package stacks_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/19-stacks"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = stacks.New()
	}
}
