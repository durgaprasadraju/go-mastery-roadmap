package arrays_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/13-arrays"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = arrays.New()
	}
}
