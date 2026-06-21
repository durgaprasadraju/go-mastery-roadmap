package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/146-circuit-breaker/exercises/solutions"
)

func BenchmarkExercise1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cb := solutions.NewCircuitBreaker(3); cb.Call(func() error { return nil })
	}
}
