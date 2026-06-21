package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/113-cqrs/exercises/solutions"
)

func BenchmarkExercise1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := solutions.NewUserStore(); s.Handle(solutions.CreateUserCommand{ID: "1", Email: "a@b.com"})
	}
}
