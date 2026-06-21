package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/147-production-projects/exercises/solutions"
)

func BenchmarkExercise1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(solutions.ModuleProgress{Module: "b", Percent: 100}).Completed()
	}
}
