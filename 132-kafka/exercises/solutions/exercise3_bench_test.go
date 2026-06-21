package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/132-kafka/exercises/solutions"
)

func BenchmarkExercise1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solutions.NewMessage("t","k",nil)
	}
}
