package searching_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/32-searching"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = searching.New()
	}
}
