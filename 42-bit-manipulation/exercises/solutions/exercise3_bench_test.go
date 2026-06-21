package solutions_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/42-bit-manipulation/exercises/solutions"
)

func BenchmarkExercise1Core(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = solutions.Exercise1Core(data)
	}
}

func BenchmarkExercise1Transform(b *testing.B) {
	s := "benchmark string for Bit Manipulation"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = solutions.Exercise1Transform(s)
	}
}
