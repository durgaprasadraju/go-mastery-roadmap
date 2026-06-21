package strings_runes_bytes_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/16-strings-runes-bytes"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings_runes_bytes.New()
	}
}
