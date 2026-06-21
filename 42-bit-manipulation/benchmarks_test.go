package bit_manipulation_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/42-bit-manipulation"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = bit_manipulation.New()
	}
}
