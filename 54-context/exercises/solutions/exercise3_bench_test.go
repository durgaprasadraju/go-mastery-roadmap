package solutions_test

import (
	"testing"
	"context"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/54-context/exercises/solutions"
)

func BenchmarkExercise1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solutions.WithTimeoutOp(context.Background(), 1)
	}
}
