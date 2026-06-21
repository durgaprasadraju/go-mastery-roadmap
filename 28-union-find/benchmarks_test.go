package union_find_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/28-union-find"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = union_find.New()
	}
}
