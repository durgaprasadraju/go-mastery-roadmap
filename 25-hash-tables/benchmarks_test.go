package hash_tables_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/25-hash-tables"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = hash_tables.New()
	}
}
