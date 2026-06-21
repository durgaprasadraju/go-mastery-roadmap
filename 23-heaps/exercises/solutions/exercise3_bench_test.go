package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/23-heaps/exercises/solutions"
)

func BenchmarkExercise1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := solutions.NewStack(); s.Push(1)
	}
}
