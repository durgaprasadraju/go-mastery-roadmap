package maps_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/15-maps"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = maps.New()
	}
}
