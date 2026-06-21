package sliding_window_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/33-sliding-window"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sliding_window.New()
	}
}
