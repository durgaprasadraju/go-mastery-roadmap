package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/129-github-actions/exercises/solutions"
)

func BenchmarkExercise1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solutions.Liveness()
	}
}
