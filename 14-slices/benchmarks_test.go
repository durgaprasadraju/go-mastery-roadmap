package slices_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/14-slices"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = slices.New()
	}
}
