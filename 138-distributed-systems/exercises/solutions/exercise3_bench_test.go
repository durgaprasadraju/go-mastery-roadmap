package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/138-distributed-systems/exercises/solutions"
)

func BenchmarkExercise1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solutions.Encode(3)
	}
}
