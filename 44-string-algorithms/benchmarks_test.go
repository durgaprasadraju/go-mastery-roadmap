package string_algorithms_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/44-string-algorithms"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string_algorithms.New()
	}
}
